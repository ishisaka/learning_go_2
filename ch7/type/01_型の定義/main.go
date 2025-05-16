package main

// 構造体の定義
type Person struct {
	LastName  string
	FirstName string
	Age       int
}

// さまざまな型の定義
type Score int                    // intを基底型としてScoreを定義
type Converter func(string) Score // stringを引数、Scoreを戻り値とする関数型Converterを定義
type TeamScores map[string]Score  // stringをScoreにマップするmap、TeamScoresを定義
