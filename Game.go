package main

import (
	"github.com/gen2brain/raylib-go/physics"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct{}

func (g *Game) run() {
	rl.InitWindow(1280, 720, "ez")
	rl.SetTargetFPS(120)
	rl.SetExitKey(rl.KeyKp0)
	physics.Init()
	current_scene = &GameScene{}

	if current_scene != nil {
		current_scene.create_game_objects()
	}

	if current_scene != nil {
		current_scene.init()
	}

	for !rl.WindowShouldClose() {
		physics.Update()
		if current_scene != nil {
			current_scene.update(float64(rl.GetFrameTime()))
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(20, 20)
		for i, body := range physics.GetBodies() {
			vertexCount := physics.GetShapeVerticesCount(i)
			for j := 0; j < vertexCount; j++ {
				// Get physics bodies shape vertices to draw lines
				// NOTE: GetShapeVertex() already calculates rotation transformations
				vertexA := body.GetShapeVertex(j)

				jj := 0
				if j+1 < vertexCount { // Get next vertex or first to close the shape
					jj = j + 1
				}

				vertexB := body.GetShapeVertex(jj)

				rl.DrawLineV(vertexA, vertexB, rl.Black) // Draw a line between two vertex positions
			}
		}
		if current_scene != nil {
			current_scene.render()
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
