package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type Vector struct {
	x int
	y int
}

func parseLine(line string) Vector {
	parts := strings.Split(line, " ")
	amt, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("this should never happen")
	}
	if parts[0] == "forward" {
		return Vector{
			x: amt,
			y: 0,
		}
	} else if parts[0] == "up" {
		return Vector{
			x: 0,
			y: -amt,
		}
	}
	return Vector{
		x: 0,
		y: amt,
	}
}

func main() {
	lines := strings.Split(getInput(), "\n")

	pos := Vector{0, 0}
	for _, line := range lines {
		amount := parseLine(line)
		pos.x += amount.x
		pos.y += amount.y
	}
	fmt.Printf("point: %+v\n", pos)
}
