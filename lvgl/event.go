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
	"sync"
	"unsafe"
)

var (
	mu    sync.Mutex
	cbidx int
	cbFns = make(map[int]EventCallbackFn)
)

// EventCallbackFn is the callback function prototype
type EventCallbackFn func(*LVObj, LVEvent)

// LVEvent represents lv_event_t
type LVEvent C.struct__lv_event_t

// register registers a callback function in the map
func register(fn EventCallbackFn) int {
	mu.Lock()
	defer mu.Unlock()
	cbidx++
	for cbFns[cbidx] != nil {
		cbidx++
	}
	cbFns[cbidx] = fn
	return cbidx
}

// unregister deletes a callback function from the map
func unregister(i int) {
	mu.Lock()
	defer mu.Unlock()
	delete(cbFns, i)
}

// lookupCallback returns a callbackfunction from the map
func lookupCallback(i int) EventCallbackFn {
	mu.Lock()
	defer mu.Unlock()
	return cbFns[i]
}

// SetEventCallback registers an event handler on a LVObject
func (obj *LVObj) SetEventCallback(cb EventCallbackFn) error {

	_, err := C.lv_obj_set_event_cb((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), nil)
	return err
}
