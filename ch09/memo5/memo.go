package memo

// Func はメモ化される関数の型です。
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // resが設定されたら閉じられる
}

// requestは、Funcがkeyへ適用されることを要求するメッセージです。
type request struct {
	key      string
	response chan<- result //クライアントは結果を一つだけ望んでいます。
}

type Memo struct {
	requests chan request
}

// Newはfのメモ化を返します。クライアントは後でCloseを呼び出さなければなりません。
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Getは並行的に安全です
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// これは、このkeyに対する最初のリクエスト。
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) //call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	//関数を評価する。
	e.res.value, e.res.err = f(key)
	//用意ができたことをブロードキャストする。
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	//用意ができるのを待つ
	<-e.ready
	//結果をクライアントへ送信する。
	response <- e.res
}
