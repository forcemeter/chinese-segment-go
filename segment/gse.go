package segment

import (
	"github.com/go-ego/gse"
	"strings"
)

type Gse struct {
	c gse.Segmenter
}

func (sg *Gse) Init() {
	sg.c = gse.Segmenter{}

	sg.c.Dict = gse.NewDict()
	sg.c.Init()
	if err := sg.c.LoadDictEmbed("zh"); err != nil {
		return
	}
	sg.c.LoadNoFreq = true
}

func (sg *Gse) LoadDictWords(words []string) error {
	return sg.c.LoadDictStr(strings.Join(words, "\n"))
}

func (sg *Gse) Cut(text string) (res Tokens) {
	res.Words = sg.c.Cut(text, true)
	return
}

func (sg *Gse) Tag(text string) (res Tokens) {
	return
}
