package layershell

// #cgo pkg-config: gtk4 gtk4-layer-shell-0
// #include <gtk/gtk.h>
// #include "gtk4-layer-shell.h"
import "C"
import (
	//  "github.com/brotholo/gotk4/gdk"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"

	//  "github.com/gotk3/gotk3/gtk"
	"unsafe"
)

type LayerShellEdgeFlags int

const (
	LAYER_SHELL_EDGE_LEFT LayerShellEdgeFlags = iota
	LAYER_SHELL_EDGE_RIGHT
	LAYER_SHELL_EDGE_TOP
	LAYER_SHELL_EDGE_BOTTOM
)

type LayerShellLayerFlags int

const (
	LAYER_SHELL_LAYER_BACKGROUND LayerShellLayerFlags = iota
	LAYER_SHELL_LAYER_BOTTOM
	LAYER_SHELL_LAYER_TOP
	LAYER_SHELL_LAYER_OVERLAY
)

type LayerShellKeyboardMode int

const (
	LAYER_SHELL_KEYBOARD_MODE_NONE LayerShellKeyboardMode = iota
	LAYER_SHELL_KEYBOARD_MODE_EXCLUSIVE
	LAYER_SHELL_KEYBOARD_MODE_ON_DEMAND
	LAYER_SHELL_KEYBOARD_MODE_ENTRY_NUMBER
)

func nativeWindow(window *gtk.Window) *C.GtkWindow {
	//  w := window
	//  var w string
	//  C.pass_pointer(pointer.Save(&s))
	w := window.Widget.Native()
	//  wp := (*C.GtkWindow)(pointer.Save(w))
	wp := (*C.GtkWindow)(unsafe.Pointer(&w))
	return wp
}

//  func InitForWindow(window *gtk.Window) {
//  //  w := nativeWindow(window)
//  w := window.Widget.Native()
//  //  wp := (*C.GtkWindow)(unsafe.Pointer(w))
//  wp := (*C.GtkWindow)(pointer.Save(w))
//  C.gtk_layer_init_for_window(wp)

func nativeMonitor(monitor *gdk.Monitor) *C.GdkMonitor {
	m := monitor
	mp := (*C.GdkMonitor)(unsafe.Pointer(m))
	return mp
}

func boolConv(b bool) C.int {
	if b {
		return C.int(1)
	}
	return C.int(0)
}

//  #include <stdio.h>

//  typedef int* pInt;

//  void foo(pInt p[]) { // you probably wanna pass a len to the function.
//  *p[0] = 100;
//  printf("foo()\n");
//  }

//  */
//  import "C"
//  import "unsafe"

//  func main() {
//  var (
//  i, sz  = 0, 2
//  arr    = (*C.pInt)(C.malloc(C.size_t(sz)))
//  ps     = (*[100000]C.pInt)(unsafe.Pointer(arr))[:sz:sz]
//  p1, p2 = (C.pInt)(unsafe.Pointer(&i)), (C.pInt)(unsafe.Pointer(&i))
//  )
//  ps[0], ps[1] = p1, p2
//  C.foo(arr)
//  C.free(unsafe.Pointer(arr))
//  println("i", i)
//  }

//  typedef int* pInt;

//  void foo(pInt p[]) {
//  printf("foo()\n");
//  }
//  */
//  import "C"
//  import "unsafe"

// func main() {
// var i C.int
// var p1 C.pInt = (*C.int)(unsafe.Pointer(&i))
// var p2 C.pInt = (*C.int)(unsafe.Pointer(&i))
// var ps []C.pInt = []C.pInt{p1, p2}
// C.foo(unsafe.Pointer(&ps[0]))
// }
//
//	func InitForWindow(window *gtk.Window) *gtk.Window {
func InitForWindow(window *gtk.Window) {
	w := nativeWindow(window)
	C.gtk_layer_init_for_window(w)
	//  return pointer.Restore(w)
}

func SetLayer(window *gtk.Window, layer LayerShellLayerFlags) {
	w := nativeWindow(window)
	C.gtk_layer_set_layer(w, C.GtkLayerShellLayer(layer))
}

func AutoExclusiveZoneEnable(window *gtk.Window) {
	w := nativeWindow(window)
	C.gtk_layer_auto_exclusive_zone_enable(w)
}

func SetExclusiveZone(window *gtk.Window, size int) {
	w := nativeWindow(window)
	C.gtk_layer_set_exclusive_zone(w, C.int(size))
}

func SetAnchor(window *gtk.Window, side LayerShellEdgeFlags, pinned bool) {
	w := nativeWindow(window)
	C.gtk_layer_set_anchor(w, C.GtkLayerShellEdge(side), boolConv(pinned))
}

func SetMargin(window *gtk.Window, side LayerShellEdgeFlags, margin int) {
	w := nativeWindow(window)
	C.gtk_layer_set_margin(w, C.GtkLayerShellEdge(side), C.int(margin))
}

func SetMonitor(window *gtk.Window, monitor *gdk.Monitor) {
	C.gtk_layer_set_monitor(nativeWindow(window), nativeMonitor(monitor))
}

func SetKeyboardMode(window *gtk.Window, mode LayerShellKeyboardMode) {
	w := nativeWindow(window)
	C.gtk_layer_set_keyboard_mode(w, C.GtkLayerShellKeyboardMode(mode))
}
