package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(sprite.main, nil)
}
