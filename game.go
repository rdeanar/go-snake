package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

var keyboardEventsChannel = make(chan termbox.Key)

func main() {
	fieldWidth = 20
	fieldHeight = 14

	setInitValues()
	draw()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKeyboard(keyboardEventsChannel)

loop:
	for {
		select {
		case key := <-keyboardEventsChannel:
			switch key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowUp:
				gameState.pushDirection(UP)
			case termbox.KeyArrowRight:
				gameState.pushDirection(RIGHT)
			case termbox.KeyArrowDown:
				gameState.pushDirection(DOWN)
			case termbox.KeyArrowLeft:
				gameState.pushDirection(LEFT)
			}
		default:
			time.Sleep(200 * time.Millisecond)
			gameState.move()
			draw()
		}
	}
}

func listenToKeyboard(channel chan termbox.Key) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			channel <- ev.Key
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y := 0; y < fieldHeight; y++ {
		for x := 0; x < fieldWidth; x++ {
			termbox.SetCell(x*2, y, '·', termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	for _, s := range gameState.snake {
		termbox.SetCell(s.x*2, s.y, '×', termbox.ColorDefault, termbox.ColorDefault)
	}

	termbox.SetCell(gameState.apple.x*2, gameState.apple.y, '○', termbox.ColorGreen, termbox.ColorDefault)
	termbox.Flush()
}
