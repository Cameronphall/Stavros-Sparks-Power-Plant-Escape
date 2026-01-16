package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerState int

const (
	IdleFront PlayerState = iota
	WalkFront
	IdleBack
	WalkBack
	WalkToward
)

type Player struct {
	X, Y  float32
	Flip  bool
	State PlayerState
	Anims map[PlayerState]*Animation
}

func NewPlayer() Player {

	idleFrontTex := rl.LoadTexture("Sprites/StavrosIdle.png")
	walkFrontTex := rl.LoadTexture("Sprites/Stavros walk.png")
	idleBackTex := rl.LoadTexture("Sprites/IdleBack.png")
	walkBackTex := rl.LoadTexture("Sprites/walkback.png")
	walkTowardTex := rl.LoadTexture("Sprites/WalkForward.png")

	p := Player{
		X:    200,
		Y:    650,
		Flip: false,
		Anims: map[PlayerState]*Animation{
			IdleFront: &Animation{
				Texture:     idleFrontTex,
				FrameWidth:  32,
				FrameHeight: 32,
				Frames:      2,
				Speed:       4,
			},
			WalkFront: &Animation{
				Texture:     walkFrontTex,
				FrameWidth:  32,
				FrameHeight: 32,
				Frames:      3,
				Speed:       8,
			},
			IdleBack: &Animation{
				Texture:     idleBackTex,
				FrameWidth:  32,
				FrameHeight: 32,
				Frames:      2,
				Speed:       4,
			},
			WalkBack: &Animation{
				Texture:     walkBackTex,
				FrameWidth:  32,
				FrameHeight: 32,
				Frames:      4,
				Speed:       8,
			},
			WalkToward: &Animation{
				Texture:     walkTowardTex,
				FrameWidth:  32,
				FrameHeight: 32,
				Frames:      4,
				Speed:       8,
			},
		},
	}

	return p
}

func (p *Player) Update() {

	speed := float32(200) * rl.GetFrameTime()

	moving := false

	if rl.IsKeyDown(rl.KeyA) {
		p.X -= speed
		p.Flip = true
		p.State = WalkFront
		moving = true
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.X += speed
		p.Flip = false
		p.State = WalkFront
		moving = true
	}

	if rl.IsKeyDown(rl.KeyW) {
		p.State = WalkBack
		moving = true
		p.Y -= speed
	}

	if rl.IsKeyDown(rl.KeyS) {
		p.State = WalkToward
		moving = true
		p.Y += speed
	}

	if !moving {
		if p.State == WalkBack || p.State == IdleBack {
			p.State = IdleBack
		} else {
			p.State = IdleFront
		}
	}

	if p.X < 0 {
		p.X = 0
	}

	if p.X+p.Width() > float32(rl.GetScreenWidth()) {
		p.X = float32(rl.GetScreenWidth()) - p.Width()
	}

	if p.Y < 625 {
		p.Y = 625
	}

	if p.Y+p.Height() > 775 {
		p.Y = 775 - p.Height()
	}

	p.Anims[p.State].Update()
}

func (p *Player) UpdateComp() {

	speed := float32(200) * rl.GetFrameTime()

	moving := false

	if rl.IsKeyDown(rl.KeyA) {
		p.X -= speed
		p.Flip = true
		p.State = WalkFront
		moving = true
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.X += speed
		p.Flip = false
		p.State = WalkFront
		moving = true
	}

	if rl.IsKeyDown(rl.KeyW) {
		p.State = WalkBack
		moving = true
		p.Y -= speed
	}

	if rl.IsKeyDown(rl.KeyS) {
		p.State = WalkToward
		moving = true
		p.Y += speed
	}

	if !moving {
		if p.State == WalkBack || p.State == IdleBack {
			p.State = IdleBack
		} else {
			p.State = IdleFront
		}
	}

	if p.X < 0 {
		p.X = 0
	}

	if p.X+p.Width() > float32(rl.GetScreenWidth()) {
		p.X = float32(rl.GetScreenWidth()) - p.Width()
	}

	if p.Y < 230 {
		p.Y = 230
	}

	if p.Y+p.Height() > float32(rl.GetScreenHeight()) {
		p.Y = float32(rl.GetScreenHeight()) - p.Height()
	}

	p.Anims[p.State].Update()
}

func (p *Player) Draw() {
	p.Anims[p.State].Draw(p.X, p.Y, p.Flip)
}

func (p *Player) Width() float32 {
	return float32(p.Anims[p.State].FrameWidth) * 3
}

func (p *Player) Height() float32 {
	return float32(p.Anims[p.State].FrameHeight) * 3
}

func (p *Player) UpdateFacing() {

	if rl.IsKeyDown(rl.KeyW) {
		p.State = WalkBack
		return
	}

	if rl.IsKeyDown(rl.KeyS) {
		p.State = WalkFront
		return
	}
}
