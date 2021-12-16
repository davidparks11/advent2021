package coordinate

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("{x:%d, y:%d}", p.X, p.Y)
}
