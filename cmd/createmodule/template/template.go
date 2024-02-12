package template

import "strings"

func getModuleShort(module string) string {
	moduleSplitted := strings.Split(module, " ")
	short := ""

	for _, word := range moduleSplitted {
		short += strings.ToLower(string(word[0]))
	}

	return short
}
