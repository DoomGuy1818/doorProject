package models

type Category struct {
	id    string
	title string
}

func (c Category) GetId() string {
	return c.id
}

func (c Category) GetName() string {
	return c.title
}

func (c Category) SetCategory(category string) {
	c.title = category
}
