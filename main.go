package main

import (
	"fmt"

	"github.com/tkanata/go-training-neko/black"
	"github.com/tkanata/go-training-neko/model"
	"github.com/tkanata/go-training-neko/white"
)

type Neko interface {
	GetColor() string
	SetColor(color string)
}

func main() {
	neko1 := white.NewWhiteNeko()
	neko2 := black.NewBlackNeko()

	generateNeko(neko1, "しろ")
	generateNeko(neko2, "くろ")

	fmt.Println(neko1.GetColor())
	fmt.Println(neko2.GetColor())
}

// whiteNeko構造体もblackNeko構造体もNekoインタフェースを満たしているので、この関数はどちらの構造体も受け取ることができる。
func generateNeko(n model.Neko, color string) model.Neko {
	n.SetColor(color)
	return n
}
