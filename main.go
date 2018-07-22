package main

import (
	"./engine"
	"./zhenai/parser"
	"./scheduler"
	"./persist"
)

func main()  {
    e := engine.ConcurrentEngin{
    	Scheduler:  &scheduler.QueuedScheduler{},
    	WorkerCount: 100,
    	Itemchan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/haerbin",
	//	ParserFunc: parser.ParseCity,
	//})
}


