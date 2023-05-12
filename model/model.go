package model

// Nekoを抽象化したインタフェース
type Neko interface {
	GetColor() string
	SetColor(color string)
}
