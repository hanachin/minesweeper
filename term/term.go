package term

// #include <stdio.h>
import "C"

import (
	"fmt"
	"github.com/k0kubun/go-termios"
)

const (
	ColorBlack = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

func Clear() {
	fmt.Print("\x1B[2J")
}

func SetCursor(x, y int) {
	fmt.Printf("\x1B[%d;%dH", y, x)
}

func ResetColor() {
	fmt.Print("\x1B[m")
}

func SetForegroundColor(c int) {
	fmt.Printf("\x1B[%dm", c + 30)
}

func SetBackgroundColor(c int) {
	fmt.Printf("\x1B[%dm", c + 40)
}

func Print(s string) {
	fmt.Print(s)
}

func Println(s string) {
	fmt.Println(s)
}

func WithGameMode(f func()) {
	var originalTerm, gameTerm termios.Termios
	if err := originalTerm.GetAttr(termios.Stdin); err != nil {
		panic(err)
	}
	gameTerm = originalTerm
	gameTerm.LFlag ^= termios.ICANON
	gameTerm.LFlag ^= termios.ECHO
	gameTerm.CC[termios.VMIN] = 1
	gameTerm.CC[termios.VTIME] = 0
	if err := gameTerm.SetAttr(termios.Stdin, termios.TCSANOW); err != nil {
		panic(err)
	}
	defer func() {
		if err := originalTerm.SetAttr(termios.Stdin, termios.TCSANOW); err != nil {
			panic(err)
		}
	}()
	f()
}

func Getc() byte {
	return byte(C.fgetc(C.stdin))
}
