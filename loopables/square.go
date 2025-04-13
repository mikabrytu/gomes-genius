package loopables

import (
	"fmt"
	"littlejumbo/genius/utils"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
)

type Square struct {
	name  string
	rect  render.RectSpecs
	color render.Color
}

func NewSquare(name string, rect render.RectSpecs, color render.Color) {
	square := &Square{
		name:  name,
		rect:  rect,
		color: color,
	}

	lifecycle.Register(lifecycle.Loopable{
		Init:   square.init,
		Update: square.update,
	})
}

func (s *Square) init() {
	events.Subscribe(events.INPUT_MOUSE_CLICK_DOWN, func(params ...any) error {
		position := params[0].([]any)
		click := utils.Click{
			X: position[0].([]any)[0].(int32),
			Y: position[0].([]any)[1].(int32),
		}

		if utils.IsClickInsideRect(click, s.rect) {
			fmt.Printf("Click inside square %v\n", s.name)
		}

		return nil
	})
}

func (s *Square) update() {
	render.DrawSimpleShapes(s.rect, s.color)
}
