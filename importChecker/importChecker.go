package main

import (
	"os"
	"strings"
)

func main() {
	size_buf := 128

	f_source, _ := os.Open("source.txt")
	defer f_source.Close()
	sourceText := ""
	for {
		buf := make([]byte, size_buf)
		n, _ := f_source.Read(buf)

		if n != 0 {
			new_str := string(buf)
			sourceText = strings.Join([]string{sourceText, new_str}, "")
		} else {
			break
		}
	}
	sourceTokens := TokensAtAsource(sourceText)

	f_dest, _ := os.Open("destination.txt")
	defer f_dest.Close()
	destText := ""
	for {
		buf := make([]byte, size_buf)
		n, _ := f_dest.Read(buf)

		if n != 0 {
			new_str := string(buf)
			destText = strings.Join([]string{destText, new_str}, "")
		} else {
			break
		}
	}

	found := ""
	not_found := ""
	for _, token := range sourceTokens {
		if strings.Contains(destText, "defs."+token) {
			found += ("\t" + token + "\n")
		} else {
			not_found += ("\t" + token + "\n")
		}
	}

	println("Tokens found in the destination:")
	print(found)
	println()
	println("Tokens NOT found in the destination:")
	print(not_found)
	println()

	println("Press enter to quit")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func TokensAtAsource(source string) []string {
	var output_buf []string
	is_comment := false

	lines := strings.Lines(source)
	for single_line := range lines {
		if strings.HasPrefix(single_line, "\"\"\"") {
			if !strings.HasSuffix(single_line, "\"\"\"\r\n") {
				is_comment = true
			}
			continue
		}
		if is_comment {
			if strings.HasSuffix(single_line, "\"\"\"\r\n") {
				is_comment = false
			}
			continue
		}
		space_trimmed := strings.ReplaceAll(single_line, " ", "")
		space_trimmed = strings.ReplaceAll(space_trimmed, "\b", "")
		split_eq := strings.Split(space_trimmed, "=")
		keyword := (strings.Split(split_eq[0], ":")[0])
		if strings.HasPrefix(keyword, "tag_") ||
			strings.HasPrefix(keyword, "hedr_") ||
			strings.HasPrefix(keyword, "opt_") ||
			strings.HasPrefix(keyword, "list_") ||
			strings.HasPrefix(keyword, "dict_") {
			output_buf = append(output_buf, keyword)
		}
	}
	return output_buf
}
