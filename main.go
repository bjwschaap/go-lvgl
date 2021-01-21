package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bjwschaap/go-lvgl/lvgl"
	log "github.com/sirupsen/logrus"
)

var (
	style               lvgl.LVStyle
	scr, tv, t1, t2, t3 *lvgl.LVObj
	lbl1, lbl2, lbl3    *lvgl.LVObj
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		DisableSorting:  true,
	})
}

func main() {
	// Start
	log.Info("Starting go-lvgl")

	// Handle OS interrupts
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGHUP)

	// Create a context for cancellation
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Make the GUI
	log.Debug("construct GUI")
	createScreen()

	// Start LVGL's task handler
	log.Debug("starting lvgl taskhandler")
	lvgl.StartTaskHandler(ctx)

	select {
	case s := <-interrupt:
		log.WithField("signal", s.String()).Info("received OS interrupt")
		cancel()
	}

	// Clear the display
	_ = scr.Clean()
	_ = lvgl.RefreshNow()

	// Give goroutines some time to cleanup/finish
	time.Sleep(1 * time.Second)

	log.Info("exit")
	os.Exit(0)
}

// createScreen assembles the GUI through LVGL API
// errors are ignored for readability
func createScreen() {
	style.Init()
	style.SetBgColor(lvgl.StateDefault, lvgl.ColorBlack)

	scr, _ = lvgl.GetActiveScreen()
	scr.AddStyle(lvgl.ObjMaskPartMain, &style)

	tv, _ = lvgl.TabviewCreate(scr, nil)

	t1, _ = tv.AddTab("Controls")
	t2, _ = tv.AddTab("Visuals")
	t3, _ = tv.AddTab("Selectors")

	lbl1, _ = lvgl.Label(t1, nil)
	lbl1.SetText("This is tab 1")
	lbl2, _ = lvgl.Label(t2, nil)
	lbl2.SetText("This is tab 2")
	lbl3, _ = lvgl.Label(t3, nil)
	lbl3.SetText("This is tab 3")

}
