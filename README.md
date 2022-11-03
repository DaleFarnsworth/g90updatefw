## Xiegu G90 and Xiego G106 Firmware Updater

The command-line programs `g90updatefw` (alias `g106updatefw`) uploads
firmware updates to the Xiegu G90 and Xiegu G106 radios.

Note that g90updatefw can be used to update firmware on either
the G90 or the G106 radios.
g106updatefw and g90updatefw are two names for the same program.

Here is the output of `g90updatefw --help`:

    g90updatefw version 1.5

    This program is designed to write a firmware file to a Xiegu radio.
    It can be used to update either the main unit or the display unit.

	Usage: g90updatefw <firmware_file> <serial_device>
	  or   g90updatefw -h or g90updatefw --help
	  or   g90updatefw -v or g90updatefw --version

    where <firmware_file> is the name of a firmware file for either the
    main unit or for the display unit and <serial_device> is the name of
    the serial port connected to the Xiegu radio.  On non-windows machines
    the <serial_device> is typically similar to /dev/ttyUSB2. On windows
    machines it will be similar to COMM2.

    Specifying -h or --help produces this help message.
    Specifying -v or --version prints the program version.

    You should start the program with the programming cable plugged in
    and the power disconnected from the radio.

The output from g106updatefw --help is extremely similar.

Source code and additional information about `g90updatefw` and `g106updatefw` may be found at
[https://github.com/DaleFarnsworth/g90updatefw](
https://github.com/DaleFarnsworth/g90updatefw).

A Windows/386 executable (suitable for 32-bit and 64-bit Windows) may be downloaded
from [downloads/windows/386](
downloads/windows/386).
This program needs to be run from a cmd.exe window that has administrator privileges.

Some people have reported receiving a virus warning when downloading the above Windows/386
version of g90updatefw.  They have also reported that the 64-bit executable, which is
fine for most current windows system, does not trigger the virus report.

A Windows/amd64 executable (suitable 64-bit Windows) may be downloaded
from [downloads/windows/amd64](
downloads/windows/amd64).
This program needs to be run from a cmd.exe window that has administrator privileges.

A Linux/amd64 executable may be downloaded from
[downloads/linux/amd64](
downloads/linux/amd64).

A MacOS/amd64 executable may be downloaded from
[downloads/darwin/amd64](
downloads/darwin/amd64).

A MacOS/arm64 executable may be downloaded from
[downloads/darwin/arm64](
downloads/darwin/arm64).

A Linux/arm executable (suitable for the Raspberry Pi with a 32-bit OS) may be downloaded from
[downloads/linux/arm](
downloads/linux/arm).

A Linux/arm64 executable (suitable for the Raspberry Pi with a 64-bit OS) may be downloaded from
[downloads/linux/arm64](
downloads/linux/arm64).

Executables for Linux on PPC, RISCV and S390X can be found below
[downloads/linux](
downloads/linux).

Executables for all supported OS and Architecture variants can be found starting at 
[downloads](
downloads).

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
