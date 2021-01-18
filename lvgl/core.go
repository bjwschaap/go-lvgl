package lvgl

// The lv_obj Base Object
// https://docs.lvgl.io/v7/en/html/widgets/obj.html#overview

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
import "unsafe"

// LVObj is the base object that implements the basic
// properties of widgets on a screen.
type LVObj C.struct__lv_obj_t

// GetChild helps to iterate through the children of an object
func GetChild(obj, child *LVObj) (*LVObj, error) {
	var p1, p2 *C.struct__lv_obj_t
	p1 = (*C.struct__lv_obj_t)(unsafe.Pointer(obj))
	if child != nil {
		p2 = (*C.struct__lv_obj_t)(unsafe.Pointer(child))
	} else {
		p2 = nil
	}

	lvobj, err := C.lv_obj_get_child(p1, p2)
	if err != nil {
		return nil, err
	}

	return (*LVObj)(unsafe.Pointer(lvobj)), nil
}

// Align ...
func (obj *LVObj) Align(base *LVObj, align C.uchar, x int, y int) error {
	var ba *C.struct__lv_obj_t
	ba = (*C.struct__lv_obj_t)(unsafe.Pointer(base))
	_, err := C.lv_obj_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), ba, align, C.short(x), C.short(y))
	return err
}
