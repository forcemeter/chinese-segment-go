package segment

import (
	"github.com/yanyiwu/gojieba"
)

type Jieba struct {
	c *gojieba.Jieba
}

func (sg *Jieba) Init() {
	sg.c = gojieba.NewJieba()
}

func (sg *Jieba) LoadDictWords(words []string) error {
	for _, word := range words {
		if word != "" && len(word) > 0 {
			sg.c.AddWord(word)
		}
	}

	return nil
}

func (sg *Jieba) Cut(text string) (res Tokens) {
	res.Words = sg.c.Cut(text, true)
	return
}

func (sg *Jieba) Tag(text string) (res Tokens) {
	res.Words = sg.c.Tag(text)
	return
}
