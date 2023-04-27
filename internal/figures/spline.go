package figures

import (
	"fmt"
	"github.com/cnkei/gospline"
)

type Spline struct {
	spline  gospline.Spline
	Counter int
	x       []float64
	y       []float64
}

func (s *Spline) AddPoint(point Point) {
	s.x = append(s.x, float64(point.X))
	s.y = append(s.y, float64(point.Y))
	s.Counter++
}

func (s *Spline) GetVertices() *[]float32 {
	if len(s.x) <= 1 {
		return &[]float32{}
	}
	xStart := s.x[0]
	xEnd := s.x[len(s.x)-1]
	s.spline = gospline.NewCubicSpline(s.x, s.y)
	fmt.Println(len(s.spline.Range(s.x[0], s.x[len(s.x)-1], 200)))
	ln := (xEnd - xStart) / 100

	var rs []float32
	for i := 99; i > 0; i-- {
		rs = append(rs, float32(xStart+ln*float64(i)), float32(s.spline.At(xStart+ln*float64(i))), 0)
	}

	fmt.Println(rs[0], rs[1])
	fmt.Println(rs[len(rs)-3], rs[len(rs)-2])
	return &rs
}
