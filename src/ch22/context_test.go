package ch22_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func watch1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("test1 will done")
			return
		default:
			if value := ctx.Value("minitest"); value != nil {
				fmt.Printf("ctx.value %v\n", value)
			}
			time.Sleep(time.Millisecond * 10)
		}
	}

}

func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("test2 will done")
			return
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}
func watch3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("test3 will done")
			return
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
}
func TestCancelContext(t *testing.T) {
	//父context
	parentCtx, cancel := context.WithCancel(context.Background())
	go watch1(parentCtx)

	//子的context
	childCtx, _ := context.WithCancel(parentCtx)
	go watch2(childCtx)

	//孙子context
	grandsonCtx, _ := context.WithCancel(childCtx)
	go watch3(grandsonCtx)

	time.Sleep(time.Millisecond * 100)
	cancel()
	time.Sleep(time.Millisecond * 20)
}

func TestTimeoutCancel(t *testing.T) {

	//cancel的时间会置为最大的 1 5  10
	parentctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	go watch1(parentctx)

	childCtx, _ := context.WithTimeout(parentctx, time.Second*5)
	go watch2(childCtx)

	grandsonCtx, _ := context.WithTimeout(childCtx, time.Second*8)
	go watch3(grandsonCtx)

	time.Sleep(time.Second * 8)
}

func TestWithValueCtx(t *testing.T) {
	ctx := context.WithValue(context.Background(), "minitest", "this is test value")
	go watch1(ctx)

	time.Sleep(time.Microsecond * 10)
}

func TestWithDeadline(t *testing.T) {
	//withDeadline与WithTimeout类似,只是入参不同，具有定时功能
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	go watch1(ctx)

	time.Sleep(time.Second * 4)
}
