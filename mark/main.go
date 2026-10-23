// Command mark writes the go-autodiff mark, svg/color/go-autodiff.svg:
//
//	cd mark && go run . > ../svg/color/go-autodiff.svg
//
// go-autodiff is automatic differentiation, and the mark is the partial
// derivative sign, the symbol of the thing it computes: a bowl, and a stem
// that rises from the bowl's right side and curls back over it.
//
// Nothing is placed by hand. The bowl is an ellipse given by its centre and
// two radii; the stem leaves the bowl at its rightmost point, which the
// program computes, and the curl is two cubic segments whose points are
// offsets from the top of the stem. Change the radii and the sign redraws
// itself, still joined.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#A5D65A" // the gradient's light stop
	dark  = "#3F6B08" // its dark stop

	stroke = 13.0 // the family's line width
	centre = 128.0

	rx = 34.0 // horizontal radius of the bowl, between stroke centres
	ry = 40.0 // vertical radius of the bowl

	rise = 88.0 // how far the arch rises above the centre of the bowl
	drop = 24.0 // how far the free end of the arch hangs below its top
	arch = 54.0 // how far the free end of the arch lies left of the stem
)

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	// The bowl is centred horizontally and sits so that the whole sign, bowl
	// and arch, is centred vertically on the square.
	cx := centre
	cy := centre + (rise-ry)/2

	right := cx + rx // the stem leaves the bowl here
	top := cy - rise // the top of the arch

	// The arch: up the stem, over the top (both tangents horizontal at the
	// summit, so the two segments join smoothly), then down a little.
	summit := right - arch/2
	bowl := fmt.Sprintf(`<ellipse cx="%s" cy="%s" rx="%s" ry="%s"/>`, num(cx), num(cy), num(rx), num(ry))
	stem := fmt.Sprintf(`<path d="M%s %sC%s %s %s %s %s %sC%s %s %s %s %s %s"/>`,
		num(right), num(cy), // from the bowl's rightmost point
		num(right), num(top+30),
		num(right-6), num(top),
		num(summit), num(top), // the summit
		num(summit-16), num(top),
		num(right-arch), num(top+10),
		num(right-arch), num(top+drop)) // the free end

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-autodiff">
  <defs><linearGradient id="adgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#adgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
    %s
  </g>
</svg>
`, light, dark, num(stroke), bowl, stem)
}
