package cli

import "fmt"

// Flags holds common CLI flags
type Flags struct {
	DebugFlag *bool
}

//BuildVersion version of caolila
var BuildVersion = ""

//BuildTime time of build
var BuildTime = ""

//GitRevision revision of git at time of build
var GitRevision = ""

// ShowVersion prints out the version information on stdout
func ShowVersion() {
	fmt.Printf("app version : %s\n", BuildVersion)
	fmt.Printf("git revision: %s\n", GitRevision)
	fmt.Printf("build time  : %s\n", BuildTime)
}
