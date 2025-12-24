package endpoint

import (
	"sync"
	"sync/atomic"
	"time"
)

// LatencySlideWindows 延迟滑动窗口
type LatencySlideWindows struct {
	counters []*LatencyCounter
	index    int32
	stopCh   chan struct{}
	stopOnce sync.Once
}

// LatencyCounter 延迟计数器
type LatencyCounter struct {
	count int64
	spend int64
}

// NewLatencySlideWindows 创建延迟滑动窗口
func NewLatencySlideWindows(bucketSize int, windowTimePerBucket int64) *LatencySlideWindows {
	counters := make([]*LatencyCounter, bucketSize)
	for i := 0; i < bucketSize; i++ {
		counters[i] = &LatencyCounter{}
	}

	windows := &LatencySlideWindows{
		counters: counters,
		index:    0,
		stopCh:   make(chan struct{}),
	}

	// 启动滑动定时器
	go windows.startSliding(windowTimePerBucket)

	return windows
}

// startSliding 开始滑动
func (w *LatencySlideWindows) startSliding(windowTimePerBucket int64) {
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
func (w *LatencySlideWindows) slideToNextBucket() {
	currentIndex := atomic.LoadInt32(&w.index)
	nextIndex := (currentIndex + 1) % int32(len(w.counters))

	// 重置下一个桶
	w.counters[nextIndex].reset()

	// 更新索引
	atomic.StoreInt32(&w.index, nextIndex)
}

// Increment 增加延迟统计
func (w *LatencySlideWindows) Increment(spendMs int64) {
	currentIndex := atomic.LoadInt32(&w.index)
	w.counters[currentIndex].increment(spendMs)
}

// GetLatency 获取平均延迟
func (w *LatencySlideWindows) GetLatency() float64 {
	var sum, count int64

	for _, counter := range w.counters {
		c, s := counter.getCounters()
		count += c
		sum += s
	}

	if count == 0 {
		return 0
	}

	return float64(sum) / float64(count)
}

// Stop 停止滑动窗口
func (w *LatencySlideWindows) Stop() {
	w.stopOnce.Do(func() {
		close(w.stopCh)
	})
}

// increment 增加计数和耗时
func (c *LatencyCounter) increment(spendMs int64) {
	atomic.AddInt64(&c.count, 1)
	atomic.AddInt64(&c.spend, spendMs)
}

// reset 重置计数器
func (c *LatencyCounter) reset() {
	atomic.StoreInt64(&c.count, 0)
	atomic.StoreInt64(&c.spend, 0)
}

// getCounters 获取当前计数值
func (c *LatencyCounter) getCounters() (count, spend int64) {
	return atomic.LoadInt64(&c.count), atomic.LoadInt64(&c.spend)
}
