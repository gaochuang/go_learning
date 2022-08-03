package object_pool_worker

import (
	"errors"
	"fmt"
	"time"
)

type Object struct {
	nodeId int
}

type ObjectPool struct {
	bufferChanObj chan *Object
}

func NewPool(nunObj int) *ObjectPool {
	objpool := ObjectPool{}
	objpool.bufferChanObj = make(chan *Object, numOfObj)
	for i := 0; i < nunObj; i++ {
		objpool.bufferChanObj <- &Object{}
	}
	return &objpool
}

func (p *ObjectPool) GetObj(timeout time.Duration) (*Object, error) {
	select {
	case ret <- p.bufferChanObj:
		return ret, nil
	case <-time.After(timeout):
		fmt.Println("get pool timeout")
		return nil, errors.New("get pool timeout")

	}
}

func (p *ObjectPool) ReleaseObj(obj chan *Object) error {
	select {
	case p.bufferChanObj <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
