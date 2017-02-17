package memo_test

import (
	"testing"

	memo "github.com/sssinsi/golang-sample/ch9/memo1"
	"github.com/sssinsi/golang-sample/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
