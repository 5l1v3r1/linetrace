package linetrace

import "math"

const minGainRatio = 0.3

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

// Cost computes the inverse correlation between the
// path's image and c.Image.
// This is 0 when the paths are perfectly correlated.
func (c *CorrelationCost) Cost(p Path) float64 {
	if len(p) < 2 {
		return 1
	}
	return 1 - c.Image.Correlation(p.Image(c.Image.Size))
}

// SearchPath searches for the path which best satisfies
// a cost function.
// The spacing specifies how many pixels along each axis
// the path's points may be.
// stride defines how many pixels between starting coords
// the search places.
// minX, minY, maxX, maxY define the bounding box in
// which the path may begin, where the mins are inclusive
// and the maxes are exclusive.
func SearchPath(cost CostFunc, spacing, stride, minX, minY, maxX, maxY int) Path {
	var bestPath Path
	var bestCost float64
	for y := minY; y <= maxY; y += stride {
		for x := minX; x <= maxX; x += stride {
			p := searchOnePath(cost, stride, x, y)
			c := cost.Cost(p)
			if bestPath == nil || c < bestCost {
				bestCost = c
				bestPath = p
			}
		}
	}
	return bestPath
}

func searchOnePath(cost CostFunc, stride, curX, curY int) Path {
	path := Path{PathNode{float64(curX), float64(curY)}}
	curCost := cost.Cost(path)
	initCost := curCost

	for {
		var bestContX, bestContY int
		bestContCost := math.Inf(1)
		for y := -stride; y <= stride; y++ {
			for x := -stride; x <= stride; x++ {
				newPath := append(path, PathNode{float64(curX + x), float64(curY + y)})
				c := cost.Cost(newPath)
				if c < bestContCost {
					bestContCost = c
					bestContX = x
					bestContY = y
				}
			}
		}
		var avgGains float64
		if len(path) > 0 {
			avgGains = (initCost - curCost) / float64(len(path))
		}
		gains := curCost - bestContCost
		if gains <= avgGains*minGainRatio {
			break
		}
		curCost = bestContCost
		curX += bestContX
		curY += bestContY
		path = append(path, PathNode{float64(curX), float64(curY)})
	}

	return path
}
