package format

import (
	"math"
	"strings"

	"github.com/muesli/reflow/truncate"
)

/*
Receives the content to center and the height of the container,
returns the text padded with new lines
*/
func CenterVertical(content string, height int) string {
	contentHeight := CountNewlines(content)
	// debug_log.LogToFile(fmt.Sprintf("container height:%d", height))
	// debug_log.LogToFile(fmt.Sprintf("content height:%d", contentHeight))
	padTop := math.Ceil(float64(height-contentHeight) / 2)
	padBottom := math.Floor(float64(height-contentHeight) / 2)
	// debug_log.LogToFile(fmt.Sprintf("padTop:%f\npadBottom:%f\n", padTop, padBottom))

	result := ""

	for i := 0; i < int(padTop)-1; i++ {
		result += "\n"
	}
	result += content
	for i := 0; i < int(padBottom); i++ {
		result += "\n"
	}
	return result
}

func TruncateHorizontal(text string, width int) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {

		// new_line := []rune(line)
		// line_len := len(new_line)
		// if line_len > width {

		// 	new_line = new_line[:width]
		// }
		lines[i] = truncate.String(line, uint(width))
	}
	return strings.Join(lines, "\n")
}

func CountNewlines(s string) int {
	count := 0
	for _, c := range s {
		if c == '\n' {
			count++
		}
	}
	return count
}
