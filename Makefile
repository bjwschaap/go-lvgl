BIN = golvgl
GOCMD = go
GOBUILD = ${GOCMD} build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"'
GOCLEAN = ${GOCMD} clean

build:
	${GOBUILD} -o ${BIN} -v *.go

clean:
	${GOCLEAN}

init:
	git submodule init
	git submodule update
	mkdir -p include/obj/lvgl/lv_drivers/display
	mkdir -p include/obj/lvgl/lv_drivers/gtkdrv
	mkdir -p include/obj/lvgl/lv_drivers/indev
	cd include/lvgl; ln -s ../lv_drivers lv_drivers

lvgl:
	$(MAKE) -C include all

.PHONY: build clean init lvgl
