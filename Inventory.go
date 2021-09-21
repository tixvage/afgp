package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sword struct {
	texture rl.Texture2D
	pos     rl.Vector2
}

func (s *Sword) load_texture() {
	s.texture = rl.LoadTexture("assets/sword.png")
}
func (s *Sword) ez() {

}
func (s *Sword) get_texture() *rl.Texture2D {
	return &s.texture
}

func (s *Sword) get_pos() rl.Vector2 {
	return s.pos
}

func (s *Sword) set_pos(a rl.Vector2) {
	s.pos = a
}

type Item interface {
	load_texture()
	get_texture() *rl.Texture2D
	get_pos() rl.Vector2
	set_pos(rl.Vector2)
}

type Inventory struct {
	Component
	texture       rl.Texture2D
	current_place rl.Texture2D
	place         int32
	items         []Item
}

func (p *Inventory) init_items() {
	p.items = append(p.items, &Sword{})

}

func (p *Inventory) load_item_textures() {
	for i := 0; i < len(p.items); i++ {
		p.items[i].load_texture()
	}
}

func (p *Inventory) render_items() {
	for i := 0; i < len(p.items); i++ {
		rl.DrawTextureEx(*p.items[i].get_texture(), rl.Vector2{X: 270 + float32((i+1)*80), Y: 595}, 0, 5, rl.White)
	}
}

func (p *Inventory) init() {
	fmt.Println("inventory init", p.self.name)
	p.place = 1
	p.init_items()
	p.load_item_textures()
	p.texture = rl.LoadTexture("assets/inventory.png")
	p.current_place = rl.LoadTexture("assets/bruh.png")
}

func (p *Inventory) update(dt float64) {
	if rl.IsKeyDown(rl.KeyOne) {
		p.place = 1
	} else if rl.IsKeyDown(rl.KeyTwo) {
		p.place = 2
	} else if rl.IsKeyDown(rl.KeyThree) {
		p.place = 3
	} else if rl.IsKeyDown(rl.KeyFour) {
		p.place = 4
	} else if rl.IsKeyDown(rl.KeyFive) {
		p.place = 5
	} else if rl.IsKeyDown(rl.KeySix) {
		p.place = 6
	} else if rl.IsKeyDown(rl.KeySeven) {
		p.place = 7
	}

	if p.place != 1 || rl.GetMouseWheelMove() != -1 {
		if p.place != 7 || rl.GetMouseWheelMove() != 1 {
			p.place += rl.GetMouseWheelMove()
		}
	}
}

func (p *Inventory) render() {
	rl.DrawTextureEx(p.texture, rl.Vector2{X: 360, Y: 600}, 0, 5, rl.White)
	rl.DrawTextureEx(p.current_place, rl.Vector2{X: 270 + float32(p.place*80), Y: 595}, 0, 5, rl.White)
	p.render_items()
}

func (p *Inventory) get_name() string {
	return "inventory"
}

type DragDropInventory struct {
	Component
	texture               rl.Texture2D
	selected_item_texture rl.Texture2D
	items                 [9][3]Item
	selected_item         Item
	selected_item_pos     rl.Vector2
	last_pos              rl.Vector2
	active                bool
}

func (d *DragDropInventory) draw_inventory() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			rl.DrawTextureEx(d.texture, rl.Vector2{X: 230 + float32(j*95), Y: 150 + float32(i*95)}, 0, 5, rl.White)
			if d.items[j][i] != nil {
				rl.DrawTextureEx(*d.items[j][i].get_texture(), rl.Vector2{X: 230 + (d.items[j][i].get_pos().X * 95), Y: 150 + (d.items[j][i].get_pos().Y * 95)}, 0, 5, rl.White)
			}
		}
	}
}

func (d *DragDropInventory) add_item(item Item) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			if d.items[j][i] == nil {
				fmt.Println(i, j)
				item.set_pos(rl.Vector2{X: float32(j), Y: float32(i)})
				d.items[j][i] = item
				return
			}
		}
	}
}

func (d *DragDropInventory) load_items() {
	d.add_item(&Sword{})
	d.add_item(&Sword{})
	d.add_item(&Sword{})
	d.add_item(&Sword{})

	for i := 0; i < 9; i++ {
		for j := 0; j < 3; j++ {
			if d.items[i][j] != nil {
				d.items[i][j].load_texture()
			}
		}
	}
}

func (d *DragDropInventory) init() {
	d.texture = rl.LoadTexture("assets/dragdrop.png")
	d.selected_item_texture = rl.LoadTexture("assets/bruh.png")
	d.load_items()
	d.selected_item = nil
	d.active = false
}

func (d *DragDropInventory) update(dt float64) {
	if !d.active {
		if rl.IsKeyPressed(rl.KeyE) {
			d.active = true
			return
		}
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		temp := &Sword{}
		temp.load_texture()
		d.add_item(temp)
	}
	if d.active {
		if rl.IsKeyPressed(rl.KeyE) {
			d.active = false
			return
		}
		//dont look here
		for i := 0; i < 3; i++ {
			for j := 0; j < 9; j++ {
				if d.items[j][i] != nil {
					if d.selected_item == nil {
						if ismouseonitem(d.items[j][i]) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								d.selected_item = d.items[j][i]
								d.selected_item_pos.X = 225 + (float32(j) * 95)
								d.selected_item_pos.Y = 145 + (float32(i) * 95)
								d.last_pos = rl.Vector2{X: float32(int((rl.GetMouseX() - 230) / 95)), Y: float32(int((rl.GetMouseY() - 150) / 95))}
								return
							}
						}
					}
				}
			}
		}
		//move items
		if d.selected_item != nil {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				if d.items[(int((rl.GetMouseX() - 230) / 95))][int((rl.GetMouseY()-150)/95)] == nil {
					d.items[int(d.last_pos.X)][int(d.last_pos.Y)] = nil
					d.selected_item.set_pos(rl.Vector2{X: float32(int((rl.GetMouseX() - 230) / 95)), Y: float32(int((rl.GetMouseY() - 150) / 95))})
					d.items[int(d.selected_item.get_pos().X)][int(d.selected_item.get_pos().Y)] = d.selected_item
					d.selected_item = nil
				}
			}
		}
	}
}

func (d *DragDropInventory) render() {
	if d.active {
		d.draw_inventory()
		if d.selected_item != nil {
			rl.DrawTextureEx(d.selected_item_texture, d.selected_item_pos, 0, 5, rl.White)
		}
	}

	//rl.DrawText(fmt.Sprint(rl.GetMousePosition()), 100, 20, 20, rl.Black)
	//rl.DrawText("press e to see inventory\nand press enter to add item", 20, 50, 20, rl.Black)
}

func (d *DragDropInventory) get_name() string {
	return "dragdrop"
}

func ismouseonitem(item Item) bool {
	return rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(230+(item.get_pos().X*95), 150+(item.get_pos().Y*95), float32(item.get_texture().Width*5), float32(item.get_texture().Height*5)))
}
