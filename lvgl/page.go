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

const (
	PagePartBG         uint8 = C.LV_PAGE_PART_BG
	PagePartScrollbar  uint8 = C.LV_PAGE_PART_SCROLLBAR
	PagePartEdge       uint8 = C.LV_PAGE_PART_EDGE_FLASH
	PagePartScrollable uint8 = C.LV_PAGE_PART_SCROLLABLE
)
