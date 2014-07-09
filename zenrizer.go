package zenra

import (
	"github.com/ikawaha/kagome/tokenizer"
	"log"
)

// ZENRA phrase
const ZENRA = "全裸で"

// Zenrize returns zenrized text
func Zenrize(input string) (result string) {
	t := tokenizer.NewTokenizer()
	morphs, err := t.Tokenize(input)
	if err != nil {
		log.Fatal(err)
	}

	var verb = false
	result = ""
	for i := range morphs {
		m := morphs[len(morphs)-i-1]
		c, err := m.Content()
		if err != nil {
			log.Fatal(err)
		}
		if verb {
			insert := true
			if c.Pos == "名詞" || c.Pos == "副詞" || c.Pos == "動詞" {
				insert = false
			} else if (c.Pos == "助詞" || c.Pos == "助動詞") && i < len(morphs)-1 {
				prev, err := morphs[len(morphs)-i-2].Content()
				if err != nil {
					log.Fatal(err)
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
		if c.Pos == "動詞" {
			verb = true
		}
		if m.Class != tokenizer.DUMMY {
			result = m.Surface + result
		}
	}
	if verb {
		result = ZENRA + result
	}
	return result
}
