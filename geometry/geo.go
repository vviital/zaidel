package geometry

import (
	"math"
)

// Coordinate struct
type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Coordinates type
type Coordinates []Coordinate

// FindByX returns points by X
func (pts Coordinates) FindByX(x float64) (int, Coordinate) {
	for i, p := range pts {
		if p.X == x {
			return i, p
		}
	}
	return -1, Coordinate{}
}

// GetYs returns slice of Y geometry.Coordinates for points
func (pts Coordinates) GetYs() (ys []float64) {
	if pts == nil {
		return
	}

	for _, p := range pts {
		ys = append(ys, p.Y)
	}

	return
}

// GetMaxPointByY returns max point by Y value
func (pts Coordinates) GetMaxPointByY() Coordinate {
	var index = 0

	for i, point := range pts {
		if point.Y > pts[index].Y {
			index = i
		}
	}

	return pts[index]
}

// TriangleArea func
func TriangleArea(a, b, c Coordinate) float64 {
	return math.Abs((a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2)
}

// CalculateArea returns area of
func CalculateArea(l, r int, points, backgroundData Coordinates) (area float64) {
	for ai := l; ai < r; ai++ {
		if ai == l {
			area += TriangleArea(points[ai], points[ai+1], backgroundData[ai+1])
		} else if ai+1 == r {
			area += TriangleArea(points[ai-1], points[ai], backgroundData[ai-1])
		} else {
			area += TriangleArea(points[ai], backgroundData[ai], backgroundData[ai+1])
			area += TriangleArea(points[ai], points[ai+1], backgroundData[ai+1])
		}
	}
	return
}

// CalculateSpectrumArea returns spectrum area
func CalculateSpectrumArea(points, background Coordinates) (area float64) {
	var calculate bool
	var leftIndex int

	for i := 1; i < len(background); i++ {
		index := i - 1
		heightBefore := points[index].Y - background[index].Y
		heightNow := points[i].Y - background[i].Y
		if heightBefore >= 0 && !calculate {
			leftIndex = i
			calculate = true
		}
		if heightNow <= 0 && calculate {
			area += CalculateArea(leftIndex, i, points, background)
			calculate = false
		}
	}
	return
}
