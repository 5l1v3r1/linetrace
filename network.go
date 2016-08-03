package linetrace

import (
	"log"
	"math"

	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/sgd"
	"github.com/unixpickle/weakai/neuralnet"
	"github.com/unixpickle/weakai/rnn"
)

const (
	stoppingThreshold = 2
	networkStateSize  = 128
	networkHiddenSize = 128

	networkStepSize  = 1e-3
	networkBatchSize = 20
)

// A Network uses a recurrent neural network to trace
// the path of a line drawing.
type Network struct {
	Block rnn.StackedBlock
}

// NewNetwork creates a randomly initialized network.
func NewNetwork() *Network {
	outLayer := neuralnet.Network{
		&neuralnet.DenseLayer{
			InputCount:  networkStateSize,
			OutputCount: networkHiddenSize,
		},
		neuralnet.Sigmoid{},
		&neuralnet.DenseLayer{
			InputCount:  networkHiddenSize,
			OutputCount: 3,
		},
	}
	outLayer.Randomize()
	return &Network{
		Block: rnn.StackedBlock{
			rnn.NewLSTM(sampleImageSize*sampleImageSize+2, networkStateSize),
			rnn.NewLSTM(networkStateSize, networkStateSize),
			rnn.NewNetworkBlock(outLayer, 0),
		},
	}
}

// Apply uses the network to produce a path for the
// given image, where the image's bounds are defined
// by the corners (0,0) and (1,1).
func (n *Network) Apply(img *Image) Path {
	var res Path

	runner := rnn.Runner{Block: n.Block}
	curVec := make(linalg.Vector, len(img.Values)+2)
	copy(curVec, img.Values)
	for {
		output := runner.StepTime(curVec)
		if math.Exp(output[2]) > stoppingThreshold {
			break
		}
		curVec = make(linalg.Vector, len(curVec))
		copy(curVec, img.Values)
		curVec[len(img.Values)] = output[0]
		curVec[len(img.Values)+1] = output[1]
		res = append(res, PathNode{
			X: output[0],
			Y: output[1],
		})
	}

	return res
}

// Train trains the network on a sample set.
func (n *Network) Train(training, validation sgd.SampleSet) {
	var costFunc CostFunc
	gradienter := &sgd.Adam{
		Gradienter: &rnn.BPTT{
			Learner:  n.Block,
			CostFunc: costFunc,
		},
	}
	var i int
	sgd.SGDInteractive(gradienter, training, networkStepSize, networkBatchSize, func() bool {
		runner := rnn.Runner{
			Block: n.Block,
		}
		log.Printf("Epoch %d: cost=%f cross=%f", i,
			runner.TotalCost(networkBatchSize, training, costFunc),
			runner.TotalCost(networkBatchSize, validation, costFunc))
		i++
		return true
	})
}
