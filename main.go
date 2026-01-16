package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1024, 1024, "Stavros Sparks' Power Plant Escape")
	rl.SetExitKey(0)
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	background := rl.LoadTexture("Sprites/Mainroom.png")
	defer rl.UnloadTexture(background)
	switchSound := rl.LoadSound("Sounds/79237__joedeshon__electrical_switch_01.wav")
	//https://freesound.org/people/joedeshon/sounds/79237/
	rl.SetSoundVolume(switchSound, 0.25)
	electric := rl.LoadSound("Sounds/82745__sikoses__stm1_some_bass (1).mp3")
	//https://freesound.org/people/Sikoses/sounds/82745/
	rl.SetSoundVolume(electric, 0.25)
	boxMove := rl.LoadSound("Sounds/BoxMove.wav")
	defer rl.UnloadSound(boxMove)
	//https://freesound.org/people/mushusito/sounds/472378/
	rl.SetSoundVolume(boxMove, 2)
	elevatorChime := rl.LoadSound("Sounds/253982__kwahmah_02__lift-doors-opening-01.wav")
	defer rl.UnloadSound(elevatorChime)
	//https://freesound.org/people/kwahmah_02/sounds/253982/
	rl.SetSoundVolume(elevatorChime, 0.25)

	batteryEmpty := rl.LoadTexture("Sprites/batteryEmpty.png")
	batteryFull := rl.LoadTexture("Sprites/batteryFull.png")
	backgroundMusic := rl.LoadMusicStream("Sounds/639799__m-murray__the-stranger-theme-stranger-things-inspired.wav")
	
	rl.PlayMusicStream(backgroundMusic)
	rl.SetMusicVolume(backgroundMusic, 0.15)
	crateTex := rl.LoadTexture("Sprites/crate.png") 
	defer rl.UnloadTexture(crateTex)
	sokobanBackground := rl.LoadTexture("Sprites/roomLeftBackground.png")
	defer rl.UnloadTexture(sokobanBackground)
	rightRoomBackground := rl.LoadTexture("Sprites/ComputerRoom.png")
	defer rl.UnloadTexture(rightRoomBackground)
	winBackground := rl.LoadTexture("Sprites/GameWin.png")
	defer rl.UnloadTexture(winBackground)
	pauseBackground := rl.LoadTexture("Sprites/Paused.png")
	defer rl.UnloadTexture(pauseBackground)
	menuBackground := rl.LoadTexture("Sprites/Menu.png")

	player := NewPlayer()
	panelDoor := NewDoor(700, 455, 100, 150)
	leftDoor := NewDoor(20, 570, 100, 150)
	rightDoor := NewDoor(870, 570, 100, 150)
	computerDoor := NewDoor(470, 180, 100, 150)
	elevator := NewDoor(400, 540, 220, 180)
	game := NewGame()
	game.Scene = SceneMenu
	progress := NewProgress()
	box := NewBreakerBox()
	panel := NewStatementPanel()
	batteryHUD := NewBatteryHUD(batteryEmpty, batteryFull)
	textBox := TextBox{}
	sokoban := NewSokoban(crateTex, boxMove)

	restartButton := rl.NewRectangle(200, 840, 320, 60)
	exitButton := rl.NewRectangle(550, 840, 320, 60)

	pauseContinue := rl.NewRectangle(275, 455, 475, 70)
	pauseRestart := rl.NewRectangle(275, 570, 475, 70)
	pauseExit := rl.NewRectangle(275, 685, 475, 70)

	menuStart := rl.NewRectangle(90, 830, 370, 130)
	menuExit := rl.NewRectangle(565, 830, 370, 130)
	showTip := true

	var finalGame *FinalGame
	var previousScene Scene

	defer rl.UnloadTexture(batteryEmpty)
	defer rl.UnloadTexture(batteryFull)

	defer rl.UnloadTexture(player.Anims[IdleFront].Texture)
	defer rl.UnloadTexture(player.Anims[WalkFront].Texture)
	defer rl.UnloadTexture(player.Anims[IdleBack].Texture)
	defer rl.UnloadTexture(player.Anims[WalkBack].Texture)
	defer rl.UnloadTexture(player.Anims[WalkToward].Texture)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(backgroundMusic)
		if game.Scene != SceneLeftRoom && game.Scene != SceneComputerRoom {
			player.Update()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsKeyPressed(rl.KeyEscape) && game.Scene != ScenePause && game.Scene != SceneWin {
			previousScene = game.Scene
			game.Scene = ScenePause
		}

		switch game.Scene {

		case SceneMain:
			if showTip{
				textBox.Visible = true
				textBox.Text = "Search the power plant for clues to get the power back on\n Press F to dismiss"
			}else{
				textBox.Visible = false
			}
			
			if rl.IsKeyPressed(rl.KeyF){
				showTip = false
			}
			rl.DrawTexture(background, 0, 0, rl.White)
			batteryHUD.Draw(progress)
			player.Draw()



			

			if !progress.BreakerCompleted && panelDoor.checkProximity(player) {
				textBox.Visible = true
				textBox.Text = "Hmmmm... something must have tripped a breaker. Press F to inspect the breaker box."
			}
			if !progress.BreakerCompleted && panelDoor.checkProximity(player) && rl.IsKeyPressed(rl.KeyF) {
				game.Scene = SceneBreaker
			}
			if !progress.LeftCompleted && leftDoor.checkProximity(player) {
				textBox.Visible = true
				textBox.Text = "This door leads deeper into the plant... Press F to enter."
			}
			if !progress.LeftCompleted && leftDoor.checkProximity(player) && rl.IsKeyPressed(rl.KeyF) {
				game.Scene = SceneLeftRoom
			}
			if !progress.ComputerRoomCompleted && rightDoor.checkProximity(player) {
				textBox.Visible = true
				textBox.Text = "This door leads deeper into the plant... Press F to enter."
			}
			if rightDoor.checkProximity(player) && rl.IsKeyPressed(rl.KeyF) {
				game.Scene = SceneComputerRoom
			}

			if elevator.checkProximity(player) && progress.Batteries() < 3 {
				textBox.Visible = true
				textBox.Text = "The elevator has no power... It looks like it needs 3 batteries."
			}
			if elevator.checkProximity(player) && progress.Batteries() == 3 {
				textBox.Visible = true
				textBox.Text = "All batteries are connected.\nPress F to power the elevator and escape."
			}
			if elevator.checkProximity(player) && progress.Batteries() == 3 && rl.IsKeyPressed(rl.KeyF) {
				game.Scene = SceneWin
				rl.PlaySound(elevatorChime)
			}

		case SceneBreaker:
			textBox.Visible = false
			textBox.Text = ""

			box.DrawPanelBackground()

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				mouse := rl.GetMousePosition()
				box.HandleClicks(mouse, switchSound)

				states := box.GetStates()
				panel.EvaluateRules(states, electric)
			}

			box.DrawAllBreakers()
			panel.Draw()

			if panel.checkAllCorrect() {
				progress.BreakerCompleted = true
				game.Scene = SceneMain
			}

			if rl.IsKeyPressed(rl.KeyEscape) {
				game.Scene = SceneMain
			}
		case SceneLeftRoom:

			rl.DrawTexture(sokobanBackground, 0, 0, rl.White)

			sokoban.Update()
			player.UpdateFacing()
			textBox.Visible = true
			textBox.Text = "Move the crate onto the power switch to restore power to the room. \nPress R to reset the crate's position if needed."

			sokoban.Draw()
			px, py := sokoban.GetPlayerWorldPos()
			player.Draw()

			player.X = px + float32(sokoban.tileSize)/2 - player.Width()/2
			player.Y = py + float32(sokoban.tileSize)/2 - player.Height()/2

			if sokoban.IsSolved() {
				progress.LeftCompleted = true
				game.Scene = SceneMain
				rl.PlaySound(electric)
			}

			if rl.IsKeyPressed(rl.KeyEscape) {
				game.Scene = SceneMain
			}
		case SceneComputerRoom:
			textBox.Visible = true
			textBox.Text = "Maybe that computer can help get the power back on..."
			if !progress.ComputerRoomCompleted && computerDoor.checkProximityComputer(player) {
				textBox.Visible = true
				textBox.Text = "Press F to use the computer"
				if rl.IsKeyPressed(rl.KeyF) {
					finalGame = NewFinalGame()
					game.Scene = SceneFinalGame
				}
			}

			rl.DrawTexture(rightRoomBackground, 0, 0, rl.White)

			player.UpdateComp()
			player.Draw()

			if rl.IsKeyPressed(rl.KeyEscape) {
				game.Scene = SceneMain
			}

		case SceneFinalGame:
			textBox.Text = "Use the numbers and operators provided to make 48\n You may only use each number once and must use all numbers."
			rl.ClearBackground(rl.Black)

			finalGame.Update()
			finalGame.Draw()

		
			if finalGame.IsSolved {
				progress.ComputerRoomCompleted = true
			}

			if rl.IsKeyPressed(rl.KeyEscape) {
				game.Scene = SceneMain
			}
			if finalGame.JustSolved {

				rl.PlaySound(electric) 
				progress.ComputerRoomCompleted = true
				progress.RightCompleted = true
				finalGame.JustSolved = false
				game.Scene = SceneMain
			}
		case SceneWin:
			textBox.Visible = false
			rl.DrawTexture(winBackground, 0, 0, rl.White)


			drawWinButton(restartButton, "RESTART")
			drawWinButton(exitButton, "EXIT")


			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				mouse := rl.GetMousePosition()

				if rl.CheckCollisionPointRec(mouse, restartButton) {
					progress = NewProgress()
					finalGame = nil
					game.Scene = SceneMain
				}

				if rl.CheckCollisionPointRec(mouse, exitButton) {
					rl.CloseWindow()
				}
			}
		case ScenePause:

			rl.DrawTexture(pauseBackground, 0, 0, rl.White)

			mouse := rl.GetMousePosition()

			continueHover := rl.CheckCollisionPointRec(mouse, pauseContinue)
			restartHover := rl.CheckCollisionPointRec(mouse, pauseRestart)
			exitHover := rl.CheckCollisionPointRec(mouse, pauseExit)

			drawPauseButton(pauseContinue, continueHover)
			drawPauseButton(pauseRestart, restartHover)
			drawPauseButton(pauseExit, exitHover)

			drawButtonLabel(pauseContinue, "CONTINUE")
			drawButtonLabel(pauseRestart, "RESTART")
			drawButtonLabel(pauseExit, "EXIT")

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				if continueHover {
					game.Scene = previousScene
				}

				if restartHover {
					progress = NewProgress()
					finalGame = nil
					game.Scene = SceneMain
				}

				if exitHover {
					rl.CloseWindow()
				}
			}
		case SceneMenu:

			rl.DrawTexture(menuBackground, 0, 0, rl.White)

			drawMenuButton(menuStart, "START")
			drawMenuButton(menuExit, "EXIT")

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				mouse := rl.GetMousePosition()


				if rl.CheckCollisionPointRec(mouse, menuStart) {
					game.Scene = SceneMain
				}

				if rl.CheckCollisionPointRec(mouse, exitButton) {
					rl.CloseWindow()
				}
			}
		default:
			game.Scene = SceneMain
		}
		textBox.Draw()

		rl.EndDrawing()

	}
}

