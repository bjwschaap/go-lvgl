package lvgl

/*
  #cgo CFLAGS: -I../include -I../include/lvgl -I../include/lv_drivers/display -I../include/lv_drivers/indev
  #include "lv_conf.h"
  #include "lv_drv_conf.h"
  #include "lvgl.h"
  #include "fbdev.h"
  #include "evdev.h"

  #cgo LDFLAGS: -L../include -llvgl
*/
import "C"

import (
	"unsafe"
)

// GetActiveScreen returns the active screen object
func GetActiveScreen() *LVObj {
	scr := C.lv_scr_act()
	return (*LVObj)(unsafe.Pointer(scr))
}

// LoadScreen loads an LVObject as screen
func LoadScreen(scr *LVObj) {
	C.lv_scr_load((*C.struct__lv_obj_t)(unsafe.Pointer(scr)))
}

// LoadScreenAnim loads an LVObject as screen, with animation
func LoadScreenAnim(scr *LVObj, anim LVLoadScreenAnimation, time uint, delay uint, autoDel bool) {
	C.lv_scr_load_anim(
		(*C.struct__lv_obj_t)(unsafe.Pointer(scr)),
		(C.lv_scr_load_anim_t)(anim),
		C.uint(time),
		C.uint(delay),
		C.bool(autoDel))
}

// GetScreen returns the screen for an object
func (obj *LVObj) GetScreen() *LVObj {
	scr := C.lv_obj_get_screen((*C.struct__lv_obj_t)(unsafe.Pointer(obj)))
	return (*LVObj)(unsafe.Pointer(scr))
}
