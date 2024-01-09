package nova

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X, Y float64
}

func ZeroVector() Vector2D {
	return Vector2D{X: 0, Y: 0}
}

func NewVector2D(x, y float64) Vector2D {
	fmt.Println("New Vector X:", x, "Y:", y)
	return Vector2D{X: x, Y: y}
}

func (v Vector2D) Set(x, y float64) Vector2D {
	v.X = x
	v.Y = y
	return v
}

func (v Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v.X + v2.X, v.Y + v2.Y}
}
func (v Vector2D) Sub(v2 Vector2D) Vector2D {
	return Vector2D{v.X - v2.X, v.Y - v2.Y}
}
func (v Vector2D) Scale(ratio float64) Vector2D {

	return Vector2D{v.X * ratio, v.Y * ratio}
}
func (v Vector2D) Normalize() Vector2D {
	mag := 1 / float64(math.Sqrt(v.X*v.X+v.Y*v.Y))
	return Vector2D{mag * v.X, mag * v.Y}
}

type Position struct {
}
