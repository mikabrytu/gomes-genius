package loopables

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
)

func DrawSquare(rect render.RectSpecs, color render.Color) {
	lifecycle.Register(lifecycle.Loopable{
		Update: func() {
			render.DrawSimpleShapes(rect, color)
		},
	})
}
