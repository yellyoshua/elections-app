package main

import (
	"github.com/yellyoshua/elections-app/setups"
)

func main() {
	setups.Folders()
	setups.Environments()
	setups.ServerAndDatabase(false)
}
