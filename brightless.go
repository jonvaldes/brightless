package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Constants
var minBrightness float64 = 0.1
var maxBrightness float64 = 1.0
var deltaBrightness float64 = 0.1

// Utility functions
func checkUserErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		usage()
		os.Exit(1)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func findFirst(lines []string, text string) (string, error) {
	for _, l := range lines {
		if strings.Contains(l, text) {
			return l, nil
		}
	}
	return "", fmt.Errorf("Error: Can't find text \"%s\"", text)
}

func usage() {
	fmt.Printf(
		`
Usage: brightness [DELTA]
Modify the monitor brightness, adding DELTA to it.
This program allows brightness to be between %.2f and %.2f.

Examples
  - To increase brightness by 1/10th
      brightness 0.1

  - To dim monitor by 1/10th
      brightness -0.1

  - To set full brightness
      brightness 1

  - To set full dimming
      brightness -1
`, minBrightness, maxBrightness)
}

func main() {

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	delta, err := strconv.ParseFloat(os.Args[1], 64)
	checkUserErr(err)

	out, err := exec.Command("xrandr", "--verbose").Output()
	checkErr(err)

	lines := strings.Split(string(out), "\n")

	// Get main monitor name
	monitorLine, err := findFirst(lines, "connected primary")
	checkErr(err)
	monitor := strings.Split(monitorLine, " ")[0]

	// Get current brightness
	brightnessLine, err := findFirst(lines, "Brightness")
	checkErr(err)

	stringB := strings.Split(brightnessLine, ":")[1][1:]
	brightness, err := strconv.ParseFloat(stringB, 64)
	checkErr(err)

	// Calculate new brightness
	newBrightness := math.Min(math.Max(brightness+delta, minBrightness), maxBrightness)

	// Set brightness
	checkErr(exec.Command("xrandr", "--output", monitor, "--brightness", fmt.Sprintf("%f", newBrightness)).Run())
	return
}
