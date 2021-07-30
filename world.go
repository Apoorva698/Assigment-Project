package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//World holds the grid on which the cells are placed as well as a few of it's attributes.
type World struct {
	generation   int
	xSize, ySize int
	cells        []bool
}

//NewWorld initializes a new world (constructor) with an all-dead grid
func NewWorld(x, y int) *World {
	cells := make([]bool, x*y)
	return &World{generation: 0, xSize: x, ySize: y, cells: cells}
}

//RandomiseAndInitialise sets a given percentage of the cells in the grid to "alive".
//The grid is then randomized, leaving a random grid with a defined amount of living cells.
func (w *World) RandomiseAndInitialise(percentage int) {
	//Calculate number of living cells
	numAlive := percentage * len(w.cells) / 100
	//Insert living cells at the beginning
	for i := 0; i < numAlive; i++ {
		w.cells[i] = true
	}
	//Randomize slice
	vals := w.cells
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(vals); n > 0; n-- {
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
	}
	w.cells = vals
}

//GetCurrentCoordinates retrieves a cell's state by it's x and y coordinates (alive: true, dead: false)
func (w *World) GetCurrentCoordinates(x, y int) bool {
	if !InBounds(x, y, w.xSize, w.ySize) {
		log.Fatal("Invalid Coordinate")
	}
	return w.cells[y*(w.xSize)+x]
}

//PrintGrid displays a ASCII representation of the grid to stout
func (w *World) PrintGrid() {
	//Top margin
	fmt.Print("╔")
	for x := 1; x <= w.xSize; x++ {
		fmt.Print("══")
	}
	fmt.Println("╗")
	//Rows
	for y := 0; y < w.ySize; y++ {
		fmt.Print("║")
		//Columns
		for x := 0; x < w.xSize; x++ {
			if w.GetCurrentCoordinates(x, y) {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("║")
	}
	//Bottom margin
	fmt.Print("╚")
	for x := 1; x <= w.xSize; x++ {
		fmt.Print("══")
	}
	fmt.Println("╝")
}

//AliveNeighbours returns the number of alive neighbours of a given cell
func (w *World) AliveNeighbours(x, y int) int {
	count := 0
	arr := []int{-1, 0, 1}

	for _, v1 := range arr {
		for _, v2 := range arr {
			if InBounds(x+v1, y+v2, w.xSize, w.ySize) {
				if w.GetCurrentCoordinates(x+v1, y+v2) && !(v1 == 0 && v2 == 0) {
					count++
				}
			}
		}
	}
	return count
}

//SetCoordinates sets a cell defined by it's x and y coordinates to a given state (alive: true, dead: false)
func (w *World) SetCoordinates(x, y int, val bool) {
	if !InBounds(x, y, w.xSize, w.ySize) {
		log.Fatal("Invalid Coordinate")
	}

	w.cells[y*(w.xSize)+x] = val
}

//Iterate implements the rules for the problem statement. It takes the current grid, applies the 4 rules and sets the grid to it's new state.
func (w *World) Iterate() {
	gbOld := NewWorld(w.xSize, w.ySize)
	copy(gbOld.cells, w.cells)
	for y := 0; y < w.ySize; y++ {
		for x := 0; x < w.xSize; x++ {
			//Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
			if !gbOld.GetCurrentCoordinates(x, y) && gbOld.AliveNeighbours(x, y) == 3 {
				w.SetCoordinates(x, y, true)
				continue
			}
			//Any live cell with fewer than two live neighbors dies, as if by underpopulation.
			if gbOld.GetCurrentCoordinates(x, y) && gbOld.AliveNeighbours(x, y) < 2 {
				w.SetCoordinates(x, y, false)
				continue
			}
			//Any live cell with two or three live neighbors lives on to the next generation.
			if gbOld.GetCurrentCoordinates(x, y) && ((gbOld.AliveNeighbours(x, y) == 2) || (gbOld.AliveNeighbours(x, y) == 3)) {
				//No need to set, already alive
				continue
			}
			//Any live cell with more than three live neighbors dies, as if by overpopulation.
			if gbOld.GetCurrentCoordinates(x, y) && (gbOld.AliveNeighbours(x, y) > 3) {
				w.SetCoordinates(x, y, false)
				continue
			}
		}
	}
}

//AreWorldsEqual is a helper function to compare a GameBoard to another. It returns true if size, dimensions and the cell's states are identical
func (w *World) AreWorldsEqual(gb2 *World) bool {
	if w.xSize != gb2.xSize || w.ySize != gb2.ySize {
		return false
	}
	if len(w.cells) != len(gb2.cells) {
		return false
	}
	for k := range w.cells {
		if w.cells[k] != gb2.cells[k] {
			return false
		}
	}
	return true
}
