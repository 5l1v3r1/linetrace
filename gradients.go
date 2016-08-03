package linetrace

// A CostFunc computes how well paths approximate an
// image.
type CostFunc interface {
	Cost(p Path) float64
}

// CorrelationCost is a cost function which turns paths
// into images and measures their correlation with a
// pre-determined image.
type CorrelationCost struct {
	Image *Image
}

// Cost computes the correlation between the path's
// image and c.Image.
func (c *CorrelationCost) Cost(p Path) float64 {
	return c.Image.Correlation(p.Image(c.Image.Size))
}

// Gradient computes a Path representing the rate of
// change of some cost function with respect to each
// element in the given Path p.
func Gradient(p Path, c CostFunc) Path {
	tempPath := make(Path, len(p))
	copy(tempPath, p)
	res := make(Path, len(p))
	for i := range p {
		tempPath[i].X += 1
		right := c.Cost(tempPath)
		tempPath[i].X -= 2
		left := c.Cost(tempPath)
		tempPath[i].X += 1
		res[i].X = (right - left) / 2

		tempPath[i].Y += 1
		right = c.Cost(tempPath)
		tempPath[i].Y -= 2
		left = c.Cost(tempPath)
		tempPath[i].Y += 1
		res[i].Y = (right - left) / 2
	}
	return res
}

// AddGradient adds a gradient path to another path.
// For gradient descent, stepSize should be negative.
func AddGradient(p, grad Path, stepSize float64) {
	for i, g := range grad {
		p[i].X += g.X * stepSize
		p[i].Y += g.Y * stepSize
	}
}
