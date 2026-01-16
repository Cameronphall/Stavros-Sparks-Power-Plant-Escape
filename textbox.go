package main

import rl "github.com/gen2brain/raylib-go/raylib"

type TextBox struct {
    Text    string
    Visible bool
}


func (t *TextBox) Draw() {
    if !t.Visible || t.Text == "" {
        return
    }

    screenW := rl.GetScreenWidth()
    screenH := rl.GetScreenHeight()

    padding := int32(20)
    height := int32(90)
    y := int32(screenH) - height - padding

    box := rl.NewRectangle(
        float32(padding),
        float32(y),
        float32(int32(screenW) - padding*2),
        float32(height),
    )


    rl.DrawRectangleRounded(box, 0.1, 10, rl.Color{R: 20, G: 20, B: 20, A: 220})

    rl.DrawRectangleRoundedLines(box, 0.1, 10, rl.Color{R: 90, G: 90, B: 90, A: 255})

    rl.DrawText(
        t.Text,
        padding+20,
        y+25,
        20,
        rl.RayWhite,
    )
}