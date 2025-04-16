package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// window dimensions
	const (
		screenWidth  = 800
		screenHeight = 450
	)
	rl.InitWindow(screenWidth, screenHeight, "Axe Game Go")
	defer rl.CloseWindow()

	// circle coordinates
	circleX := 200
	circleY := 200
	circleRadius := 25
	// circle edges
	lCircleX := circleX - circleRadius
	rCircleX := circleX + circleRadius
	uCircleY := circleY - circleRadius
	bCircleY := circleY + circleRadius

	// axe coordinates
	axeX := 400
	axeY := 0
	axeLength := 50
	// axe edges
	lAxeX := axeX
	rAxeX := axeX + axeLength
	uAxeY := axeY
	bAxeY := axeY + axeLength

	direction := 10

	collisionWithAxe := (bAxeY >= uCircleY) && (uAxeY <= bCircleY) && (lAxeX <= rCircleX) && (rAxeX >= lCircleX)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if collisionWithAxe {
			rl.DrawText("Game Over!", 400, 200, 20, rl.Red)
		} else {
			// game logic begin

			// update the edges
			lCircleX = circleX - circleRadius
			rCircleX = circleX + circleRadius
			uCircleY = circleY - circleRadius
			bCircleY = circleY + circleRadius

			lAxeX = axeX
			rAxeX = axeX + axeLength
			uAxeY = axeY
			bAxeY = axeY + axeLength

			collisionWithAxe = (bAxeY >= uCircleY) && (uAxeY <= bCircleY) && (lAxeX <= rCircleX) && (rAxeX >= lCircleX)

			rl.DrawCircle(int32(circleX), int32(circleY), float32(circleRadius), rl.Blue)
			rl.DrawRectangle(int32(axeX), int32(axeY), int32(axeLength), int32(axeLength), rl.Red)

			// move the axe
			axeY += direction
			if axeY > screenHeight || axeY < 0 {
				direction = -direction
			}

			if rl.IsKeyDown(rl.KeyD) && circleX < screenWidth {
				circleX += 10
			}
			if rl.IsKeyDown(rl.KeyA) && circleX > 25 {
				circleX -= 10
			}

		}

		rl.EndDrawing()

	}
}
