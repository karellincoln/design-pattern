package main

import "github.com/karellincoln/design-pattern/structural/adapter/outsideServer"

func main() {

	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windowsMachine := &outsideServer.Windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
