package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/Ice3man543/hawkeye/core"
	"github.com/Ice3man543/hawkeye/utils"
)

func main() {
	fmt.Printf(utils.Banner)
	fmt.Printf("\n[%s%s%s] HawkEye : An advance filesystem analysis tool", utils.Info, utils.Now(), utils.Reset)
	fmt.Printf("\n[%s%s%s] Written By : @Ice3man", utils.Info, utils.Now(), utils.Reset)
	fmt.Printf("\n[%s%s%s] Github : https://github.com/Ice3man543\n\n", utils.Info, utils.Now(), utils.Reset)

	state := utils.ParseArguments()

	if state.Directory == "" {
		fmt.Printf("\nhawkeye: no directory specified\n")
		fmt.Printf("For Usage instructions, use -h flag\n")
		os.Exit(1)
	}

	var OutputArray []*utils.Output
	if state.Directory != "" {
		var wg, wg2 sync.WaitGroup

		pathChan := make(chan string)
		wg.Add(state.Threads)
		resultChan := make(chan *utils.Output)
		wg2.Add(1)

		for i := 0; i < state.Threads; i++ {
			go func() {
				defer wg.Done()
				core.WorkPath(pathChan, resultChan)
			}()
		}

		go func() {
			core.ProcessDirectory(state.Directory, state, pathChan)
			close(pathChan)
		}()

		go func() {
			for result := range resultChan {
				fmt.Printf("\n[%s%s%s] %s", utils.Info, result.Description, utils.Reset, result.Path)

				OutputArray = append(OutputArray, result)
			}

			wg2.Done()
		}()

		wg.Wait()
		close(resultChan)
		wg2.Wait()
	}

	if state.Output != "" {
		utils.WriteOutput(OutputArray, state)
	}

	fmt.Printf("\n\n[%s%s%s] Enjoy the hunt! \\o/", utils.Info, utils.Now(), utils.Reset)
}
