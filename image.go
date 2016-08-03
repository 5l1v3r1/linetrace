package linetrace

import (
	"image"
	"image/color"
	"math"

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
			_, _, _, a := scaledImg.At(x, y).RGBA()
			val := float64(a) / 0xffff
			res.Values[idx] = val
			idx++
		}
	}
	return res
}

func (i *Image) At(x, y int) float64 {
	return i.Values[x+y*i.Size]
}

func (i *Image) Correlation(img1 *Image) float64 {
	var res float64
	var iMag float64
	var img1Mag float64
	for j, x := range img1.Values {
		res += x * i.Values[j]
		iMag += i.Values[j] * i.Values[j]
		img1Mag += x * x
	}
	return res / math.Sqrt(iMag*img1Mag)
}

func (i *Image) GoImage() image.Image {
	output := image.NewGray(image.Rect(0, 0, i.Size, i.Size))
	for y := 0; y < i.Size; y++ {
		for x := 0; x < i.Size; x++ {
			val := uint8(i.Values[x+y*i.Size]*0xff + 0.5)
			output.SetGray(x, y, color.Gray{Y: val})
		}
	}
	return output
}
