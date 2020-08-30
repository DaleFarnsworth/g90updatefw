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
	"time"

	"github.com/tarm/serial"
)

type Serial struct {
	port    *serial.Port
	timeout time.Duration
}

const minTimeout = time.Millisecond * 10 // must be greater than 1

func SerialOpen(name string, baud int) (*Serial, error) {
	c := &serial.Config{Name: name, Baud: baud, ReadTimeout: minTimeout}
	port, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	return &Serial{port, minTimeout}, nil
}

func (serial *Serial) Close() error {
	return serial.port.Close()
}

func (serial *Serial) Read(buf []byte) (int, error) {
	iterations := int((serial.timeout + minTimeout - 1*time.Millisecond) / minTimeout)

	index := 0
	for i := 0; i < iterations; i++ {
		c, err := serial.port.Read(buf[index:])
		index += c
		if err != nil {
			return index, err
		}
	}

	return index, nil
}

func (serial *Serial) SetReadTimeout(timeout time.Duration) error {
	serial.timeout = timeout

	return nil
}

func (serial *Serial) Write(buf []byte) (int, error) {
	return serial.port.Write(buf)
}

func (serial *Serial) Flush() error {
	return serial.port.Flush()
}
