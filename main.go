package main

import (
	"fmt"
)

const (
	N          = 10  // Ant and Antlion number
	LOOP_COUNT = 100 // Max loop count
	AREA_RANGE = 1   // Range where ants and antlions can move. When AREA_RANGE = 1 then ant can move in [-1, 1]
	DIMENSION  = 3   // Ant and Antlion dimension count and also the number of arguments of the target function
)

func main() {
	fmt.Println(LOOP_COUNT + AREA_RANGE)
}
