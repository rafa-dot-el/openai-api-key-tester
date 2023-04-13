package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const modelsURL = "https://api.openai.com/v1/models"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	validCount := 0
	invalidCount := 0

	for scanner.Scan() {
		key := strings.TrimSpace(scanner.Text())
		fmt.Printf("[*] Checking API key: %s\n", key)

		client := &http.Client{}
		req, err := http.NewRequest("GET", modelsURL, nil)
		if err != nil {
			fmt.Println("[-] Error creating request:", err)
			continue
		}
		req.Header.Set("Authorization", "Bearer "+key)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("[-] Error sending request:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("\033[32m[+] Valid API key: %s\033[0m\n", key)
			validCount++
		} else {
			fmt.Printf("\033[31m[-] Invalid API key: %s\033[0m\n", key)
			invalidCount++
		}

		time.Sleep(5 * time.Second)
	}

	fmt.Printf("[*] Summary: %d valid keys, %d invalid keys\n", validCount, invalidCount)
}
