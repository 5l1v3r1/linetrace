package linetrace

import (
	"github.com/unixpickle/autofunc"
	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/weakai/neuralnet"
)

// CostFunc is a cost function useful for training
// a line tracing network.
//
// The vectors must have a length divisible by three,
// since each triple is (x, y, shouldStop), where x
// and y are raw numbers for regression and shouldStop
// is a flag to be fed through a sigmoid and then
// through cross-entropy loss.
type CostFunc struct{}

func (_ CostFunc) Cost(x linalg.Vector, a autofunc.Result) autofunc.Result {
	if len(x)%3 != 0 {
		panic("vectors must have lengths divisible by 3")
	}
	return autofunc.Pool(a, func(actual autofunc.Result) autofunc.Result {
		expectedStop := make(linalg.Vector, 0, len(x)/3)
		expectedCoords := make(linalg.Vector, 0, 2*len(x)/3)
		actualStop := make([]autofunc.Result, 0, len(x)/3)
		actualCoords := make([]autofunc.Result, 0, 2*len(x)/3)
		for i, xVal := range x {
			aVal := autofunc.Slice(actual, i, i+1)
			if i%3 == 2 {
				expectedStop = append(expectedStop, xVal)
				actualStop = append(actualStop, aVal)
			} else {
				expectedCoords = append(expectedCoords, xVal)
				actualCoords = append(actualCoords, aVal)
			}
		}
		msc := neuralnet.MeanSquaredCost{}
		sce := neuralnet.SigmoidCECost{}
		return autofunc.Add(msc.Cost(expectedCoords, autofunc.Concat(actualCoords...)),
			sce.Cost(expectedStop, autofunc.Concat(actualStop...)))
	})
}

func (_ CostFunc) CostR(v autofunc.RVector, x linalg.Vector,
	a autofunc.RResult) autofunc.RResult {
	if len(x)%3 != 0 {
		panic("vectors must have lengths divisible by 3")
	}
	return autofunc.PoolR(a, func(actual autofunc.RResult) autofunc.RResult {
		expectedStop := make(linalg.Vector, 0, len(x)/3)
		expectedCoords := make(linalg.Vector, 0, 2*len(x)/3)
		actualStop := make([]autofunc.RResult, 0, len(x)/3)
		actualCoords := make([]autofunc.RResult, 0, 2*len(x)/3)
		for i, xVal := range x {
			aVal := autofunc.SliceR(actual, i, i+1)
			if i%3 == 2 {
				expectedStop = append(expectedStop, xVal)
				actualStop = append(actualStop, aVal)
			} else {
				expectedCoords = append(expectedCoords, xVal)
				actualCoords = append(actualCoords, aVal)
			}
		}
		msc := neuralnet.MeanSquaredCost{}
		sce := neuralnet.SigmoidCECost{}
		return autofunc.AddR(msc.CostR(v, expectedCoords, autofunc.ConcatR(actualCoords...)),
			sce.CostR(v, expectedStop, autofunc.ConcatR(actualStop...)))
	})
}
