//main file for go-life program, explore the game of life

package main

import "fmt"

const rows = 10
const cols = 10

type cell struct {
	rowNum       int
	colNumb      int
	cellAlive    bool
	madeAliveCtr int
	madeDeadCtr  int
}

func (c cell) print() {

	fmt.Printf("Cell: Row:%d Col:%d Alive:%t Made Live:%d Made Dead:%d \n", c.rowNum, c.colNumb, c.cellAlive, c.madeAliveCtr, c.madeDeadCtr)
}

func (c cell) symbolize() string {
	if c.cellAlive {
		return "@"
	} else {
		return "-"
	}
}

func main() {
	fmt.Println("welcome to the game of life")
	c := cell{
		rowNum:       1,
		colNumb:      1,
		cellAlive:    false,
		madeAliveCtr: 0,
		madeDeadCtr:  0,
	}

	c.print()
	fmt.Println(c.symbolize())
	c.cellAlive = true
	fmt.Println(c.symbolize())

	//define grid as rowxcol array. can use array since size known at compile time and doesnt change during execution
	grid := [rows][cols]cell{}

	//initialize grid
	for kr, r := range grid {
		for kc, c := range r {
			grid[kr][kc].cellAlive = true
			grid[kr][kc].colNumb = kc
			grid[kr][kc].rowNum = kr
			grid[kr][kc].madeAliveCtr = 0
			grid[kr][kc].madeDeadCtr = 0

		}
	}
	fmt.Println(grid)

	// end of program
	fmt.Println("end of program")

}
