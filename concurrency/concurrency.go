package concurrency

func mockWebsiteChecker(url string) bool {
	if url == "aaa" {
		return false
	}
	return true
}

type WebsiteChecker func(string) bool

type result struct {
	url string
	ret bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.ret
	}
	return results
}
