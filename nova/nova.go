package nova

import (
	"fmt"
	"time"
)

func Test() {
	// Create Time
	t := CreateTime()
	t.Start()

	a := new(Object2D)
	target := Vector2D{100, 100}
	for {
		time.Sleep(time.Millisecond * 1200)
		fmt.Printf("%s [Position] %v\r", t.String(), a.Position)

		a.MoveTo(a.Position.Add(target.Sub(a.Position).Normalize().Scale(2)))
	}

}
