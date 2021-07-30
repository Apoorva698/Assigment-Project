package main

//InBounds is a helper function to check if a coordinate tupel is inside the grid.
func InBounds(x, y, xSize, ySize int) bool {
	return x >= 0 && x < xSize && y >= 0 && y < ySize
}
