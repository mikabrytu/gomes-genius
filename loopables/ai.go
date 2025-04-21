package loopables

import (
	"fmt"
	event_names "littlejumbo/genius/events"
	"math/rand"
	"time"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
)

var squares []*Square

const WAIT_TIME = 500

func NewAi() {
	println("AI initialized")

	lifecycle.Register(lifecycle.Loopable{
		Init:   _init,
		Update: update,
	})
}

func LoadSquares(sList []*Square) {
	println("AI squares loaded")

	squares = sList
}

func NewAISequence(size int) {
	println("Generating AI sequence")

	if size == 0 {
		panic("Size of the sequence cannot be 0")
	}

	sequence := make([]int, size)

	for i := range size {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(squares))
		sequence[i] = index
	}

	fmt.Printf("AI sequence: %v\n", sequence)

	events.Emit(event_names.GENIUS_AI_SEQUENCE_FINISHED, sequence)
}

func PlayAINote(note int) {
	fmt.Printf("Playing AI note %d\n", note)

	squares[note].Click()

	time.AfterFunc(WAIT_TIME*time.Millisecond, func() {
		println("Time waited before emitting event")

		events.Emit(event_names.GENIUS_AI_SINGLE_NOTE_FINISHED)
	})
}

func _init() {
	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_F, func(params ...any) error {
		NewAISequence(1)
		return nil
	})
}

func update() {

}
