## Xiegu G90 and Xiego G106 Firmware Updater

The command-line programs `g90updatefw` (alias `g106updatefw`) uploads
firmware updates to the Xiegu G90 and Xiegu G106 radios.

Note that g90updatefw can be used to update firmware on either
the G90 or the G106 radios.

Here is the output of `g90updatefw --help`:

    g90updatefw version 1.4

    This program is designed to write a firmware file to a Xiegu radio.
    It can be used to update either the main unit or the display unit.

	Usage: g90updatefw [options] <firmware_file> <serial_device>
	  or   g90updatefw -h or g90updatefw --help
	  or   g90updatefw -v or g90updatefw --version

    where <firmware_file> is the name of a firmware file for either the
    main unit or for the display unit and <serial_device> is the name of
    the serial port connected to the Xiegu radio.  On non-windows machines
    the <serial_device> is typically /dev/ttyUSB0.

    Specifying -h or --help produces this help message.
    Specifying -v or --version prints the program version.

    Options:
	-g90, --g90, -G90, --G90
	    Specifies that the target radio is an original Xiegu G90
	    that requires typing a character to interrupt the
	    bootloader to enable loading new firmware.
	    This is the default if the firmware filename or the program
	    name contains "g90" or "G90".

	-g90v, --g90v, -G90V, --G90V
	    Specifies that the target radio is a newer Xiegu G90V
	    that requires holding in the volume control while
	    powering on to enable loading new firmware.

	-g106, --g106, -G106, --G106
	    Specifies that the target radio is a Xiegu G106.
	    This is the default if the firmware filename or the program
	    name contains "g106" or "G106".

    To update a G90 radio, specify the --g90 option or ensure that the
    firmware filename or the program name contains the string "g90" or "G90".

    Newer G90 radios may require specifying the --g90v option. This is needed
    if the G90 display or main unit requires holding in the volume control while
    powering on the unit to enable loading new firmware.

    To update a G106 radio, specify the --g106 option or ensure that the
    firmware filename or the program name contains the string "g106" or "G106".

    You should start the program with the programming cable plugged in
    and the power disconnected from the radio.

Source code and additional information about `g90updatefw` may be found at
[https://github.com/DaleFarnsworth/g90updatefw](
https://github.com/DaleFarnsworth/g90updatefw).

A Windows/386 executable (suitable for 32-bit and 64-bit Windows) may be downloaded
from [https://www.farnsworth.org/dale/g90updatefw/downloads/windows/386](
https://www.farnsworth.org/dale/g90updatefw/downloads/windows/386).
This program needs to be run from a cmd.exe window that has administrator privileges.

Some people have reported receiving a virus warning when downloading the above Windows/386
version of g90updatefw.  They have also reported that the 64-bit executable, which is
fine for most current windows system, does not trigger the virus report.

A Windows/amd64 executable (suitable 64-bit Windows) may be downloaded
from [https://www.farnsworth.org/dale/g90updatefw/downloads/windows/amd64](
https://www.farnsworth.org/dale/g90updatefw/downloads/windows/amd64).
This program needs to be run from a cmd.exe window that has administrator privileges.

A Linux/amd64 executable may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/linux/amd64](
https://www.farnsworth.org/dale/g90updatefw/downloads/linux/amd64).

A MacOS/amd64 executable may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/darwin/amd64](
https://www.farnsworth.org/dale/g90updatefw/downloads/darwin/amd64).

A MacOS/arm64 executable may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/darwin/arm64](
https://www.farnsworth.org/dale/g90updatefw/downloads/darwin/arm64).

Executables for other OS and Architecture variants can be found starting at 
[https://www.farnsworth.org/dale/g90updatefw/downloads](
https://www.farnsworth.org/dale/g90updatefw/downloads).

A Linux/arm executable (suitable for the Raspberry Pi with a 32-bit OS) may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/linux/arm](
https://www.farnsworth.org/dale/g90updatefw/downloads/linux/arm).

A Linux/arm64 executable (suitable for the Raspberry Pi with a 64-bit OS) may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/linux/arm64](
https://www.farnsworth.org/dale/g90updatefw/downloads/linux/arm64).

Executables for Linux on PPC, RISCV and S390X can be found below
[https://www.farnsworth.org/dale/g90updatefw/downloads/linux](
https://www.farnsworth.org/dale/g90updatefw/downloads/linux).

I'll entertain requests to make executables for other OS/Architecture
combinations available.

NOTE: You will likely have to set execute permissions on the binary file after
downloading.  On Linux this may be done by "chmod 755 <filename>".

Dale Farnsworth dale@farnsworth.org

P.S. Andrew, KB0OTY, made a video about using g90updatefw on the Raspberry Pi.
[https://www.youtube.com/watch?v=tnU0LtWxqOs](
https://www.youtube.com/watch?v=tnU0LtWxqOs).

Also Steve, KM9G, a video about using g90updatefw on Windows.
[https://www.youtube.com/watch?v=NIUbbSjRXsU](
https://www.youtube.com/watch?v=NIUbbSjRXsU).

Thanks Andrew and Steve!
