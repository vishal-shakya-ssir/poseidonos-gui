package handler

import (
	"A-module/log"
)

var bmcSendChan chan string
var bmcReceiveChan chan string

func init() {
	bmcSendChan = make(chan string, 1)
}

func ConnectToBMC() error {
	return nil
}

func GetBMCResponse() string {
	return <-bmcReceiveChan
}

func SendBMC(bmcMsg string) {
	bmcSendChan <- bmcMsg
	log.Printf("SendBMC : Message -> bmcSendChan\n%s", bmcMsg)
}

func writeToBMCSomething(bmcMsg string) {
}
