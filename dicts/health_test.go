package dicts

import "testing"

func TestHealthDict(t *testing.T) {
	words := LoadHealthDict()
	t.Log(words[0:100])
}
