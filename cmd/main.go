package main

import (
	"flag"
	"fmt"
	"os"
)

// These are populated at build time
var (
	// Version is the version string at which the CLI is built.
	Version string
	// BuildDate is the date on which this CLI binary was built
	BuildDate string
	// Commit is the git commit from which this CLI binary was built.
	Commit string
	//BuiltBy is the release program that built this binary
	BuiltBy string
	//Os the Operating System for which this binary is built
	Os string
	//Arch the Architecture for which this binary is compatible
	Arch string
)

func main() {
	fmt.Println("Jai Guru!")

	versionInfo := flag.Bool("version", false, "display version information")
	flag.Parse()

	if *versionInfo {
		fmt.Printf("Version:      %s\n", Version)
		fmt.Printf("Build Date:   %s\n", BuildDate)
		fmt.Printf("Git Revision: %s\n", Commit)
		fmt.Printf("Built-By: %s\n", BuiltBy)
		fmt.Printf("OS: %s\n", Os)
		fmt.Printf("Arch: %s\n", Arch)
		os.Exit(0)
	}
}
