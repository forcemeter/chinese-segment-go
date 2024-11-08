package dicts

import (
	_ "embed"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/text/gstr"
)

//go:embed health.txt
var txt string

func LoadHealthDict() []string {
	rows := gstr.Explode("\n", txt)

	words := gset.NewStrSet()
	for _, row := range rows {
		if row == "" {
			continue
		}
		word := gstr.Explode("\t", row)[0]

		if word == "" || gstr.LenRune(word) <= 2 {
			continue
		}

		words.Add(word)
	}

	return words.Slice()
}
