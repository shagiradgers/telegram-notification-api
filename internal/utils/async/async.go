package async

import "sync"

type Dispatcher struct {
	wg    *sync.WaitGroup
	limit int
	lock  *sync.Mutex

	jobs []job
}

type job func() error

func (d *Dispatcher) AddJob(f job) {

	d.jobs = append(d.jobs, f)
}

func (d *Dispatcher) Run() error {
	var err error

	jobCh := make(chan func() error, d.limit)
	d.wg.Add(d.limit)

	go func() {
		for _, job := range d.jobs {
			jobCh <- job
		}
		close(jobCh)
		d.jobs = make([]job, 0)
	}()

	go func() {
		for i := 0; i < d.limit; i++ {
			go func() {
				defer d.wg.Done()
				for job := range jobCh {
					if localErr := job(); localErr != nil {
						d.lock.Lock()
						err = localErr
						d.lock.Unlock()
					}
				}
			}()
		}
	}()
	d.wg.Wait()
	return err
}

func NewAsyncDispatcher(limit int) *Dispatcher {
	return &Dispatcher{
		limit: limit,
		wg:    &sync.WaitGroup{},
		jobs:  make([]job, 0),
		lock:  &sync.Mutex{},
	}
}
