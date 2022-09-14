package snippets

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url        string
	err        error
	latency    time.Duration
	statuscode int
}

func fetch(url string, ch chan<- result) {

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- result{url, err, 0, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t, resp.StatusCode}
		resp.Body.Close()
	}
}

func main() {

	results_ch := make(chan result)
	targetUrls := []string{
		"https://kyledinh.com",
		"https://mockingbox.com",
		"https://github.com",
		"https://nytimes.com",
		"https://datawasher.txferry.com",
		"https://datawasher.txferry.com/create?limit=10&first_name=MOX_RFN&last_name=MOX_RLN&email=MOX_EMAIL&addr=MOX_RSA&code=MOX_RI_1000&state=MOX_STATE&sex=MOX_RSMF",
		"https://datawasher.txferry.com/random_contact",
		"https://datawasher.txferry.com/contacts",
	}

	// each fetch in a go routine
	for _, url := range targetUrls {
		go fetch(url, results_ch)
	}

	// keeps chan open til all done
	for range targetUrls {
		r := <-results_ch

		if r.err != nil {
			log.Printf("%-20s %s \n", r.url, r.err)
		} else {
			log.Printf("%-20s %s %v \n", r.url, r.latency, r.statuscode)
		}
	}

}

