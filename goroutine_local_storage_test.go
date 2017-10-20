package gls

import (
	"sync"
	"testing"
	"time"
)

func TestGLS(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		v := Get("key")
		if v != nil {
			t.Fail()
		}

		Set("k1", "1")

		time.Sleep(200 * time.Millisecond)

		v1 := Get("k1")
		if v1 == nil {
			t.Fail()
		}

		v2 := Get("k2")
		if v2 != nil {
			t.Fail()
		}

		wg.Done()
	}()

	go func() {
		v := Get("key")
		if v != nil {
			t.Fail()
		}

		Set("k2", "2")

		time.Sleep(200 * time.Millisecond)

		v1 := Get("k1")
		if v1 != nil {
			t.Fail()
		}

		v2 := Get("k2")
		if v2 == nil {
			t.Fail()
		}

		wg.Done()
	}()

	wg.Wait()

	Clear()
}
