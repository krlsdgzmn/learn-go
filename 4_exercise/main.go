package main

import "fmt"

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
	fmt.Printf("%s is added to your inventory\n", i.Name)
}

func (p *Player) DropItem(i Item) {
	for index, value := range p.Inventory {
		if value.Name == i.Name {
			p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)
			fmt.Printf("You dropped %s\n", i.Name)
			return
		}
	}
	fmt.Printf("Item %s not found\n", i.Name)
}

func (p *Player) UseItem(i Item) {
	for index, value := range p.Inventory {
		if value.Name == i.Name {
			p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)
			fmt.Printf("Item %s is used\n", i.Name)
			return
		}
	}
	fmt.Printf("Item %s not found\n", i.Name)
}

func main() {
	p := Player{Name: "John Karlos"}
	item1 := Item{Name: "Attack Potion"}
	item2 := Item{Name: "Health Potion"}

	p.PickUpItem(item1)
	p.PickUpItem(item2)
	fmt.Printf("%s: %v\n\n", p.Name, p.Inventory)

	p.DropItem(item1)
	fmt.Printf("%s: %v\n\n", p.Name, p.Inventory)

	p.UseItem(item2)
	fmt.Printf("%s: %v\n\n", p.Name, p.Inventory)

	p.UseItem(item1)
	fmt.Printf("%s: %v\n\n", p.Name, p.Inventory)
}
