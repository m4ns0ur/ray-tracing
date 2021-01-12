package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(0)
	runtime.GOMAXPROCS(runtime.NumCPU())
	render()
}

func render() {
	// Image.
	const aspectRatio = 3.0 / 2.0
	const imageWidth = 1200
	const imageHeight = int(float64(imageWidth) / aspectRatio)
	const samplesPerPixel = 1 //500
	const maxDepth = 50

	// World.
	world := randomScene()

	// Camera.
	lookFrom := &Point3{13, 2, 3}
	lookAt := &Point3{0, 0, 0}
	up := &Vec3{0, 1, 0}
	aperture := 0.1
	distToFocus := 10.0
	cam := NewCamera(lookFrom, lookAt, up, 20, aspectRatio, aperture, distToFocus)

	// Render.
	name := "render-" + strconv.FormatInt(time.Now().Unix(), 10) + ".ppm"
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	if err != nil {
		log.Fatal(err)
	}

	width := float64(imageWidth - 1)
	height := float64(imageHeight - 1)

	cr := make(chan *Ray, samplesPerPixel)
	cc := make(chan *Color, samplesPerPixel)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Printf("\rScanlines remaining: %d ", j)
		for i := 0; i < imageWidth; i++ {
			c := &Color{0, 0, 0}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / width
				v := (float64(j) + rand.Float64()) / height
				go cam.Ray(u, v, cr)
				go rayColor(<-cr, world, maxDepth, cc)
				c = c.Add(<-cc)
			}
			c.Write(f, samplesPerPixel)
		}
	}

	close(cr)
	close(cc)

	fmt.Printf("\nDone.\n")
}

func rayColor(r *Ray, world *HitterList, depth int, cc chan *Color) {
	// If we've exceeded the ray bounce limit, no more light is gathered.
	if depth <= 0 {
		cc <- &Color{0, 0, 0}
		return
	}
	if ok, rec := world.Hit(r, 0.001, Infinity, new(HitRecord)); ok {
		if ok, attenuation, scattered := rec.Mat.Scatter(r, rec); ok {
			go rayColor(scattered, world, depth-1, cc)
			cc <- attenuation.Mul(<-cc)
			return
		}
		cc <- &Color{0, 0, 0}
		return
	}
	unitVecDir := r.Dir.UnitVector()
	t := 0.5 * (unitVecDir.Y + 1.0)
	cw := &Color{1.0, 1.0, 1.0}
	cb := &Color{0.5, 0.7, 1.0}
	cc <- cw.Mult(1.0 - t).Add(cb.Mult(t))
	return
}

func randomScene() *HitterList {
	world := new(HitterList)
	groundMaterial := &Lambertian{&Color{0.5, 0.5, 0.5}}
	world.Add(&Sphere{new(HitRecord), &Point3{0, -1000, 0}, 1000, groundMaterial})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := &Point3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Sub(&Point3{4, 0.2, 0}).ToVec3().Len() > 0.9 {
				if chooseMat < 0.8 {
					// Diffuse.
					albedo := RandomColor().Mul(RandomColor())
					sphereMaterial := &Lambertian{albedo}
					world.Add(&Sphere{new(HitRecord), center, 0.2, sphereMaterial})
				} else if chooseMat < 0.95 {
					// Metal.
					albedo := RandomColorInRange(0.5, 1)
					fuzz := RandFloat(0, 0.5)
					sphereMaterial := NewMetal(albedo, fuzz)
					world.Add(&Sphere{new(HitRecord), center, 0.2, sphereMaterial})
				} else {
					// Glass.
					sphereMaterial := &Dielectric{1.5}
					world.Add(&Sphere{new(HitRecord), center, 0.2, sphereMaterial})
				}
			}
		}
	}

	material1 := &Dielectric{1.5}
	world.Add(&Sphere{new(HitRecord), &Point3{0, 1, 0}, 1.0, material1})

	material2 := &Lambertian{&Color{0.4, 0.2, 0.1}}
	world.Add(&Sphere{new(HitRecord), &Point3{-4, 1, 0}, 1.0, material2})

	material3 := NewMetal(&Color{0.7, 0.6, 0.5}, 0.0)
	world.Add(&Sphere{new(HitRecord), &Point3{4, 1, 0}, 1.0, material3})

	return world
}
