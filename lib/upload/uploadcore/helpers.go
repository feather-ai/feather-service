package uploadcore

import "strings"

func ExtractPackedFilenames(packedFilenames string) []string {
	raw := strings.Split(packedFilenames, "\n")

	ret := make([]string, 0, len(raw))
	for _, f := range raw {
		if len(f) > 0 {
			ret = append(ret, f)
		}
	}
	return ret
}
