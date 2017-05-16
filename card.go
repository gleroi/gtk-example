package main

import "C"
import "github.com/gotk3/gotk3/gtk"

type Card struct {
	*gtk.Box
}

func CardNew() (*Card, error) {
	widget, err := content()
	if err != nil {
		return nil, err
	}
	return &Card{
		widget,
	}, nil
}

func content() (*gtk.Box, error) {
	container, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	label, _ := gtk.LabelNew("1")
	container.Add(label)
	container.SetChildPacking(label, true, true, 0, gtk.PACK_START)
	return container, nil
}
