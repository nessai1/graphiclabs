package figures

type Pixels struct {
	vertices []float32
	Diameter float32
}

func (p *Pixels) AddPixel(point Point) {
	p.vertices =
		append(p.vertices,
			point.X-p.Diameter/2, point.Y+p.Diameter/2, 0,
			point.X-p.Diameter/2, point.Y-p.Diameter/2, 0,
			point.X+p.Diameter/2, point.Y+p.Diameter/2, 0,
			point.X+p.Diameter/2, point.Y+p.Diameter/2, 0,
			point.X+p.Diameter/2, point.Y-p.Diameter/2, 0,
			point.X-p.Diameter/2, point.Y-p.Diameter/2, 0,
		)
}

func (p *Pixels) GetVertices() *[]float32 {
	return &p.vertices
}
