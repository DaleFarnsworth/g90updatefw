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

// +build linux darwin

package main

import (
	"time"

	"github.com/pkg/term"
)

type Serial struct {
	term *term.Term
}

func SerialOpen(name string, baud int) (*Serial, error) {
	term, err := term.Open(name, term.Speed(baud), term.RawMode)
	if err != nil {
		return nil, err
	}

	return &Serial{term}, nil
}

func (serial *Serial) Close() {
	serial.term.Close()
}

func (serial *Serial) Read(buf []byte) (int, error) {
	return serial.term.Read(buf)
}

func (serial *Serial) SetReadTimeout(timeout time.Duration) {
	serial.term.SetReadTimeout(timeout)
}

func (serial *Serial) Write(buf []byte) (int, error) {
	return serial.term.Write(buf)
}

func (serial *Serial) Flush() {
	serial.term.Flush()
}
