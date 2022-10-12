package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Checks if line of text is markdown hotizontal rule.
// Returns true if text is horizontal rule
func CheckHorizontalRule(text string) bool {
	acceptedPrefixes := [3]string{"---", "***", "___"}

	for _, prefix := range acceptedPrefixes {
		if strings.HasPrefix(text, prefix) {
			currentChar := rune(prefix[0])
			for _, char := range text {
				if char != currentChar {
					return false
				}
			}
			return true
		}
	}

	return false
}

// Checks if line of text has markdown prefixes, currently used for checking headings. 
// Returns true if any of the appropriate prefixes are found
func CheckMarkdownPrefix(text string) (string, bool) {
	acceptedPrefixes := [2]string{"# ", "## "}

	for _, prefix := range acceptedPrefixes {
		if strings.HasPrefix(text, prefix) {
			return prefix, true
		}
	}

	return "", false
}

// Replaces inline markdown features into HTML, currently used for `code blocks`.
// Takes a line of text and returns that line processed into HTML (if found) 
func GenerateInlineMarkdownHtml(text string) string {
    text = regexp.MustCompile(`\*\*(.*?)\*\*`).ReplaceAllString(text, `<b>$1</b>`)
    text = regexp.MustCompile(`__(.*?)__`).ReplaceAllString(text, `<b>$1</b>`)
    text = regexp.MustCompile(`_(.*?)_`).ReplaceAllString(text, `<i>$1</i>`)
    text = regexp.MustCompile(`\*(.*?)\*`).ReplaceAllString(text, `<i>$1</i>`)
	text = regexp.MustCompile("`(.*?)`").ReplaceAllString(text, `<code>$1</code>`)
    text = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`).ReplaceAllString(text, `<a href="$2">$1</a>`)
	return text
}

// Replaces markdown prefixes into HTML tags.
// Takes a prefix (#) & line of text and returns that line proccessed into HTML
func GeneratePrefixMarkdownHtml(prefix string, text string) string {
	prefixesHtmlFormatStrings := map[string]string{
		"# ":  "<h1>%s</h1>",
		"## ": "<h2>%s</h2>",
	}

	if formatString, found := prefixesHtmlFormatStrings[prefix]; found {
		return fmt.Sprintf(formatString, strings.Replace(text, prefix, "", 1)) + "\n"
	}

	return text
}
