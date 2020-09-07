package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

// Search queries the Stack Exchange API for a question matching the search
// criteria
func Search(search string, site string, tags string) ([]Answer, error) {

	// trim the search string
	search = strings.TrimSpace(search)

	// disallow an empty search string
	if search == "" {
		return []Answer{}, fmt.Errorf("<question> must not be empty")
	}

	// initialize an HTTP client
	url := "https://api.stackexchange.com/2.2/search/advanced"
	cl := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Answer{}, fmt.Errorf("failed to init http request: %v", err)
	}

	// assemble the query string parameters
	q := req.URL.Query()
	q.Add("answers", "1") // this is the minimum - not the total
	q.Add("filter", "!2uDdBA_UQti85J*)fMajDUwad2zPVKEdobB3yIR1nL")
	q.Add("order", "desc")
	q.Add("q", search)
	q.Add("site", site)
	q.Add("sort", "relevance")
	if tags != "" {
		q.Add("tagged", tags)
	}
	req.URL.RawQuery = q.Encode()

	// make the HTTP request
	res, err := cl.Do(req)
	if err != nil {
		return []Answer{}, fmt.Errorf("failed to make http request: %v", err)
	}

	// unpack the JSON response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Answer{}, fmt.Errorf("failed to read response: %v", err)
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []Answer{}, fmt.Errorf("failed to unmarshall json: %v", err)
	}

	// return early if no answers were returned
	if len(response.Items) == 0 {
		return []Answer{}, nil
	}

	question := response.Items[0]
	answers := make([]Answer, 0)
	for _, a := range question.Answers {
		if a.Score >= 1 {
			answers = append(answers, a)
		}
	}

	// sort the answers by score from highest to lowest
	sort.SliceStable(answers, func(i, j int) bool {
		return answers[i].Score > answers[j].Score
	})

	return answers, nil
}
