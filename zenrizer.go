package zenra

import (
	"github.com/ikawaha/kagome/dic"
	"github.com/ikawaha/kagome/tokenizer"
)

// ZENRA phrase
const ZENRA = "全裸で"

// Zenrize returns zenrized text
func Zenrize(input string) (result string, err error) {
	t := tokenizer.NewTokenizer()
	morphs, err := t.Tokenize(input)
	if err != nil {
		return
	}

	var verb = false
	result = ""
	for i := range morphs {
		var (
			curr dic.Content
			prev dic.Content
		)
		morph := morphs[len(morphs)-i-1]
		curr, err = morph.Content()
		if err != nil {
			return
		}
		if verb {
			insert := true
			if curr.Pos == "名詞" || curr.Pos == "副詞" || curr.Pos == "動詞" {
				insert = false
			} else if (curr.Pos == "助詞" || curr.Pos == "助動詞") && i < len(morphs)-1 {
				prev, err = morphs[len(morphs)-i-2].Content()
				if err != nil {
					return
				}
				if prev.Katuyoukei == "連用形" {
					insert = false
				}
			}
			if insert {
				result = ZENRA + result
				verb = false
			}
		}
		if curr.Pos == "動詞" {
			verb = true
		}
		if morph.Class != tokenizer.DUMMY {
			result = morph.Surface + result
		}
	}
	if verb {
		result = ZENRA + result
	}
	return result, nil
}
