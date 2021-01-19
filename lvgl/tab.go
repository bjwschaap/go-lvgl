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
	"errors"
	"unsafe"
)

// TabviewCreate creates a new tabview
// @param scr is the screen to put the label on
// @param copy is another label to copy the styling from
// @return the created tabview object
func TabviewCreate(scr, copy *LVObj) (*LVObj, error) {
	if scr == nil {
		return nil, errors.New("screen can't be nil")
	}
	p1 := (*C.struct__lv_obj_t)(unsafe.Pointer(scr))
	p2 := (*C.struct__lv_obj_t)(unsafe.Pointer(copy))

	label, err := C.lv_tabview_create(p1, p2)
	if err != nil {
		return nil, err
	}

	return (*LVObj)(unsafe.Pointer(label)), nil
}

// AddTab adds a new tab with the given name to the tabview
// @param tabview pointer to Tab view object where to ass the new tab
// @param name the text on the tab button
// @return pointer to the created page object (lv_page). You can create your content here
func (tv *LVObj) AddTab(name string) (*LVObj, error) {
	n := C.CString(name)
	defer C.free((unsafe.Pointer)(n))
	page, err := C.lv_tabview_add_tab((*C.struct__lv_obj_t)(unsafe.Pointer(tv)), n)
	return (*LVObj)(unsafe.Pointer(page)), err
}
