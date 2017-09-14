package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/docopt/docopt-go"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

// Make a timestampped "horizontal rule" to separate output into groups.
func horizontalRule(title string) string {
	now := time.Now().String()
	newTitle := fmt.Sprintf("%s %s", now, title)
	result := "----- " + newTitle + " " + strings.Repeat("-", 71-len(newTitle))
	return result
}

func main() {
	usage := `
Usage:
    go-stdin-sleep-stdout [--version] [--help] [--sleep=<seconds>] [--debug]

Options:
   -h, --help
   --sleep=<seconds>  How many seconds to sleep.
   --debug            Additional output for debugging
`
	// DocOpt processing.

	commandVersion := fmt.Sprintf("%s %s-%s", programName, buildVersion, buildIteration)
	args, _ := docopt.Parse(usage, nil, true, commandVersion, true)

	// Configure output log.

	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds | log.LUTC)

	// Parse command-line options.

	isDebug := args["--debug"].(bool)
	sleepTime := 0
	if args["--sleep"] != nil {
		sleepTime, _ = strconv.Atoi(args["--sleep"].(string))
	}

	// Read entire input from 'stdin'

	if isDebug {
		log.Println(horizontalRule("Begin go-stdin-sleep-stdout"))
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("ioutil.ReadAll(os.Stdin) failed.  Error: %+v\n", err)
	}

	if isDebug {
		log.Printf("%s\n", input)
	}

	// Sleep.

	if isDebug {
		log.Printf("Sleep(%d)\n", sleepTime)
	}
	time.Sleep(time.Duration(sleepTime) * time.Second)

	// Write output (aka input) to 'stdout'.

	if isDebug {
		log.Println(horizontalRule("End go-stdin-sleep-stdout"))
	}

	fmt.Printf("%s", input)

}
