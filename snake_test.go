package main

import (
	"testing"
)

func TestCoordPlus(t *testing.T) {

	testCases := []struct {
		base   Coord
		add    Coord
		width  int
		height int
		expect Coord
	}{
		{Coord{2, 3}, Coord{1, 1}, 10, 5, Coord{3, 4}},
		{Coord{9, 3}, Coord{1, 1}, 10, 5, Coord{0, 4}},
		{Coord{2, 4}, Coord{1, 1}, 10, 5, Coord{3, 0}},
		{Coord{0, 0}, Coord{-1, -1}, 10, 5, Coord{9, 4}},
		{Coord{0, 0}, Coord{22, 22}, 10, 5, Coord{2, 2}},
		{Coord{0, 0}, Coord{-22, -22}, 10, 5, Coord{8, 3}},
	}

	for _, tc := range testCases {

		fieldWidth = tc.width
		fieldHeight = tc.height

		actual := tc.base.plus(tc.add)
		if actual != tc.expect {
			t.Errorf("FAIL. got %d expect %d", actual, tc.expect)
		} else {
			t.Logf("OK. %d plus %d with dimensions %d equals %d", tc.base, tc.add, []int{tc.width, tc.height}, tc.expect)
		}
	}
}

func TestPushDirection(t *testing.T) {

	testCases := []struct {
		predDirection Coord
		direction     Coord
		expect        Coord
		snakeLength   int
	}{
		/* Single length snake */
		// Same direction
		{UP, UP, UP, 1},
		{RIGHT, RIGHT, RIGHT, 1},

		// Opposite direction
		{UP, DOWN, DOWN, 1},
		{RIGHT, LEFT, LEFT, 1},
		{DOWN, UP, UP, 1},
		{LEFT, RIGHT, RIGHT, 1},

		// Change direction
		{DOWN, RIGHT, RIGHT, 1},
		{LEFT, DOWN, DOWN, 1},

		/* Long length snake */
		// Same direction
		{UP, UP, UP, 2},
		{RIGHT, RIGHT, RIGHT, 2},

		// Opposite direction
		{UP, DOWN, UP, 2},
		{RIGHT, LEFT, RIGHT, 2},
		{DOWN, UP, DOWN, 2},
		{LEFT, RIGHT, LEFT, 2},

		// Change direction
		{DOWN, RIGHT, RIGHT, 2},
		{LEFT, DOWN, DOWN, 2},
	}

	state := &State{}
	for _, tc := range testCases {

		state.prevStepDirection = tc.predDirection
		state.direction = state.prevStepDirection
		state.pushDirection(tc.direction)
		state.snake = Path{}
		for i := 0; i < tc.snakeLength; i++ {
			state.snake = append(state.snake, getRandomCoord())
		}

		actual := state.direction
		if actual != tc.expect {
			t.Errorf("FAIL. prev: %d, push: %d, length: %d, got %d, expect %d", tc.predDirection, tc.direction, tc.snakeLength, actual, tc.expect)
		} else {
			t.Logf("OK. prev: %d, push: %d, length: %d, results: %d", tc.predDirection, tc.direction, tc.snakeLength, tc.expect)
		}
	}
}

func TestEatApple(t *testing.T) {

	state := &State{}
	state.apple = Coord{1, 3}
	state.snake = Path{Coord{0, 1}}
	state.prevStepDirection = DOWN

	moves := []Coord{
		RIGHT,
		DOWN,
		DOWN,
	}

	for _, direction := range moves {
		state.pushDirection(direction)
		state.move()
	}

	actual := len(state.snake)
	expected := 2
	if actual != expected {
		t.Errorf("FAIL. got %d expect %d", actual, expected)
	} else {
		t.Logf("OK. snake lench equals %d as expected", expected)
	}
}

func TestEatSelf(t *testing.T) {

	state := &State{}
	state.snake = Path{
		Coord{0, 0},
		Coord{0, 1},
		Coord{0, 2},
		Coord{1, 2},
		Coord{1, 1},
	}
	state.prevStepDirection = UP
	state.pushDirection(LEFT)
	state.move()

	actual := state.crashed
	expected := true
	if actual != expected {
		t.Errorf("FAIL. got %v expect %v", actual, expected)
	} else {
		t.Logf("OK. snake crashed as expected")
	}
}
