package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	inputX = kingpin.Flag("xsize", "The width of the grid").Short('x').Default("25").Int()
	inputY = kingpin.Flag("ysize", "The height of the grid").Short('y').Default("25").Int()
	inputP = kingpin.Flag("percentage", "Percentage of living cells at the start").Short('p').Default("50").Int()
	inputI = kingpin.Flag("iterations", "Number of iterations. Any negative number will use the default, infinity").Short('i').Default("-1").Int()
	inputF = kingpin.Flag("fps", "Frames per second, how log to wait until the next iteration is displayed").Short('f').Default("10").Int()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()
	w := NewWorld(*inputX, *inputY)
	w.RandomiseAndInitialise(*inputP)
	sleepTime := time.Duration(1000 / *inputF) * time.Millisecond
	//main loop
	i := *inputI
	for {
		if i == 0 {
			break
		}
		i--
		cmd := exec.Command("clear")
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/c", "cls")
		}
		cmd.Stdout = os.Stdout
		cmd.Run()
		w.PrintGrid()
		fmt.Printf("Generation: %d\n", -i)
		w.Iterate()
		time.Sleep(sleepTime)
	}
}
