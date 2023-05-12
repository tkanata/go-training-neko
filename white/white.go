package white

import "github.com/tkanata/go-training-neko/model"

type whiteNeko struct {
	Color string
}

// NewWhiteNekoは新しい白猫を生成する
func NewWhiteNeko() model.Neko {
	return &whiteNeko{}
}

func (n *whiteNeko) GetColor() string {
	return n.Color
}

func (n *whiteNeko) SetColor(color string) {
	n.Color = color
}
