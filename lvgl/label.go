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
import (
	"unsafe"
)

// Label returns a new Label object
// @param scr is the screen to put the label on
// @param copy is another label to copy the styling from
// @return the created label object
func Label(scr, copy *LVObj) *LVObj {
	p1 := (*C.struct__lv_obj_t)(unsafe.Pointer(scr))
	p2 := (*C.struct__lv_obj_t)(unsafe.Pointer(copy))

	label := C.lv_label_create(p1, p2)
	return (*LVObj)(unsafe.Pointer(label))
}

// SetText ...
func (obj *LVObj) SetText(str string) {
	txt := C.CString(str)
	defer C.free(unsafe.Pointer(txt))
	C.lv_label_set_text((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), txt)
}
