package main

import (
	"math"
	"math/rand"
)

// Materialer is any material that might scatter the light.
type Materialer interface {
	Scatter(rIn *Ray, rec *HitRecord) (ok bool, attenuation *Color, scattered *Ray)
}

// Lambertian is a lambertian reflectance material (like a solid object).
type Lambertian struct {
	Albedo *Color
}

// Scatter returns always true since lambertian reflectance material scatters the light.
func (l *Lambertian) Scatter(rIn *Ray, rec *HitRecord) (ok bool, attenuation *Color, scattered *Ray) {
	scatterDirection := rec.Normal.Add(RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}

	return true, l.Albedo, &Ray{rec.P, scatterDirection}
}

// Metal is a metalic material.
type Metal struct {
	Albedo *Color
	Fuzz   float64
}

// NewMetal creates a new metal.
func NewMetal(c *Color, f float64) *Metal {
	if f > 1 {
		f = 1
	}
	return &Metal{
		Albedo: c,
		Fuzz:   f,
	}
}

// Scatter returns true if metal material scatters the light.
func (m *Metal) Scatter(rIn *Ray, rec *HitRecord) (ok bool, attenuation *Color, scattered *Ray) {
	reflected := rIn.Dir.UnitVector().Reflect(rec.Normal)
	scattered = &Ray{rec.P, reflected.Add(RandomInUnitSphere().Mult(m.Fuzz))}
	return scattered.Dir.Dot(rec.Normal) > 0, m.Albedo, scattered
}

// Dielectric is a material with refraction.
type Dielectric struct {
	// Index of refraction.
	Ir float64
}

// Scatter returns always true since dielectric scatters the light.
func (d *Dielectric) Scatter(rIn *Ray, rec *HitRecord) (ok bool, attenuation *Color, scattered *Ray) {
	var refractionRatio float64
	if rec.FrontFace {
		refractionRatio = 1.0 / d.Ir
	} else {
		refractionRatio = d.Ir
	}
	unitDir := rIn.Dir.UnitVector()
	cosTheta := math.Min(unitDir.Neg().Dot(rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1.0

	var dir *Vec3
	if cannotRefract || reflectance(cosTheta, refractionRatio) > rand.Float64() {
		dir = unitDir.Reflect(rec.Normal)
	} else {
		dir = unitDir.Refract(rec.Normal, refractionRatio)
	}

	return true, &Color{1.0, 1.0, 1.0}, &Ray{rec.P, dir}
}

func reflectance(cosine, refIdx float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
