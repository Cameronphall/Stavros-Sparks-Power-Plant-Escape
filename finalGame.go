package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type FinalGame struct {
	IsSolved   bool
	JustSolved bool
	Buttons    [8]rl.Rectangle

	Numbers   [4]int
	Operators [4]string

	Current     *float64
	PendingOp   string
	Target      float64
	UsedNumbers [4]bool
	ResetButton rl.Rectangle
}

func NewFinalGame() *FinalGame {
	g := &FinalGame{}
	g.Numbers = [4]int{2, 6, 4, 12}
	g.Operators = [4]string{"+", "-", "*", "/"}
	g.Target = 48
	g.Current = nil
	g.PendingOp = ""
	g.initButtons()
	return g
}

func (g *FinalGame) Reset() {
	g.Current = nil
	g.PendingOp = ""
	g.IsSolved = false
	g.JustSolved = false

	for i := range g.UsedNumbers {
		g.UsedNumbers[i] = false
	}
}


func (g *FinalGame) Draw() {

	rl.ClearBackground(rl.Black)

	rl.DrawText("TARGET: 48", 320, 120, 26, rl.Green)
	rl.DrawText("CURRENT:", 560, 120, 26, rl.RayWhite)

	currentText := "None"
	if g.Current != nil {
		currentText = fmt.Sprintf("%.0f", *g.Current)
	}

	rl.DrawText(currentText, 720, 120, 26, rl.Yellow)

	for i, rect := range g.Buttons {

		var color rl.Color

		if i < 4 {

			if g.UsedNumbers[i] {
				color = rl.DarkGray 
			} else {
				color = rl.Red 
			}

		} else { 

			op := g.Operators[i-4]

			if g.PendingOp == op {
				color = rl.Blue 
			} else {
				color = rl.DarkBlue
			}
		}

		rl.DrawRectangleRec(rect, color)
		rl.DrawRectangleLinesEx(rect, 2, rl.Black)

		var label string

		if i < 4 {
			label = fmt.Sprintf("%d", g.Numbers[i])
		} else {
			label = g.Operators[i-4]
		}

		textWidth := rl.MeasureText(label, 24)

		rl.DrawText(
			label,
			int32(rect.X+rect.Width/2-float32(textWidth)/2),
			int32(rect.Y+rect.Height/2-12),
			24,
			rl.White,
		)
	}
	
	rl.DrawRectangleRec(g.ResetButton, rl.DarkGreen)
	rl.DrawRectangleLinesEx(g.ResetButton, 2, rl.Black)

	label := "RESET"
	textWidth := rl.MeasureText(label, 24)

	rl.DrawText(
		label,
		int32(g.ResetButton.X+g.ResetButton.Width/2-float32(textWidth)/2),
		int32(g.ResetButton.Y+g.ResetButton.Height/2-12),
		24,
		rl.White,
	)

	if g.IsSolved {
		rl.DrawText("SOLVED", 450, 220, 26, rl.Green)

	}
}

func (g *FinalGame) initButtons() {

	buttonWidth := float32(90)
	buttonHeight := float32(50)
	spacing := float32(12)

	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	textBoxHeight := float32(80)
	bottomMargin := float32(20)
	offset := float32(40)
	y := screenHeight - textBoxHeight - buttonHeight - bottomMargin - offset

	totalWidth := float32(8)*buttonWidth + float32(7)*spacing
	startX := (screenWidth - totalWidth) / 2

	for i := 0; i < 8; i++ {
		g.Buttons[i] = rl.NewRectangle(
			startX+float32(i)*(buttonWidth+spacing),
			y,
			buttonWidth,
			buttonHeight,
		)
	}
	resetWidth := g.Buttons[0].Width*2 + 12
	resetHeight := g.Buttons[0].Height

	resetX := g.Buttons[3].X - g.Buttons[0].Width + 90
	resetY := g.Buttons[0].Y - resetHeight - 15

	g.ResetButton = rl.NewRectangle(resetX, resetY, resetWidth, resetHeight)

}

func applyOperation(left float64, op string, right float64) float64 {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		return left
	}
}

func (g *FinalGame) Update() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mouse := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(mouse, g.ResetButton) {
			g.Reset()
			return
		}
	}
	if g.IsSolved {
		return
	}

	if !rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		return
	}

	mouse := rl.GetMousePosition()

	for i, rect := range g.Buttons {
		if !rl.CheckCollisionPointRec(mouse, rect) {
			continue
		}

		
		if i < 4 {

			
			if g.UsedNumbers[i] {
				return
			}

			value := float64(g.Numbers[i])

			if g.Current == nil {
				g.Current = &value
				g.UsedNumbers[i] = true 
				return
			}

			if g.PendingOp == "" {
				return
			}

			result := applyOperation(*g.Current, g.PendingOp, value)
			g.Current = &result
			g.PendingOp = ""
			g.UsedNumbers[i] = true 

			if int(*g.Current) == int(g.Target) && g.allNumbersUsed() {
				g.IsSolved = true
				g.JustSolved = true
			}

			return
		}

		if g.Current == nil {
			return
		}

		g.PendingOp = g.Operators[i-4]
		return
	}
}

func (g *FinalGame) allNumbersUsed() bool {
	for _, used := range g.UsedNumbers {
		if !used {
			return false
		}
	}
	return true
}
