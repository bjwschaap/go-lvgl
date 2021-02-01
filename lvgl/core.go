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
	DisplaySizeSmall      LVDisplaySize = C.LV_DISP_SIZE_SMALL
	DisplaySizeMedium     LVDisplaySize = C.LV_DISP_SIZE_MEDIUM
	DisplaySizeLarge      LVDisplaySize = C.LV_DISP_SIZE_LARGE
	DisplaySizeExtraLarge LVDisplaySize = C.LV_DISP_SIZE_EXTRA_LARGE

	StateDefault  LVState = C.LV_STATE_DEFAULT
	StateChecked  LVState = C.LV_STATE_CHECKED
	StateFocused  LVState = C.LV_STATE_FOCUSED
	StateEdited   LVState = C.LV_STATE_EDITED
	StateHovered  LVState = C.LV_STATE_HOVERED
	StatePressed  LVState = C.LV_STATE_PRESSED
	StateDisabled LVState = C.LV_STATE_DISABLED

	ObjMaskPartMain     uint8 = C.LV_OBJMASK_PART_MAIN
	ObjPartMain         uint8 = C.LV_OBJ_PART_MAIN
	ObjPartAll          uint8 = C.LV_OBJ_PART_ALL
	ObjPartVirtualFirst uint8 = C._LV_OBJ_PART_VIRTUAL_FIRST
	ObjPartVirtualLast  uint8 = C._LV_OBJ_PART_VIRTUAL_LAST
	ObjPartRealFirst    uint8 = C._LV_OBJ_PART_REAL_FIRST
	ObjPartRealLast     uint8 = C._LV_OBJ_PART_REAL_LAST
)

// LVObj is the base object that implements the basic
// properties of widgets on a screen.
type LVObj C.struct__lv_obj_t

// LVDisplay is the display type
type LVDisplay C.struct__disp_t

// LVDisplaySize represents the display size category
type LVDisplaySize C.lv_disp_size_t

// LVUserData represents lv_obj user data
type LVUserData C.lv_obj_user_data_t

// LVState describes the state of an object
type LVState C.lv_state_t

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

// ObjCreate creates a new lv_obj
func ObjCreate(parent *LVObj, copy *LVObj) *LVObj {
	obj := C.lv_obj_create((*C.struct__lv_obj_t)(unsafe.Pointer(parent)), (*C.struct__lv_obj_t)(unsafe.Pointer(copy)))
	return (*LVObj)(obj)
}

// Delete destroys an object
func (obj *LVObj) Delete() {
	C.lv_obj_del((*C.struct__lv_obj_t)(unsafe.Pointer(obj)))
}

// Align ...
func (obj *LVObj) Align(base *LVObj, align uint8, x int16, y int16) {
	ba := (*C.struct__lv_obj_t)(unsafe.Pointer(base))
	C.lv_obj_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), ba, C.uchar(align), C.short(x), C.short(y))
}

// GetSizeCategory returns the size category of the display based on it's hor. res. and dpi.
// @param disp pointer to a display (NULL to use the default display)
// @return DisplaySizeSmall/DisplaySizeMedium/DisplaySizeLarge/DisplaySizeExtraLarge
func (disp *LVDisplay) GetSizeCategory() LVDisplaySize {
	d := (*C.struct__disp_t)(unsafe.Pointer(disp))
	cat := C.lv_disp_get_size_category(d)
	return LVDisplaySize(cat)
}

// SetDefault makes the display the default display
func (disp *LVDisplay) SetDefault() {
	C.lv_disp_set_default((*C.struct__disp_t)(unsafe.Pointer(disp)))
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

// SetSize configures the size of the object
func (obj *LVObj) SetSize(w, h int16) {
	C.lv_obj_set_size((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.short(w), C.short(h))
}

// SetPosition configures the position of the object
func (obj *LVObj) SetPosition(x, y int16) {
	C.lv_obj_set_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.short(x), C.short(y))
}

// SetX sets the object's X position
func (obj *LVObj) SetX(x int16) {
	C.lv_obj_set_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.short(x))
}

// SetY sets the object's X position
func (obj *LVObj) SetY(y int16) {
	C.lv_obj_set_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.short(y))
}

// SetState overrides/forces the current state of an object
func (obj *LVObj) SetState(state LVState) {
	C.lv_obj_set_state((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (C.lv_state_t)(state))
}

// AddState adds a given state or states to the object. The other state bits will remain unchanged.
// If specified in the styles a transition animation will be started
// the previous state to the current
// @param obj pointer to an object
// @param state the state bits to add. E.g `LV_STATE_PRESSED | LV_STATE_FOCUSED`
func (obj *LVObj) AddState(state LVState) {
	C.lv_obj_add_state((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (C.lv_state_t)(state))
}

// ClearState removes a given state or states to the object. The other state bits will remain unchanged.
// If specified in the styles a transition animation will be started
// from the previous state to the current
// @param obj pointer to an object
// @param state the state bits to remove. E.g `LV_STATE_PRESSED | LV_STATE_FOCUSED`
func (obj *LVObj) ClearState(state LVState) {
	C.lv_obj_clear_state((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (C.lv_state_t)(state))
}
