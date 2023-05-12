package black

import "github.com/tkanata/go-training-neko/model"

type blackNeko struct {
	Color string
}

func NewBlackNeko() model.Neko {
	return &blackNeko{}
}

func (n *blackNeko) GetColor() string {
	return n.Color
}

func (n *blackNeko) SetColor(color string) {
	n.Color = color
}
