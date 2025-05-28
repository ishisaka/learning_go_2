package cmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestCreatePersonはCreatePerson関数が期待通りのPersonオブジェクトを生成するかをテストします。
func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  37,
	}
	result := CreatePerson("Dennis", 37)
	// 期待値と結果を比較する
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff) // DateAddedで差違が出る
	}
}

// TestCreatePersonIgnoreDateはCreatePerson関数が正しいNameとAgeを設定するかを検証します。
func TestCreatePersonIgnoreDate(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  37,
	}
	result := CreatePerson("Dennis", 37)
	// DateAddedでは必ず差違が出るのでそこは比較しないように比較関数を用意する
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})
	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
	if result.DateAdded.IsZero() {
		t.Error("DateAdded wasn't assigned")
	}
}
