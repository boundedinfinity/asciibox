package comment_type

//go:generate enumer -path=main.go

type CommentType string

const (
	DoubleBackslash CommentType = "double_backslash"
	Pound           CommentType = "pound"
)
