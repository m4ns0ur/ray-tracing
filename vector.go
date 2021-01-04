package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Vec3 is a vector of three items X, Y, Z.
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// Random generate a random Vec3.
func Random() *Vec3 {
	return &Vec3{
		X: rand.Float64(),
		Y: rand.Float64(),
		Z: rand.Float64(),
	}
}

// RandomInRange generate a random Vec3 in range of min and max.
func RandomInRange(min, max float64) *Vec3 {
	return &Vec3{
		X: RandFloat(min, max),
		Y: RandFloat(min, max),
		Z: RandFloat(min, max),
	}
}

// RandomInUnitSphere generate a random vector in unit sphere.
func RandomInUnitSphere() *Vec3 {
	for true {
		p := RandomInRange(-1, 1)
		if p.SquaredLen() >= 1 {
			continue
		}
		return p
	}
	return nil
}

// RandomUnitVector generates a random unit vector.
func RandomUnitVector() *Vec3 {
	return RandomInUnitSphere().UnitVector()
}

// RandomInHemisphere generates a random vector in hemisphere.
func RandomInHemisphere(normal *Vec3) *Vec3 {
	inUnitSphere := RandomInUnitSphere()
	// In the same hemisphere as the normal
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.Neg()
}

// ToPoint3 converts p to a Vec3.
func (v *Vec3) ToPoint3() *Point3 {
	return &Point3{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}

// ToColor converts v to a Color.
func (v *Vec3) ToColor() Color {
	return Color{
		R: v.X,
		G: v.Y,
		B: v.Z,
	}
}

// Neg negatives items of vector.
func (v *Vec3) Neg() *Vec3 {
	return &Vec3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

// Add adds v2 items to vector items.
func (v *Vec3) Add(v2 *Vec3) *Vec3 {
	return &Vec3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

// Mul multiplies v and v2 items.
func (v *Vec3) Mul(v2 *Vec3) *Vec3 {
	return &Vec3{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
		Z: v.Z * v2.Z,
	}
}

// Mult multiplies vector items by t.
func (v *Vec3) Mult(t float64) *Vec3 {
	return &Vec3{
		X: v.X * t,
		Y: v.Y * t,
		Z: v.Z * t,
	}
}

// Div divides vector items by t.
func (v *Vec3) Div(t float64) *Vec3 {
	return v.Mult(1 / t)
}

// SquaredLen returns the sum of all vector squared items.
func (v *Vec3) SquaredLen() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Len returns the sum of all vector items.
func (v *Vec3) Len() float64 {
	return math.Sqrt(v.SquaredLen())
}

func (v Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)
}

// Dot calculates u and v dot product.
func (v *Vec3) Dot(v2 *Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// Cross calculates u and v cross product.
func (v *Vec3) Cross(v2 *Vec3) *Vec3 {
	return &Vec3{
		X: v.Y*v2.Z - v.Z*v2.Y,
		Y: v.Z*v2.X - v.X*v2.Z,
		Z: v.X*v2.Y - v.Y*v2.X,
	}
}

// UnitVector returns unit vector of v.
func (v *Vec3) UnitVector() *Vec3 {
	return v.Div(v.Len())
}
