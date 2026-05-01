package main

import (
	"iter"
	"os"
	"strings"
)

func main() {

	f, _ := os.Open("source.txt")
	defer f.Close()

	text := ""
	for {
		size_buf := 128
		buf := make([]byte, size_buf)
		n, _ := f.Read(buf)

		if n != 0 {
			new_str := string(buf)
			text = strings.Join([]string{text, new_str}, "")
		} else {
			break
		}
	}

	lines := strings.Lines(text)

	slice_vars := Purify(lines)

	out_buffer := ""
	count := 0
	for _, single_var := range slice_vars {
		if count%2 == 0 {
			out_buffer = strings.Join([]string{out_buffer, single_var, " : "}, "")
		} else {
			out_buffer = strings.Join([]string{out_buffer, single_var, ",\n"}, "")
		}
		count++
	}
	out_buffer = strings.TrimSuffix(out_buffer, ",\n")
	out_buffer = strings.Join([]string{"{", out_buffer, "}"}, "")
	print(out_buffer)
}

func Purify(lines iter.Seq[string]) []string {
	var pure []string
	is_comment := false
	for single_line := range lines {

		single_line = strings.TrimSpace(single_line)
		single_line = strings.ReplaceAll(single_line, " ", "")
		if single_line == "" {
			continue
		}

		frags := strings.Split(single_line, "#")
		if len(frags) >= 2 {
			single_line = frags[0]
			if single_line == "" {
				continue
			}
		}

		frags = strings.Split(single_line, "\"\"\"")
		if len(frags) == 3 {
			//Skip, beginning and completing comment line
			continue
		} else if len(frags) == 2 && !is_comment {
			//like """<something>
			is_comment = true
			continue
		} else if len(frags) == 2 && is_comment {
			//like <something>"""(EOL)
			is_comment = false
			continue
		} else if len(frags) == 1 && is_comment {
			//skip, a line part of a multi-line comment
			continue
		}
		//coming here means it's not a comment line

		frags = strings.Split(frags[0], "=")
		frags = strings.Split(frags[0], ":")
		pure = append(pure, frags[0])
	}
	return pure
}
