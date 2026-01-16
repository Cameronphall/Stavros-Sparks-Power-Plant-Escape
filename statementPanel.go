package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"strings"
)
type Statement struct {
	Text  string
	Valid bool
}

type StatementPanel struct {
	Pos      rl.Vector2
	Width    float32
	Height   float32
	Items    []Statement
}

func NewStatementPanel() StatementPanel {
	return StatementPanel{
		Pos:   rl.NewVector2(520, 40),
		Width: 440,
		Height: 920,
		Items: []Statement{
			{Text: "Breaker 1 must be ON", Valid: false},
			{Text: "Breakers 2 and 3 must match", Valid: false},
			{Text: "At least 4 breakers must be ON", Valid: false},
			{Text: "Either breaker 7 or breaker 8 must be on", Valid: false},
			{Text: "No 3 consecutive breakers may be on", Valid: false},
			{Text: "Breaker 10 must be on if and only if breaker 9 is off", Valid: false},
			{Text: "No more than 6 breakers may be on", Valid: false},
			{Text: "At least 2 consecutive beakers must be off", Valid: false},
			{Text: "If a breaker number is a multiple of 3 at least one breaker next to it must be on", Valid: false},
			{Text: "If breaker 4 is on, then breaker 5 must be off", Valid: false},
		},
	}
}

func (p *StatementPanel) DrawPanelBackground() {
	panel := rl.NewRectangle(
		p.Pos.X,
		p.Pos.Y,
		p.Width,
		p.Height,
	)

	
	rl.DrawRectangleRounded(panel, 0.05, 8, rl.Color{R: 45, G: 45, B: 45, A: 255})

	
	rl.DrawRectangleRoundedLines(
		panel,
		0.05,
		8,
		rl.Color{R: 80, G: 80, B: 80, A: 255},
	)

	
}


func (p *StatementPanel) Draw() {
	p.DrawPanelBackground()
	p.DrawStatementsWrapped()
}

func (p *StatementPanel) DrawStatementsWrapped() {
    fontSize := 20
    padding := float32(20)
    lineSpacing := float32(8)

    
    totalStatements := len(p.Items)
    spacePerStatement := (p.Height - padding*2) / float32(totalStatements)

    for i, item := range p.Items {
        
        blockY := p.Pos.Y + padding + float32(i)*spacePerStatement

        
        statusColor := rl.Red
        if item.Valid {
            statusColor = rl.Green
        }
        rl.DrawCircle(int32(p.Pos.X+20), int32(blockY+8), 8, statusColor)

        
        maxWidth := int(p.Width - 60)
        wrapped := WrapText(item.Text, maxWidth, fontSize)

       
        for j, line := range wrapped {
            lineY := blockY + float32(j)*(float32(fontSize)+lineSpacing)
            rl.DrawText(line, int32(p.Pos.X+40), int32(lineY), int32(fontSize), rl.RayWhite)
        }
    }
}

func WrapText(text string, maxWidth int, fontSize int) []string {
    words := strings.Split(text, " ")
    var lines []string
    var currentLine string

    for _, word := range words {
        testLine := currentLine
        if testLine != "" {
            testLine += " " + word
        } else {
            testLine = word
        }

        if rl.MeasureText(testLine, int32(fontSize)) > int32(maxWidth) {
            lines = append(lines, currentLine)
            currentLine = word
        } else {
            currentLine = testLine
        }
    }

    if currentLine != "" {
        lines = append(lines, currentLine)
    }

    return lines
}

func (p *StatementPanel) EvaluateRules(states []bool, electric rl.Sound) {
    
    p.Items[0].Valid = states[0]


    p.Items[1].Valid = states[1] == states[2]


    countOn := 0
    for _, s := range states {
        if s {
            countOn++
        }
    }
    p.Items[2].Valid = countOn >= 4

  
    p.Items[3].Valid = states[6] || states[7]


    noThree := true
    for i := 0; i < len(states)-2; i++ {
        if states[i] && states[i+1] && states[i+2] {
            noThree = false
            break
        }
    }
    p.Items[4].Valid = noThree


    p.Items[5].Valid = states[9] == (!states[8])


    p.Items[6].Valid = countOn <= 6

  
    twoOff := false
    for i := 0; i < len(states)-1; i++ {
        if !states[i] && !states[i+1] {
            twoOff = true
            break
        }
    }
    p.Items[7].Valid = twoOff

    multiplesValid := true

    checkMultiples := []int{2, 5, 8}

    for _, idx := range checkMultiples {
        left := idx-1 >= 0 && states[idx-1]
        right := idx+1 < len(states) && states[idx+1]
        if !(left || right) {
            multiplesValid = false
            break
        }
    }

    p.Items[8].Valid = multiplesValid

    
    if states[3] {
        p.Items[9].Valid = !states[4]
    } else {
        p.Items[9].Valid = true
    }
	if p.checkAllCorrect(){
		rl.PlaySound(electric)
	}
	
}

func (p StatementPanel) checkAllCorrect() bool {
	allCorrect := true
	for _,item := range p.Items{
		if !item.Valid{
			allCorrect = false
		}
	}
	return allCorrect
}