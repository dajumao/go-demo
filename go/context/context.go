package context

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Context interface {
	Deadline() (dealine time.Time,ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
//Canceled是上下文返回的错误。取消上下文时出错。
var Canceled = errors.New("context canceled")
//当上下文的截止日期已过时,返回错误
var DeadlneExceeded error = deadlineExceededError{}
type deadlineExceededError struct {}

func (deadlineExceededError)Error() string  {
	return "context deadline exceed"
}

func (deadlineExceededError)Timeout() bool  {
	return true
}

func (deadlineExceededError)Temporary() bool  {
	return true
}

//EmptyCtx从不取消，没有值，也没有截止时间。它不是
//结构，因为此类型的变量必须具有不同的地址
type emptyCtx int

func (*emptyCtx)Deadline()(deadline time.Time,ok bool)  {
	return 
}

func (*emptyCtx)Done() <-chan struct{}{
	return nil
}

func (*emptyCtx)Err() error  {
	return nil
}

func (*emptyCtx)Value(key interface{}) interface{}  {
	return nil
}

func (e *emptyCtx)String() string  {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unkonwn empty Context"
}

var (
	background = new(emptyCtx)
	todo	   = new(emptyCtx)
)

//background返回非零的空上下文。从来没有取消过，没有
//值，并且没有截止时间。它通常由主功能使用，
//初始化和测试，并作为传入的顶级上下文
//请求。
func Background() Context {
	return background
}

//TODO返回非零的空上下文。代码应在以下情况下使用context.todo
//不清楚要使用哪个上下文，或者它还不可用（因为
//尚未扩展周围函数以接受上下文
//参数）。
func TODO() Context  {
	return todo
}

//cancelfunc通知操作放弃其工作。
//CancelFunc不等待工作停止。
//第一次调用后，对cancelfunc的后续调用不做任何操作。
type CancelFunc func()

func WithCancel(parent Context) (ctx Context,cancel CancelFunc)  {
	c := newCancelCtx(parent)
	propagateCancel(parent,&c)
	return &c, func() {
		c.cancel(true,Canceled)
	}
}

//newCancelCTX返回初始化的CancelCTX。
func newCancelCtx(parent Context) cancelCtx  {
	return cancelCtx{Context:parent}
}

func propagateCancel(parent Context,child canceler)  {
	if parent.Done() == nil {
		return
	}
	if p,ok := parentCancelCtx(parent);ok {
		p.mu.Lock()
		if p.err != nil {
			child.cancel(false,p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false,parent.Err())
			case <-child.Done():

			}
		}()
	}
}

func parentCancelCtx(parent Context)(*cancelCtx,bool)  {
	for {
		switch c:=parent.(type) {
		case *cancelCtx:
			return c,true
		case *timerCtx:
			return &c.cancelCtx,true
		case *valueCtx:
			parent = c.Context
		default:
			return nil,false
		}
	}
}

func removeChild(parent Context,child canceler)  {
	p,ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children,child)
	}
	p.mu.Unlock()
}

type canceler interface {
	cancel(removeFromParent bool,err error)
	Done() <-chan struct{}
}

var closedchan = make(chan struct{})

func init()  {
	close(closedchan)
}

type cancelCtx struct {
	Context
	mu 			sync.Mutex
	done		chan struct{}
	children 	map[canceler]struct{}
	err			error	//第一次取消调用时设置为非零
}

func (c *cancelCtx)Done() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	c.mu.Unlock()
	return d
}

func (c *cancelCtx)Err() error  {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func (c *cancelCtx)String() string  {
	return fmt.Sprintf("%v.WithCancel",c.Context)
}

func (c *cancelCtx)cancel(removeFromParent bool,err error)  {
	if err == nil {
		panic("context: internal error :missing cancel ")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return
	}
	c.err = err

	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}
	for child := range c.children{
		child.cancel(false,err)
	}
	c.children = nil
	c.mu.Unlock()
	if removeFromParent {
		removeChild(c.Context, c)
	}
}

func WithDeadline(parent Context,d time.Time) (Context,CancelFunc) {
	if cur,ok := parent.Deadline();ok&&cur.Before(d) {
		return WithCancel(parent)
	}
	c := &timerCtx{
		cancelCtx:newCancelCtx(parent),
		deadline:d,
	}
	propagateCancel(parent,c)
	dur := time.Until(d)
	if dur <= 0 {
		c.cancel(true,DeadlneExceeded)
		return c, func() {
			c.cancel(false,Canceled)
		}
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		c.timer = time.AfterFunc(dur, func() {
			c.cancel(true,DeadlneExceeded)
		})
	}
	return c, func() {
		c.cancel(true,Canceled)
	}
}


//timerctx包含计时器和截止时间。它将cancelctx嵌入到
//实现done和err。它通过停止计时器来实现取消，然后
//委托cancelctx.cancel。
type timerCtx struct {
	cancelCtx
	timer 		*time.Timer
	deadline	time.Time
}

func (c *timerCtx)Deadline()(deadline time.Time,ok bool)  {
	return c.deadline,true
}

func (c *timerCtx)String() string  {
	return fmt.Sprintf("%v.WithDeadline(%s [%s])", c.cancelCtx.Context, c.deadline, time.Until(c.deadline))
}

func (c *timerCtx)cancel(removeFromParent bool,err error)  {
	c.cancelCtx.cancel(false,err)
	if removeFromParent {
		removeChild(c.cancelCtx.Context,c)
	}
	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}

func WithTimeout(parent Context,timeout time.Duration)(Context,CancelFunc)  {
	return WithDeadline(parent,time.Now().Add(timeout))
}

func WithValue(parent Context,key,val interface{}) Context  {
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return  &valueCtx{parent,key,val}
}

type valueCtx struct {
	Context
	key,val	interface{}
}

func (c *valueCtx)String() string  {
	return fmt.Sprintf("%v.WithValue(%#v, %#v)", c.Context, c.key, c.val)
}

func (c *valueCtx)Value(key interface{}) interface{}  {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}

