package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"	
	"strings"
)

func main() {
	var (
		sum     map[string]int // total visits per domain
		domains []string       // unique domain names
		total   int            // total visits to all domains		
	)

	sum = make(map[string]int)

    fmt.Println("Enter EXIT to end the input")
	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {		
		domain := in.Text()
		
		if (domain=="EXIT") {
		break;
		}

		// Collect the unique domains
		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		// Keep track of total and per domain visits
		total += 1
		sum[domain] += 1		
	}

	// Print the visits per domain
	sort.Strings(domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

}