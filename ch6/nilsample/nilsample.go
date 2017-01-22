package nilsample

//IntLint は整数のリンクリストです。
//nilの*IntLintは空リストを表します。
//レシーバ値としてnilを返すmethodを持つ型を定義する場合には、箒のようにドキュメンテーションコメントで明示的のそのことを指摘しておくのが良いです。
type IntList struct {
	Value int
	Tail  *IntList
}

//Sum はリンスと要素の合計値を返します。
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
