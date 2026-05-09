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
	size_buf := 128
	for {
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

	out_buffer := Import(lines)

	// for _, single_var := range slice_vars {
	// 	out_buffer = strings.Join([]string{out_buffer, single_var, ",\n"}, "")
	// }
	// out_buffer = strings.TrimSuffix(out_buffer, ",\n")
	// out_buffer = strings.Join([]string{"[", out_buffer, "]"}, "")
	print(out_buffer)

	byte_buffer := []byte(out_buffer)

	f_new, _ := os.Create("out.txt")
	defer f_new.Close()
	f_new.Write(byte_buffer)
}

func Import(lines iter.Seq[string]) string {
	is_comment := false
	skip_comment := false
	var output_buf string
	for single_line := range lines {
		if strings.HasPrefix(single_line, "\"\"\"") {
			if !skip_comment {
				output_buf += single_line
			}
			if !strings.HasSuffix(single_line, "\"\"\"\r\n") {
				is_comment = true
			} else if skip_comment {
				skip_comment = false
			}
			continue
		}
		if is_comment {
			if !skip_comment {
				output_buf += single_line
			}
			if strings.HasSuffix(single_line, "\"\"\"\r\n") {
				is_comment = false
				skip_comment = false
			}
			continue
		}
		space_trimmed := strings.ReplaceAll(single_line, " ", "")
		split_eq := strings.Split(space_trimmed, "=")
		if len(split_eq) >= 2 {
			skip_comment = false
			keyword := (strings.Split(split_eq[0], ":")[0])
			if !strings.HasPrefix(keyword, "tag") && !strings.HasPrefix(keyword, "hedr") {
				skip_comment = true
				continue
			}
			to_join := []string{keyword, " = defs.", keyword, "\n"}
			output_buf += strings.Join(to_join, "")
		}
	}
	return output_buf

}
