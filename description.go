package check

import "strings"

func joinDescriptions(descriptions ...string) string {
	if len(descriptions) == 0 {
		return ""
	}

	return strings.Join(descriptions, ": ") + ": "
}
