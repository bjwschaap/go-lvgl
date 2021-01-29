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

const (
	DisplaySizeSmall      uint8 = C.LV_DISP_SIZE_SMALL
	DisplaySizeMedium     uint8 = C.LV_DISP_SIZE_MEDIUM
	DisplaySizeLarge      uint8 = C.LV_DISP_SIZE_LARGE
	DisplaySizeExtraLarge uint8 = C.LV_DISP_SIZE_EXTRA_LARGE

	StateDefault  uint8 = C.LV_STATE_DEFAULT
	StateChecked  uint8 = C.LV_STATE_CHECKED
	StateFocused  uint8 = C.LV_STATE_FOCUSED
	StateEdited   uint8 = C.LV_STATE_EDITED
	StateHovered  uint8 = C.LV_STATE_HOVERED
	StatePressed  uint8 = C.LV_STATE_PRESSED
	StateDisabled uint8 = C.LV_STATE_DISABLED

	ObjMaskPartMain uint8 = C.LV_OBJMASK_PART_MAIN
)

// LVObj is the base object that implements the basic
// properties of widgets on a screen.
type LVObj C.struct__lv_obj_t

// LVDisplay is the display type
type LVDisplay C.struct__disp_t

// LVUserData represents lv_obj user data
type LVUserData C.lv_obj_user_data_t

// GetChild helps to iterate through the children of an object
func GetChild(obj, child *LVObj) *LVObj {
	var p1, p2 *C.struct__lv_obj_t
	p1 = (*C.struct__lv_obj_t)(unsafe.Pointer(obj))
	if child != nil {
		p2 = (*C.struct__lv_obj_t)(unsafe.Pointer(child))
	} else {
		p2 = nil
	}

	lvobj := C.lv_obj_get_child(p1, p2)
	return (*LVObj)(unsafe.Pointer(lvobj))
}

// Align ...
func (obj *LVObj) Align(base *LVObj, align uint8, x int, y int) {
	ba := (*C.struct__lv_obj_t)(unsafe.Pointer(base))
	C.lv_obj_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), ba, C.uchar(align), C.short(x), C.short(y))
}

// GetDisplaySizeCategory returns the size category of the display based on it's hor. res. and dpi.
// @param disp pointer to a display (NULL to use the default display)
// @return DisplaySizeSmall/DisplaySizeMedium/DisplaySizeLarge/DisplaySizeExtraLarge
func GetDisplaySizeCategory(disp *LVDisplay) uint8 {
	d := (*C.struct__disp_t)(unsafe.Pointer(disp))
	cat := C.lv_disp_get_size_category(d)
	return uint8(cat)
}

// Clean resets/clears an lv_obj back to init state
func (obj *LVObj) Clean() {
	C.lv_obj_clean((*C.struct__lv_obj_t)(unsafe.Pointer(obj)))
}

// UserData returns the lv_obj's user data
func (obj *LVObj) UserData() LVUserData {
	u := C.lv_obj_get_user_data((*C.struct__lv_obj_t)(unsafe.Pointer(obj)))
	return LVUserData(u)
}

// SetUserData sets the user data on a lv_obj
func (obj *LVObj) SetUserData(u unsafe.Pointer) {
	C.lv_obj_set_user_data((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (C.lv_obj_user_data_t)(u))
}
