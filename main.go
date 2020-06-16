// Copyright 2020 Dale Farnsworth. All rights reserved.

// Dale Farnsworth
// 1007 W Mendoza Ave
// Mesa, AZ  85210
// USA
//
// dale@farnsworth.org

// This program is free software: you can redistribute it and/or modify
// it under the terms of version 3 of the GNU General Public License
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/dalefarnsworth/go-xmodem/xmodem"
	"github.com/pkg/term"
)

const (
	versionMajor = 1
	versionMinor = 0

	waitTimeout   = 20 * time.Millisecond
	uploadTimeout = 10 * time.Second

	banner = "Hit a key to abort"
	menu   = "1.Update FW"
	waitFW = "Wait FW file"

	attentionGrabber = " "
	menuSelector     = "1"

	buflen = 16 * 1024
)

var progname string

func readString(term *term.Term) string {
	buf := make([]byte, buflen)

	i := 0
	for {
		n, err := term.Read(buf[i:])
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			if i == 0 {
				continue
			}
			break
		}
		syscall.Write(syscall.Stdout, buf[i:i+n])
		i += n
	}

	if i >= len(buf) {
		log.Fatal(errors.New("Read buffer overrun"))
	}

	return string(buf[0:i])
}

func expectSend(term *term.Term, expect, send string) {
	fmt.Printf("> Waiting for '%s'...\n\n", expect)

	str := readString(term)
	if !strings.Contains(str, expect) {
		log.Fatalf("'%s' not found.", expect)
	}

	if len(send) != 0 {
		_, err := term.Write([]byte(send))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func updateG90(term *term.Term, data []byte) {
	term.Flush()

	term.SetReadTimeout(waitTimeout)
	expectSend(term, banner, attentionGrabber)
	fmt.Println()

	term.SetReadTimeout(waitTimeout)
	expectSend(term, menu, menuSelector)
	fmt.Println()

	term.SetReadTimeout(waitTimeout)
	expectSend(term, waitFW, "")
	fmt.Printf("\n\n> Uploading %d bytes.\n", len(data))

	term.SetReadTimeout(uploadTimeout)
	counter := 0
	previousBlock := -1
	callback := func(block int) {
		if counter%40 == 0 {
			if counter != 0 {
				fmt.Print("\n")
			}
			fmt.Print("> ")
		}
		marker := "."
		if block != previousBlock+1 {
			marker = "R"
		}
		fmt.Print(marker)
		counter++
		previousBlock = block
	}
	err := xmodem.ModemSend1K(term, data, callback)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n> Upload complete.")

	term.Flush()
}

func usage(strs ...string) {
	if len(strs) > 0 {
		for _, str := range strs {
			fmt.Fprintln(os.Stderr, str)
		}
		fmt.Fprintln(os.Stderr)
	}
	fmt.Fprintf(os.Stderr, "Usage: %s <firmware_file> <serial_device>\n", progname)
	fmt.Fprintf(os.Stderr, "  or   %s --help\n", progname)
	fmt.Fprintf(os.Stderr, "  or   %s --version\n", progname)
	os.Exit(1)
}

func help() {
	help := `
This program is designed to write a firmware file to the Xiegu G90
radio.  It can be used to update either the main unit or the display unit.

    Usage: %s <firmware_file> <serial_device>
      or   %s --help
      or   %s --version

where <firmware_file> is the name of a firmware file for either the
main unit or for the display unit and <serial_device> is the name of
the serial port connected to the Xiegu G90.  The <serial_device> is
typically /dev/ttyUSB0.

You should start the program with the programming cable plugged in
and the power disconnected from the radio.  After starting the program,
reconnect the power cable and power-on the radio.  The program runs
without any user interaction.`

	fmt.Printf(help, progname, progname, progname)
	fmt.Println()
}

func instructions() {
	instructions := `
> 1. Disconnect power cable from radio.
> 2. Reconnect power cable to radio.
> 3. Power-on the radio.
`

	fmt.Println(instructions[1:])
}

func version() {
	fmt.Printf("%s version %d.%d\n", progname, versionMajor, versionMinor)
}

func main() {
	progname = filepath.Base(os.Args[0])
	log.SetPrefix(progname + ": ")
	log.SetFlags(log.Lshortfile)

	if len(os.Args) != 3 {
		if len(os.Args) == 2 {
			switch os.Args[1] {
			case "-h", "--help":
				help()

			case "-v", "--version":
				version()
			}
			os.Exit(0)
		}
		usage()
	}

	fwFilename := os.Args[1]
	devName := os.Args[2]

	term, err := term.Open(devName, term.Speed(115200), term.RawMode)
	if err != nil {
		usage(err.Error())
	}
	defer term.Close()

	data, err := ioutil.ReadFile(fwFilename)
	if err != nil {
		usage(err.Error())
	}

	instructions()

	updateG90(term, data)
}
