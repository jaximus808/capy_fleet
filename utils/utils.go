package utils

import "math"

type Vec2 struct {
	x float64
	y float64
}

func Vector2(_x float64, _y float64) *Vec2 {
	return &Vec2{
		x: _x,
		y: _y,
	}
}

func (v *Vec2) X() float64 {
	return v.x
}

func (v *Vec2) Y() float64 {
	return v.y
}

func (v *Vec2) Set(other_v *Vec2) {
	v.x = other_v.x
	v.y = other_v.y
}

func (v *Vec2) Add(other_v *Vec2) *Vec2 {
	v.x += other_v.x
	v.y += other_v.y
	return v
}

// adds and creates a copy, does not modify org
func (v *Vec2) Addc(other_v *Vec2) *Vec2 {
	return &Vec2{
		x: v.x + other_v.x,
		y: v.y + other_v.y,
	}
}

func (v *Vec2) Sub(other_v *Vec2) *Vec2 {
	v.x -= other_v.x
	v.y -= other_v.y
	return v
}
func (v *Vec2) Subc(other_v *Vec2) *Vec2 {
	return &Vec2{
		x: v.x - other_v.x,
		y: v.y - other_v.y,
	}
}
func (v *Vec2) Scale(scaler float64) *Vec2 {
	v.x *= scaler
	v.y *= scaler
	return v
}

// makes a copy vector of v scaled
func (v *Vec2) Scalec(scaler float64) *Vec2 {
	return &Vec2{
		x: v.x * scaler,
		y: v.y * scaler,
	}
}

func (v *Vec2) Dot(other_v *Vec2) float64 {
	return other_v.x*v.x + other_v.y*v.y
}
func (v *Vec2) Mag() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *Vec2) Norm() *Vec2 {
	mag := v.Mag()
	t_v := &Vec2{
		x: v.x / mag,
		y: v.y / mag,
	}
	if math.IsNaN(t_v.x) {
		t_v.x = 0
	}
	if math.IsNaN(t_v.y) {
		t_v.y = 0
	}
	return t_v
}

// creates a normalized vector that points to another vec
func (v *Vec2) VecTowards(other_v *Vec2) *Vec2 {
	non_norm := Vector2(
		other_v.x-v.x,
		other_v.y-v.y,
	)
	return non_norm.Norm()
}
func (v *Vec2) Dist(other_v *Vec2) float64 {
	_x := v.x - other_v.x
	_y := v.y - other_v.y
	return math.Sqrt(_x*_x + _y*_y)
}

func (v *Vec2) Equals(other_v *Vec2) bool {
	return v.x == other_v.x && v.y == other_v.y
}
