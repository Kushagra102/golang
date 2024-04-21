package main

import "fmt"

type Player struct {
	health int
}

func (player *Player) takeDamageFromGoblin(dmg int) {
	fmt.Println("Player is taking damage from goblin")
	player.health -= dmg
}

func takeDamageFromGoblin(player *Player, dmg int) {
	fmt.Println("Player is taking damage from goblin")	
	player.health -= dmg
}

func main() {
	player := &Player {
		health: 100,
	}
	fmt.Printf("Before taking damage %+v\n", player)
	player.takeDamageFromGoblin(50)
	takeDamageFromGoblin(player, 20)
	fmt.Printf("After taking damage %+v\n", player)
}
