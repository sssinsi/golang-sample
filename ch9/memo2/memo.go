package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // cacheを保護する
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

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		//2つのクリティカルセッションの間で、幾つかのゴルーチンがf(key)の計算で
		//競合してマップを更新するかもしれない
		memo.mu.Unlock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
