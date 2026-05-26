package core

import "strings"

func (a *AsciiService) GenerateAscii(inputSlice []string, banner string) (string, error) {
	// to avoid memory overhead using a string builder would be better
	var sb strings.Builder
	file, err := a.Loadbanner(banner)
	if err != nil {
		return "", nil
	}
	for i := 0; i < len(inputSlice); i++ {
		// An empty segment in the input represents a newline in the original string
		if inputSlice[i] == "" {
			sb.WriteString("\n")
			continue
		}
		// Iterate over each of the 8 pixel rows that make up a character
		for j := 1; j <= 8; j++ {
			// For each character in the current word, print its j-th font row
			for _, val := range inputSlice[i] {
				if !(val >= 32 && val <= 127) {
					return "", INVALID_CHAR
				}
				vals := (int(val-32) * 9) + j
				sb.WriteString(file[vals])

			}
			sb.WriteString("\n")
		}
	}
	return sb.String(), nil
}
func (a *AsciiService) Loadbanner(s string) ([]string, error) {
	if val, ok := a.Banners[s]; ok {
		return val, nil
	}
	return nil, Banner_Not_Found
}
