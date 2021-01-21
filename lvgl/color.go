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

var (
	ColorWhite   LVColor = LVColor(C.LV_COLOR_WHITE)
	ColorSilver  LVColor = LVColor(C.LV_COLOR_SILVER)
	ColorGray    LVColor = LVColor(C.LV_COLOR_GRAY)
	ColorBlack   LVColor = LVColor(C.LV_COLOR_BLACK)
	ColorRed     LVColor = LVColor(C.LV_COLOR_RED)
	ColorMaroon  LVColor = LVColor(C.LV_COLOR_MAROON)
	ColorYellow  LVColor = LVColor(C.LV_COLOR_YELLOW)
	ColorOlive   LVColor = LVColor(C.LV_COLOR_OLIVE)
	ColorLime    LVColor = LVColor(C.LV_COLOR_LIME)
	ColorGreen   LVColor = LVColor(C.LV_COLOR_GREEN)
	ColorCyan    LVColor = LVColor(C.LV_COLOR_CYAN)
	ColorAqua    LVColor = LVColor(C.LV_COLOR_AQUA)
	ColorTeal    LVColor = LVColor(C.LV_COLOR_TEAL)
	ColorBlue    LVColor = LVColor(C.LV_COLOR_BLUE)
	ColorNavy    LVColor = LVColor(C.LV_COLOR_NAVY)
	ColorMagenta LVColor = LVColor(C.LV_COLOR_MAGENTA)
	ColorPurple  LVColor = LVColor(C.LV_COLOR_PURPLE)
	ColorOrange  LVColor = LVColor(C.LV_COLOR_ORANGE)
)

// LVColor represents a 16 bit color value
type LVColor C.lv_color_t

// ColorHex returns a LVColor object by hex code
func ColorHex(c uint32) LVColor {
	color := C.lv_color_hex(C.uint32_t(c))
	return LVColor(color)
}
