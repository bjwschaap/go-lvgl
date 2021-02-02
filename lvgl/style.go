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
import "unsafe"

//go:generate go run gen.go

const (
	StyleRadius     LVStyleProperty = C.LV_STYLE_RADIUS
	StyleClipCorner LVStyleProperty = C.LV_STYLE_CLIP_CORNER
	StyleSize       LVStyleProperty = C.LV_STYLE_SIZE

	StyleTransformWidth  LVStyleProperty = C.LV_STYLE_TRANSFORM_WIDTH
	StyleTransformHeight LVStyleProperty = C.LV_STYLE_TRANSFORM_HEIGHT
	StyleTransformAngle  LVStyleProperty = C.LV_STYLE_TRANSFORM_ANGLE
	StyleTransformZoom   LVStyleProperty = C.LV_STYLE_TRANSFORM_ZOOM

	StyleOpacityScale LVStyleProperty = C.LV_STYLE_OPA_SCALE

	StylePadTop    LVStyleProperty = C.LV_STYLE_PAD_TOP
	StylePadLeft   LVStyleProperty = C.LV_STYLE_PAD_LEFT
	StylePadRight  LVStyleProperty = C.LV_STYLE_PAD_RIGHT
	StylePadInner  LVStyleProperty = C.LV_STYLE_PAD_INNER
	StylePadBottom LVStyleProperty = C.LV_STYLE_PAD_BOTTOM

	StyleMarginTop    LVStyleProperty = C.LV_STYLE_MARGIN_TOP
	StyleMarginBottom LVStyleProperty = C.LV_STYLE_MARGIN_BOTTOM
	StyleMarginLeft   LVStyleProperty = C.LV_STYLE_MARGIN_LEFT
	StyleMarginRight  LVStyleProperty = C.LV_STYLE_MARGIN_RIGHT
	StyleBGBlendMode  LVStyleProperty = C.LV_STYLE_BG_BLEND_MODE

	StyleBGMainStop          LVStyleProperty = C.LV_STYLE_BG_MAIN_STOP
	StyleBGGradientStop      LVStyleProperty = C.LV_STYLE_BG_GRAD_STOP
	StyleBGGradientDirection LVStyleProperty = C.LV_STYLE_BG_GRAD_DIR
	StyleBGColor             LVStyleProperty = C.LV_STYLE_BG_COLOR
	StyleBGGradientColor     LVStyleProperty = C.LV_STYLE_BG_GRAD_COLOR
	StyleBGOpacity           LVStyleProperty = C.LV_STYLE_BG_OPA

	StyleBorderWidth     LVStyleProperty = C.LV_STYLE_BORDER_WIDTH
	StyleBorderSide      LVStyleProperty = C.LV_STYLE_BORDER_SIDE
	StyleBorderBlendMode LVStyleProperty = C.LV_STYLE_BORDER_BLEND_MODE
	StyleBorderPost      LVStyleProperty = C.LV_STYLE_BORDER_POST
	StyleBorderColor     LVStyleProperty = C.LV_STYLE_BORDER_COLOR
	StyleBorderOpacity   LVStyleProperty = C.LV_STYLE_BORDER_OPA

	StyleOutlineWidth     LVStyleProperty = C.LV_STYLE_OUTLINE_WIDTH
	StyleOutlinePad       LVStyleProperty = C.LV_STYLE_OUTLINE_PAD
	StyleOutlineBlendMode LVStyleProperty = C.LV_STYLE_OUTLINE_BLEND_MODE
	StyleOutlineColor     LVStyleProperty = C.LV_STYLE_OUTLINE_COLOR
	StyleOutlineOpacity   LVStyleProperty = C.LV_STYLE_OUTLINE_OPA

	StyleShadowWidth     LVStyleProperty = C.LV_STYLE_SHADOW_WIDTH
	StyleShadowOffsetX   LVStyleProperty = C.LV_STYLE_SHADOW_OFS_X
	StyleShadowOffsetY   LVStyleProperty = C.LV_STYLE_SHADOW_OFS_Y
	StyleShadowSpread    LVStyleProperty = C.LV_STYLE_SHADOW_SPREAD
	StyleShadowBlendMode LVStyleProperty = C.LV_STYLE_SHADOW_BLEND_MODE
	StyleShadowColor     LVStyleProperty = C.LV_STYLE_SHADOW_COLOR
	StyleShadowOpacity   LVStyleProperty = C.LV_STYLE_SHADOW_OPA

	StylePatternBlendMode      LVStyleProperty = C.LV_STYLE_PATTERN_BLEND_MODE
	StylePatternRepeat         LVStyleProperty = C.LV_STYLE_PATTERN_REPEAT
	StylePatternRecolor        LVStyleProperty = C.LV_STYLE_PATTERN_RECOLOR
	StylePatternOpacity        LVStyleProperty = C.LV_STYLE_PATTERN_OPA
	StylePatternRecolorOpacity LVStyleProperty = C.LV_STYLE_PATTERN_RECOLOR_OPA
	StylePatternImage          LVStyleProperty = C.LV_STYLE_PATTERN_IMAGE

	StyleValueLetterSpace LVStyleProperty = C.LV_STYLE_VALUE_LETTER_SPACE
	StyleValueLineSpace   LVStyleProperty = C.LV_STYLE_VALUE_LINE_SPACE
	StyleValueBlendMode   LVStyleProperty = C.LV_STYLE_VALUE_BLEND_MODE
	StyleValueOffsetX     LVStyleProperty = C.LV_STYLE_VALUE_OFS_X
	StyleValueOffsetY     LVStyleProperty = C.LV_STYLE_VALUE_OFS_Y
	StyleValueAlign       LVStyleProperty = C.LV_STYLE_VALUE_ALIGN
	StyleValueColor       LVStyleProperty = C.LV_STYLE_VALUE_COLOR
	StyleValueOpacity     LVStyleProperty = C.LV_STYLE_VALUE_OPA
	StyleValueFont        LVStyleProperty = C.LV_STYLE_VALUE_FONT
	StyleValueString      LVStyleProperty = C.LV_STYLE_VALUE_STR

	StyleTextLetterSpace     LVStyleProperty = C.LV_STYLE_TEXT_LETTER_SPACE
	StyleTextLineSpace       LVStyleProperty = C.LV_STYLE_TEXT_LINE_SPACE
	StyleTextDecor           LVStyleProperty = C.LV_STYLE_TEXT_DECOR
	StyleTextBlendMode       LVStyleProperty = C.LV_STYLE_TEXT_BLEND_MODE
	StyleTextColor           LVStyleProperty = C.LV_STYLE_TEXT_COLOR
	StyleTextSelectedColor   LVStyleProperty = C.LV_STYLE_TEXT_SEL_COLOR
	StyleTextSelectedBGColor LVStyleProperty = C.LV_STYLE_TEXT_SEL_BG_COLOR
	StyleTextOpacity         LVStyleProperty = C.LV_STYLE_TEXT_OPA
	StyleTextFont            LVStyleProperty = C.LV_STYLE_TEXT_FONT

	StyleLineWidth     LVStyleProperty = C.LV_STYLE_LINE_WIDTH
	StyleLineBlendMode LVStyleProperty = C.LV_STYLE_LINE_BLEND_MODE
	StyleLineDashWidth LVStyleProperty = C.LV_STYLE_LINE_DASH_WIDTH
	StyleLineDashGap   LVStyleProperty = C.LV_STYLE_LINE_DASH_GAP
	StyleLineRounded   LVStyleProperty = C.LV_STYLE_LINE_ROUNDED
	StyleLineColor     LVStyleProperty = C.LV_STYLE_LINE_COLOR
	StyleLineOpacity   LVStyleProperty = C.LV_STYLE_LINE_OPA

	StyleImageBlendMode      LVStyleProperty = C.LV_STYLE_IMAGE_BLEND_MODE
	StyleImageRecolor        LVStyleProperty = C.LV_STYLE_IMAGE_RECOLOR
	StyleImageOpacity        LVStyleProperty = C.LV_STYLE_IMAGE_OPA
	StyleImageRecolorOpacity LVStyleProperty = C.LV_STYLE_IMAGE_RECOLOR_OPA

	StyleTransitionTime      LVStyleProperty = C.LV_STYLE_TRANSITION_TIME
	StyleTransitionDelay     LVStyleProperty = C.LV_STYLE_TRANSITION_DELAY
	StyleTransitionProperty1 LVStyleProperty = C.LV_STYLE_TRANSITION_PROP_1
	StyleTransitionProperty2 LVStyleProperty = C.LV_STYLE_TRANSITION_PROP_2
	StyleTransitionProperty3 LVStyleProperty = C.LV_STYLE_TRANSITION_PROP_3
	StyleTransitionProperty4 LVStyleProperty = C.LV_STYLE_TRANSITION_PROP_4
	StyleTransitionProperty5 LVStyleProperty = C.LV_STYLE_TRANSITION_PROP_5
	StyleTransitionProperty6 LVStyleProperty = C.LV_STYLE_TRANSITION_PROP_6
	StyleTransitionPath      LVStyleProperty = C.LV_STYLE_TRANSITION_PATH

	StyleScaleWidth          LVStyleProperty = C.LV_STYLE_SCALE_WIDTH
	StyleScaleBorderWidth    LVStyleProperty = C.LV_STYLE_SCALE_BORDER_WIDTH
	StyleScaleEndBorderWidth LVStyleProperty = C.LV_STYLE_SCALE_END_BORDER_WIDTH
	StyleScaleEndLineWidth   LVStyleProperty = C.LV_STYLE_SCALE_END_LINE_WIDTH
	StyleScaleGradientColor  LVStyleProperty = C.LV_STYLE_SCALE_GRAD_COLOR
	StyleScaleEndColor       LVStyleProperty = C.LV_STYLE_SCALE_END_COLOR
)

