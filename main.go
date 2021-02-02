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
	style, tstyle  lvgl.LVStyle
	scr            *lvgl.LVObj
	tv, t1, t2, t3 *lvgl.LVObj
	// lbl1, lbl2, lbl3    *lvgl.LVObj
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
	scr.Clean()
	lvgl.RefreshNow()

	// Give lvgl and goroutines some time to cleanup/finish
	time.Sleep(2 * time.Second)

	log.Info("exit")
	os.Exit(0)
}

// createScreen assembles the GUI through LVGL API
func createScreen() {

	// Make a Black active screen. so when we clear it
	// it will be a 'blank' screen
	style.Init()
	style.SetBgColor(lvgl.StateDefault, lvgl.ColorBlack)
	scr = lvgl.GetActiveScreen()
	scr.AddStyle(lvgl.ObjMaskPartMain, &style)

	// Make a tabview bar
	tv = scr.Tabview(nil)

	// Register event handler
	tv.RegisterEventCallback(MyCallback)

	// Add 3 tabs (+pages) to the tabview
	t1 = tv.AddTab("Controls")
	t2 = tv.AddTab("Visuals")
	t3 = tv.AddTab("Selectors")

	// Style tabpages
	tstyle.Init()
	tstyle.SetBgColor(lvgl.StateDefault, lvgl.ColorBlue)
	tstyle.SetPadInner(lvgl.StateDefault, 5)
	t1.AddStyle(lvgl.PagePartBG, &tstyle)

	// Add some labels to the tab pages
	lbl1 := t1.LabelCreate(nil)
	lbl1.SetText("This is tab 1")
	lbl2 := t2.LabelCreate(nil)
	lbl2.SetText("This is tab 2")
	lbl3 := t3.LabelCreate(nil)
	lbl3.SetText("This is tab 3")
}

// MyCallback is a test callback function
func MyCallback(obj *lvgl.LVObj, event lvgl.LVEvent) {
	log.WithField("event", event).Debug("MyCallback received event")
	switch event {
	case lvgl.EventValueChanged:
		log.Debug("Value changed event")
	case lvgl.EventDelete:
		log.Debug("Object deleted event")
	}
}
