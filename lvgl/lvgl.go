package lvgl

/*
  #cgo CFLAGS: -I../include -I../include/lvgl -I../include/lv_drivers/display -I../include/lv_drivers/indev
  #include "lvgl.h"
  #include "fbdev.h"
  #include "evdev.h"

  #cgo LDFLAGS: -L../include -llvgl
  #include <stdlib.h>
  #include <unistd.h>
  #include <stddef.h>

  #define LV_BUF_SIZE 384000

  static lv_disp_t * disp;
  static lv_disp_buf_t disp_buf;
  static lv_color_t lvbuf1[LV_BUF_SIZE];
  static lv_color_t lvbuf2[LV_BUF_SIZE];

  void init()
  {
	// LVGL init
	lv_init();

	// Linux framebuffer device init
	fbdev_init();

	// Touch pointer device init
	evdev_init();

	// Initialize display device
	lv_disp_buf_init(&disp_buf, lvbuf1, lvbuf2, LV_BUF_SIZE);
	lv_disp_drv_t disp_drv;
    lv_disp_drv_init(&disp_drv);
	disp_drv.flush_cb = fbdev_flush;
	disp_drv.buffer = &disp_buf;
	disp = lv_disp_drv_register(&disp_drv);

	// Initialize pointer device
	lv_indev_drv_t indev_drv;
	lv_indev_drv_init(&indev_drv);
	indev_drv.type = LV_INDEV_TYPE_POINTER;
	indev_drv.read_cb = evdev_read;
	lv_indev_drv_register(&indev_drv);
  }

  void handle_tick(uint32_t t)
  {
	lv_tick_inc(t);
	lv_task_handler();
  }

  void refresh_now(void)
  {
	  lv_refr_now(disp);
  }

*/
import "C"
import (
	"context"
	"time"
)

func init() {
	C.init()
}

// StartTaskHandler starts a go routine that
// increments the ticks, and periodically
// calls the taskhandler (tickless mode)
func StartTaskHandler(ctx context.Context) {
	// start a timer that ticks every second
	ticker := time.NewTicker(5000 * time.Microsecond)

	// start the Go routine
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				C.handle_tick(C.uint32_t(5))
			}
		}
	}()
}

// RefreshNow forces lvgl to refresh/flush the (default) display
func RefreshNow() {
	C.refresh_now()
}
