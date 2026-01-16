package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)


type Door struct{
	DoorRec rl.Rectangle
	Color rl.Color
}

func NewDoor(x, y, width, height float32)Door{
	return  Door{
		DoorRec: rl.NewRectangle(x, y, width, height),
		Color: rl.Green,
	}
}

func (d *Door) DrawDoor(){
	rl.DrawRectangle(int32(d.DoorRec.X), int32(d.DoorRec.Y), int32(d.DoorRec.Width), int32(d.DoorRec.Height), d.Color)
}

func (d *Door)checkProximity(p Player) bool{
	if p.X >= d.DoorRec.X && p.X <= d.DoorRec.X + d.DoorRec.Width && p.Y <= 635 {
		d.Color = rl.Green
		return true
	}else{
		d.Color = rl.Blue
		return false
	}
}

func (d *Door)checkProximityComputer(p Player) bool{
	if p.X >= d.DoorRec.X && p.X <= d.DoorRec.X + d.DoorRec.Width && p.Y <= 250 {
		d.Color = rl.Green
		return true
	}else{
		d.Color = rl.Blue
		return false
	}
}