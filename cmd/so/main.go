package main

// TODO: relevance/score filter

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/docopt/docopt.go"

	"github.com/cheat/so/internal/api"
)

const version = "1.0.0"

func main() {

	// specify the docopts
	// NB: the --json flag is currently ignored: `so` will currently only output
	// JSON. This flag is being left here to make it easier to potentially add
	// other output formats in the future.
	usage := `Usage:
  so [options] <question>

  Options:
  -j --json             Output JSON
  -n --number=<number>  Maximum number of answers to return
  -t --tag=<tag>        Filter by <tag>`

	// initialize options
	opts, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		// panic here, because this should never happen
		fmt.Fprintf(os.Stderr, "docopt failed to parse: %v\n", err)
		os.Exit(1)
	}

	// optionally assign a --tag value
	tags := ""
	if opts["--tag"] != nil {
		tags = opts["--tag"].(string)
	}

	// query the search API
	answers, err := api.Search(opts["<question>"].(string), "stackoverflow", tags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to query api: %v\n", err)
		os.Exit(1)
	}

	// decorate the (raw) answers with `Code` members
	for i := range answers {
		answers[i].Code, err = answers[i].Extract()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to extract code blocks: %v\n", err)
			os.Exit(1)
		}
	}

	// filter out answers that contain no <code>
	var filtered []api.Answer
	for _, a := range answers {
		if a.Code != "" {
			filtered = append(filtered, a)
		}
	}
	answers = filtered

	// optionally constrain the response to --number answers
	if opts["--number"] != nil {

		// parse --number into an integer (`limit`)
		limit, err := strconv.ParseInt(opts["--number"].(string), 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse --number: %v\n", err)
			os.Exit(1)
		}

		// assert that `answers` <= `limit`
		if len(answers) > int(limit) {
			answers = answers[0:limit]
		}
	}

	// output JSON and exit
	j, err := json.Marshal(answers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal answers: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(j))
	os.Exit(0)
}
