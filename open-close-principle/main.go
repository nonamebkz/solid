package main

import (
	"fmt"
)

const (
	spear = "spear"
)

type Weapon string

type King struct {
	weapon Weapon
}

func (k King) useWeapon() {
	fmt.Println("king use weapon " + k.weapon)
}

func main() {
	k := &King{spear}
	k.useWeapon()

}
