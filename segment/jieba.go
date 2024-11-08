package segment

import (
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
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

func (sg *Jieba) Pos(text string) (res []Pos) {
	pos := sg.c.Tag(text)
	for _, text := range pos {
		p := gstr.Explode("/", text)
		if len(p) == 1 {
			glog.Info(nil, text, p)
			continue
		}
		res = append(res, Pos{
			Word: p[0],
			Type: p[1],
		})
	}
	return
}

func (sg *Jieba) Ner(text string, ps ...string) (res []Ner) {
	if len(ps) == 0 {
		ps = NerPosFilter
	}
	pos := sg.Pos(text)
	for _, p := range pos {
		// 将 CJK 字符串转为 rune 数组并获取其长度
		if gstr.LenRune(p.Word) == 1 {
			continue
		}

		if p.Type == "" {
			p.Type = "nz"
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