func drawWinButton(rect rl.Rectangle, label string) {

	mouse := rl.GetMousePosition()
	hovered := rl.CheckCollisionPointRec(mouse, rect)

	color := rl.NewColor(12, 20, 37, 255)
	if hovered {
		color = rl.NewColor(40, 52, 78, 255)

	}

	rl.DrawRectangleRec(rect, color)
	rl.DrawRectangleLinesEx(rect, 3, rl.Black)

	textSize := 26
	textWidth := rl.MeasureText(label, int32(textSize))

	rl.DrawText(
		label,
		int32(rect.X+rect.Width/2-float32(textWidth)/2),
		int32(rect.Y+rect.Height/2-float32(textSize)/2),
		int32(textSize),
		rl.White,
	)
}

func drawPauseButton(rect rl.Rectangle, hovered bool) {

	color := rl.NewColor(22, 28, 42, 255)
	if hovered {
		color = rl.NewColor(40, 52, 78, 255)

	}

	rl.DrawRectangleRec(rect, color)
	rl.DrawRectangleLinesEx(rect, 2, rl.Black)
}

func drawButtonLabel(rect rl.Rectangle, text string) {

	fontSize := int32(26)
	textWidth := rl.MeasureText(text, fontSize)

	x := int32(rect.X + rect.Width/2 - float32(textWidth)/2)
	y := int32(rect.Y + rect.Height/2 - float32(fontSize)/2)

	rl.DrawText(text, x, y, fontSize, rl.White)
}

func drawMenuButton(rect rl.Rectangle, label string) {

	mouse := rl.GetMousePosition()
	hovered := rl.CheckCollisionPointRec(mouse, rect)

	baseColor := rl.NewColor(50, 60, 50, 255)
	hoverColor := rl.NewColor(110, 104, 68, 255)

	color := baseColor
	if hovered {
		color = hoverColor
	}

	rl.DrawRectangleRec(rect, color)

	rl.DrawRectangleLinesEx(rect, 3, rl.Black)


	textSize := 28
	textWidth := rl.MeasureText(label, int32(textSize))

	rl.DrawText(
		label,
		int32(rect.X+rect.Width/2-float32(textWidth)/2),
		int32(rect.Y+rect.Height/2-float32(textSize)/2),
		int32(textSize),
		rl.White,
	)
}
