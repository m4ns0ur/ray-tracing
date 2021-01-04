package main

// HitRecord has info about hit point.
type HitRecord struct {
	P         *Point3
	Normal    *Vec3
	T         float64
	FrontFace bool
}

// Hitter is any object which ray tracing hit it.
type Hitter interface {
	Hit(r *Ray, tMin, tMax float64, rec *HitRecord) bool
}

// SetFaceNormal sets the face normal.
func (h *HitRecord) SetFaceNormal(r *Ray, outwardNormal *Vec3) {
	h.FrontFace = r.Dir.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Neg()
	}
}

// HitterList is a list of hitters.
type HitterList struct {
	*HitRecord
	Hitters []Hitter
}

// Add adds h to hittter list.
func (hl *HitterList) Add(h Hitter) {
	hl.Hitters = append(hl.Hitters, h)
}

// Clear empty the hittter list.
func (hl *HitterList) Clear() {
	hl.Hitters = nil
}

// Hit returns true if ray trace hits the hitter list.
func (hl *HitterList) Hit(r *Ray, tMin, tMax float64, rec *HitRecord) (bool, *HitRecord) {
	tempRec := &HitRecord{}
	hitAnything := false
	closestSoFar := tMax

	for _, hitter := range hl.Hitters {
		if hitter.Hit(r, tMin, closestSoFar, tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			rec = tempRec
		}
	}

	return hitAnything, rec
}
