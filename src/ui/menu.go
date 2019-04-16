package ui

import (
	"fmt"

	c "github.com/CooChristmas/Christmas/src/color"
	"github.com/CooChristmas/Christmas/src/version"
)

func ShowMenu() {
	fmt.Println("")
	fmt.Println(c.YELLOW, "                                     Christmas v"+version.GetVersion())
	fmt.Println(c.CYAN, "                               	by CooChristmas")
	fmt.Println(c.CYAN, "                               	Blog: https://CooChristmas.github.io")
	fmt.Println("")
	fmt.Println(c.GREEN, " Please use `tab` to autocomplete commands,")
	fmt.Println(c.GREEN, " or type `exit` to quit this program.")
	fmt.Println("")
}
