package main

import (
	"fmt"
	"github.com/karellincoln/design-pattern/structural/adapter/outsideServer"
)

type windowsAdapter struct {
	windowMachine *outsideServer.Windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.InsertIntoUSBPort()
}
