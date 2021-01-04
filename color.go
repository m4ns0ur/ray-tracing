package main

import (
	"fmt"
	"math"
)

// Color is an RGB color.
type Color struct {
	R float64
	G float64
	B float64
}

// ToVec3 converts c to a Vec3.
func (c Color) ToVec3() *Vec3 {
	return &Vec3{
		X: c.R,
		Y: c.G,
		Z: c.B,
	}
}

// Write writes color RGB values in stdout.
func (c Color) Write(samplesPerPixel int) {
	r := c.R
	g := c.G
	b := c.B

	// Divide the color by the number of samples and gamma-correct for gamma=2.0.
	scale := 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(scale * r)
	g = math.Sqrt(scale * g)
	b = math.Sqrt(scale * b)

	// Write the translated [0,255] value of each color component.
	fmt.Printf("%d %d %d\n", int(256*Clamp(r, 0.0, 0.999)), int(256*Clamp(g, 0.0, 0.999)), int(256*Clamp(b, 0.0, 0.999)))
}

// Mult multiplies c items by t, and returns a new Color.
func (c Color) Mult(t float64) Color {
	return Color{
		R: c.R * t,
		G: c.G * t,
		B: c.B * t,
	}
}

// Add adds c1 and c2 items, and returns a new Color.
func (c Color) Add(c2 Color) Color {
	return Color{
		R: c.R + c2.R,
		G: c.G + c2.G,
		B: c.B + c2.B,
	}
}
