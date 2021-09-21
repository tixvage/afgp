package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/physics"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerComponent struct {
	Component
	pos   rl.Vector2
	body  *physics.Body
	engel *physics.Body
}

func (p *PlayerComponent) init() {
	fmt.Println("playercomponent init", p.self.name)
	p.body = physics.NewBodyRectangle(rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2), 50, 50, 1)
	p.body.FreezeOrient = true
	p.engel = physics.NewBodyRectangle(rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())), float32(rl.GetScreenWidth()), 100, 10)
	p.engel.Enabled = false
}

func (p *PlayerComponent) update(dt float64) {

	if rl.IsKeyPressed(rl.KeyR) {
		p.body.Position = rl.NewVector2(0, 0)
		p.body.Velocity = rl.NewVector2(0, 0)
		p.body.SetRotation(0)
	}

	if rl.IsKeyDown(rl.KeyRight) {
		p.body.Velocity.X = 0.5
	} else if rl.IsKeyDown(rl.KeyLeft) {
		p.body.Velocity.X = -0.5
	}
}

func (p *PlayerComponent) render() {
	rl.DrawRectangleV(p.body.Position, rl.NewVector2(50, 50), rl.Brown)
	rl.DrawRectangle(int32(rl.GetScreenWidth()/2)-640, int32(rl.GetScreenHeight())-25, int32(rl.GetScreenWidth()), 100, rl.Blue)
}

func (p *PlayerComponent) get_name() string {
	return "playercomponent"
}
