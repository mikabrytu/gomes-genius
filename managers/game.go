package managers

import (
	"fmt"
	event_names "littlejumbo/genius/events"
	"littlejumbo/genius/loopables"
	"littlejumbo/genius/utils"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/render"
)

var red, green, blue, yellow *loopables.Square
var sequence []int
var playerNoteCount int

func Game() {
	prepareSquares()
	prepareAI()

	events.Subscribe(event_names.GENIUS_AI_SEQUENCE_FINISHED, func(params ...any) error {
		sequence = params[0].([]interface{})[0].([]interface{})[0].([]int)
		fmt.Printf("AI Sequence: %v\n", sequence)

		enablePlayer(true)
		return nil
	})

	events.Subscribe(event_names.GENIUS_PLAYER_SINGLE_NOTE_FINISHED, func(params ...any) error {
		note := params[0].([]interface{})[0].([]interface{})[0].(int)
		fmt.Printf("Player note: %d\n", note)

		playerNoteCount++
		if playerNoteCount == len(sequence) {
			playerNoteCount = 0
			enablePlayer(false)
			loopables.EnablePlay(true)
		}

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
	loopables.EnablePlay(true)
	loopables.LoadSquares([]*loopables.Square{red, green, blue, yellow})
}

func enablePlayer(enable bool) {
	red.EnablePlay(enable)
	green.EnablePlay(enable)
	blue.EnablePlay(enable)
	yellow.EnablePlay(enable)
}
