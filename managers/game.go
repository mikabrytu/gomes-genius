package managers

import (
	"littlejumbo/genius/loopables"

	"github.com/mikabrytu/gomes-engine/render"
)

func Game() {
	var size int32 = 200
	var offset int32 = 10
	rect := render.RectSpecs{
		PosX:   offset,
		PosY:   offset,
		Width:  size,
		Height: size,
	}

	red := rect
	green := rect
	blue := rect
	yellow := rect

	green.PosX += size + offset
	blue.PosY += size + offset
	yellow.PosX += size + offset
	yellow.PosY += size + offset

	loopables.DrawSquare(red, render.Red)
	loopables.DrawSquare(green, render.Green)
	loopables.DrawSquare(blue, render.Blue)
	loopables.DrawSquare(yellow, render.Yellow)
}
