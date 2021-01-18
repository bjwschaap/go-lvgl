package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bjwschaap/go-lvgl/lvgl"
)

func main() {
	scr, err := lvgl.GetActiveScreen()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lbl, err := lvgl.Label(scr, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = lbl.SetText("hello world!")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = lbl.Align(nil, lvgl.AlignCenter, 0, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for true {
		lvgl.TickInc(5)
		lvgl.TaskHandler()
		time.Sleep(5000)
	}

	os.Exit(0)
}
