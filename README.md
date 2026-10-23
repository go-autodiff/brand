# go-autodiff — brand assets

Official logos for the **go-autodiff** organization.

<img src="avatar/go-autodiff.png" width="160" alt="go-autodiff">

The mark is the partial derivative sign, the symbol of what the library computes: a bowl, and a stem that rises from its right side and arches back over it. It is written by `mark/main.go`, which computes the layout from a
few numbers; nothing is placed by hand. Its gradient (`#A5D65A` to `#3F6B08`)
is a lime green that sets go-autodiff apart from the violet of go-numerics, the emerald of go-ndarray and go-fft, the blue of go-plot, the magenta of go-bignum, the brick red of go-sparse, the orange of go-approx and the rose of go-dsp; its dark end reads
on white at 6.33:1 and its light end on a dark page at 11.37:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job runs `mark/main.go` again and fails if `svg/color/go-autodiff.svg` differs.

---
*Generated from `svg/color/go-autodiff.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
