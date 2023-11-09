//main file for go-life program, explore the game of life

package main

import "fmt"

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

}
