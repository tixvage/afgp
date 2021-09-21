package main

var current_scene IScene = nil

type Scene struct {
	entities []*Entity
}

func (s *Scene) add_entity(entity *Entity) {
	s.entities = append(s.entities, entity)
}

func (s *Scene) get_entity(name string) *Entity {
	for i := 0; i < len(s.entities); i++ {
		if s.entities[i].name == name {
			return s.entities[i]
		}
	}

	return nil
}

func (s *Scene) init() {
	for i := 0; i < len(s.entities); i++ {
		s.entities[i].init()
	}
}

func (s *Scene) update(dt float64) {
	for i := 0; i < len(s.entities); i++ {
		s.entities[i].update(dt)
	}
}

func (s *Scene) render() {
	for i := 0; i < len(s.entities); i++ {
		s.entities[i].render()
	}
}

type IScene interface {
	get_base_scene() *Scene
	create_game_objects()
	init()
	update(float64)
	render()
}

func change_scene(new_scene IScene) {
	current_scene = new_scene
	current_scene.init()
}
