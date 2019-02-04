package point

import "strconv"

type Point struct {
	X int
	Y int
}

func PrintPoint (p Point) string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}
