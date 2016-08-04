package main

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
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

	cost := &linetrace.CorrelationCost{Image: path.Image(24)}
	solution := linetrace.SearchPath(cost, 12, 5, 0, 0, 120, 120)

	fmt.Println("Solution cost:", cost.Cost(solution))

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
