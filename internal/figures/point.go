package figures

type Point struct {
	X float32
	Y float32
}

func GeneratePointByWindow(height int, width int, cursorX float64, cursorY float64) Point {
	hHalf := float64(height) / 2
	wHalf := float64(width) / 2

	x := (cursorX - wHalf) / wHalf
	y := -1 * (cursorY - hHalf) / hHalf
	return Point{X: float32(x), Y: float32(y)}
}
