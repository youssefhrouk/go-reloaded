package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	reloaded "reloaded/packages"
)

func bintoDec(s string) string {
	dec, _ := strconv.ParseInt(s, 2, 64)
	return fmt.Sprint(dec)
}

func hextoDec(s string) string {
	hex, _ := strconv.ParseInt(s, 16, 64)
	return fmt.Sprint(hex)
}

func isValidHex(s string) bool {
	_, err := strconv.ParseInt(s, 16, 64)
	return err == nil
}

func isValidBin(s string) bool {
	_, err := strconv.ParseInt(s, 2, 64)
	return err == nil
}

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Usage: program input.txt output.txt")
		os.Exit(1)
	}
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	if filepath.Ext(inputFileName) != ".txt" {
		fmt.Println("Input file must have a .txt extension.")
		return
	}
	if filepath.Ext(outputFileName) != ".txt" {
		fmt.Println("Output file must have a .txt extension.")
		return
	}
	s, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	text := string(s)
	var modifiedLines []string
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		var words []string
		words = strings.Fields(line)
		for i := 0; i < len(words); i++ {
			if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ",") {
				directive := words[i]
				var nbStr string
				if i != len(words)-1 {
					nbStr = words[i+1]
				}
				if !strings.HasSuffix(nbStr, ")") {
					fmt.Printf("Invalid syntax for directive: %s\n", directive)
					continue
				}
				nbStr = nbStr[:len(nbStr)-1] // Remove trailing ")"
				nb, err := strconv.Atoi(nbStr)
				if err != nil {
					fmt.Printf("Invalid syntax for directive: %s\n", directive)
					continue
				}
				if nb > i {
					fmt.Printf("Didn't find %v words to process for directive: %s\n", nb, directive)
					words = append(words[:i],words[i+2:]...)
					continue
				}
				switch directive {
				case "(up,":
					for k := 1; k <= nb; k++ {
						words[i-k] = strings.ToUpper(words[i-k])
					}
				case "(low,":
					for k := 1; k <= nb; k++ {
						words[i-k] = strings.ToLower(words[i-k])
					}
				case "(cap,":
					for k := 1; k <= nb; k++ {
						words[i-k] = reloaded.Capitalize(words[i-k])
					}
				}
				words = append(words[:i], words[i+2:]...)
				i--
			} else if words[i] == "(bin)" {
				if i != 0 {
					if isValidBin(words[i-1]) {
						words[i-1] = bintoDec(words[i-1])
						words = append(words[:i], words[i+1:]...)
						i-- // Adjust index after removing element
					}
				} else {
					words[i] = ""
					words[i] = strings.Trim(words[i], " ")
					continue
				}
			} else if words[i] == "(hex)" {
				if i != 0 {
					if isValidHex(words[i-1]) {
						words[i-1] = hextoDec(words[i-1])
						words = append(words[:i], words[i+1:]...)
						i--
					}
				} else {
					words[i] = ""
					words[i] = strings.Trim(words[i], " ")
					continue
				}
			} else if words[i] == "(up)" {
				if i != 0 {
					words[i-1] = strings.ToUpper((words[i-1]))
					words = append(words[:i], words[i+1:]...)
					i--
				} else {
					words[i] = ""
					words[i] = strings.Trim(words[i], " ")
					continue
				}
			} else if words[i] == "(cap)" {
				if i != 0 {
					words[i-1] = reloaded.Capitalize((words[i-1]))
					words = append(words[:i], words[i+1:]...)
					i--
				} else {
					words[i] = ""
					words[i] = strings.Trim(words[i], " ")
					continue
				}
			} else if words[i] == "(low)" {
				if i != 0 {
					words[i-1] = strings.ToLower((words[i-1]))
					words = append(words[:i], words[i+1:]...)
					i--
				} else {
					words[i] = ""
					words[i] = strings.Trim(words[i], " ")
					continue
				}
			}
		}
		modifiedSentence := reloaded.Punctuation(words)
		modifiedLines = append(modifiedLines, modifiedSentence)
	}
	outputContent := strings.Join(modifiedLines, "\n")
	outputContent = reloaded.Apostrophe(outputContent)

	file, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(outputContent)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
}
