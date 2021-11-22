## Xiegu G90 Firmware Updater

The command-line program `g90updatefw` uploads
firmware updates to the Xiegu G90 radio.

Here is the output of `g90updatefw --help`:

    This program is designed to write a firmware file to the Xiegu G90
    radio.  It can be used to update either the main unit or the display unit.
    
        Usage: g90updatefw <firmware_file> <serial_device>
          or   g90updatefw --help
          or   g90updatefw --version
    
    where <firmware_file> is the name of a firmware file for either the
    main unit or for the display unit and <serial_device> is the name of
    the USB serial port connected to the Xiegu G90.  On Linux, the <serial_device> is
    typically /dev/ttyUSB0.  On Windows, com4 works for me as the serial device,
    but your system may be different
    
    You should start the program with the programming cable plugged in
    and the power disconnected from the radio.  After starting the program,
    reconnect the power cable and power-on the radio.  The program runs
    without any user interaction.

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

A Linux/arm executable (suitable for the Raspberry Pi with a 32-bit OS) may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/linux/arm](
https://www.farnsworth.org/dale/g90updatefw/downloads/linux/arm).

A Linux/arm64 executable (suitable for the Raspberry Pi with a 64-bit OS) may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/linux/arm64](
https://www.farnsworth.org/dale/g90updatefw/downloads/linux/arm64).

Executables for Linux on PPC, RISCV and S390X can be found below
[https://www.farnsworth.org/dale/g90updatefw/downloads/linux](
https://www.farnsworth.org/dale/g90updatefw/downloads/linux).

A MacOS/amd64 executable may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/darwin/amd64](
https://www.farnsworth.org/dale/g90updatefw/downloads/darwin/amd64).

A MacOS/arm64 executable may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/darwin/arm64](
https://www.farnsworth.org/dale/g90updatefw/downloads/darwin/arm64).

Executables for other OS and Architecture variants can be found starting at 
[https://www.farnsworth.org/dale/g90updatefw/downloads](
https://www.farnsworth.org/dale/g90updatefw/downloads).

I'll entertain requests to make executables for other OS/Architecture
combinations available.

NOTE: You will likely have to set execute permissions on the binary file after
downloading.  On Linux this may be done by "chmod 755 <filename>".

Dale Farnsworth dale@farnsworth.org

P.S.  KB0OTY made a video about using g90updatefw on the Raspberry Pi.
[https://www.youtube.com/watch?v=tnU0LtWxqOs](
https://www.youtube.com/watch?v=tnU0LtWxqOs).
Thanks Andrew!
