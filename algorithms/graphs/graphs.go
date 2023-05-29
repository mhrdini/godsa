package graphs

import "fmt"

type Color int

const (
	White = Color(iota)
	Gray
	Black
)

func (c Color) String() string {
	switch c {
	case White:
		return "white"
	case Gray:
		return "gray"
	case Black:
		return "black"
	default:
		return "error"
	}
}

type Vertex struct {
	Color
	Value  int
	Dist   float64
	Parent *Vertex
}

func (v *Vertex) String() string {
	switch v.Parent {
	case nil:
		return fmt.Sprintf("[%v %v %v -1]", v.Value, v.Color, v.Dist)
	default:
		return fmt.Sprintf("[%v %v %v %v]", v.Value, v.Color, v.Dist, v.Parent.Value)
	}
}
