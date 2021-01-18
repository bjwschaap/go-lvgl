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
	"errors"
	"unsafe"
)

// GetActiveScreen returns the active screen object
func GetActiveScreen() (*LVObj, error) {
	scr, err := C.lv_scr_act()
	if err != nil {
		return nil, err
	}

	return (*LVObj)(unsafe.Pointer(scr)), nil
}

// LoadScreen loads an LVObject as screen
func LoadScreen(scr *LVObj) error {
	if scr == nil {
		return errors.New("scr can't be nil")
	}

	p1 := (*C.struct__lv_obj_t)(unsafe.Pointer(scr))
	_, err := C.lv_scr_load(p1)
	return err
}

// GetScreen returns the screen for an object
func GetScreen(obj *LVObj) (*LVObj, error) {
	if obj == nil {
		return nil, errors.New("obj can't be nil")
	}

	p1 := (*C.struct__lv_obj_t)(unsafe.Pointer(obj))
	scr, err := C.lv_obj_get_screen(p1)
	if err != nil {
		return nil, err
	}

	return (*LVObj)(unsafe.Pointer(scr)), nil
}
