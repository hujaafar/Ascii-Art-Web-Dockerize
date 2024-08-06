package functions

import (
	"errors"
	"os"
	"strings"
)

func Ascii(txt, format string) (string, error) {
	str := ""
	txt = strings.ReplaceAll(txt, "\r\n", "\n")
	textSlice := strings.Split(txt, "\n")

	if !charValidation(txt) {
		return "", errors.New("error: invalid character")
	}

	file, err := os.ReadFile("banners/" + format + ".txt")
	if err != nil {
		return "", errors.New("error: reading file")
	}

	slice := strings.Split(string(file), "\n")

	for j, txt := range textSlice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					str += slice[firstLine]
				}
				str += "\n"
			}
		} else if j != len(textSlice)-1 {
			str += "\n"
		}
	}
	return str, nil
}
