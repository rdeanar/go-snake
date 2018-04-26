package main

import (
	"math/rand"
	"time"
	"math"
)

type Coord struct {
	x int
	y int
}

type Path []Coord

type State struct {
	snake             Path
	apple             Coord
	direction         Coord
	prevStepDirection Coord
	crashed           bool
}

var UP Coord = Coord{0, -1}
var RIGHT Coord = Coord{1, 0}
var DOWN Coord = Coord{0, 1}
var LEFT Coord = Coord{-1, 0}

var gameState State

var fieldWidth int
var fieldHeight int

func (s *State) pushDirection(d Coord) {
	if s.prevStepDirection.x+d.x != 0 || s.prevStepDirection.y+d.y != 0 || len(s.snake) <= 1 {
		s.direction = d
	}
}

func willCrash(head Coord, list Path) bool {
	if len(list) < 5 { // only snake with length >= 5 can be self eaten
		return false
	}
	for _, v := range list {
		if v == head {
			return true
		}
	}
	return false
}

func (s *State) move() {

	// if already crashed, re-run game
	if s.crashed {
		setInitValues()
	}

	nextHead := s.snake.getNextCoord(s.direction)

	if willCrash(nextHead, s.snake) {
		s.crashed = true
		return
	}

	s.snake = append(s.snake, s.snake.getNextCoord(s.direction))

	if nextHead != s.apple {
		s.snake = s.snake[1:]
	} else {
		// generate new apple
		s.apple = getRandomCoord() //TODO exclude snake path
	}
	s.prevStepDirection = s.direction
}

func (p Path) getNextCoord(direction Coord) Coord {
	headPosition := p[len(p)-1]
	return headPosition.plus(direction)
}

func (base Coord) plus(add Coord) Coord {
	// TODO make it simpler, less universal
	base.x = int(math.Mod(float64(base.x+int(math.Mod(float64(add.x), float64(fieldWidth)))+fieldWidth), float64(fieldWidth)))
	base.y = int(math.Mod(float64(base.y+int(math.Mod(float64(add.y), float64(fieldHeight)))+fieldHeight), float64(fieldHeight)))

	return base
}

func getRandomCoord() Coord {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	return Coord{
		x: random.Intn(fieldWidth - 1),
		y: random.Intn(fieldHeight - 1),
	}
}

func setInitValues() {
	gameState.snake = Path{getRandomCoord()}
	gameState.apple = getRandomCoord()
	gameState.prevStepDirection = LEFT
	gameState.direction = LEFT
	gameState.crashed = false
}
