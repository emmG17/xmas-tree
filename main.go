package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	bold        = "\033[1m"
)

var ornamentColors = []string{colorRed, colorYellow, colorBlue, colorPurple, colorCyan}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func getRandomOrnament() string {
	ornaments := []string{"o", "*", "O", "°", "●"}
	colors := ornamentColors
	return colors[rand.Intn(len(colors))] + ornaments[rand.Intn(len(ornaments))] + colorReset
}

func drawTree(frame int) {
	tree := []string{
		"           " + colorYellow + "*" + colorReset,
		"          /|\\",
		"         / | \\",
		"        /  o  \\",
		"       /   |   \\",
		"      /  o   o  \\",
		"     /     |     \\",
		"    /   o  |  o   \\",
		"   /       |       \\",
		"  /  o   o   o   o  \\",
		" /         |         \\",
		"/    o  o  |  o  o    \\",
		"       " + colorRed + "   |||" + colorReset,
		"       " + colorRed + "   |||" + colorReset,
	}

	// Animate by replacing 'o' with random colored ornaments
	for i, line := range tree {
		newLine := ""
		for _, char := range line {
			if char == 'o' && frame%2 == 0 {
				newLine += getRandomOrnament()
			} else if char == 'o' {
				newLine += ornamentColors[rand.Intn(len(ornamentColors))] + "o" + colorReset
			} else {
				newLine += string(char)
			}
		}
		if i > 0 && i < len(tree)-2 {
			fmt.Println(colorGreen + newLine + colorReset)
		} else {
			fmt.Println(newLine)
		}
	}
}

func drawMessage() {
	message := []string{
		"",
		"  " + colorRed + bold + "╔══════════════════════════════╗" + colorReset,
		"  " + colorRed + bold + "║                              ║" + colorReset,
		"  " + colorRed + bold + "║   " + colorWhite + "🎄 MERRY CHRISTMAS! 🎄" + colorRed + "     ║" + colorReset,
		"  " + colorRed + bold + "║                              ║" + colorReset,
		"  " + colorRed + bold + "║      " + colorYellow + "Happy Holidays!" + colorRed + "         ║" + colorReset,
		"  " + colorRed + bold + "║                              ║" + colorReset,
		"  " + colorRed + bold + "╚══════════════════════════════╝" + colorReset,
	}

	for _, line := range message {
		fmt.Println(line)
	}
}

func main() {
	fmt.Println(colorWhite + "  Press Ctrl+C to exit\n" + colorReset)
	time.Sleep(2 * time.Second)

	for frame := range 50 {
		clearScreen()
		fmt.Println()
		drawTree(frame)
		drawMessage()
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}

	clearScreen()
	fmt.Println()
	drawTree(0)
	drawMessage()
	fmt.Println(colorWhite + "\n  🎅 Ho ho ho! 🎅\n" + colorReset)
}

