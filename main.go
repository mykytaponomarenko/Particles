package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth   = 1080
	screenHeight  = 720
	particleCount = 1000
)

type Particle struct {
	x, y   float64
	vx, vy float64
	color  uint8
	t      int
}

type ParticleManager struct {
	particles []*Particle
}

func NewParticleManager() *ParticleManager {
	p := make([]*Particle, particleCount)
	for i := range p {
		p[i] = &Particle{
			x:  rand.Float64() * screenWidth,
			y:  rand.Float64() * screenHeight,
			vx: (rand.Float64() - 0.5) * 2,
			vy: (rand.Float64() - 0.5) * 2,
			t:  int(rand.Intn(3)),
		}
	}
	return &ParticleManager{
		particles: p,
	}
}

func (pm *ParticleManager) Update() {
	for _, p := range pm.particles {
		p.x += p.vx
		p.y += p.vy
		if p.x < 0 || p.x >= screenWidth {
			p.vx *= -1
		}
		if p.y < 0 || p.y >= screenHeight {
			p.vy *= -1
		}
	}
}

func (pm *ParticleManager) Draw(screen *ebiten.Image) {
	for _, p := range pm.particles {
		ebitenutil.DrawRect(screen, p.x, p.y, 4, 4, getColor(*p))
	}
}

type Game struct {
	particles *ParticleManager
}

func NewGame() *Game {
	return &Game{
		particles: NewParticleManager(),
	}
}

func (g *Game) Update() error {
	g.particles.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.particles.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func getColor(p Particle) color.RGBA {
	switch p.t {
	case 1:
		return color.RGBA{0x00, 0xff, 0x00, 0xff} // Green
	case 2:
		return color.RGBA{0xff, 0xff, 0x00, 0xff} // Yellow
	case 3:
		return color.RGBA{0x00, 0x00, 0xff, 0xff} // Blue
	default:
		return color.RGBA{0xff, 0xff, 0xff, 0xff} // White (default)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Particles")

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
