package main

import (
	"math"
	"math/rand"
)

// Infinity is maximum float64.
const Infinity = math.MaxFloat64

// ToRadians converts degrees to radians.
func ToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// RandFloat returns a random float64 between min and max.
func RandFloat(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

// Clamp clamps x to min, max range.
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
