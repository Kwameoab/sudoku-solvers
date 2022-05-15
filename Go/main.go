package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Board struct {
	board [9][9]int
}

func (this *Board) print() {
	res := ""
	for i := 0; i < 9; i++ {
		if i == 3 || i == 6 {
			res += "------+-------+------\n"
		}
		for j := 0; j < 9; j++ {
			val := this.board[i][j]
			if val == 0 {
				res += "  "
			} else {
				res += fmt.Sprint(val)
				res += " "
			}

			if j == 2 || j == 5 {
				res += "| "
			}
		}
		res += "\n"
	}
	fmt.Println(res)
}

func (this *Board) importFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		for index, value := range split {
			num, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			this.board[i][index] = num
		}
		i++
	}
}

func (this Board) useableNum(num int, i int, j int) bool {
	for row := 0; row < 9; row++ {
		if num == this.board[i][row] {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if num == this.board[col][j] {
			return false
		}
	}

	boxi := (i / 3) * 3
	boxj := (j / 3) * 3
	for bi := 0; bi < 3; bi++ {
		for bj := 0; bj < 3; bj++ {
			if this.board[boxi+bi][boxj+bj] == num {
				return false
			}
		}
	}
	return true
}

func (this *Board) solve() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// fmt.Println("Here", this.board[i][j])
			if this.board[i][j] == 0 {

				for try := 1; try < 10; try++ {
					// fmt.Println("try", try)
					if this.useableNum(try, i, j) {
						this.board[i][j] = try
						// this.print()
						if this.solve() {
							return true
						}
						this.board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	b := Board{[9][9]int{}}

	filename := "../txt/test-2.txt"

	b.importFile(filename)

	fmt.Printf("Importing from file %s gave this board", filename)
	b.print()

	fmt.Println("\nStarting to solve...")

	start := time.Now()

	if b.solve() {
		duration := time.Since(start)
		fmt.Println("\nWas able to solve board")
		b.print()
		fmt.Printf("It took: %d milliseconds", duration.Milliseconds())
	} else {
		fmt.Println("Could not solve board")
	}
}
