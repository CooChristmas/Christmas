package cmd

import (
	"github.com/CooChristmas/Christmas/src/ui"
	"github.com/CooChristmas/Christmas/src/util"
)

func Start() {
	util.DetectOS()
	util.ClearScreen()
	ui.StartMenu()
}
