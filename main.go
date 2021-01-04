package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(256, 256)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	// dc.SavePNG("out.png")

	render()
}

func render() {
	// Image.
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samplesPerPixel := 100
	maxDepth := 50

	// World.
	world := &HitterList{}
	world.Add(NewSphere(&Point3{0, 0, -1}, 0.5))
	world.Add(NewSphere(&Point3{0, -100.5, -1}, 100))

	// Camera.
	cam := NewCamera()

	// Render.
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", j)
		for i := 0; i < imageWidth; i++ {
			c := Color{0, 0, 0}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)
				r := cam.Ray(u, v)
				c = c.Add(rayColor(r, world, maxDepth))
			}
			c.Write(samplesPerPixel)
		}
	}

	fmt.Fprintf(os.Stderr, "\nDone.\n")
}

func rayColor(r *Ray, world *HitterList, depth int) Color {
	// If we've exceeded the ray bounce limit, no more light is gathered.
	if depth <= 0 {
		return Color{0, 0, 0}
	}
	if ok, rec := world.Hit(r, 0.001, Infinity, &HitRecord{}); ok {
		target := rec.P.Addv(RandomInHemisphere(rec.Normal))
		return rayColor(&Ray{rec.P, target.Sub(rec.P).ToVec3()}, world, depth-1).Mult(0.5)
	}
	unitVecDir := r.Dir.UnitVector()
	t := 0.5 * (unitVecDir.Y + 1.0)
	return Color{1.0, 1.0, 1.0}.Mult(1.0 - t).Add(Color{0.5, 0.7, 1.0}.Mult(t))
}

func hitSphere(center *Point3, radius float64, r *Ray) float64 {
	oc := r.Orig.Sub(center)
	a := r.Dir.SquaredLen()
	halfB := oc.Dot(r.Dir.ToPoint3())
	c := oc.ToVec3().SquaredLen() - radius*radius
	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return -1.0
	}
	return (-halfB - math.Sqrt(discriminant)) / a
}
