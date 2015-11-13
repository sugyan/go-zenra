package zenra

import (
	"github.com/ikawaha/kagome/tokenizer"
)

// ZENRA phrase
const ZENRA = "全裸で"

// Zenrizer type
type Zenrizer struct {
	tokenizer tokenizer.Tokenizer
}

// NewZenrizer returns Zenrizer instance
func NewZenrizer() *Zenrizer {
	return &Zenrizer{
		tokenizer: tokenizer.New(),
	}
}

// Zenrize returns zenrized text
func (z *Zenrizer) Zenrize(input string) (result string) {
	var verb = false
	tokens := z.tokenizer.Tokenize(input)
	for i := range tokens {
		// 末尾からさかのぼる
		token := tokens[len(tokens)-i-1]
		if token.Class == tokenizer.DUMMY {
			continue
		}
		features := token.Features()
		// フラグが立っていれば「全裸で」を挿入
		if verb {
			insert := true
			// 名詞／副詞／動詞のときはまだ挿入しない
			// また、連用形の動詞→助(動)詞の場合も挿入しない
			if features[0] == "名詞" || features[0] == "副詞" || features[0] == "動詞" {
				insert = false
			} else if (features[0] == "助詞" || features[0] == "助動詞") && i < len(tokens)-1 {
				prev := tokens[len(tokens)-i-2].Features()
				if prev[5] == "連用形" || prev[5] == "連用タ接続" {
					insert = false
				}
			}
			if insert {
				result = ZENRA + result
				verb = false
			}
		}
		// 出力の連結
		result = token.Surface + result
		if features[0] == "動詞" {
			verb = true
		}
	}
	if verb {
		result = ZENRA + result
	}
	return result
}
