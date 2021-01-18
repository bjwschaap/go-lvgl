package lvgl

/*
  #cgo CFLAGS: -I../include -I../include/lvgl -I../include/lv_drivers/display -I../include/lv_drivers/indev
  #include "lv_conf.h"
  #include "lv_drv_conf.h"
  #include "lvgl.h"
  #include "fbdev.h"
  #include "evdev.h"

  #cgo LDFLAGS: -L../include -llvgl
  #include <stdlib.h>

  #define LV_BUF_SIZE 384000

  static lv_disp_buf_t disp_buf;
  static lv_color_t lvbuf1[LV_BUF_SIZE];
  static lv_color_t lvbuf2[LV_BUF_SIZE];

  void disp_init(){
	// Initialize display device
	lv_disp_buf_init(&disp_buf, lvbuf1, lvbuf2, LV_BUF_SIZE);
	lv_disp_drv_t disp_drv;
    lv_disp_drv_init(&disp_drv);
	disp_drv.flush_cb = fbdev_flush;
	disp_drv.buffer = &disp_buf;
	lv_disp_drv_register(&disp_drv);

	// Initialize pointer device
	lv_indev_drv_t indev_drv;
	lv_indev_drv_init(&indev_drv);
	indev_drv.type = LV_INDEV_TYPE_POINTER;
	indev_drv.read_cb = evdev_read;
	lv_indev_drv_register(&indev_drv);
  }

*/
import "C"

func init() {
	C.lv_init()
	C.fbdev_init()
	C.evdev_init()
	C.disp_init()
}

// TickInc ...
func TickInc(tick int) {
	C.lv_tick_inc(C.uint(tick))
}

// TaskHandler ...
func TaskHandler() {
	C.lv_task_handler()
}
