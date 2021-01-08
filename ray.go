package main

// Ray is a ray line in 3D dimension.
type Ray struct {
	Orig *Point3
	Dir  *Vec3
}

// At calculate ray position with parameter t.
func (r *Ray) At(t float64) *Point3 {
	return r.Orig.Addv(r.Dir.Mult(t))
}
