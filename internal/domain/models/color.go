package models

type Color struct {
	id    string
	color string
}

func (c Color) GetId() string {
	return c.id
}

func (c Color) GetName() string {
	return c.color
}

func (c Color) SetColor(color string) {
	c.color = color
}
