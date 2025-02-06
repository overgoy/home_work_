package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func analyzeLogs(filePath, level string) (map[string]int, error) {
	file, newErr := os.Open(filePath)
	if newErr != nil {
		return nil, newErr
	}
	defer file.Close()

	stats := make(map[string]int)
	logPattern := regexp.MustCompile(`\[(INFO|ERROR|DEBUG|WARN)\]`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if matches := logPattern.FindStringSubmatch(line); matches != nil {
			logLevel := matches[1]
			if level == "" || strings.EqualFold(logLevel, level) {
				stats[logLevel]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func writeStats(stats map[string]int, outputPath string) error {
	var output *os.File
	var err error

	if outputPath != "" {
		output, err = os.Create(outputPath)
		if err != nil {
			return err
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	for level, count := range stats {
		fmt.Fprintf(output, "%s: %d\n", level, count)
	}
	return nil
}

func main() {
	filePath := flag.String("file", os.Getenv("LOG_ANALYZER_FILE"), "Path to log file")
	logLevel := flag.String("level", os.Getenv("LOG_ANALYZER_LEVEL"), "Log level to analyze (default: INFO)")
	outputPath := flag.String("output", os.Getenv("LOG_ANALYZER_OUTPUT"), "Output file for statistics")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Error: Log file path is required.")
		flag.Usage()
		os.Exit(1)
	}

	stats, analyzeErr := analyzeLogs(*filePath, *logLevel)
	if analyzeErr != nil {
		fmt.Printf("Error analyzing logs: %v\n", analyzeErr)
		os.Exit(1)
	}

	if outputErr := writeStats(stats, *outputPath); outputErr != nil {
		fmt.Printf("Error writing stats: %v\n", outputErr)
		os.Exit(1)
	}
}