// LVStyle contains an object's styling information
type LVStyle C.struct__lv_style_t

// LVStyleProperty references a specific style attribute
type LVStyleProperty C.lv_style_property_t

// LVStyleList represents lv_style_list_t
type LVStyleList C.struct__lv_style_list_t

// LVOpacity represents lv_opa_t
type LVOpacity C.lv_opa_t

// LVBlendMode represents lv_blend_mode_t
type LVBlendMode C.lv_blend_mode_t

// LVGradient represents lv_grad_dir_t
type LVGradient C.lv_grad_dir_t

// LVBorderSide represents lv_border_side_t
type LVBorderSide C.lv_border_side_t

// LVFont represents lv_font_t
type LVFont C.struct__lv_font_struct

// LVTextDecor represents lv_text_decor_t
type LVTextDecor C.lv_text_decor_t

// Init initializes a new lv_style_t object
func (style *LVStyle) Init() {
	s := (*C.struct___7)(unsafe.Pointer(style))
	C.lv_style_init(s)
}

// Init initializes a new lv_style_list_init
func (list *LVStyleList) Init() {
	C.lv_style_list_init((*C.struct___5)(unsafe.Pointer(list)))
}

// Copy copies a stylelist to another stylelist
func (list *LVStyleList) Copy(dest *LVStyleList) {
	s := (*C.struct___5)(unsafe.Pointer(list))
	d := (*C.struct___5)(unsafe.Pointer(dest))
	C.lv_style_list_copy(d, s)
}

