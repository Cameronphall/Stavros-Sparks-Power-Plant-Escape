package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	Texture     rl.Texture2D
	FrameWidth  int
	FrameHeight int
	Frames      int
	Speed       float32
	Timer       float32
	Index       int
}

func NewAnimation(tex rl.Texture2D, frameWidth, frameHeight, frames int, speed float32) Animation {
	return Animation{
		Texture:     tex,
		FrameWidth:  frameWidth,
		FrameHeight: frameHeight,
		Frames:      frames,
		Speed:       speed,
		Timer:       0,
		Index:       0,
	}
}

func (a *Animation) Update() {
	a.Timer += rl.GetFrameTime()
	if a.Timer >= 1.0/a.Speed {
		a.Index++
		if a.Index >= a.Frames {
			a.Index = 0
		}
		a.Timer = 0
	}
}

func (a *Animation) Reset() {
	a.Index = 0
	a.Timer = 0
}
func (a *Animation) Draw(x, y float32, flip bool) {
    src := rl.Rectangle{
        X:      float32(a.Index * a.FrameWidth),
        Y:      0,
        Width:  float32(a.FrameWidth),
        Height: float32(a.FrameHeight),
    }

    
    if flip {
        src.Width = -src.Width
    }

    
    scale := float32(3)

    dest := rl.Rectangle{
        X:      x,
        Y:      y,
        Width:  float32(a.FrameWidth) * scale,
        Height: float32(a.FrameHeight) * scale,
    }

    origin := rl.Vector2{0, 0}

    rl.DrawTexturePro(a.Texture, src, dest, origin, 0, rl.White)
}

