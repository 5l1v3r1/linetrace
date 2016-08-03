package linetrace

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

type PathNode struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Path []PathNode

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
