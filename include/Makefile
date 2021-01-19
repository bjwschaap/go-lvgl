BIN = liblvgl.so
ABIN = liblvgl.a
DESTDIR = /usr
PREFIX = /local

CC=gcc
CXX=g++
AR=ar
CFLAGS = -Wall -Wshadow -Wundef -Wmaybe-uninitialized
CFLAGS += -O3 -g3 -I./
CFLAGS += -fPIC

# directory for local libs
LDFLAGS = -L$(DESTDIR)$(PREFIX)/lib
LIBS += -lstdc++ -lm

VPATH =

$(info LDFLAGS ="$(LDFLAGS)")

LVGL_DIR = lvgl

include $(LVGL_DIR)/lvgl.mk
include $(LVGL_DIR)/lv_drivers/lv_drivers.mk

# folder for object files
OBJDIR = ./obj

CSRCS += $(wildcard *.c)

COBJS = $(patsubst %.c,$(OBJDIR)/%.o,$(CSRCS))

SRCS = $(CSRCS)
OBJS = $(COBJS)

#.PHONY: clean

all: default

$(OBJDIR)/%.o: %.c
	@mkdir -p $(OBJDIR)
	@$(CC)  $(CFLAGS) -c $< -o $@
	@echo "CC $<"

default: $(OBJS)
	$(CC) -shared -o $(BIN) $(OBJS) $(LDFLAGS) $(LIBS)
	$(AR) q ${ABIN} $(OBJS)

#	nothing to do but will print info
nothing:
	$(info OBJS ="$(OBJS)")
	$(info SRCS ="$(SRCS)")
	$(info DONE)


clean:
	rm -f $(OBJS)
	rm -f $(BIN)
	rm -f $(ABIN)
