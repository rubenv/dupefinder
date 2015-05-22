package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/rubenv/dupefinder"
)

const (
	generateHelp = `Usage: dupefinder -generate filename folder...
    Generates a catalog file at filename based on one or more folders`

	detectHelp = `Usage: dupefinder -detect [-dryrun / -rm] filename folder...
    Detects duplicates using a catalog file in on one or more folders`
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	detect := true
	generate := false

	flag.BoolVar(&detect, "detect", false, "Detect duplicate files using a catalog")
	flag.BoolVar(&generate, "generate", false, "Generate a catalog file")

	// Detection flags
	dryrun := true
	rm := false

	flag.BoolVar(&dryrun, "dryrun", false, "Print what would be deleted")
	flag.BoolVar(&rm, "rm", false, "Delete detected duplicates (at your own risk!)")

	flag.Usage = func() {
		fmt.Println(generateHelp)
		fmt.Println()
		fmt.Println(detectHelp)
		fmt.Println()
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()

	if !detect && !generate {
		fmt.Println("Either -generate or -detect should be specified")
		os.Exit(1)
	}

	if detect && generate {
		fmt.Println("Only one of -generate or -detect should be specified")
		os.Exit(1)
	}

	if generate {

		if len(args) < 2 {
			fmt.Println(generateHelp)
			os.Exit(1)
		}

		err := dupefinder.Generate(args[0], args[1:]...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if detect {
		if !dryrun && !rm {
			fmt.Println("Either -rm or -dryrun should be specified")
			os.Exit(1)
		}

		if dryrun && rm {
			fmt.Println("Only one of -rm or -dryrun should be specified")
			os.Exit(1)
		}

		if len(args) < 2 {
			fmt.Println(detectHelp)
			os.Exit(1)
		}

		err := dupefinder.Detect(args[0], dryrun, rm, args[1:]...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
