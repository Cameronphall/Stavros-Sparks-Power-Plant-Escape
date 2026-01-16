package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Breaker struct {
	Pos        rl.Vector2
	OffRec     rl.Rectangle
	OnRec      rl.Rectangle
	SwitchBase rl.Rectangle
	Switch     rl.Rectangle

	LeftZone  rl.Rectangle
	RightZone rl.Rectangle

	IsOn bool
}

func (b *Breaker) DrawBreaker(number int) {

	rl.DrawText(
		fmt.Sprintf("%d", number),
		int32(b.Pos.X-30), 
		int32(b.Pos.Y+20), 
		24,               
		rl.RayWhite,
	)

	rl.DrawRectangleRounded(
		rl.NewRectangle(b.Pos.X, b.Pos.Y, b.OffRec.Width, b.OffRec.Height),
		0.2,
		8,
		rl.DarkGray,
	)

	track := rl.NewRectangle(
		b.Pos.X+10,
		b.Pos.Y+20,
		b.OffRec.Width-20,
		b.OffRec.Height-40,
	)
	rl.DrawRectangleRounded(track, 0.2, 8, rl.Gray)

	rockerWidth := track.Width / 2
	rockerHeight := track.Height

	var rockerX float32
	if b.IsOn {
		rockerX = track.X + rockerWidth
	} else {
		rockerX = track.X
	}

	rocker := rl.NewRectangle(
		rockerX,
		track.Y,
		rockerWidth,
		rockerHeight,
	)

	rl.DrawRectangleRounded(rocker, 0.3, 8, rl.LightGray)
	rl.DrawRectangleRoundedLines(rocker, 0.3, 8, rl.DarkGray)

	offColor := rl.Color{R: 200, G: 100, B: 100, A: 255}
	onColor := rl.Color{R: 100, G: 200, B: 100, A: 255}

	if !b.IsOn {
		rl.DrawText("OFF", int32(b.Pos.X+15), int32(b.Pos.Y+5), 16, offColor)
		rl.DrawText("ON", int32(int32(b.Pos.X)+int32(b.OffRec.Width)-50), int32(b.Pos.Y+5), 16, rl.Gray)
	} else {
		rl.DrawText("OFF", int32(b.Pos.X+15), int32(b.Pos.Y+5), 16, rl.Gray)
		rl.DrawText("ON", int32(int32(b.Pos.X)+int32(b.OffRec.Width)-50), int32(b.Pos.Y+5), 16, onColor)
	}
}

func (b *Breaker) HandleClick(mouse rl.Vector2, switchSound rl.Sound) {

	if rl.CheckCollisionPointRec(mouse, b.LeftZone) {
		b.IsOn = false
		rl.PlaySound(switchSound)
		return
	}

	if rl.CheckCollisionPointRec(mouse, b.RightZone) {
		b.IsOn = true
		rl.PlaySound(switchSound)
		return
	}
}
