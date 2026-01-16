package main

import rl "github.com/gen2brain/raylib-go/raylib"

type BatteryHUD struct {
	Empty rl.Texture2D
	Full  rl.Texture2D
	PosX  int32
	PosY  int32
	Gap   int32
}

func NewBatteryHUD(emptyTex, fullTex rl.Texture2D) BatteryHUD {
	return BatteryHUD{
		Empty: emptyTex,
		Full:  fullTex,
		PosX:  0,
		PosY:  20,
		Gap:   40,
	}
}

func (b BatteryHUD) Draw(progress GameProgress) {

	earned := progress.Batteries()

	for i := 0; i < 3; i++ {
		startX := int32(rl.GetScreenWidth() - (3 * int(b.Gap)) - 20)
		x := startX + int32(i)*b.Gap

		rl.DrawTexture(b.Empty, x, b.PosY, rl.White)

		if i < earned {
			rl.DrawTexture(b.Full, x, b.PosY, rl.White)
		}
	}
}
