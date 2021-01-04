package main

// Camera is a camera.
type Camera struct {
	Origin          *Point3
	Horizontal      *Vec3
	Vertical        *Vec3
	LowerLeftCorner *Point3
}

// NewCamera creats a new camera.
func NewCamera() *Camera {
	aspectRatio := 16.0 / 9.0
	viewPortHeight := 2.0
	viewPortWidth := aspectRatio * viewPortHeight
	focalLen := 1.0

	origin := &Point3{0, 0, 0}
	horizontal := &Vec3{viewPortWidth, 0, 0}
	vertical := &Vec3{0, viewPortHeight, 0}
	lowerLeftCorner := origin.Subv(horizontal.Div(2)).Subv(vertical.Div(2)).Subv(&Vec3{0, 0, focalLen})

	return &Camera{
		Origin:          origin,
		Horizontal:      horizontal,
		Vertical:        vertical,
		LowerLeftCorner: lowerLeftCorner,
	}
}

// Ray generates camera ray by u, v.
func (c *Camera) Ray(u, v float64) *Ray {
	return &Ray{
		Orig: c.Origin,
		Dir:  c.LowerLeftCorner.Addv(c.Horizontal.Mult(u)).Addv(c.Vertical.Mult(v)).Sub(c.Origin).ToVec3(),
	}
}
