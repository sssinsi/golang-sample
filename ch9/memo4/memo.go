package memo4

import "sync"

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} //res が設定されたら閉じられる
}

//Funcはmemo化される関数の型です
type Func func(key string) (interface{}, error)

type Memo struct {
	f     Func
	mu    sync.Mutex //cacheを保護する
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()       //ロック取得
	e := memo.cache[key] //entry
	if e == nil {
		//これはkeyに対する最初のリクエスト
		//このゴルーチンは値を計算し、readyの状態をブロードキャストする責任を持つ
		e = &entry{ready: make(chan struct{})} //entryを新規作成
		memo.cache[key] = e                    //keyに新規entryを設定
		memo.mu.Unlock()                       //ロック解放
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) //readyの状態をブロードキャストする

	} else {
		//これはkeyに対する繰り返しのリクエスト
		memo.mu.Unlock() //ロック解放
		<-e.ready        //readyの状態を待つ。チャネルが閉じられるまで待つ
	}
	return e.res.value, e.res.err
}
