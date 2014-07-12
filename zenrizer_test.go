package zenra

import (
	"testing"
)

func TestZenrize(t *testing.T) {
	var (
		result string
		err    error
	)

	// 城の崎にて
	result, err = Zenrize("山の手線に跳ね飛ばされて怪我をした、その後養生に、一人で但馬の城崎温泉へ出掛けた。")
	if err != nil {
		t.Error(err)
	}
	if result != "山の手線に全裸で跳ね飛ばされて怪我を全裸でした、その後養生に、一人で但馬の城崎温泉へ全裸で出掛けた。" {
		t.Errorf("城の崎にて: %s", result)
	}

	// ジョジョ
	result, err = Zenrize("お前は今まで食ったパンの枚数を覚えているのか？")
	if err != nil {
		t.Error(err)
	}
	if result != "お前は今まで全裸で食ったパンの枚数を全裸で覚えているのか？" {
		t.Errorf("ジョジョ: %s", result)
	}

	// スラムダンク
	result, err = Zenrize("あきらめたらそこで試合終了だよ")
	if err != nil {
		t.Error(err)
	}
	if result != "全裸であきらめたらそこで試合終了だよ" {
		t.Errorf("スラムダンク: %s", result)
	}
}
