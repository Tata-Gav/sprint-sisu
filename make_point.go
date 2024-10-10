package sprint

type Point struct {
	X float32
	Y float32
	text string
}

func MokePoint(x, y float32, text string) Point {
	return Point{
		X: x,
		Y: y,
		Text: text,
	}
}