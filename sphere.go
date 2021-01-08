package main

import (
	"math"
)

// Sphere is an sphere.
type Sphere struct {
	*HitRecord
	Center *Point3
	Radius float64
	Mat    Materialer
}

// Hit returns true if ray hits the sphere.
func (s *Sphere) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool {
	oc := r.Orig.Sub(s.Center)
	a := r.Dir.SquaredLen()
	halfB := oc.Dot(r.Dir.ToPoint3())
	c := oc.ToVec3().SquaredLen() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.Sub(s.Center).Div(s.Radius).ToVec3()
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = s.Mat
	return true
}
