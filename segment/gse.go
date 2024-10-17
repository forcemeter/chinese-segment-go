package segment

import (
	"github.com/go-ego/gse"
	"github.com/gogf/gf/v2/text/gstr"
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

func (sg *Gse) Pos(text string) (res []Pos) {
	pos := sg.c.Pos(text, false)
	for _, p := range pos {
		res = append(res, Pos{
			Word: p.Text,
			Type: p.Pos,
		})
	}
	return
}

func (sg *Gse) Ner(text string, ps ...string) (res []Ner) {
	if len(ps) == 0 {
		ps = NerPosFilter
	}
	pos := sg.Pos(text)
	for _, p := range pos {
		// 将 CJK 字符串转为 rune 数组并获取其长度
		if gstr.LenRune(p.Word) == 1 {
			continue
		}

		// 只需要指定的词性
		if !gstr.InArray(ps, p.Type) {
			continue
		}

		res = append(res, Ner{
			Word: p.Word,
			Type: p.Type,
		})
	}
	return
}
