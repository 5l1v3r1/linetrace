package linetrace

import (
	"image"
	"image/color"
	"math"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

type PathNode struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Path []PathNode

// Copy creates a copy of this path.
func (p Path) Copy() Path {
	res := make(Path, len(p))
	copy(res, p)
	return res
}

func (p Path) Image(size int) *Image {
	newImage := image.NewRGBA(image.Rect(0, 0, underlyingImageSize, underlyingImageSize))
	if len(p) < 2 {
		return imageFromGoImage(newImage, size)
	}

	ctx := draw2dimg.NewGraphicContext(newImage)
	ctx.SetLineCap(draw2d.RoundCap)
	ctx.SetLineJoin(draw2d.RoundJoin)
	ctx.SetLineWidth(12)
	ctx.SetStrokeColor(color.Gray{})
	ctx.BeginPath()
	ctx.MoveTo(p[0].X, p[0].Y)
	for _, item := range p[1:] {
		ctx.LineTo(item.X, item.Y)
	}
	ctx.Stroke()

	return imageFromGoImage(newImage, size)
}

func (p Path) EvenInterpolation(pointDist float64) Path {
	res := Path{p[0]}

	var dist float64
	curPath := p.Copy()
	for len(curPath) > 1 {
		segLen := math.Sqrt(math.Pow(curPath[1].X-curPath[0].X, 2) +
			math.Pow(curPath[1].Y-curPath[0].Y, 2))
		if segLen < pointDist-dist {
			dist += segLen
			curPath = curPath[1:]
		} else {
			xDiff := curPath[1].X - curPath[0].X
			yDiff := curPath[1].Y - curPath[0].Y
			ratio := (pointDist - dist) / segLen
			dist = 0
			curPath[0].X += ratio * xDiff
			curPath[0].Y += ratio * yDiff
			res = append(res, curPath[0])
		}
	}

	return res
}
