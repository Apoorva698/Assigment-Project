package main

import (
	"fmt"
	"testing"
)

func TestWorld_Neighbours(t *testing.T) {

	w0 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w1 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, false, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w2 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, false, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w3 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, false, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w4 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w5 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w6 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, false, false, false,
		false, false, false, false, false,
	}}

	w7 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, false, false,
		false, false, false, false, false,
	}}

	w8 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, false, false, false, false,
	}}

	//corner top left

	w9 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w10 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, true, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w11 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, true, false, false, false,
		true, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	w12 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		true, true, false, false, false,
		true, true, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}}

	//corner bottom right
	w13 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, true,
	}}

	w14 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, true, false,
		false, false, false, false, true,
	}}

	w15 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, true, true,
		false, false, false, false, true,
	}}

	w16 := &World{generation: 0, xSize: 5, ySize: 5, cells: []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, true, true,
		false, false, false, true, true,
	}}

	tests := []struct {
		name string
		w    *World
		x    int
		y    int
		want int
	}{
		{"0 Neighbours", w0, 2, 2, 0},
		{"1 Neighbours", w1, 2, 2, 1},
		{"2 Neighbours", w2, 2, 2, 2},
		{"3 Neighbours", w3, 2, 2, 3},
		{"4 Neighbours", w4, 2, 2, 4},
		{"5 Neighbours", w5, 2, 2, 5},
		{"6 Neighbours", w6, 2, 2, 6},
		{"7 Neighbours", w7, 2, 2, 7},
		{"8 Neighbours", w8, 2, 2, 8},

		{"0 Neighbours, corner top left", w9, 0, 0, 0},
		{"1 Neighbours, corner top left", w10, 0, 0, 1},
		{"2 Neighbours, corner top left", w11, 0, 0, 2},
		{"3 Neighbours, corner top left", w12, 0, 0, 3},

		{"0 Neighbours, corner bottom right", w13, 4, 4, 0},
		{"1 Neighbours, corner bottom right", w14, 4, 4, 1},
		{"2 Neighbours, corner bottom right", w15, 4, 4, 2},
		{"3 Neighbours, corner bottom right", w16, 4, 4, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.AliveNeighbours(tt.x, tt.y); got != tt.want {
				tt.w.PrintGrid()
				t.Errorf("World.AliveNeighbours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_InBounds(t *testing.T) {

	w := NewWorld(5, 5)
	tests := []struct {
		name string
		w    *World
		x    int
		y    int
		want bool
	}{
		{"Top Left", w, 0, 0, true},
		{"Top Right", w, 4, 0, true},
		{"Bottom Left", w, 0, 4, true},
		{"Bottom Right", w, 4, 4, true},

		{"Out Left", w, -1, 0, false},
		{"Out Right", w, 5, 0, false},
		{"Out Top", w, 0, -1, false},
		{"Out Bottom", w, 5, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.w
			if got := InBounds(tt.x, tt.y, w.xSize, w.ySize); got != tt.want {
				t.Errorf("InBounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_Iterate(t *testing.T) {

	// Any live cell with fewer than two live neighbors dies, as if by underpopulation
	w0 := NewWorld(3, 3)
	/*
		0 0 0
		0 1 0
		0 0 1
	*/

	w0.SetCoordinates(1, 1, true)
	w0.SetCoordinates(2, 2, true)

	w0want := NewWorld(3, 3)

	/*
		0 0 0
		0 0 0
		0 0 0
	*/

	// Any live cell with two or three live neighbors lives on to the next generation
	w1 := NewWorld(3, 3)

	/*
		0 0 0
		1 1 0
		0 1 1
	*/

	w1.SetCoordinates(0, 1, true)
	w1.SetCoordinates(1, 1, true)
	w1.SetCoordinates(1, 2, true)
	w1.SetCoordinates(2, 2, true)

	w1want := NewWorld(3, 3)

	/*
		0 0 0
		1 1 1
		1 1 1
	*/

	w1want.SetCoordinates(0, 1, true)
	w1want.SetCoordinates(0, 2, true)
	w1want.SetCoordinates(1, 1, true)
	w1want.SetCoordinates(1, 2, true)
	w1want.SetCoordinates(2, 1, true)
	w1want.SetCoordinates(2, 2, true)

	// Any live cell with more than three live neighbors dies, as if by overpopulation
	w2 := NewWorld(3, 3)

	/*
		1 1 1
		1 1 1
		1 1 1
	*/

	w2.SetCoordinates(0, 0, true)
	w2.SetCoordinates(1, 0, true)
	w2.SetCoordinates(2, 0, true)
	w2.SetCoordinates(0, 1, true)
	w2.SetCoordinates(1, 1, true)
	w2.SetCoordinates(2, 1, true)
	w2.SetCoordinates(0, 2, true)
	w2.SetCoordinates(1, 2, true)
	w2.SetCoordinates(2, 2, true)

	w2want := NewWorld(3, 3)

	/*
		1 0 1
		0 0 0
		1 0 1
	*/

	w2want.SetCoordinates(0, 0, true)
	w2want.SetCoordinates(2, 0, true)
	w2want.SetCoordinates(0, 2, true)
	w2want.SetCoordinates(2, 2, true)

	// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction
	w3 := NewWorld(3, 3)

	/*
		0 0 0
		0 0 0
		1 1 1
	*/

	w3.SetCoordinates(0, 2, true)
	w3.SetCoordinates(1, 2, true)
	w3.SetCoordinates(2, 2, true)

	w3want := NewWorld(3, 3)

	/*
		0 0 0
		0 1 0
		0 1 0
	*/

	w3want.SetCoordinates(1, 1, true)
	w3want.SetCoordinates(1, 2, true)

	tests := []struct {
		name string
		w    *World
		want *World
	}{
		{"Any live cell with fewer than two live neighbors dies, as if by underpopulation.", w0, w0want},
		{"Any live cell with two or three live neighbors lives on to the next generation.", w1, w1want},
		{"Any live cell with more than three live neighbors dies, as if by overpopulation.", w2, w2want},
		{"Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.", w3, w3want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Iterate()

			if !tt.w.AreWorldsEqual(tt.want) {
				t.Errorf("World.Iterate() = %v, want %v", tt.w, tt.want)
				fmt.Println("Got:")
				tt.w.PrintGrid()
				fmt.Println("Want:")
				tt.want.PrintGrid()
			}
		})
	}
}

func TestWorld_Equal(t *testing.T) {

	// Equal, all dead
	w0a := NewWorld(5, 5)
	w0b := NewWorld(5, 5)

	// Different xSize
	w1a := NewWorld(4, 5)
	w1b := NewWorld(5, 5)

	// Different ySize
	w2a := NewWorld(5, 5)
	w2b := NewWorld(5, 6)

	// Different xSize and ySize
	w3a := NewWorld(5, 3)
	w3b := NewWorld(3, 5)

	// Different cell length
	w4a := NewWorld(5, 5)
	w4b := NewWorld(5, 5)
	w4b.cells = []bool{false, false, false}

	// Different cell(s)
	w5a := NewWorld(5, 5)
	w5b := NewWorld(5, 5)
	w5b.SetCoordinates(0, 0, true)

	// Equal, some alive
	w6a := NewWorld(5, 5)
	w6a.SetCoordinates(0, 0, true)
	w6a.SetCoordinates(1, 0, true)
	w6a.SetCoordinates(2, 4, true)
	w6a.SetCoordinates(3, 2, true)

	w6b := NewWorld(5, 5)
	w6b.SetCoordinates(0, 0, true)
	w6b.SetCoordinates(1, 0, true)
	w6b.SetCoordinates(2, 4, true)
	w6b.SetCoordinates(3, 2, true)

	tests := []struct {
		name string
		wa   *World
		wb   *World
		want bool
	}{
		{"Equal, all dead", w0a, w0b, true},
		{"Different xSize", w1a, w1b, false},
		{"Different ySize", w2a, w2b, false},
		{"Different xSize and ySize", w3a, w3b, false},
		{"Different cell length", w4a, w4b, false},
		{"Different cells", w5a, w5b, false},
		{"Equal, some alive ", w6a, w6b, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wa.AreWorldsEqual(tt.wb); got != tt.want {
				t.Errorf("World.AreWorldsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_RandInit(t *testing.T) {
	tests := []struct {
		name       string
		w          *World
		percentage int
		wantAlive  int
	}{
		{"0 percent", NewWorld(10, 5), 0, 0},
		{"30 percent", NewWorld(10, 5), 30, 15},
		{"50 percent", NewWorld(10, 5), 50, 25},
		{"100 percent", NewWorld(10, 5), 100, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.RandomiseAndInitialise(tt.percentage)

			count := 0
			for _, v := range tt.w.cells {
				if v {
					count++
				}
			}

			if count != tt.wantAlive {
				t.Errorf("World.RandomiseAndInitialise() ran, %v, cells alive,  want %v", count, tt.wantAlive)
			}
		})
	}
}

func TestWorld_Set(t *testing.T) {

	// Set dead cell to alive
	w0 := NewWorld(3, 3)
	w0.cells = []bool{
		false, false, false,
		false, false, false,
		false, false, false,
	}

	w0want := NewWorld(3, 3)
	w0want.cells = []bool{
		false, false, false,
		false, true, false,
		false, false, false,
	}

	// Set dead cell to dead
	w1 := NewWorld(3, 3)
	w1.cells = []bool{
		false, false, false,
		false, false, false,
		false, false, false,
	}

	w1want := NewWorld(3, 3)
	w1want.cells = []bool{
		false, false, false,
		false, false, false,
		false, false, false,
	}
	// Set living cell to alive
	w2 := NewWorld(3, 3)
	w2.cells = []bool{
		false, false, false,
		false, true, false,
		false, false, false,
	}

	w2want := NewWorld(3, 3)
	w2want.cells = []bool{
		false, false, false,
		false, true, false,
		false, false, false,
	}
	// Set living cell to dead
	w3 := NewWorld(3, 3)
	w3.cells = []bool{
		false, false, false,
		false, true, false,
		false, false, false,
	}

	w3want := NewWorld(3, 3)
	w3want.cells = []bool{
		false, false, false,
		false, false, false,
		false, false, false,
	}

	tests := []struct {
		name string
		w    *World
		want *World
		x    int
		y    int
		val  bool
	}{
		{"Set dead cell to alive", w0, w0want, 1, 1, true},
		{"Set dead cell to dead", w1, w1want, 1, 1, false},
		{"Set living cell to alive", w2, w2want, 1, 1, true},
		{"Set living cell to dead", w3, w3want, 1, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.w
			w.SetCoordinates(tt.x, tt.y, tt.val)

			if !w.AreWorldsEqual(tt.want) {
				t.Errorf("World.SetCoordinates() = %v, want %v", w, tt.want)
			}
		})
	}
}

func TestWorld_Get(t *testing.T) {

	w0 := NewWorld(5, 3)
	w0.cells = []bool{
		false, true, true, true, true,
		true, true, true, true, true,
		true, true, true, true, true,
	}

	w1 := NewWorld(5, 3)
	w1.cells = []bool{
		true, true, true, true, false,
		true, true, true, true, true,
		true, true, true, true, true,
	}

	w2 := NewWorld(5, 3)
	w2.cells = []bool{
		true, true, true, true, true,
		true, true, true, true, true,
		false, true, true, true, true,
	}

	w3 := NewWorld(5, 3)
	w3.cells = []bool{
		true, true, true, true, true,
		true, true, true, true, true,
		true, true, true, true, false,
	}

	w4 := NewWorld(5, 3)
	w4.cells = []bool{
		true, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, false,
	}

	w5 := NewWorld(5, 3)
	w5.cells = []bool{
		false, false, false, false, true,
		false, false, false, false, false,
		false, false, false, false, false,
	}

	w6 := NewWorld(5, 3)
	w6.cells = []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		true, false, false, false, false,
	}

	w7 := NewWorld(5, 3)
	w7.cells = []bool{
		false, false, false, false, false,
		false, false, false, false, false,
		false, false, false, false, true,
	}

	tests := []struct {
		name string
		w    *World
		x    int
		y    int
		want bool
	}{
		{"Top left, dead", w0, 0, 0, false},
		{"Top right, dead", w1, 4, 0, false},
		{"Bottom left, dead", w2, 0, 2, false},
		{"Bottom right, dead", w3, 4, 2, false},
		{"Top left, alive", w4, 0, 0, true},
		{"Top right, alive", w5, 4, 0, true},
		{"Bottom left, alive", w6, 0, 2, true},
		{"Bottom right, alive", w7, 4, 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.w
			if got := w.GetCurrentCoordinates(tt.x, tt.y); got != tt.want {
				t.Errorf("World.GetCurrentCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_Print(t *testing.T) {

	w0 := NewWorld(5, 5)
	w0.RandomiseAndInitialise(30)

	tests := []struct {
		name string
		w    *World
	}{
		{"Print randomly initialized board", w0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.PrintGrid()
		})
	}
}
