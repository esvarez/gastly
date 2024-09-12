package blur

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	gophersImage *ebiten.Image
)

type Game struct{}

func (g Game) Update() error {
	//TODO implement me
	panic("implement me")
}

func (g Game) Draw(screen *ebiten.Image) {
	//TODO implement me
	panic("implement me")
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//TODO implement me
	panic("implement me")
}

func New() {
	img, _, err := image.Decode(bytes.NewReader(images.FiveYears_jpg))
	if err != nil {
		log.Fatal("error decoding image", err)
	}

	gophersImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Blur (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
