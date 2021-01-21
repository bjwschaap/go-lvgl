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

// Tabview creates a new tabview on a screen
// @param copy is another label to copy the styling from
// @return the created tabview object
func (obj *LVObj) Tabview(copy *LVObj) *LVObj {
	p1 := (*C.struct__lv_obj_t)(unsafe.Pointer(obj))
	p2 := (*C.struct__lv_obj_t)(unsafe.Pointer(copy))

	label := C.lv_tabview_create(p1, p2)
	return (*LVObj)(unsafe.Pointer(label))
}

// AddTab adds a new tab with the given name to the tabview
// @param tabview pointer to Tab view object where to ass the new tab
// @param name the text on the tab button
// @return pointer to the created page object (lv_page). You can create your content here
func (obj *LVObj) AddTab(name string) *LVObj {
	n := C.CString(name)
	defer C.free((unsafe.Pointer)(n))
	page := C.lv_tabview_add_tab((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), n)
	return (*LVObj)(unsafe.Pointer(page))
}
