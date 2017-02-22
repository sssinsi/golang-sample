package memo5

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

//requestは、FUncがkeyへ適用されることを要求するメッセージです
type request struct {
	key      string
	response chan<- result //クライアントは結果を１つだけ望んでいます
}
type Memo struct {
	requests chan request
}

// Newはfのメモ化を返します。クライアントは後でCloseを呼びださなければなりません
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) //call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	//evaluate the function
	e.res.value, e.res.err = f(key)
	//broadcast the ready condition
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	//wait for the ready condition
	<-e.ready
	//send the result to the client
	response <- e.res
}