// Copy copies a style to a destination style
// @param dest The destination style
// @return error
func (style *LVStyle) Copy(dest *LVStyle) {
	s := (*C.struct___7)(unsafe.Pointer(style))
	d := (*C.struct___7)(unsafe.Pointer(dest))
	C.lv_style_copy(d, s)
}

// Reset clears all properties from a style and all allocated memories
// @param dest The destination style
// @return error
func (style *LVStyle) Reset() {
	s := (*C.struct___7)(unsafe.Pointer(style))
	C.lv_style_reset(s)
}

// AddStyle adds a style to a LVObj
func (obj *LVObj) AddStyle(part uint8, style *LVStyle) {
	s := (*C.struct___7)(unsafe.Pointer(style))
	C.lv_obj_add_style((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uchar(part), s)
}

// RemoveProperty removes a property from a style
// @param style pointer to a style
// @param prop  a style property ORed with a state.
// E.g. `LV_STYLE_BORDER_WIDTH | (LV_STATE_PRESSED << LV_STYLE_STATE_POS)`
// @return true: the property was found and removed; false: the property wasn't found
func (style *LVStyle) RemoveProperty(prop LVStyleProperty) bool {
	s := (*C.struct___7)(unsafe.Pointer(style))
	return bool(C.lv_style_remove_prop(s, (C.lv_style_property_t)(prop)))
}

// ResetStyleList resets all styles within the style list of an object
// Release all used memories and cancel pending related transitions.
// Also notifies the object about the style change.
// @param obj pointer to an object
// @param part the part of the object which style list should be reseted.
// E.g. `LV_OBJ_PART_MAIN`, `LV_BTN_PART_MAIN`, `LV_SLIDER_PART_KNOB`
func (obj *LVObj) ResetStyleList(part uint8) {
	C.lv_obj_reset_style_list((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uchar(part))
}

// RefreshStyle refreshed one or more styles on an object
func (obj *LVObj) RefreshStyle(part uint8, prop LVStyleProperty) {
	C.lv_obj_refresh_style((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uchar(part), (C.lv_style_property_t)(prop))
}
