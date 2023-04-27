package figures

import "math"

type CircleFragment struct {
	Point           Point
	Rotation        int
	Radius          float32
	currentVertices []float32
}

func (cf *CircleFragment) GetVertices() *[]float32 {
	if len(cf.currentVertices) == 0 {
		cf.RegenerateVertices()
	}

	return &cf.currentVertices
}

func (cf *CircleFragment) RegenerateVertices() {
	cf.currentVertices = generateQuarterCircleVertices(cf.Point, cf.Radius, 100, cf.Rotation)
}

func generateQuarterCircleVertices(point Point, radius float32, segments int, offset int) []float32 {
	var vertices []float32

	// центр окружности
	center := []float32{0.0 + point.X, 0.0 + point.Y, 0.0}

	// угол шага между сегментами
	angleStep := math.Pi / (2.0 * float64(segments))

	// генерация вершин для четверти круга
	for i := offset; i <= segments+offset; i++ {
		angle := float64(i) * angleStep

		x := -radius * float32(math.Cos(angle))
		y := radius * float32(math.Sin(angle))

		vertices = append(vertices, center[0], center[1], center[2])
		vertices = append(vertices, x+point.X, y+point.Y, 0.0)

		if i != segments {
			nextAngle := float64(i+1) * angleStep
			nextX := -radius * float32(math.Cos(nextAngle))
			nextY := radius * float32(math.Sin(nextAngle))

			vertices = append(vertices, nextX+point.X, nextY+point.Y, 0.0)
		}
	}

	return vertices
}
