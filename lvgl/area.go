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
	AlignCenter         C.uchar = C.LV_ALIGN_CENTER
	AlignInTopLeft      C.uchar = C.LV_ALIGN_IN_TOP_LEFT
	AlignInTopMid       C.uchar = C.LV_ALIGN_IN_TOP_MID
	AlignInTopRight     C.uchar = C.LV_ALIGN_IN_TOP_RIGHT
	AlignInBottomLeft   C.uchar = C.LV_ALIGN_IN_BOTTOM_LEFT
	AlignInBottomMid    C.uchar = C.LV_ALIGN_IN_BOTTOM_MID
	AlignInBottomRight  C.uchar = C.LV_ALIGN_IN_BOTTOM_RIGHT
	AlignInLeftMid      C.uchar = C.LV_ALIGN_IN_LEFT_MID
	AlignInRightMid     C.uchar = C.LV_ALIGN_IN_RIGHT_MID
	AlignOutTopLeft     C.uchar = C.LV_ALIGN_OUT_TOP_LEFT
	AlignOutTopMid      C.uchar = C.LV_ALIGN_OUT_TOP_MID
	AlignOutTopRight    C.uchar = C.LV_ALIGN_OUT_TOP_RIGHT
	AlignOutBottomLeft  C.uchar = C.LV_ALIGN_OUT_BOTTOM_LEFT
	AlignOutBottomMid   C.uchar = C.LV_ALIGN_OUT_BOTTOM_MID
	AlignOutBottomRight C.uchar = C.LV_ALIGN_OUT_BOTTOM_RIGHT
	AlignOutLeftTop     C.uchar = C.LV_ALIGN_OUT_LEFT_TOP
	AlignOutLeftMid     C.uchar = C.LV_ALIGN_OUT_LEFT_MID
	AlignOutLeftBottom  C.uchar = C.LV_ALIGN_OUT_LEFT_BOTTOM
	AlignOutRightTop    C.uchar = C.LV_ALIGN_OUT_RIGHT_TOP
	AlignOutRightMid    C.uchar = C.LV_ALIGN_OUT_RIGHT_MID
	AlignOutRightBottom C.uchar = C.LV_ALIGN_OUT_RIGHT_BOTTOM
)
