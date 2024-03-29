package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var commandLineOptions Options

	flag.BoolVar(&commandLineOptions.printBytes, "c", false, "Count bytes")
	flag.BoolVar(&commandLineOptions.printLines, "l", false, "Count lines")
	flag.BoolVar(&commandLineOptions.printWords, "w", false, "Count words")
	flag.BoolVar(&commandLineOptions.printChars, "m", false, "Count characters")

	flag.Parse()

	if !commandLineOptions.printBytes &&
		!commandLineOptions.printChars &&
		!commandLineOptions.printLines &&
		!commandLineOptions.printWords {
		commandLineOptions.printBytes = true
		commandLineOptions.printChars = true
		commandLineOptions.printLines = true
		commandLineOptions.printWords = true
	}

	filenames := flag.CommandLine.Args()

	run(filenames, commandLineOptions)
}

func run(filenames []string, options Options) {
	if len(filenames) == 0 {
		reader := bufio.NewReader(os.Stdin)
		fileStats := CalculateStats(reader)
		fmt.Println(formatStats(options, fileStats, ""))
	} else {
		CalculateStatsForFiles(filenames, options)
	}
}
