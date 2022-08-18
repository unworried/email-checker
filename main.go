package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/unworried/email-checker/checker"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		out, err := checker.CheckDomain(scanner.Text())
		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		fmt.Printf(out)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v", err)
	}
}
