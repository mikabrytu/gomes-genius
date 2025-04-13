package managers

import (
	"littlejumbo/genius/loopables"
	"littlejumbo/genius/utils"

	"github.com/mikabrytu/gomes-engine/render"
)

type Note struct {
	id   int8
	path string
}

var red, green, blue, yellow *loopables.Square

func Game() {
	instantiateSquares()
}

func instantiateSquares() {
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

	red = loopables.NewSquare("Red", rrect, render.Red, utils.NOTE_C_PATH)
	green = loopables.NewSquare("Green", grect, render.Green, utils.NOTE_E_PATH)
	blue = loopables.NewSquare("Blue", brect, render.Blue, utils.NOTE_G_PATH)
	yellow = loopables.NewSquare("Yellow", yrect, render.Yellow, utils.NOTE_B_PATH)
}
