package main

import (
	"bytes"
	_ "embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var playerPng []byte
var playerImg *ebiten.Image

//go:embed assets/sanic.png
var imgPng []byte
var img *ebiten.Image

//go:embed assets/sanic.wav
var sanicWav []byte
var audioContext *audio.Context

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(imgPng))
	if err != nil {
		log.Fatal(err)
	}

	if audioContext == nil {
		audioContext = audio.NewContext(44100)
	}

	sanicSound, err := wav.DecodeWithoutResampling(bytes.NewReader(sanicWav))
	if err != nil {
		log.Fatal(err)
	}
	ssPlayer, err := audioContext.NewPlayer(sanicSound)
	if err != nil {
		log.Fatal(err)
	}
	ssPlayer.Play()
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 746, 696
}

func main() {
	ebiten.SetWindowSize(746, 696)
	ebiten.SetWindowTitle("sanic.exe")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
