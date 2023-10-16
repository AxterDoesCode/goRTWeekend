package Vec3

import (
	"fmt"
	"math"
)

type Vec3 struct {
	e [3]float64
}

func NewVec3() Vec3 {
	return Vec3{e: [3]float64{0, 0, 0}}
}

func NewVec3WithValues(X, Y, Z float64) Vec3 {
	return Vec3{e: [3]float64{X, Y, Z}}
}

func (v *Vec3) X() float64 {
	return v.e[0]
}

func (v *Vec3) Y() float64 {
	return v.e[1]
}

func (v *Vec3) Z() float64 {
	return v.e[2]
}

func (v *Vec3) Neg() Vec3 {
	return NewVec3WithValues(-v.X(), -v.Y(), -v.Z())
}

func (v *Vec3) At(i int) float64 {
	return v.e[i]
}

func (v *Vec3) Add(vec Vec3) Vec3 {
	v.e[0] += vec.e[0]
	v.e[1] += vec.e[1]
	v.e[2] += vec.e[2]
	return *v
}

func (v *Vec3) Multiply(t float64) Vec3 {
	v.e[0] *= t
	v.e[1] *= t
	v.e[2] *= t
	return *v
}

func (v *Vec3) Divide(t float64) Vec3 {
	v.Multiply(1 / t)
	return *v
}

func (v *Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) String() string {
	return fmt.Sprintf("%v %v %v", v.e[0], v.e[1], v.e[2])
}

func Dot(u, v Vec3) float64 {
	return u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2]
}

func Cross(u, v Vec3) Vec3 {
	return Vec3{[3]float64{u.e[1]*v.e[2] - u.e[2]*v.e[1],
		u.e[2]*v.e[0] - u.e[0]*v.e[2],
		u.e[0]*v.e[1] - u.e[1]*v.e[0]}}
}

func UnitVector(v Vec3) Vec3 {
	return v.Divide(v.Length())
}

type Point3 = Vec3
