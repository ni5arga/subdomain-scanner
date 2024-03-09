package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	targetDomain   string
	wordlistPath   string
	outputFilePath string
)

func init() {
	flag.StringVar(&targetDomain, "domain", "", "The target domain which will be scanned for subdomains. Example: example.com")
	flag.StringVar(&wordlistPath, "wordlist", "subdomains.txt", "path to wordlist file, uses subdomains.txt if none provided")
	flag.StringVar(&outputFilePath, "output-file", "found-subdomains.txt", "output file to write found subdomains to, uses found-subdomains.txt if none provided")
	flag.Parse()
}

func main() {
	if targetDomain == "" {
		fmt.Println("Please give a target domain using the -domain flag.")
		return
	}

	// get the list from wordlist
	subdomains, err := readWordlist(wordlistPath)
	if err != nil {
		fmt.Println("Error reading wordlist:", err)
		return
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	fmt.Fprintln(outputFile, "Discovered Subdomains with Status Code 200")

	channel := make(chan int, len(subdomains))

	for _, subdomain := range subdomains {
		url := fmt.Sprintf("http://%s.%s", subdomain, targetDomain)
		go scan(url, channel)
	}

	for _, subdomain := range subdomains {
		url := fmt.Sprintf("http://%s.%s", subdomain, targetDomain)

		statusCode := <-channel
		if statusCode == 200 {
			fmt.Printf("[✅] Discovered subdomain: %s\n", url)
			// write the subdomain to file
			fmt.Fprintln(outputFile, url)
		}
	}
}

func readWordlist(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var subdomains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		subdomains = append(subdomains, scanner.Text())
	}

	return subdomains, scanner.Err()
}

func scan(url string, channel chan int) {
	resp, err := http.Get(url)

	if resp == nil || err != nil {
		channel <- 0
		return
	}

	channel <- resp.StatusCode
	defer resp.Body.Close()
}
