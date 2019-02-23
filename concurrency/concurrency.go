package concurrency

// WebsiteChecker checks the status of a url input
type WebsiteChecker func(string) bool

// CheckWebsites takes an array of urls and checks the status
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
