package utils

import "github.com/mikabrytu/gomes-engine/render"

// TODO: Incorporate this into the engine

type Click struct {
	X int32
	Y int32
}

func IsClickInsideRect(click Click, rect render.RectSpecs) bool {
	return click.X >= rect.PosX && click.X <= rect.PosX+rect.Width &&
		click.Y >= rect.PosY && click.Y <= rect.PosY+rect.Height
}

func LerpColor(original, target render.Color, time float64) render.Color {
	return render.Color{
		R: uint8(float64(original.R)*(1-time) + float64(target.R)*time),
		G: uint8(float64(original.G)*(1-time) + float64(target.G)*time),
		B: uint8(float64(original.B)*(1-time) + float64(target.B)*time),
		A: uint8(float64(original.A)*(1-time) + float64(target.A)*time),
	}
}
