package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/unixpickle/linetrace"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "data.json output_rnn")
		os.Exit(1)
	}
	sampleData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read samples:", err)
		os.Exit(1)
	}
	var paths []linetrace.Path
	if err := json.Unmarshal(sampleData, &paths); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse samples:", err)
		os.Exit(1)
	}
	sampleSet := &linetrace.SampleSet{Paths: paths}

	trainingSet := sampleSet.Subset(0, sampleSet.Len()/2)
	validationSet := sampleSet.Subset(sampleSet.Len()/2, sampleSet.Len())
	log.Printf("Using %d training and %d validation samples ...", trainingSet.Len(),
		validationSet.Len())

	network := linetrace.NewNetwork()
	network.Train(trainingSet, validationSet)

	outData, err := network.Block.Serialize()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to serialize RNN:", err)
		os.Exit(1)
	}
	if err := ioutil.WriteFile(os.Args[2], outData, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to save RNN:", err)
		os.Exit(1)
	}
}
