package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TileEmpty = iota
	TileWall
)

type Sokoban struct {
	grid             [][]int
	cols, rows       int
	tileSize         int32
	originX, originY int32

	playerX, playerY int
	crateX, crateY   int
	goalX, goalY     int

	crateTex     rl.Texture2D
	blockers     []rl.Rectangle
	drawGrid     bool
	startCrateX  int
	startCrateY  int
	startPlayerX int
	startPlayerY int
	crateSound rl.Sound
}

type Blocker struct {
	Rect rl.Rectangle
}

func NewSokoban(crateTex rl.Texture2D, crateSound rl.Sound) *Sokoban {
	blockers := []rl.Rectangle{
		rl.NewRectangle(0, 0, 1024, 120),
		rl.NewRectangle(0, 900, 1024, 124),
		rl.NewRectangle(0, 0, 64, 1024),
		rl.NewRectangle(960, 0, 64, 1024),

		rl.NewRectangle(690, 390, 110, 110),
		rl.NewRectangle(375, 670, 110, 110),
		rl.NewRectangle(240, 390, 110, 110),
	}

	grid := make([][]int, 16)

	for y := 0; y < 16; y++ {
		grid[y] = make([]int, 16)
	}

	return &Sokoban{
		grid:     grid,
		rows:     len(grid),
		cols:     len(grid[0]),
		tileSize: 64,
		originX:  0,
		originY:  0,

		playerX: 1, playerY: 3,
		crateX: 3, crateY: 3,
		goalX: 11, goalY: 11,

		crateTex:     crateTex,
		blockers:     blockers,
		drawGrid:     false,
		
		startCrateX:  3,
		startCrateY:  3,
		crateSound: crateSound,

	}
}

func (s *Sokoban) isWalkable(x, y int) bool {

	if y < 0 || y >= s.rows || x < 0 || x >= s.cols {
		return false
	}

	if s.grid[y][x] == TileWall {
		return false
	}

	return true
}

func (s *Sokoban) Update() {
	if rl.IsKeyPressed(rl.KeyR) {
		s.crateX = s.startCrateX
		s.crateY = s.startCrateY
		return
	}
	dx, dy := 0, 0

	if rl.IsKeyPressed(rl.KeyW) {
		dy = -1
	} else if rl.IsKeyPressed(rl.KeyS) {
		dy = 1
	} else if rl.IsKeyPressed(rl.KeyA) {
		dx = -1
	} else if rl.IsKeyPressed(rl.KeyD) {
		dx = 1
	}

	if dx == 0 && dy == 0 {
		return
	}

	nextPX := s.playerX + dx
	nextPY := s.playerY + dy

	nextPlayerRect := s.playerWorldRect(nextPX, nextPY)
	crateRect := s.crateWorldRect(s.crateX, s.crateY)

	if rl.CheckCollisionRecs(nextPlayerRect, crateRect) {

		nextCX := s.crateX + dx
		nextCY := s.crateY + dy

		if !s.isWalkable(nextCX, nextCY) || s.crateBlockedAt(nextCX, nextCY) {
			
			return
		}

		s.crateX = nextCX
		s.crateY = nextCY
		rl.PlaySound(s.crateSound)
	}

	if !s.isWalkable(nextPX, nextPY) ||
		s.rectBlocked(s.playerWorldRect(nextPX, nextPY)) {
		return
	}
	s.playerX = nextPX
	s.playerY = nextPY
}

func (s *Sokoban) IsSolved() bool {
	crate := s.crateWorldRect(s.crateX, s.crateY)
	goal := s.goalWorldRect()
	return rl.CheckCollisionRecs(crate, goal)
}

func (s *Sokoban) gridToWorld(x, y int) (int32, int32) {
	wx := s.originX + int32(x)*s.tileSize
	wy := s.originY + int32(y)*s.tileSize
	return wx, wy
}

func (s *Sokoban) Draw() {

	if s.drawGrid {

		for x := 0; x <= s.cols; x++ {
			px := s.originX + int32(x)*s.tileSize
			rl.DrawLine(px, s.originY,
				px, s.originY+int32(s.rows)*s.tileSize,
				rl.Fade(rl.Lime, 0.6))
		}

		for y := 0; y <= s.rows; y++ {
			py := s.originY + int32(y)*s.tileSize
			rl.DrawLine(s.originX, py,
				s.originX+int32(s.cols)*s.tileSize,
				py,
				rl.Fade(rl.Lime, 0.6))
		}
		for _, r := range s.blockers {
			rl.DrawRectangleLinesEx(r, 2, rl.Red)
		}
	}

	gx, gy := s.gridToWorld(s.goalX, s.goalY)
	rl.DrawRectangle(gx, gy, s.tileSize, s.tileSize, rl.Color{R: 230, G: 180, B: 40, A: 200})

	cx, cy := s.gridToWorld(s.crateX, s.crateY)

	scale := float32(3)

	w := float32(s.crateTex.Width) * scale
	h := float32(s.crateTex.Height) * scale

	centerX := float32(cx) + float32(s.tileSize)/2
	centerY := float32(cy) + float32(s.tileSize)/2

	drawX := centerX - w/2
	drawY := centerY - h/2

	rl.DrawTextureEx(
		s.crateTex,
		rl.NewVector2(drawX, drawY),
		0,
		scale,
		rl.White,
	)

}


func (s *Sokoban) GetPlayerWorldPos() (float32, float32) {
	wx := float32(s.originX + int32(s.playerX)*s.tileSize)
	wy := float32(s.originY + int32(s.playerY)*s.tileSize)
	return wx, wy
}

func (s *Sokoban) crateWorldRect(x, y int) rl.Rectangle {

	scale := float32(3)

	worldX := float32(s.originX + int32(x)*s.tileSize)
	worldY := float32(s.originY + int32(y)*s.tileSize)

	crateW := float32(s.crateTex.Width) * scale
	crateH := float32(s.crateTex.Height) * scale

	cx := worldX + float32(s.tileSize)/2
	cy := worldY + float32(s.tileSize)/2

	left := cx - crateW/2
	top := cy - crateH/2

	inset := float32(12)

	return rl.NewRectangle(
		left+inset,
		top+inset,
		crateW-inset*2,
		crateH-inset*2,
	)
}

func (s *Sokoban) crateBlockedAt(x, y int) bool {
	crateRect := s.crateWorldRect(x, y)

	for _, b := range s.blockers {
		if rl.CheckCollisionRecs(crateRect, b) {
			return true
		}
	}
	return false
}

func (s *Sokoban) goalWorldRect() rl.Rectangle {
	wx := float32(s.originX + int32(s.goalX)*s.tileSize)
	wy := float32(s.originY + int32(s.goalY)*s.tileSize)

	size := float32(s.tileSize)

	return rl.NewRectangle(wx, wy, size, size)
}

func (s *Sokoban) playerWorldRect(x, y int) rl.Rectangle {
	wx := float32(s.originX + int32(x)*s.tileSize)
	wy := float32(s.originY + int32(y)*s.tileSize)

	pw := float32(s.tileSize) * 0.8
	ph := float32(s.tileSize) * 0.8

	cx := wx + float32(s.tileSize)/2 - pw/2
	cy := wy + float32(s.tileSize)/2 - ph/2

	return rl.NewRectangle(cx, cy, pw, ph)
}

func (s *Sokoban) rectBlocked(r rl.Rectangle) bool {
	for _, b := range s.blockers {
		if rl.CheckCollisionRecs(r, b) {
			return true
		}
	}
	return false
}
