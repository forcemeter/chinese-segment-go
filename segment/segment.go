package segment

type Segment interface {
	Init()
	LoadDictWords(words []string) error
	Cut(text string) Tokens
	Tag(text string) Tokens
}

type Tokens struct {
	Words []string `json:"words"`
}

type Ner struct {
	Word string
	Type string
}
