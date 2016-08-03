package linetrace

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/nfnt/resize"
	"github.com/unixpickle/num-analysis/linalg"
)

const underlyingImageSize = 120

type PathNode struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Image struct {
	Size   int
	Values linalg.Vector
}

func imageFromGoImage(img image.Image, size int) *Image {
	scaledImg := resize.Resize(uint(size), uint(size), img, resize.Bilinear)
	res := &Image{
		Size:   size,
		Values: make(linalg.Vector, size*size),
	}
	var idx int
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			r, g, b, _ := scaledImg.At(x, y).RGBA()
			val := float64(r+g+b) / (0xffff * 3)
			res.Values[idx] = val
			idx++
		}
	}
	return res
}

func (i *Image) At(x, y int) float64 {
	return i.Values[x+y*i.Size]
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
	ctx.SetStrokeColor(color.Gray{Y: 0xff})
	ctx.BeginPath()
	ctx.MoveTo(p[0].X, p[0].Y)
	for _, item := range p[1:] {
		ctx.LineTo(item.X, item.Y)
	}
	ctx.Stroke()

	return imageFromGoImage(newImage, size)
}
