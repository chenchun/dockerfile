package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"
)

const (
	kb            = 1024
	mb            = kb * 1024
	golandDefault = 10
)

func mockTest() {
	for _, use := range []struct {
		mem uint64
		cpu int
	}{
		{20, 20},
		{30, 40},
		{50, 80},
		{30, 60},
		{10, 10},
	} {
		stress(use.cpu, use.mem, time.Minute)
	}
}

func stress(cpu int, mem uint64, duration time.Duration) {
	mem = mem/1024/1024
	log.Printf("cpu %d0m, mem %dMB", cpu, mem)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		stressMem(mem, ctx)
	}()
	for i := 0; i < cpu/100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stressCpu(100, ctx)
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		stressCpu(cpu%100, ctx)
	}()
	time.Sleep(duration)
	cancel()
	wg.Wait()
	runtime.GC()
}

func stressMem(memMB uint64, ctx context.Context) {
	if memMB < golandDefault {
		return
	} else {
		memMB -= golandDefault
	}
	arr := make([]int, memMB*mb/4)
	arr[0] = 1
	select {
	case <-ctx.Done():
		return
	default:
	}
}

func stressCpu(percentage int, ctx context.Context) {
	// sleep percentage msec in 100 msec
	runDuration := time.Millisecond * time.Duration(percentage)
	for {
		start := time.Now()
		for {
			if time.Now().Sub(start) < runDuration {
				continue
			}
			time.Sleep(time.Duration(100-percentage) * time.Millisecond)
			break
		}
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
