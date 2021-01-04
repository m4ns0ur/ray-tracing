package main

// Point3 is a point in 3D dimension.
type Point3 struct {
	X float64
	Y float64
	Z float64
}

// ToVec3 converts p to a Vec3.
func (p *Point3) ToVec3() *Vec3 {
	return &Vec3{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
	}
}

// Mult multiplies v items by t, and returns a new Point3.
func (p *Point3) Mult(t float64) *Point3 {
	return &Point3{
		X: p.X * t,
		Y: p.Y * t,
		Z: p.Z * t,
	}
}

// Addv adds p and v items, and returns a new Point3.
func (p *Point3) Addv(v *Vec3) *Point3 {
	return &Point3{
		X: p.X + v.X,
		Y: p.Y + v.Y,
		Z: p.Z + v.Z,
	}
}

// Sub subtracts p1 and p2 items, and returns a new Point3.
func (p *Point3) Sub(p2 *Point3) *Point3 {
	return &Point3{
		X: p.X - p2.X,
		Y: p.Y - p2.Y,
		Z: p.Z - p2.Z,
	}
}

// Subv subtracts p and v items, and returns a new Point3.
func (p *Point3) Subv(v *Vec3) *Point3 {
	return &Point3{
		X: p.X - v.X,
		Y: p.Y - v.Y,
		Z: p.Z - v.Z,
	}
}

// Dot calculates p1 and p2 dot product.
func (p *Point3) Dot(p2 *Point3) float64 {
	return p.X*p2.X + p.Y*p2.Y + p.Z*p2.Z
}

// Dotv calculates p and v dot product.
func (p *Point3) Dotv(v *Vec3) float64 {
	return p.X*v.X + p.Y*v.Y + p.Z*v.Z
}

// Div divides point items by t.
func (p *Point3) Div(t float64) *Point3 {
	return p.Mult(1 / t)
}
