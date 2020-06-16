package pipline

import (
	"fmt"
	"sync"
)

func (pp *Pipline) Run() {
	for _, v := range pp.Pper.Do() {
		t, _ := pp.Tfer.Do(v)
		g := pp.Geer.Do(t)
		pp.Eter.Write(g)
	}
}

func (pp *Pipline) RunAsync() {
	pumpCh := pp.Pper.DoAsync()
	transfCh1 := pp.Tfer.DoAsync(pumpCh)
	transfCh2 := pp.Tfer.DoAsync(pumpCh)
	GenCh1 := pp.Geer.DoAsync(transfCh1)
	GenCh2 := pp.Geer.DoAsync(transfCh2)
	GenCh3 := pp.Geer.DoAsync(transfCh1)
	GenCh4 := pp.Geer.DoAsync(transfCh2)
	ExpCh1 := pp.Eter.WriteAsync(GenCh1)
	ExpCh2 := pp.Eter.WriteAsync(GenCh2)
	ExpCh3 := pp.Eter.WriteAsync(GenCh3)
	ExpCh4 := pp.Eter.WriteAsync(GenCh4)
	for s := range Merge(ExpCh1, ExpCh2, ExpCh3, ExpCh4) {
		println(s)
	}
}

// 输出结果合成器
func Merge(chs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	collect := func(in <-chan string) {
		defer wg.Done()
		for n := range in {
			out <- fmt.Sprintf("%s has built successfully", n)
		}
	}
	wg.Add(len(chs))
	// FAN-IN
	for _, c := range chs {
		go collect(c)
	}
	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
