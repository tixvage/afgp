package mainzort

type IComponent interface {
	init()
	update(float64)
	render()
	get_name() string
	get_entity() *Entity
	set_entity(*Entity)
}

type Component struct {
	self *Entity
}

func (c *Component) get_entity() *Entity {
	return c.self
}

func (c *Component) set_entity(e *Entity) {
	c.self = e
}

type Entity struct {
	comps []IComponent
	name  string
}

func (e *Entity) add_component(comp IComponent) {
	comp.set_entity(e)
	e.comps = append(e.comps, comp)
}

func (e *Entity) get_component(name string) IComponent {
	for i := 0; i < len(e.comps); i++ {
		if e.comps[i].get_name() == name {
			return e.comps[i]
		}
	}
	return nil
}

func (e *Entity) init() {
	for i := 0; i < len(e.comps); i++ {
		e.comps[i].init()
	}
}

func (e *Entity) update(dt float64) {
	for i := 0; i < len(e.comps); i++ {
		e.comps[i].update(dt)
	}
}

func (e *Entity) render() {
	for i := 0; i < len(e.comps); i++ {
		e.comps[i].render()
	}
}

func NewEntity(name string) Entity {
	temp := Entity{}
	temp.name = name
	return temp
}
