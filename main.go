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
	styleBox            lvgl.LVStyle
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
func createScreen() {
	var err error

	scr, err = lvgl.GetActiveScreen()
	if err != nil {
		log.Fatal(err)
	}

	tv, err = lvgl.TabviewCreate(scr, nil)
	if err != nil {
		log.Fatal(err)
	}

	t1, err = tv.AddTab("Controls")
	if err != nil {
		log.Fatal(err)
	}
	t2, err = tv.AddTab("Visuals")
	if err != nil {
		log.Fatal(err)
	}
	t3, err = tv.AddTab("Selectors")
	if err != nil {
		log.Fatal(err)
	}

	lbl1, err = lvgl.Label(t1, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err = lbl1.SetText("This is tab 1"); err != nil {
		log.Fatal(err)
	}

	lbl2, err = lvgl.Label(t2, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err = lbl2.SetText("This is tab 2"); err != nil {
		log.Fatal(err)
	}

	lbl3, err = lvgl.Label(t3, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err = lbl3.SetText("This is tab 3"); err != nil {
		log.Fatal(err)
	}
}
