package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
)

// Color is an RGB color.
type Color struct {
	R float64
	G float64
	B float64
}

// RandomColor creates a color with random RGB.
func RandomColor() *Color {
	return &Color{rand.Float64(), rand.Float64(), rand.Float64()}
}

// RandomColorInRange creates a color with random RGB in min and max range.
func RandomColorInRange(min, max float64) *Color {
	return &Color{RandFloat(min, max), RandFloat(min, max), RandFloat(min, max)}
}

// ToVec3 converts c to a Vec3.
func (c *Color) ToVec3() *Vec3 {
	return &Vec3{
		X: c.R,
		Y: c.G,
		Z: c.B,
	}
}

// Write writes color RGB values in file f.
func (c *Color) Write(f *os.File, samplesPerPixel int) {
	r := c.R
	g := c.G
	b := c.B

	// Divide the color by the number of samples and gamma-correct for gamma=2.0.
	scale := 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(scale * r)
	g = math.Sqrt(scale * g)
	b = math.Sqrt(scale * b)

	// Write the translated [0,255] value of each color component.
	_, err := fmt.Fprintf(f, "%d %d %d\n", int(256*Clamp(r, 0.0, 0.999)), int(256*Clamp(g, 0.0, 0.999)), int(256*Clamp(b, 0.0, 0.999)))
	if err != nil {
		log.Fatal(err)
	}
}

// Mul multiplies c and c2 items, and returns a new Color.
func (c *Color) Mul(c2 *Color) *Color {
	return &Color{
		R: c.R * c2.R,
		G: c.G * c2.G,
		B: c.B * c2.B,
	}
}

// Mult multiplies c items by t, and returns a new Color.
func (c *Color) Mult(t float64) *Color {
	return &Color{
		R: c.R * t,
		G: c.G * t,
		B: c.B * t,
	}
}

// Add adds c and c2 items, and returns a new Color.
func (c *Color) Add(c2 *Color) *Color {
	return &Color{
		R: c.R + c2.R,
		G: c.G + c2.G,
		B: c.B + c2.B,
	}
}
