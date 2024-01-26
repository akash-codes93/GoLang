package main

import "fmt"

type Position struct {
	posX float64
	posY float64
}

type SpecialPosition struct {
	Position
}

func (p *Position) move(x, y float64) {
	p.posX += x
	p.posY += y
}

func (p *SpecialPosition) specialMove(x, y float64) {
	p.posX += x * x
	p.posY += y * y
}

type Player struct {
	*Position
}

func newPlayer() Player {
	player := Player{
		Position: &Position{},
	}

	return player
}

type SpecialPlayer struct {
	*SpecialPosition
}

func newSpecialPlayer() SpecialPlayer {
	player := SpecialPlayer{
		SpecialPosition: &SpecialPosition{},
	}

	return player
}

func main() {
	normalPlayer := newPlayer()

	normalPlayer.move(1.1, 2.1)

	fmt.Println(normalPlayer.posX, normalPlayer.posY)

	specialPLayer := newSpecialPlayer()
	specialPLayer.move(1.1, 2.1)
	specialPLayer.specialMove(1.1, 2.1)

	fmt.Println(specialPLayer.posX, specialPLayer.posY)

}
