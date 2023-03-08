package channel

//"sync"

type stageFunc func(done chan bool, source <-chan int) <-chan int

type ArrayStage struct {
	Funcs []stageFunc
}

func CreateStage(ar ...stageFunc) *ArrayStage {
	Log(CreateStage,"Create stage implement")
	ret := new(ArrayStage)
	ret.Funcs = ar
	return ret
}

func (a *ArrayStage) Run(exit chan bool, source <-chan int) <-chan int {
	Log(a.Run," Run pipe")
	var c <-chan int = source
	for _, f := range a.Funcs {
		c = f(exit, c)
	}
	return c
}
