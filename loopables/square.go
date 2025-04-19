package loopables

import (
	event_names "littlejumbo/genius/events"
	"littlejumbo/genius/utils"
	"time"

	"github.com/mikabrytu/gomes-engine/audio"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
)

type Square struct {
	id          int
	name        string
	rect        render.RectSpecs
	color       render.Color
	note        string
	blinkIn     bool
	blinkOut    bool
	blinkStart  time.Time
	blinkFinish time.Time
	canPlay     bool
}

const BLINK_TIME time.Duration = 165

func NewSquare(name string, id int, rect render.RectSpecs, color render.Color, note string) *Square {
	square := &Square{
		id:       id,
		name:     name,
		rect:     rect,
		color:    color,
		note:     note,
		blinkIn:  false,
		blinkOut: false,
		canPlay:  true,
	}

	lifecycle.Register(lifecycle.Loopable{
		Init:   square.init,
		Update: square.update,
	})

	return square
}

func (s *Square) EnablePlay(enabled bool) {
	s.canPlay = enabled
}

func (s *Square) Click() {
	s.blinkIn = true
	s.blinkStart = setBlinkTime()
	audio.PlaySFX(s.note)
}

func (s *Square) init() {
	events.Subscribe(events.INPUT_MOUSE_CLICK_DOWN, func(params ...any) error {
		position := params[0].([]any)
		click := utils.Click{
			X: position[0].([]any)[0].(int32),
			Y: position[0].([]any)[1].(int32),
		}

		if utils.IsClickInsideRect(click, s.rect) { // TODO: Try to use the Click() code
			s.blinkIn = true
			s.blinkStart = setBlinkTime()

			if s.canPlay {
				audio.PlaySFX(s.note)
			}

			events.Emit(event_names.GENIUS_PLAYER_SINGLE_NOTE_FINISHED, s.id)
		}

		return nil
	})
}

func (s *Square) update() {
	var c render.Color
	var done bool

	if s.blinkIn {
		c, done = blink(s.color, render.Transparent, s.blinkStart)
		if done {
			s.blinkIn = false
			s.blinkOut = true
			s.blinkFinish = setBlinkTime()
		}
	} else if s.blinkOut {
		c, done = blink(render.Transparent, s.color, s.blinkFinish)
		if done {
			s.blinkOut = false
		}
	} else {
		c = s.color
	}

	render.DrawSimpleShapes(s.rect, c)
}

func blink(original, target render.Color, endTime time.Time) (render.Color, bool) {
	var c render.Color
	var finished bool = false

	totalBlinkDuration := BLINK_TIME * time.Millisecond
	elapsed := time.Since(endTime.Add(-totalBlinkDuration))
	t := float64(elapsed) / float64(totalBlinkDuration)

	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	c = utils.LerpColor(original, target, t)

	if t >= 1 {
		finished = true
	}

	return c, finished
}

func setBlinkTime() time.Time {
	return time.Now().Add(BLINK_TIME * time.Millisecond)
}
