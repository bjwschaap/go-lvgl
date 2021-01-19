package lvgl

/*
  #cgo CFLAGS: -I../include -I../include/lvgl -I../include/lv_drivers/display -I../include/lv_drivers/indev
  #include "lv_conf.h"
  #include "lv_drv_conf.h"
  #include "lvgl.h"
  #include "fbdev.h"
  #include "evdev.h"
  #include <stdlib.h>

  #cgo LDFLAGS: -L../include -llvgl
*/
import "C"
import "unsafe"

// LVStyle contains an object's styling information
type LVStyle C.struct__lv_style_t

// StyleInit initializes a new lv_style_t object
func StyleInit(style *LVStyle) error {
	s := (*C.struct___6)(unsafe.Pointer(style))
	_, err := C.lv_style_init(s)
	return err
}

// Copy copies a style to a destination style
// @param dest The destination style
// @return error
func (style *LVStyle) Copy(dest *LVStyle) error {
	s := (*C.struct___6)(unsafe.Pointer(style))
	d := (*C.struct___6)(unsafe.Pointer(dest))
	_, err := C.lv_style_copy(d, s)
	return err
}

// Reset clears all properties from a style and all allocated memories
// @param dest The destination style
// @return error
func (style *LVStyle) Reset() error {
	s := (*C.struct___6)(unsafe.Pointer(style))
	_, err := C.lv_style_reset(s)
	return err
}
