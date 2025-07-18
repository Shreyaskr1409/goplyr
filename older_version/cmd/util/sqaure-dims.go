package util

import "math"

// SquareDims returns the largest (w,h) such that
// h/w = 20/9 and w ≤ maxW, h ≤ maxH.
func SquareDims(maxW, maxH int) (w, h uint) {
	// First try full-width
	hFromW := float64(maxW) * 20.0 / 9.0
	if hFromW <= float64(maxH) {
		return uint(maxW), uint(math.Round(hFromW))
	}
	// Otherwise limit by height
	wFromH := float64(maxH) * 9.0 / 20.0
	return uint(math.Round(wFromH)), uint(maxH)
}
