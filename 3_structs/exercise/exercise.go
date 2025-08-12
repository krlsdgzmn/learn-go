package exercise

import (
	"fmt"
)

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

func (p *Player) PickUpItem(i Item) {
	p.Inventory = append(p.Inventory, i)
	fmt.Printf("Item %+v is added to player %s\n", i, p)
}
