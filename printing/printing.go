package printing

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

var reset = "\033[0m"
var red = "\033[31m"
var yellow = "\033[33m"
var blue = "\033[34m"

// PrintRed prints str red in console
func PrintRed(str string) {
	fmt.Print(red + str + reset)
}

// PrintBlue prints str blue in console
func PrintBlue(str string) {
	fmt.Print(blue + str + reset)
}

// PrintYellow prints str blue in console
func PrintYellow(str string) {
	fmt.Print(yellow + str + reset)
}

// Clear clears terminal
func Clear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Unfortunately, connect474 couldn't clear the screen")
	}
}
