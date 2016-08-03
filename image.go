package linetrace

import (
	"image"

	"github.com/nfnt/resize"
	"github.com/unixpickle/num-analysis/linalg"
)

const underlyingImageSize = 120

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
