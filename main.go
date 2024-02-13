// Copyright 2020-2022 Dale Farnsworth. All rights reserved.

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
)

const (
	versionMajor = "1"
	versionMinor = "7"

	buflen = 64 * 1024
)

var progname string

func readString(serial *Serial) string {
	buf := make([]byte, buflen)

	i := 0
	lastReadZeroBytes := false
	for {
		n, err := serial.Read(buf[i:])
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			if i == 0 {
				continue
			}
			if lastReadZeroBytes {
				// only return
				break
			}
			lastReadZeroBytes = true
			continue
		}
		lastReadZeroBytes = false
		syscall.Write(syscall.Stdout, buf[i:i+n])
		i += n
	}

	if i >= len(buf) {
		log.Fatal(errors.New("Read buffer overrun"))
	}

	return string(buf[0:i])
}

func expect(serial *Serial, expects []string) (expectIndex int) {
	previousStr := ""

	for {
		str := readString(serial)
		for i, expect := range expects {
			if strings.Contains(previousStr+str, expect) {
				return i
			}
		}
		previousStr = str
	}

	panic("unreachable")
}

func expectSend(serial *Serial, expects []string, sends []string) (whichExpect string) {
	if len(sends) != len(expects) {
		panic("length of sends array does not equal length of expects array")
	}

	fmt.Printf("> Waiting for '%s'...\n\n", strings.Join(expects, "' or '"))

	expectIndex := expect(serial, expects)
	send := sends[expectIndex]

	if len(send) != 0 {
		_, err := serial.Write([]byte(send))
		if err != nil {
			log.Fatal(err)
		}
	}

	return expects[expectIndex]
}

func instructions() {
	fmt.Print("> ")
	version(os.Stdout)
	instructions := `
> 1. Disconnect power cable from the radio.
> 2. Reconnect power cable to the radio.
> 3. Press the volume button and while holding it in,
> 4. Press the power button until the radio begins erasing the existing firmware.
`
	fmt.Println(instructions[1:])
}

func updateRadio(serial *Serial, data []byte) {
	attentionTimeout := 10 * time.Millisecond
	menuTimeout := 50 * time.Millisecond
	eraseTimeout := 50 * time.Millisecond
	uploadTimeout := 10 * time.Second
	cleanupTimeout := 500 * time.Millisecond

	banner := "Hit a key to abort"
	menu := "1.Update FW"
	waitFW := "Wait FW file"

	attentionGrabber := " "
	menuSelector := "1"

	serial.Flush()

	expects := []string{banner, menu}
	sends := []string{attentionGrabber, menuSelector}

	serial.SetReadTimeout(attentionTimeout)
	found := expectSend(serial, expects, sends)
	fmt.Println()

	if found != menu {
		serial.SetReadTimeout(menuTimeout)
		expectSend(serial, []string{menu}, []string{menuSelector})
		fmt.Println()
	}

	serial.SetReadTimeout(eraseTimeout)
	expectSend(serial, []string{waitFW}, []string{""})
	fmt.Printf("\n\n> Uploading %d bytes.\n", len(data))

	serial.SetReadTimeout(uploadTimeout)
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
	err := xmodem.ModemSend1K(serial, data, callback)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n> Upload complete.")

	serial.SetReadTimeout(cleanupTimeout)
	readString(serial)
}

func usage(strs ...string) {
	version(os.Stderr)
	fmt.Fprintln(os.Stderr)

	if len(strs) > 0 {
		for _, str := range strs {
			fmt.Fprintln(os.Stderr, str)
		}
		fmt.Fprintln(os.Stderr)
	}
	fmt.Fprintf(os.Stderr, "Usage: %s <firmware_file> <serial_device>\n", progname)
	fmt.Fprintf(os.Stderr, "    or %s --help\n", progname)
	fmt.Fprintf(os.Stderr, "    or %s --version\n\n", progname)
	fmt.Fprintf(os.Stderr, "    Use the --help option for more detailed usage information.\n")
	os.Exit(1)
}

func help() {
	version(os.Stdout)
	help := `
This program is designed to write a firmware file to a Xiegu radio.
It can be used to update either the main unit or the display unit.

    Usage: %s <firmware_file> <serial_device>
      or   %s -h or %s --help
      or   %s -v or %s --version

where <firmware_file> is the name of a firmware file for either the
main unit or for the display unit and <serial_device> is the name of
the serial port connected to the Xiegu radio.  On non-windows machines
the <serial_device> is typically similar to /dev/ttyUSB2. On windows
machines it will be similar to COMM2.

Specifying -h or --help produces this help message.
Specifying -v or --version prints the program version.

You should start the program with the programming cable plugged in
and the power disconnected from the radio.
`
	fmt.Printf(help, progname, progname, progname, progname, progname)
	fmt.Println()
}

func version(f *os.File) {
	fmt.Fprintf(f, "%s version %s.%s\n", progname, versionMajor, versionMinor)
}

func main() {
	progname = filepath.Base(os.Args[0])
	log.SetPrefix(progname + ": ")
	log.SetFlags(log.Lshortfile)
	args := os.Args[1:]

	for len(args) > 0 && args[0][0] == '-' {
		switch args[0] {
		case "-h", "--help":
			help()
			os.Exit(0)

		case "-v", "--version":
			version(os.Stdout)
			os.Exit(0)

		default:
			usage("Bad option: " + args[0])
		}

		args = args[1:]
	}

	if len(args) != 2 {
		usage()
	}

	fwFilename := args[0]
	devName := args[1]

	serial, err := SerialOpen(devName, 115200)
	if err != nil {
		usage(err.Error())
	}
	defer serial.Close()

	data, err := ioutil.ReadFile(fwFilename)
	if err != nil {
		usage(err.Error())
	}

	instructions()

	updateRadio(serial, data)
}
