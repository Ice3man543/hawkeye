package core

import (
	"fmt"

	"github.com/Ice3man543/hawkeye/utils"
	"github.com/karrick/godirwalk"
)

func WorkPath(pathChan <-chan string, resultChan chan<- *utils.Output) {
	for path := range pathChan {
		matchFile := NewMatchFile(path)
		if matchFile.IsSkippable() {
			continue
		}

		for _, signature := range Signatures {
			if signature.Match(matchFile) {
				output := &utils.Output{
					Path:        path,
					Description: signature.Description(),
					Comment:     signature.Comment(),
				}

				resultChan <- output
				break
			}
		}

	}
}

// ProcessDirectory processes a directory specified
func ProcessDirectory(Directory string, state *utils.State, pathChan chan<- string) {
	_ = godirwalk.Walk(Directory, &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.IsDir() == false {
				if state.Verbose {
					fmt.Printf("[%s%s%s] [File] %s\n", utils.Blue, utils.Now(), utils.Reset, osPathname)
				}
				pathChan <- osPathname
			}
			return nil
		},

		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
	})
}
