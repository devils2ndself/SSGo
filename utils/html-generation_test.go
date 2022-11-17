package utils_test

import (
	"testing"

	utils "github.com/devils2ndself/SSGo/utils"

	. "github.com/stretchr/testify/assert"
)

// ParseText

func Test_ParseText_NoHeading(t *testing.T) {
	result, title := utils.ParseText("Test\n\nTest 2")
	Equal(t, "<p>Test </p>\n\n<p>Test 2 </p>\n", result)
	Equal(t, "", title)
}

func Test_ParseText_WithHeading(t *testing.T) {
	result, title := utils.ParseText("Test\n\n\nparagraph")
	Equal(t, "<h1>Test</h1>\n\n<p>paragraph </p>\n", result)
	Equal(t, "Test", title)
}

func Test_ParseText_NoText(t *testing.T) {
	result, title := utils.ParseText("")
	Equal(t, "", result)
	Equal(t, "", title)
}

func Test_ParseText_WhiteSpace(t *testing.T) {
	result, title := utils.ParseText("  Test  ")
	Equal(t, "<p>Test </p>\n", result)
	Equal(t, "", title)
}

// ParseMarkdown

func Test_ParseMarkdown_paragraphs(t *testing.T) {
	result := utils.ParseMarkdown([]byte("Test"))
	Equal(t, "<p>Test</p>\n", string(result))
}

func Test_ParseMarkdown_headings(t *testing.T) {
	result := utils.ParseMarkdown([]byte("# Test\n## Test\n### Test"))
	Equal(t, "<h1>Test</h1>\n\n<h2>Test</h2>\n\n<h3>Test</h3>\n", string(result))
}

func Test_ParseMarkdown_bold(t *testing.T) {
	result := utils.ParseMarkdown([]byte("**Test** __Test__"))
	Equal(t, "<p><strong>Test</strong> <strong>Test</strong></p>\n", string(result))
}

func Test_ParseMarkdown_italics(t *testing.T) {
	result := utils.ParseMarkdown([]byte("*Test* _Test_"))
	Equal(t, "<p><em>Test</em> <em>Test</em></p>\n", string(result))
}

func Test_ParseMarkdown_code(t *testing.T) {
	result := utils.ParseMarkdown([]byte("`Test`"))
	Equal(t, "<p><code>Test</code></p>\n", string(result))
}

func Test_ParseMarkdown_link(t *testing.T) {
	result := utils.ParseMarkdown([]byte("[Test](test.com)"))
	Equal(t, "<p><a href=\"test.com\">Test</a></p>\n", string(result))
}

func Test_ParseMarkdown_hr(t *testing.T) {
	result := utils.ParseMarkdown([]byte("---\n****\n____"))
	Equal(t, "<hr>\n\n<hr>\n\n<hr>\n", string(result))
}

func Test_ParseMarkdown_blockquote(t *testing.T) {
	result := utils.ParseMarkdown([]byte("> Test"))
	Equal(t, "<blockquote>\n<p>Test</p>\n</blockquote>", string(result))
}
