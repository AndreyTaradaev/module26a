package channel
import(
	//"sync"
	"os"
	"fmt"
)

type stageFunc func(done chan bool, source <- chan int) <-chan int

type ArrayStage struct {
	Funcs []stageFunc	
}

func CreateStage(ar ... stageFunc) *ArrayStage {
	for _, v := range os.Args {
		fmt.Println(v)		
	}
	ret := new(ArrayStage)
	ret.Funcs = ar	
	return ret
}


func (a *ArrayStage) Run(exit chan bool, source <- chan int) <-chan int {
	var c <-chan int = source
	for _, f := range a.Funcs {
		c = f(exit, c)
	}
	return c
}
