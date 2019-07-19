package example

//go:generate go-enum --trim "Post" --format snake
const (
	PostCreate PostType = 0
	PostRead   PostType = 2
	PostUpdate PostType = 4
	PostDelete PostType = 8
)
