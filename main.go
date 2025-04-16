package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const (
		screenWidth  = 800
		screenHeight = 450
	)
	rl.InitWindow(screenWidth, screenHeight, "Axe Game Go")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Game Over!", 400, 200, 20, rl.Red)

		rl.EndDrawing()

	}
}
