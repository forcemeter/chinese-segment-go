package segment

type Segment interface {
	Init()
	LoadDictWords(words []string) error
	Cut(text string) Tokens
	Pos(text string) []Pos
	Ner(text string, ps ...string) []Ner
}

type Tokens struct {
	Words []string `json:"words"`
}

// Ner Named Entity Recognition
type Ner struct {
	Word string
	Type string
}

// Pos Part of Speech
type Pos struct {
	Word string
	Type string
}

// NerPosFilter 通过词性简单实现实体词提取
// 人名（nr）、地名（ns）、机构名（nt）、专有名词（nz）、动词（v）
var NerPosFilter = []string{"n", "ns", "nr", "nt", "v"}
