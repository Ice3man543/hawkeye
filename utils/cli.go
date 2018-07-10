package utils

import (
	"flag"
)

// Banner is the application banner
var Banner = `
 _  _                _    ___           
| || | __ _ __ __ __| |__| __|_  _  ___ 
| __ |/ _  |\ V  V /| / /| _|| || |/ -_)
|_||_|\__,_| \_/\_/ |_\_\|___|\_, |\___|
                              |__/     
	    Analysis v1.0 - by @Ice3man
`

// ParseArguments parses the command line arguments
func ParseArguments() *State {
	s := State{}

	flag.IntVar(&s.Threads, "t", 20, "Number of threads to use")
	flag.StringVar(&s.Directory, "d", "", "Directory to search stuff in (Required)")
	flag.StringVar(&s.Output, "o", "", "File to write enumeration output to")
	flag.StringVar(&s.Signatures, "sig", "", "Signatures for type of files to find")
	flag.StringVar(&s.ExcludeSignatures, "exclude-sig", "", "Signatures to be excluded from the scan")
	flag.BoolVar(&s.ListSignatures, "l", false, "List the signatures present")
	flag.BoolVar(&s.Verbose, "v", false, "Display Verbose output")
	flag.Parse()

	s.Signature = &Sign{}
	return &s
}
