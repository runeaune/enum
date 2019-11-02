package example

//go:generate go-enum --trim "Post" --format snake
const (
	PostCreate PostType = 0
	PostRead   PostType = 2
	PostUpdate PostType = 4
	PostDelete PostType = 8
)

//go:generate go-enum --trim "Direction" --format upper
const (
	DirectionUp DirectionType = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

//go:generate go-enum --trim "Answer" --format capitalize-first --no-json --with-value
const (
	AnswerYes   YesOrNo = "Y"
	AnswerNo    YesOrNo = "N"
	AnswerMaybe YesOrNo = "M"
)
