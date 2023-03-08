package channel

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	 "time"
)

const SizeBuffer int = 10
const bufferDrainInterval time.Duration = 10 * time.Second


func ReadConsole(exit chan bool, c <-chan int) <-chan int {
	dst := make(chan int)
	// run thread read console
	fmt.Println("конвейер для обработки чисел, \nисточник -- консоль\nприемник -- кольцевой буффер \nдля выхода наберите \"exit\"")
	Log(ReadConsole, "Run gorutine for read console")
	go func() {
		defer close(exit)
		scanner := bufio.NewScanner(os.Stdin)
		var data string
		for {
			scanner.Scan()
			data = scanner.Text()
			if strings.EqualFold(data, "exit") {
				Log(ReadConsole, "get event exit, close app..... ")
				fmt.Println("Программа завершила работу!")
				return
			}
			Log(ReadConsole, "Read console input data: "+ data)
			i, err := strconv.Atoi(data)
			if err != nil {
				fmt.Println("Программа обрабатывает только целые числа!")
				Log(ReadConsole, "Skip, this not number!")
				continue
			}
			Log(ReadConsole,"Sussed  "+ data+" ,the next stage")
			dst <- i
		}
	}()
	return dst
}

func FilterNegative(exit chan bool, c <-chan int) <-chan int {
	dst := make(chan int)
	Log(FilterNegative, "Run gorutine for filter negative")
	go func() {
		for {
			select {
			case d := <-c:
				if d >= 0 {
					select {
					case dst <- d:
					case <-exit:
						return
					}
				} else {
					fmt.Println("Число отрицательное!")
					Log(FilterNegative, "This data is negative, skip")
				}
			case <-exit:
				return
			}
		}
	}()
	return dst
}

func FilterSpecial(exit chan bool, c <-chan int) <-chan int {
	dst := make(chan int)
	Log(FilterSpecial, "Run gorutine for filter divisions by 3")
	go func() {
		for {
			select {
			case d := <-c:
				if d%3 == 0 {
					select {
					case dst <- d:
					case <-exit:
						return
					}
				} else {
					fmt.Println("Число не делится на 3!")
					Log(FilterSpecial, "data not divisions by 3, skip...")
				}

			case <-exit:
				return
			}
		}
	}()
	return dst
}

func ReadBuffer(exit chan bool, c <-chan int) <-chan int{
b := CreateBuffer(SizeBuffer)
dst := make (chan int )
Log(ReadBuffer, "Run gorutine read and write ring buffer")
go func (){
	for{
		select {
		case d := <- c :  b.Push(d)
		Log(ReadBuffer, "add  ring buffer " + fmt.Sprintf("%d",d) )
		case <-exit :return
		}
	}
}()


// читаем из буффера
go func() {
	for {
		select {
		case <-time.After(bufferDrainInterval):
			bufferData := b.Get()
			// Если в кольцевом буфере что-то есть - 
			// выводим
			// содержимое построчно
			Log(ReadBuffer, "read  ring buffer " )
			if bufferData != nil {
				for _, data := range bufferData {
					select {

					case dst <- data:						
					case <-exit:
						return
					}
				}
			}
		case <-exit:
			return
		}
	}
}()	
return dst 
}


func WriteConsole(exit chan bool, c <-chan int) {
	var wg sync.WaitGroup
	wg.Add(1)	
	go func() {

		for {
			select {
			case d := <-c:
				Log(WriteConsole, "get data  "+  fmt.Sprintf("%d",d) )
				fmt.Printf("Обработаны данные: %d\n", d)				
			case <-exit:
				defer wg.Done()
				return
			}
		}
	}()
	wg.Wait()
}
