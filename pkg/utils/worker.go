package utils

type Job interface {
	Task()
}

type Pool struct {
	workerChan   chan Job
	shutdownChan chan bool
	maxWork      int
}

func NewWorkPool(maxWork, maxQueue int) *Pool {
	p := &Pool{
		workerChan:   make(chan Job, maxQueue),
		shutdownChan: make(chan bool),
		maxWork:      maxWork,
	}
	return p
}

func (p *Pool) Start() {
	for i := 0; i < p.maxWork; i++ {
		go func() {
			for {
				select {
				case worker := <-p.workerChan:
					worker.Task()
				case <-p.shutdownChan:
					close(p.workerChan)
					close(p.shutdownChan)
					return
				}
			}
		}()
	}
}

func (p *Pool) Run(j Job) {
	p.workerChan <- j
}

func (p *Pool) Shutdown() {
	p.shutdownChan <- true
}
