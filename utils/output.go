package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Output is a single blob of data returned by the tool
type Output struct {
	Path        string `json:"path"`
	Description string `json:"description"`
	Comment     string `json:"comment"`
}

// Colors used throughout the program
var (
	Que    = "\033[95m"
	Blue   = "\033[94m"
	Green  = "\033[92m"
	Yellow = "\033[93m"
	Bad    = "\033[91m"
	Info   = "\033[1m"
	Reset  = "\033[0m"
)

// Now gets the current time
func Now() string {
	now := time.Now().Format("15:04:05")
	return now
}

func WriteOutput(OutputArray []*Output, state *State) {
	_, err := os.Create(state.Output)
	if err != nil {
		if state.Verbose {
			fmt.Printf("\n[%s%s%s] %s\n", Bad, Now(), Reset, err)
		}
		os.Exit(1)
	}

	data, err := json.MarshalIndent(OutputArray, "", "    ")
	if err != nil {
		if state.Verbose {
			fmt.Printf("\n[%s%s%s] %s\n", Bad, Now(), Reset, err)
		}
		os.Exit(1)
	}

	// Write the output to file
	err = ioutil.WriteFile(state.Output, data, 0644)
	if err != nil {
		if state.Verbose {
			fmt.Printf("\n[%s%s%s] %s\n", Bad, Now(), Reset, err)
		}
		os.Exit(1)
	}

	fmt.Printf("\n[%s%s%s] Written output to %s", Info, Now(), Reset, state.Output)
}
