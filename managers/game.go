package managers

import (
	"fmt"
	event_names "littlejumbo/genius/events"
	"littlejumbo/genius/loopables"
	"littlejumbo/genius/utils"
	"time"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/render"
)

var red, green, blue, yellow *loopables.Square
var sequence []int
var playerCount int
var aiCount int

func Game() {
	prepareSquares()
	prepareAI()

	events.Subscribe(event_names.GENIUS_AI_SEQUENCE_FINISHED, func(params ...any) error {
		onNewSequence(params[0].([]interface{})[0].([]interface{})[0].([]int))
		return nil
	})

	events.Subscribe(event_names.GENIUS_AI_SINGLE_NOTE_FINISHED, func(params ...any) error {
		onAINote()
		return nil
	})

	events.Subscribe(event_names.GENIUS_PLAYER_SINGLE_NOTE_FINISHED, func(params ...any) error {
		onPlayerNote(params[0].([]interface{})[0].([]interface{})[0].(int))
		return nil
	})
}

func prepareSquares() {
	var size int32 = 200
	var offset int32 = 10
	rect := render.RectSpecs{
		PosX:   offset,
		PosY:   offset,
		Width:  size,
		Height: size,
	}

	rrect := rect
	grect := rect
	brect := rect
	yrect := rect

	grect.PosX += size + offset
	brect.PosY += size + offset
	yrect.PosX += size + offset
	yrect.PosY += size + offset

	red = loopables.NewSquare("Red", 0, rrect, render.Red, utils.NOTE_C_PATH)
	green = loopables.NewSquare("Green", 1, grect, render.Green, utils.NOTE_E_PATH)
	blue = loopables.NewSquare("Blue", 2, brect, render.Blue, utils.NOTE_G_PATH)
	yellow = loopables.NewSquare("Yellow", 3, yrect, render.Yellow, utils.NOTE_B_PATH)
}

func prepareAI() {
	loopables.NewAi()
	loopables.LoadSquares([]*loopables.Square{red, green, blue, yellow})
}

func onNewSequence(s []int) {
	sequence = s
	aiCount = 0
	playerCount = 0

	enablePlayer(false)
	onAINote()
}

func onPlayerNote(note int) {
	playerCount++

	if playerCount == len(sequence) {
		playerCount = 0
		enablePlayer(false)

		time.AfterFunc(1*time.Second, func() {
			loopables.NewAISequence(len(sequence) + 1)
		})
	}
}

func onAINote() {
	if aiCount >= len(sequence) {
		fmt.Printf("AI played %d notes. Now player will play\n", aiCount)

		enablePlayer(true)
	} else {
		time.AfterFunc(1500*time.Millisecond, func() {
			playAINote()
		})
	}
}

func playAINote() {
	println("AI will play next note...")

	loopables.PlayAINote(sequence[aiCount])
	aiCount++
}

func enablePlayer(enable bool) {
	red.EnablePlay(enable)
	green.EnablePlay(enable)
	blue.EnablePlay(enable)
	yellow.EnablePlay(enable)
}
