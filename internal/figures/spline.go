package figures

type Spline struct {
	diameter         float32
	parallelVertices []float32
	vertices         []float32
}

func (s *Spline) AddPoint(point Point) {

	s.vertices = []float32{
		0.1, 0.1, 0,
		0.3, 0.1, 0,
		0.3, 0.3, 0,
		0.1, 0.3, 0,
		0.3, 0.3, 0,
		0.1, 0.1, 0,
	}

	return
	if len(s.vertices) == 0 {
		s.vertices = []float32{point.X, point.Y, 0}
	} else if len(s.vertices) == 3 {
		s.vertices = append(s.vertices, point.X, point.Y, 0)
	}
}

func (s *Spline) GetVertices() *[]float32 {
	return &s.vertices
}
