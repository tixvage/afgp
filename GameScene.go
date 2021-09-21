package main

import (
	"github.com/gen2brain/raylib-go/physics"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/ldtkgo"
	"github.com/solarlune/ldtkgo/raylibrenderer"
)

type GameScene struct {
	Scene
	ldtkProject *ldtkgo.Project
	renderer    *raylibrenderer.RaylibRenderer
}

func (g *GameScene) create_game_objects() {
	entity1 := NewEntity("deneme")
	entity1.add_component(&PlayerComponent{})
	g.add_entity(&entity1)
	entity2 := NewEntity("inventory")
	entity2.add_component(&Inventory{})
	entity2.add_component(&DragDropInventory{})
	g.add_entity(&entity2)
}

func (g *GameScene) get_base_scene() *Scene {
	return &g.Scene
}
func (g *GameScene) init() {
	g.Scene.init()
	g.ldtkProject, _ = ldtkgo.Open("assets/demo.ldtk")
	g.renderer = raylibrenderer.NewRaylibRenderer(raylibrenderer.NewDiskLoader("assets/"))
	g.renderer.Render(g.ldtkProject.Levels[0])
	bruh := g.ldtkProject.Levels[0].Layers[0].Entities[0]
	bruhbody := physics.NewBodyRectangle(rl.NewVector2(float32(bruh.Position[0]), float32(bruh.Position[1])), float32(bruh.Width), float32(bruh.Height), 10)
	bruhbody.Enabled = false
}

func (g *GameScene) render() {
	for _, layer := range g.renderer.RenderedLayers {
		rl.DrawTextureEx(layer.Image.Texture, rl.Vector2{X: 0, Y: 0}, 0, 1, rl.White)

	}
	g.Scene.render()

}
