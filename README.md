# go-lvgl
lvgl.io Go bindings

## Description
This is an experiment trying to create a LVGL user interface using CGO bindings.
Currently this project is configured/tested for usage on a Raspberry Pi4, with
the official [Raspberry 7" touchscreen](https://www.raspberrypi.org/products/raspberry-pi-touch-display/).

## Usage
On a Raspberry Pi with 7" touchscreen:

Clone this repository, then run `make init` only once to initialize the git submodules,
and create the necessary directories and symlinks.

Next the LVGL library needs to be built, using: `make lvgl`. This will compile the
LVGL C code, and create a shared and static library for it in `include/obj`. Either
of these is needed to compile the Go code.

Finally run `make` to build the Go bindings and example code. This will create a
statically linked binary called `golvgl`. Simply run this binary to see
the GUI in operation.

## Configuration
In order to use this project/LVGL with other hardware and/or platforms, you will need
to adjust the settings in `include/lv_conf.h` and `include/lv_drv_conf.h`, and
recompile the library with `make lvgl`.
