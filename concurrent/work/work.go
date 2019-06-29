package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGorutines int) *Pool  {
	p := Pool{
		work:make(chan Worker),
	}
	p.wg.Add(maxGorutines)
	for i := 0; i < maxGorutines; i++  {
		go func() {
			for w := range p.work{
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool)Run(w Worker)  {
	p.work <- w
}

func (p *Pool)Shutdown()  {
	close(p.work)
	p.wg.Wait()
}
