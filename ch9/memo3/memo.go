package memo

import (
	"sync"
)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} //resが設定されたら閉じられる
}

//Funcはmemoかされる関数の型です
type Func func(key string) (interface{}, error)

type Memo struct {
	f     Func
	mu    sync.Mutex //cacheを保護
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		//これはkeyに対する最初のリクエストです。
		//このゴルーチンは値を計算し、readyの状態をブロードキャストする責任を持ちます
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) //readyの状態をブロードキャストする
	} else {
		//これはkeyに対する繰り返しのリクエストです
		memo.mu.Unlock()
		<-e.ready //readyの状態を待ちます。
	}
	return e.res.value, e.res.err
}
