## Xiegu G90 Firmware Updater

The command-line program `g90updatefw` runs on Linux and uploads
firmware updates to the Xiegu G90 Radio.

Here is the output of `g90updatefw --help`:

    This program is designed to write a firmware file to the Xiegu G90
    radio.  It can be used to update either the main unit or the display unit.
    
        Usage: g90updatefw <firmware_file> <serial_device>
          or   g90updatefw --help
          or   g90updatefw --version
    
    where <firmware_file> is the name of a firmware file for either the
    main unit or for the display unit and <serial_device> is the name of
    the serial port connected to the Xiegu G90.  The <serial_device> is
    typically /dev/ttyUSB0.
    
    You should start the program with the programming cable plugged in
    and the power disconnected from the radio.  After starting the program,
    reconnect the power cable and power-on the radio.  The program runs
    without any user interaction.

A Linux/amd64 executable may be downloaded from
[https://www.farnsworth.org/dale/g90updatefw/downloads/linux/amd64](
https://www.farnsworth.org/dale/g90updatefw/downloads/linux/amd64).

Executables for other Linux (and MacOS) variants can be found at
[https://www.farnsworth.org/dale/g90updatefw/downloads](
https://www.farnsworth.org/dale/g90updatefw/downloads).

I'll entertain requests to make executables for other OS/Architecture
combinations available.  Unfortunately, the serial library I'm using
has no windows support, so it's very unlikely that I'll produce a
windows version.

Dale Farnsworth dale@farnsworth.org
