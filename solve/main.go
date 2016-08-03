package main

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/unixpickle/linetrace"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "line_images original.png reproduced.png")
		os.Exit(1)
	}
	var paths []linetrace.Path
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read data:", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(data, &paths); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to unmarshal data:", err)
		os.Exit(1)
	}
	rand.Seed(time.Now().UnixNano())
	path := paths[rand.Intn(len(paths))]

	cost := &linetrace.CorrelationCost{Image: path.Image(28)}
	solution := make(linetrace.Path, 15)
	for i := range solution {
		solution[i].X = rand.Float64() * 120
		solution[i].Y = rand.Float64() * 120
	}
	var lastCost float64
	for i := 0; i < 5000; i++ {
		grad := linetrace.Gradient(solution, cost)
		linetrace.AddGradient(solution, grad, 200)
		curCost := cost.Cost(solution)
		log.Println("Epoch", i, "cost", cost.Cost(solution))
		if curCost == lastCost {
			break
		}
		lastCost = curCost
	}

	origOut, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write output:", err)
		os.Exit(1)
	}
	png.Encode(origOut, path.Image(28).GoImage())
	origOut.Close()

	reproOut, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write output:", err)
		os.Exit(1)
	}
	png.Encode(reproOut, solution.Image(28).GoImage())
	reproOut.Close()
}
