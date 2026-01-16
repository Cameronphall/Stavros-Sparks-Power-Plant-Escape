package main

import rl "github.com/gen2brain/raylib-go/raylib"

type BreakerBox struct {
	Pos      rl.Vector2
	Width    float32
	Height   float32
	Breakers []Breaker
}

func NewBreakerBox() BreakerBox {
	box := BreakerBox{
		Pos:      rl.NewVector2(40, 40),
		Width:    460,
		Height:   920,
		Breakers: make([]Breaker, 10),
	}

	breakerWidth := float32(380)
	breakerHeight := float32(70)
	gap := float32(20)

	for i := range box.Breakers {
		x := box.Pos.X + (box.Width-breakerWidth)/2
		y := box.Pos.Y + gap + float32(i)*(breakerHeight+gap)
		leftZone := rl.NewRectangle(x, y, breakerWidth/2, breakerHeight)
		rightZone := rl.NewRectangle(x+breakerWidth/2, y, breakerWidth/2, breakerHeight)
		box.Breakers[i] = Breaker{
			Pos:        rl.NewVector2(x, y),
			OffRec:     rl.NewRectangle(x, y, breakerWidth, breakerHeight),
			SwitchBase: rl.NewRectangle(x+20, y+15, breakerWidth-40, breakerHeight-30),
			Switch:     rl.NewRectangle(x+25, y+20, breakerWidth-50, breakerHeight-40),
			LeftZone:   leftZone,
			RightZone:  rightZone,
			IsOn:       false,
		}
	}

	return box
}

func (box *BreakerBox) DrawAllBreakers() {
	for i := range box.Breakers {
		box.Breakers[i].DrawBreaker(i + 1)
	}
}

func (box *BreakerBox) DrawPanelBackground() {
	panel := rl.NewRectangle(box.Pos.X, box.Pos.Y, box.Width, box.Height)

	rl.DrawRectangleRounded(panel, 0.05, 8, rl.Color{R: 40, G: 40, B: 40, A: 255})

	rl.DrawRectangleRoundedLines(panel, 0.05, 8, rl.Color{R: 70, G: 70, B: 70, A: 255})

	rl.DrawRectangleLines(
		int32(panel.X)+6,
		int32(panel.Y)+6,
		int32(panel.Width)-12,
		int32(panel.Height)-12,
		rl.Color{R: 20, G: 20, B: 20, A: 255},
	)

	screwColor := rl.Color{R: 180, G: 180, B: 180, A: 255}
	r := float32(6)

	rl.DrawCircle(int32(panel.X+15), int32(panel.Y+15), r, screwColor)
	rl.DrawCircle(int32(panel.X+panel.Width-15), int32(panel.Y+15), r, screwColor)
	rl.DrawCircle(int32(panel.X+15), int32(panel.Y+panel.Height-15), r, screwColor)
	rl.DrawCircle(int32(panel.X+panel.Width-15), int32(panel.Y+panel.Height-15), r, screwColor)
}

func (box *BreakerBox) HandleClicks(mouse rl.Vector2, switchSound rl.Sound) {
	for i := range box.Breakers {
		box.Breakers[i].HandleClick(mouse, switchSound)
	}
}

func (box *BreakerBox) GetStates() []bool {
    states := make([]bool, len(box.Breakers))
    for i := range box.Breakers {
        states[i] = box.Breakers[i].IsOn
    }
    return states
}