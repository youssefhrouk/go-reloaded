package reloaded

import (
	"regexp"
	"strings"
)

func Apostrophe(text string) string {
	lines := strings.Split(text, "\n")
	newLines := []string{}
	for _, line := range lines {
		// re1 ----> to re4 make sure there is a space before and after each apostrophe, excluding the ones between characters.
		re1 := regexp.MustCompile(`('\s+)`)
		line = re1.ReplaceAllString(line, " $1")
		re2 := regexp.MustCompile(`(\s+')`)
		line = re2.ReplaceAllString(line, "$1 ")
		re3 := regexp.MustCompile(`\A'`)
		line = re3.ReplaceAllString(line, " ' ")
		re4 := regexp.MustCompile(`'$`)
		line = re4.ReplaceAllString(line, " ' ")

		// remove spaces on the right of even number apostrophes.
		re5 := regexp.MustCompile(`'\s+`)
		count := 0
		line = re5.ReplaceAllStringFunc(line, func(match string) string {
			if count%2 == 0 {
				count++
				return "'"
			} else {
				count++
				return match
			}
		})
		count = 0
		re6 := regexp.MustCompile(`\s+'`)

		// remove spaces on the left of odd number apostrophes.
		line = re6.ReplaceAllStringFunc(line, func(match string) string {
			if count%2 == 1 {
				count++
				return "'"
			} else {
				count++
				return match
			}
		})

		// re7 ----> re9 clean any remaining spaces after or before the apostrophe.
		re7 := regexp.MustCompile(`[ ]+'`)
		line = re7.ReplaceAllString(line, " '")
		re8 := regexp.MustCompile(`'[ ]+`)
		line = re8.ReplaceAllString(line, "' ")
		re9 := regexp.MustCompile(`\A '`)
		line = re9.ReplaceAllString(line, "'")

		newLines = append(newLines, strings.TrimRight(line, " \t"))
	}
	text = strings.Join(newLines, "\n")

	return strings.Trim(text, " \t")
}
