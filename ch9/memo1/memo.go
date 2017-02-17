// memoパッケージはFunc型の関数の
//並行的に安全でないmemoかを提供します
package memo

//MemoはFunc呼び差しの結果をキャッシュする
type Memo struct {
	f     Func
	cache map[string]result
}

//Funcはmemoかされる関数の型です
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

//注意：並行的に安全ではない
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
