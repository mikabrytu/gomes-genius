package managers

import (
	event_names "littlejumbo/genius/events"
	"littlejumbo/genius/loopables"
	"littlejumbo/genius/utils"
	"time"

	"github.com/mikabrytu/gomes-engine/audio"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/render"
)

var red, green, blue, yellow *loopables.Square
var sequence []int
var playerCount int
var aiCount int

const (
	WAIT_SEQUENCE_START   = 500
	WAIT_SEQUENCE_FAIL    = 1500
	WAIT_SEQUENCE_SUCCESS = 1000
	WAIT_NEXT_AI_NOTE     = 500
)

func Game() {
	prepareSquares()
	prepareAI()

	events.Subscribe(event_names.GENIUS_AI_SEQUENCE_FINISHED, func(params ...any) error {
		onNewSequence(params[0].([]any)[0].([]any)[0].([]int))
		return nil
	})

	events.Subscribe(event_names.GENIUS_AI_SINGLE_NOTE_FINISHED, func(params ...any) error {
		onAINote()
		return nil
	})

	events.Subscribe(event_names.GENIUS_PLAYER_SINGLE_NOTE_FINISHED, func(params ...any) error {
		onPlayerNote(params[0].([]any)[0].([]any)[0].(int))
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

	time.AfterFunc(WAIT_SEQUENCE_START*time.Millisecond, func() {
		playAINote()
	})
}

func onPlayerNote(note int) {
	if note != sequence[playerCount] {
		red.Click(false)
		green.Click(false)
		blue.Click(false)
		yellow.Click(false)

		audio.PlaySFX(utils.SFX_FAIL)

		time.AfterFunc(WAIT_SEQUENCE_FAIL*time.Millisecond, func() {
			println("Wrong Sequence!")
			loopables.NewAISequence(len(sequence))
		})

		return
	}

	playerCount++

	if playerCount == len(sequence) {
		playerCount = 0
		enablePlayer(false)

		time.AfterFunc(WAIT_SEQUENCE_SUCCESS*time.Millisecond, func() {
			loopables.NewAISequence(len(sequence) + 1)
		})
	}
}

func onAINote() {
	if aiCount >= len(sequence) {
		enablePlayer(true)
	} else {
		time.AfterFunc(WAIT_NEXT_AI_NOTE*time.Millisecond, func() {
			playAINote()
		})
	}
}

func playAINote() {
	loopables.PlayAINote(sequence[aiCount])
	aiCount++
}

func enablePlayer(enable bool) {
	red.EnablePlay(enable)
	green.EnablePlay(enable)
	blue.EnablePlay(enable)
	yellow.EnablePlay(enable)
}
