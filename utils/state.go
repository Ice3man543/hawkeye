package utils

// State contains current application state
type State struct {
	Directory         string
	Threads           int
	Output            string
	Verbose           bool
	ListSignatures    bool
	Signatures        string
	ExcludeSignatures string

	Signature *Sign
}

// Sign contains Signatures to be ran
type Sign struct {
	CryptoFiles        bool
	PasswordFiles      bool
	ConfigurationFiles bool
	DatabaseFiles      bool
	MiscFiles          bool
}
