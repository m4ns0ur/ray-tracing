package main

import "math"

// Camera is a camera.
type Camera struct {
	Origin          *Point3
	LowerLeftCorner *Point3
	Horizontal      *Vec3
	Vertical        *Vec3
	U               *Vec3
	V               *Vec3
	W               *Vec3
	LensRadius      float64
}

// NewCamera creats a new camera. vFov is vertical field-of-view in degrees.
func NewCamera(lookFrom, lookAt *Point3, up *Vec3, fov, aspectRatio, aperture, focusDist float64) *Camera {
	theta := ToRadians(fov)
	h := math.Tan(theta / 2)
	viewPortHeight := 2.0 * h
	viewPortWidth := aspectRatio * viewPortHeight
	w := lookFrom.Sub(lookAt).ToVec3().UnitVector()
	u := up.Cross(w).UnitVector()
	v := w.Cross(u)
	origin := lookFrom
	horizontal := u.Mult(viewPortWidth * focusDist)
	vertical := v.Mult(viewPortHeight * focusDist)
	lowerLeftCorner := origin.Subv(horizontal.Div(2)).Subv(vertical.Div(2)).Subv(w.Mult(focusDist))
	lensRadius := aperture / 2
	return &Camera{
		Origin:          origin,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      horizontal,
		Vertical:        vertical,
		U:               u,
		V:               v,
		W:               w,
		LensRadius:      lensRadius,
	}
}

// Ray generates camera ray by s, t.
func (c *Camera) Ray(s, t float64, cr chan<- *Ray) {
	rd := NewRandomInUnitDiskVec3().Mult(c.LensRadius)
	offset := c.U.Mult(rd.X).Add(c.V.Mult(rd.Y))
	cr <- &Ray{
		Orig: c.Origin.Addv(offset),
		Dir:  c.LowerLeftCorner.Addv(c.Horizontal.Mult(s)).Addv(c.Vertical.Mult(t)).Sub(c.Origin).Sub(offset.ToPoint3()).ToVec3(),
	}
}
