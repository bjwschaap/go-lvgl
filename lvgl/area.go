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

const (
	AlignCenter         uint8 = C.LV_ALIGN_CENTER
	AlignInTopLeft      uint8 = C.LV_ALIGN_IN_TOP_LEFT
	AlignInTopMid       uint8 = C.LV_ALIGN_IN_TOP_MID
	AlignInTopRight     uint8 = C.LV_ALIGN_IN_TOP_RIGHT
	AlignInBottomLeft   uint8 = C.LV_ALIGN_IN_BOTTOM_LEFT
	AlignInBottomMid    uint8 = C.LV_ALIGN_IN_BOTTOM_MID
	AlignInBottomRight  uint8 = C.LV_ALIGN_IN_BOTTOM_RIGHT
	AlignInLeftMid      uint8 = C.LV_ALIGN_IN_LEFT_MID
	AlignInRightMid     uint8 = C.LV_ALIGN_IN_RIGHT_MID
	AlignOutTopLeft     uint8 = C.LV_ALIGN_OUT_TOP_LEFT
	AlignOutTopMid      uint8 = C.LV_ALIGN_OUT_TOP_MID
	AlignOutTopRight    uint8 = C.LV_ALIGN_OUT_TOP_RIGHT
	AlignOutBottomLeft  uint8 = C.LV_ALIGN_OUT_BOTTOM_LEFT
	AlignOutBottomMid   uint8 = C.LV_ALIGN_OUT_BOTTOM_MID
	AlignOutBottomRight uint8 = C.LV_ALIGN_OUT_BOTTOM_RIGHT
	AlignOutLeftTop     uint8 = C.LV_ALIGN_OUT_LEFT_TOP
	AlignOutLeftMid     uint8 = C.LV_ALIGN_OUT_LEFT_MID
	AlignOutLeftBottom  uint8 = C.LV_ALIGN_OUT_LEFT_BOTTOM
	AlignOutRightTop    uint8 = C.LV_ALIGN_OUT_RIGHT_TOP
	AlignOutRightMid    uint8 = C.LV_ALIGN_OUT_RIGHT_MID
	AlignOutRightBottom uint8 = C.LV_ALIGN_OUT_RIGHT_BOTTOM
)
