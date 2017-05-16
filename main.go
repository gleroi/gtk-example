package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	out, _ := os.Create("out.log")
	defer out.Close()
	log.SetOutput(out)
	log.Printf("starting...")

	// Initialize GTK without parsing any command line arguments.
	args := os.Args
	gtk.Init(&args)
	log.Printf("Initialized...")

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Gtk examples for windows")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	log.Printf("Connected...")

	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Print("Unable to create Label:", err)
	}

	b, err := gtk.ButtonNewWithLabel("Bouton Test")
	if err != nil {
		log.Print("Unable to create Button:", err)
	}

	c, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	top, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	top.SetMarginStart(11)
	top.SetMarginEnd(11)
	top.SetMarginBottom(11)
	top.SetMarginTop(11)

	reveal, err := gtk.ExpanderNew("Expander")
	hbox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	for i := 1; i < 4; i++ {
		hb, _ := gtk.ButtonNewWithLabel(fmt.Sprintf("Machin %d", i))
		hbox.Add(hb)
	}
	reveal.Add(hbox)

	isReveal := false
	b.Connect("clicked", func() {
		isReveal = !isReveal
		reveal.SetExpanded(isReveal)
	})

	close, _ := gtk.ButtonNewFromIconName("window-close", gtk.ICON_SIZE_BUTTON)
	hbox.Add(close)
	close.Connect("clicked", func() {
		gtk.MainQuit()
	})
	c.Add(reveal)

	card, _ := CardNew()

	// Add the label to the window.
	top.Add(l)
	top.Add(b)
	c.Add(card)
	c.Add(top)
	c.SetChildPacking(card, true, true, 0, gtk.PACK_START)
	win.Add(c)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()
	log.Printf("Ready for gtk.Main")
	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
