package linetrace

import (
	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/sgd"
	"github.com/unixpickle/weakai/rnn"
)

const (
	sampleImageSize = 28
)

// A SampleSet turns a list of paths into an sgd.SampleSet
// full of rnn.Sequence entries.
type SampleSet struct {
	Paths []Path
}

// Len returns the number of paths.
func (s *SampleSet) Len() int {
	return len(s.Paths)
}

// Copy creates a shallow copy of s.
func (s *SampleSet) Copy() sgd.SampleSet {
	res := &SampleSet{
		Paths: make([]Path, s.Len()),
	}
	copy(res.Paths, s.Paths)
	return res
}

// Swap swaps the paths at two indices.
func (s *SampleSet) Swap(i, j int) {
	s.Paths[i], s.Paths[j] = s.Paths[j], s.Paths[i]
}

// GetSample generates an rnn.Sequence for the path at
// the given index.
func (s *SampleSet) GetSample(i int) interface{} {
	path := s.Paths[i]
	image := path.Image(sampleImageSize)

	var sample rnn.Sequence
	for i, pathItem := range path {
		input := make(linalg.Vector, len(image.Values)+2)
		copy(input, image.Values)
		input[len(image.Values)] = pathItem.X / underlyingImageSize
		input[len(image.Values)+1] = pathItem.Y / underlyingImageSize
		sample.Inputs = append(sample.Inputs, input)

		output := make(linalg.Vector, 3)
		if i+1 == len(path) {
			output[2] = 1
		} else {
			output[0] = path[i+1].X
			output[1] = path[i+1].Y
		}
		sample.Outputs = append(sample.Outputs, output)
	}

	return sample
}

// Subset generates a copy of this sample set, sliced
// between two indices.
func (s *SampleSet) Subset(start, end int) sgd.SampleSet {
	return &SampleSet{Paths: s.Paths[start:end]}
}
