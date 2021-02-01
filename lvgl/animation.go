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
	AnimationNone       LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_NONE
	AnimationOverLeft   LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_OVER_LEFT
	AnimationOverRight  LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_OVER_RIGHT
	AnimationOverTop    LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_OVER_TOP
	AnimationOverBottom LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_OVER_BOTTOM
	AnimationMoveLeft   LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_MOVE_LEFT
	AnimationMoveRight  LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_MOVE_RIGHT
	AnimationMoveTop    LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_MOVE_TOP
	AnimationMoveBottom LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_MOVE_BOTTOM
	AnimationFadeOn     LVLoadScreenAnimation = C.LV_SCR_LOAD_ANIM_FADE_ON
)

// LVLoadScreenAnimation represents the lv_scr_load_anim_t type
type LVLoadScreenAnimation C.lv_scr_load_anim_t

// LVAnimationPath represents lv_anim_path_t
type LVAnimationPath C.struct__lv_anim_path_t
