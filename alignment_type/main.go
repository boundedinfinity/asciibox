package alignment_type

//go:generate enumer -path=main.go

type Alignment string

const (
	Left   Alignment = "left"
	Middle Alignment = "middle"
	Right  Alignment = "right"
)
