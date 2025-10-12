package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseConfig(content string) (map[string]string, error) {
	config := make(map[string]string)

	re := regexp.MustCompile(`^\s*([\w.-]+)\s*=\s*(?:'([^']*)'|"([^"]*)"|([^#\s]*))?(?:\s*#.*)?$`)

	scanner := bufio.NewScanner(strings.NewReader(content))
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		matches := re.FindStringSubmatch(trimmedLine)
		if matches == nil {
			fmt.Printf("Line %d: '%s' - Is Inavlid\n", lineNo, line)
			continue
		}
		key := matches[1]
		var value string

		if matches[2] != "" {
			value = matches[2]
		} else if matches[3] != "" {
			value = matches[3]
		} else {
			value = matches[4]
		}

		config[key] = value
	}

	return config, nil
}

func main() {

	envFileContent := `
# Application Configuration
APP_NAME="My Cool App"
APP_VERSION="1.0.2-beta" # Version with quotes
PORT=8080
DEBUG_MODE="true"
# Database Settings
DB_HOST=localhost
DB_USER = admin
DB_PASSWORD = "p@s$w Ord With Sp@ces!" # Quoted password
API_ENDPOINT = https://api.example.com/v1

# An empty value
EMPTY_KEY=
ANOTHER_KEY_NO_VALUE =`

	config, err := parseConfig(envFileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for k, v := range config {
		fmt.Printf("%s=%q\n", k, v)
	}

}
