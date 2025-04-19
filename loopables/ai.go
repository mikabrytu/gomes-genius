package loopables

import (
	"container/list"
	event_names "littlejumbo/genius/events"
	"littlejumbo/genius/utils"
	"math/rand"
	"time"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
)

var squares []*Square
var ssize int
var running bool = false
var playing bool = false

const WAIT_TIME = 500

func NewAi() {
	lifecycle.Register(lifecycle.Loopable{
		Init:   _init,
		Update: update,
	})
}

func EnablePlay(enabled bool) {
	playing = enabled
}

func LoadSquares(sList []*Square) {
	squares = sList
}

func _init() {
	ssize = 0

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_F, func(params ...any) error {
		if !running {
			running = true
			return nil
		}

		return nil
	})
}

func update() {
	if running && playing && squares != nil {
		var sequence *list.List = list.New()
		ssize++

		for range ssize {
			rand.Seed(time.Now().UnixNano())
			index := rand.Intn(len(squares))
			squares[index].Click()
			sequence.PushBack(index)
		}

		isequence := utils.ListToIntArray(sequence)

		playing = false
		events.Emit(event_names.GENIUS_AI_SEQUENCE_FINISHED, isequence)
	}
}
