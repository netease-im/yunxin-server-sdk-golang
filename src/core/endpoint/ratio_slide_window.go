package endpoint

import (
	"sync"
	"sync/atomic"
	"time"
)

// RatioSlideWindows 比率滑动窗口
// Created by caojiajun on 2024/12/9
type RatioSlideWindows struct {
	counters []*RatioCounter
	index    int32
	stopCh   chan struct{}
	stopOnce sync.Once
}

// RatioCounter 比率计数器
type RatioCounter struct {
	success int64
	fail    int64
	mu      sync.RWMutex
}

// NewRatioSlideWindows 创建比率滑动窗口
func NewRatioSlideWindows(bucketSize int, windowTimePerBucket int64) *RatioSlideWindows {
	counters := make([]*RatioCounter, bucketSize)
	for i := 0; i < bucketSize; i++ {
		counters[i] = &RatioCounter{}
	}

	windows := &RatioSlideWindows{
		counters: counters,
		index:    0,
		stopCh:   make(chan struct{}),
	}

	// 启动滑动定时器
	go windows.startSliding(windowTimePerBucket)

	return windows
}

// startSliding 开始滑动
func (w *RatioSlideWindows) startSliding(windowTimePerBucket int64) {
	ticker := time.NewTicker(time.Duration(windowTimePerBucket) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.slideToNextBucket()
		case <-w.stopCh:
			return
		}
	}
}

// slideToNextBucket 滑动到下一个桶
func (w *RatioSlideWindows) slideToNextBucket() {
	currentIndex := atomic.LoadInt32(&w.index)
	nextIndex := (currentIndex + 1) % int32(len(w.counters))

	// 重置下一个桶
	w.counters[nextIndex].reset()

	// 更新索引
	atomic.StoreInt32(&w.index, nextIndex)
}

// IncrementSuccess 增加成功计数
func (w *RatioSlideWindows) IncrementSuccess() {
	currentIndex := atomic.LoadInt32(&w.index)
	w.counters[currentIndex].incrementSuccess()
}

// IncrementFail 增加失败计数
func (w *RatioSlideWindows) IncrementFail() {
	currentIndex := atomic.LoadInt32(&w.index)
	w.counters[currentIndex].incrementFail()
}

// GetSuccessRatio 获取成功率
func (w *RatioSlideWindows) GetSuccessRatio() float64 {
	var success, fail int64

	for _, counter := range w.counters {
		s, f := counter.getCounters()
		success += s
		fail += f
	}

	total := success + fail
	if total == 0 {
		return 1.0
	}

	return float64(success) / float64(total)
}

// Stop 停止滑动窗口
func (w *RatioSlideWindows) Stop() {
	w.stopOnce.Do(func() {
		close(w.stopCh)
	})
}

// incrementSuccess 增加成功计数
func (c *RatioCounter) incrementSuccess() {
	atomic.AddInt64(&c.success, 1)
}

// incrementFail 增加失败计数
func (c *RatioCounter) incrementFail() {
	atomic.AddInt64(&c.fail, 1)
}

// reset 重置计数器
func (c *RatioCounter) reset() {
	atomic.StoreInt64(&c.success, 0)
	atomic.StoreInt64(&c.fail, 0)
}

// getCounters 获取当前计数值
func (c *RatioCounter) getCounters() (success, fail int64) {
	return atomic.LoadInt64(&c.success), atomic.LoadInt64(&c.fail)
}
