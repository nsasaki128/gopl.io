package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // cacheを保護する
	cache map[string]result
}

// Func はメモ化される関数の型です。
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Getは並行的に安全です
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// 二つのクリティカルセクションの間で、いくつかのゴルーチンが
		// f(key)の計算で競合してマップを更新するかもしれません。
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
