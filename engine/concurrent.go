package engine

type ConcurrentEngin struct {
	Scheduler Scheduler
	WorkerCount int
	Itemchan chan interface{}
}
type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
	WorkerReady (chan Request)
	Run()
}
func (e *ConcurrentEngin) Run(seeds ...Request)  {

	out := make(chan ParseResult)

	e.Scheduler.Run()

    for i := 0 ; i < e.WorkerCount; i++ {
       createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <- out

		for _, item := range result.Items {
			go func() {
				e.Itemchan <- item
			}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}


func createWorker(out chan  ParseResult, s Scheduler)  {
	in := make(chan  Request)
   go func() {
   	for {
   		// tell scheduler i`m ready
   		s.WorkerReady(in)
   		request := <- in
        result, err := worker(request)
        if err != nil {
        	continue
		}

		out <- result

	}
   }()
}