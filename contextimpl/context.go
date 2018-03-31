package contextimpl

import (
	"errors"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int
    func (ctx emptyCtx) Deadline() (deadline time.Time, ok bool) {
        return 
    }
	func (ctx emptyCtx) Done() <-chan struct{} {
        return nil
    }
	func (ctx emptyCtx) Err() error {
        return nil
    }
	func (ctx emptyCtx) Value(key interface{}) interface{} {
        return nil
    }

    var (

        background = new(emptyCtx)
        todo = new(emptyCtx)
    )
func Background() Context {

	return background
}

func TODO() Context {
	return todo
}


type CancelFunc func()

func WithCancel(parent Context) (Context, CancelFunc) {
    var Canceled =errors.New("context canceled")
}