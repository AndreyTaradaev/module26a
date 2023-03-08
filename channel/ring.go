package channel

import (
	"fmt"
	"sync"
	//"errors"
)

const ErrEmply = "Emply array"

type Buffer struct {
	buf    []int
	m      sync.Mutex
	offset int
	size int
}

// ctor 
func CreateBuffer(sizearray int) *Buffer {
	b := new(Buffer)
	b.size = sizearray
	b.buf = make([]int, 0, sizearray)
	b.offset = -1
	b.m = sync.Mutex{}	
	return b
}

func (b *Buffer) Push(a int) {
	b.m.Lock()
	defer b.m.Unlock()
	if b.offset < b.size-1 {
		b.offset++
		b.buf = append(b.buf, a) 
	} else {
		sl := b.buf[1:]
		b.buf = sl
		b.buf = append(sl, a)
	}
}

func (b *Buffer) Pop() (int, error) {
	b.m.Lock()
	defer b.m.Unlock()
	if b.offset < 0 {
		return 0, fmt.Errorf(ErrEmply)
	}
	ret := b.buf[0]
	l := len(b.buf)
	sl := b.buf[1:]
	
	b.buf = make([]int , l,l)
	 copy(b.buf,sl)	
	 b.offset--
	 return  ret,nil
}

func (b *Buffer ) Get() []int {
	if(b.offset == -1) {
		return nil
	}
	var ret = make([]int,b.offset+1)

	copy(ret,b.buf)
	b.buf = make([]int,0,b.size)
	b.offset = -1
	return ret
}


func (b *Buffer ) GetBuffer() []int {	
	ret := make ([]int,len(b.buf), len(b.buf)) 
	copy(ret,b.buf)
	return  ret 
}
