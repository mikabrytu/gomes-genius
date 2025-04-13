package utils

import "github.com/mikabrytu/gomes-engine/render"

type Click struct {
	X int32
	Y int32
}

func IsClickInsideRect(click Click, rect render.RectSpecs) bool {
	return click.X >= rect.PosX && click.X <= rect.PosX+rect.Width &&
		click.Y >= rect.PosY && click.Y <= rect.PosY+rect.Height
}
