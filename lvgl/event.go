package lvgl

/*
  #cgo CFLAGS: -I../include -I../include/lvgl -I../include/lv_drivers/display -I../include/lv_drivers/indev
  #include "lv_conf.h"
  #include "lv_drv_conf.h"
  #include "lvgl.h"
  #include "fbdev.h"
  #include "evdev.h"

  #cgo LDFLAGS: -L../include -llvgl

  // Go function prototype
  extern void go_event_callback(lv_obj_t * obj, lv_event_t event);

  // callback 'proxy'
  static inline void event_cb_proxy(lv_obj_t * obj, lv_event_t event) {
	go_event_callback(obj, event);
  }

  // Callback registration
  static void _register_callback(lv_obj_t *obj) {
	  lv_obj_set_event_cb(obj, event_cb_proxy);
  }

*/
import "C"
import (
	"sync"
	"unsafe"

	"github.com/mattn/go-pointer"
)

const (
	EventPressed           LVEvent = C.LV_EVENT_PRESSED
	EventPressing          LVEvent = C.LV_EVENT_PRESSING
	EventPressLost         LVEvent = C.LV_EVENT_PRESS_LOST
	EventShortClicked      LVEvent = C.LV_EVENT_SHORT_CLICKED
	EventLongPressed       LVEvent = C.LV_EVENT_LONG_PRESSED
	EventLongPressedRepeat LVEvent = C.LV_EVENT_LONG_PRESSED_REPEAT
	EventClicked           LVEvent = C.LV_EVENT_CLICKED
	EventReleased          LVEvent = C.LV_EVENT_RELEASED
	EventDragBegin         LVEvent = C.LV_EVENT_DRAG_BEGIN
	EventDragEnd           LVEvent = C.LV_EVENT_DRAG_END
	EventDragThrowBegin    LVEvent = C.LV_EVENT_DRAG_THROW_BEGIN
	EventGesture           LVEvent = C.LV_EVENT_GESTURE
	EventKey               LVEvent = C.LV_EVENT_KEY
	EventFocused           LVEvent = C.LV_EVENT_FOCUSED
	EventDefocused         LVEvent = C.LV_EVENT_DEFOCUSED
	EventLeave             LVEvent = C.LV_EVENT_LEAVE
	EventValueChanged      LVEvent = C.LV_EVENT_VALUE_CHANGED
	EventInsert            LVEvent = C.LV_EVENT_INSERT
	EventRefresh           LVEvent = C.LV_EVENT_REFRESH
	EventApply             LVEvent = C.LV_EVENT_APPLY
	EventCancel            LVEvent = C.LV_EVENT_CANCEL
	EventDelete            LVEvent = C.LV_EVENT_DELETE
)

// EventUserData represents event info to inject into lv_obj as user_data
type EventUserData struct {
	IDX int
}

// EventCallbackFn is the callback function prototype
type EventCallbackFn func(*LVObj, LVEvent)

// LVEvent represents lv_event_t
type LVEvent C.lv_event_t

// EventFnRegister contains a map of all event callbacks
type eventCBRegister struct {
	sync.Mutex
	index int
	fns   map[int]EventCallbackFn
}

var (
	cbReg *eventCBRegister
)

func init() {
	cbReg = newEventCBReg()
}

// initializes a new event callback register
func newEventCBReg() *eventCBRegister {
	return &eventCBRegister{
		fns: make(map[int]EventCallbackFn),
	}
}

// Add event callback function to the local callback register
func (r *eventCBRegister) add(fn EventCallbackFn) int {
	r.Lock()
	defer r.Unlock()
	r.index++
	for r.fns[r.index] != nil {
		r.index++
	}
	r.fns[r.index] = fn
	return r.index
}

// Find event callback function in the local callback register
func (r *eventCBRegister) lookup(i int) EventCallbackFn {
	r.Lock()
	defer r.Unlock()
	return r.fns[i]
}

// Remove an event callback function from the local callback register
func (r *eventCBRegister) remove(i int) {
	r.Lock()
	defer r.Unlock()
	delete(r.fns, i)
}

//export go_event_callback
//
// This 'proxy' function gets called for *all* event callbacks
// and executes the callback function through the local callback
// register.
func go_event_callback(obj *C.struct__lv_obj_t, event C.lv_event_t) {
	o := (*LVObj)(obj)
	eud := pointer.Restore(unsafe.Pointer(o.UserData())).(*EventUserData)
	fn := cbReg.lookup(eud.IDX)
	fn(o, (LVEvent)(event))
}

// RegisterEventCallback registers an Event handler function for a lv_object
func (obj *LVObj) RegisterEventCallback(fn EventCallbackFn) {
	i := cbReg.add(fn)
	obj.SetUserData(pointer.Save(&EventUserData{IDX: i}))
	C._register_callback(((*C.struct__lv_obj_t)(unsafe.Pointer(obj))))
}

// UnRegisterEventCallback de-registers an Event handler function for a lv_object
func (obj *LVObj) UnRegisterEventCallback(fn EventCallbackFn) {
	o := (*LVObj)(obj)
	eud := pointer.Restore(unsafe.Pointer(o.UserData())).(*EventUserData)
	cbReg.remove(eud.IDX)
	C._register_callback(((*C.struct__lv_obj_t)(unsafe.Pointer(obj))))
}
