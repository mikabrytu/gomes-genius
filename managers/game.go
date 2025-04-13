package managers

import (
	"littlejumbo/genius/loopables"
	"littlejumbo/genius/utils"

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

	loopables.NewSquare("Red", red, render.Red, utils.NOTE_C)
	loopables.NewSquare("Green", green, render.Green, utils.NOTE_E)
	loopables.NewSquare("Blue", blue, render.Blue, utils.NOTE_G)
	loopables.NewSquare("Yellow", yellow, render.Yellow, utils.NOTE_B)
}
