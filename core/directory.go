package core

import (
	"fmt"

	"github.com/Ice3man543/hawkeye/utils"
	"github.com/karrick/godirwalk"
)

func WorkPath(pathChan <-chan string, resultChan chan<- *utils.Output, state *utils.State) {
	for path := range pathChan {
		found := false

		matchFile := NewMatchFile(path)
		if matchFile.IsSkippable() {
			continue
		}

		if state.Signature.CryptoFiles {
			for _, signature := range CryptoFilesSignatures {
				if signature.Match(matchFile) {
					output := &utils.Output{
						Path:        path,
						Description: signature.Description(),
						Comment:     signature.Comment(),
					}

					resultChan <- output
					found = true
					break
				}
			}
		}

		if !found && state.Signature.PasswordFiles {
			for _, signature := range PasswordFileSignatures {
				if signature.Match(matchFile) {
					output := &utils.Output{
						Path:        path,
						Description: signature.Description(),
						Comment:     signature.Comment(),
					}

					resultChan <- output
					found = true
					break
				}
			}
		}

		if !found && state.Signature.ConfigurationFiles {
			for _, signature := range ConfigurationFileSignatures {
				if signature.Match(matchFile) {
					output := &utils.Output{
						Path:        path,
						Description: signature.Description(),
						Comment:     signature.Comment(),
					}

					resultChan <- output
					found = true
					break
				}
			}
		}

		if !found && state.Signature.DatabaseFiles {
			for _, signature := range DatabaseFileSignatures {
				if signature.Match(matchFile) {
					output := &utils.Output{
						Path:        path,
						Description: signature.Description(),
						Comment:     signature.Comment(),
					}

					resultChan <- output
					found = true
					break
				}
			}
		}

		if !found && state.Signature.MiscFiles {
			for _, signature := range MiscSignatures {
				if signature.Match(matchFile) {
					output := &utils.Output{
						Path:        path,
						Description: signature.Description(),
						Comment:     signature.Comment(),
					}

					resultChan <- output
					found = true
					break
				}
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
