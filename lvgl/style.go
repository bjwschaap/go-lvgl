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

// Init initializes a new lv_style_t object
func (style *LVStyle) Init() {
	s := (*C.struct___7)(unsafe.Pointer(style))
	C.lv_style_init(s)
}

// Copy copies a style to a destination style
// @param dest The destination style
// @return error
func (style *LVStyle) Copy(dest *LVStyle) {
	s := (*C.struct___7)(unsafe.Pointer(style))
	d := (*C.struct___7)(unsafe.Pointer(dest))
	C.lv_style_copy(d, s)
}

// Reset clears all properties from a style and all allocated memories
// @param dest The destination style
// @return error
func (style *LVStyle) Reset() {
	s := (*C.struct___7)(unsafe.Pointer(style))
	C.lv_style_reset(s)
}

// SetBgColor sets the backgroundcolor property on a style, for a specific state
func (style *LVStyle) SetBgColor(state uint8, color LVColor) {
	s := (*C.struct___7)(unsafe.Pointer(style))
	C.lv_style_set_bg_color(s, C.uchar(state), C.lv_color_t(color))
}

// AddStyle adds a style to a LVObj
func (obj *LVObj) AddStyle(part uint8, style *LVStyle) {
	s := (*C.struct___7)(unsafe.Pointer(style))
	C.lv_obj_add_style((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uchar(part), s)
}
