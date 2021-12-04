package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2021/utils"
)

func main() {
	input := utils.InputString("day04/input.txt")
	bingo := parseBingo(input)
	fmt.Printf("Task 1. The board first to win has score %d\n", bingo.play())
	// play() actually alters the state of the boards, so it would be best to reset them
	// in-between, but it appears to still return the right value so laziness prevails...
	fmt.Printf("Task 2. The board last to win has score %d\n", bingo.playToLastWin())
}

func parseBingo(s string) bingo {
	lines := utils.SplitLine(s)

	drawsInput := strings.Split(lines[0], ",")
	draws := make([]int, len(drawsInput))
	for i, nbr := range drawsInput {
		draws[i], _ = strconv.Atoi(nbr)
	}

	var boards []board
	startLine := 2
	for i, line := range lines[2:] {
		absoluteIndex := i + 2
		// Requires there to be an empty newline at the end of the input.
		if line == "" {
			boards = append(boards, parseBoard(lines[startLine:absoluteIndex]))
			startLine = absoluteIndex + 1
		}
	}

	return bingo{
		boards: boards,
		draws:  draws,
	}
}

func parseBoard(lines []string) board {
	board := make([][]boardNumber, len(lines))
	for row, line := range lines {
		// Can be empty since there are double spaces.
		words := strings.Split(line, " ")
		board[row] = make([]boardNumber, 0, len(words))
		for _, word := range words {
			if word == "" {
				continue
			}
			nbr, _ := strconv.Atoi(word)
			boardNumber := boardNumber{
				nbr:    nbr,
				marked: false,
			}
			board[row] = append(board[row], boardNumber)
		}
	}
	return board
}

type bingo struct {
	boards []board
	draws  []int
}

type board [][]boardNumber

type boardNumber struct {
	nbr    int
	marked bool
}

func (b *bingo) play() int {
	for _, draw := range b.draws {
		score := b.draw(draw)
		if score != 0 {
			return score
		}
	}
	return 0
}

func (b *bingo) playToLastWin() int {
	boardsThatWon := make([]bool, len(b.boards))
	winCount := 0
	for _, draw := range b.draws {
		for i, board := range b.boards {
			if !boardsThatWon[i] {
				board.draw(draw)
				if board.hasWon() {
					boardsThatWon[i] = true
					winCount++
					if winCount == len(b.boards) {
						return draw * board.sumOfUnmarked()
					}
				}
			}
		}
	}
	return 0
}

func (b *bingo) draw(drawNumber int) int {
	for _, board := range b.boards {
		board.draw(drawNumber)
		if board.hasWon() {
			return drawNumber * board.sumOfUnmarked()
		}
	}
	return 0
}

func (b *board) draw(drawNumber int) {
	for rowNumber, row := range *b {
		for colNumber, boardNumber := range row {
			if boardNumber.nbr == drawNumber {
				(*b)[rowNumber][colNumber].marked = true
			}
		}
	}
}

func (b *board) hasWon() bool {
	// Mark all as marked by default, setting them to false when there's an unmarked number.
	allInColMarked := make([]bool, len((*b)[0]))
	for i := range allInColMarked {
		allInColMarked[i] = true
	}

	for _, row := range *b {
		allInRowMarked := true
		for col, boardNumber := range row {
			if !boardNumber.marked {
				allInColMarked[col] = false
				allInRowMarked = false
				// Continue parsing to check all columns.
			}
		}
		if allInRowMarked {
			return true
		}
	}
	for _, marked := range allInColMarked {
		if marked {
			return true
		}
	}
	return false
}

func (b *board) sumOfUnmarked() int {
	sum := 0
	for _, row := range *b {
		for _, boardNumber := range row {
			if !boardNumber.marked {
				sum += boardNumber.nbr
			}
		}
	}
	return sum
}

func (b board) String() string {
	var builder strings.Builder
	for _, row := range b {
		for _, boardNumber := range row {
			if boardNumber.marked {
				builder.WriteRune('M')
			}
			s := strconv.Itoa(boardNumber.nbr)
			builder.WriteString(s)
			builder.WriteRune(' ')
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
