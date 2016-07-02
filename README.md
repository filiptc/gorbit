
#GOrbit


Simple HTTP MJPEG streaming for V4L2 webcams and pan/tilt control (for Logitech® Orbit AF only). Written in Go.


## Requirements
* Own a V4L2 compatible camera and install v4l2 packages.
* (optional) For controlling Logitech® Orbit AF, scroll down to [this section](#logitech®-orbit-af-control).

## Install
Download latest release for your architecture (Linux only). Be aware that the cross compiling for ARM devices
(e.g. Raspberry Pi) will not work for the time being, due to CGO compiler limitations. If you need to
compile for ARM devices, you will need to do so from within the device (and vice versa).

## Usage
#### CLI controls:
* Reset pan/tilt: `gorbit control -r`
* Pan (100 steps*): `gorbit control -p 100`
* Tilt (100 steps*): `gorbit control -t 100`

(* 1 step = 1 / 64 degree)

#### Web controls:
* Start server: `gorbit serve`
* Open `http://localhost:8001` on your browser
* Click on a point in image to pan/tilt, so that that point becomes the new center.

## Build from sources
Install Go on your system. Run `go install github.com/filiptc/gorbit`. Binary will be available on `$GOPATH/bin/gorbit`.


## Logitech® Orbit AF control
In order to correctly make GOrbit control pan/tilt/reset you need to install uvcdynctrl
(`sudo apt-get install uvcdynctrl` on Debian) and import Logitech® controls. This can be done running
`uvcdynctrl --import=/path/to/logitech.xml`. The default file provided by the uvcdynctrl package works for
pan/tilt but doesn't allow for pan/tilt simultaneous reset. Download the following [logitech.xml](https://raw.githubusercontent.com/llmike/v4l2-tools/master/libwebcam-src-0.2.4/uvcdynctrl/data/046d/logitech.xml)
to allow pan/tilt resets.


## The MIT License (MIT)
Copyright (c) 2016 Philip Thomas Casado

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
