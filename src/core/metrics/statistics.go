package metrics

import (
	"math"
	"sync/atomic"
)

// Statistics 统计信息收集器
type Statistics struct {
	count      int64     // 请求次数计数器
	sum        int64     // 耗时总和
	maxValue   *MaxValue // 最大值
	distribute []int64   // 分布数组
	exact      bool      // 是否精确模式
}

// MaxValue 最大值管理器
type MaxValue struct {
	value int64
}

// NewStatistics 创建统计信息收集器
func NewStatistics() *Statistics {
	return NewStatisticsWithOptions(false, -1)
}

// NewStatisticsWithOptions 创建指定选项的统计信息收集器
func NewStatisticsWithOptions(exact bool, expectMaxValue int) *Statistics {
	var distributeLength int
	if exact {
		distributeLength = expectMaxValue
	} else {
		distributeLength = 256
	}

	distribute := make([]int64, distributeLength)

	return &Statistics{
		count:      0,
		sum:        0,
		maxValue:   &MaxValue{value: 0},
		distribute: distribute,
		exact:      exact,
	}
}

// Update 更新统计数据
func (s *Statistics) Update(value int64) {
	atomic.AddInt64(&s.count, 1)
	atomic.AddInt64(&s.sum, value)
	s.maxValue.Update(value)

	if value < 0 {
		return
	}

	var index int
	if s.exact {
		if value >= int64(len(s.distribute)) {
			index = len(s.distribute) - 1
		} else {
			index = int(value)
		}
	} else {
		index = s.getDistributeIndex(value)
	}

	if index >= 0 && index < len(s.distribute) {
		atomic.AddInt64(&s.distribute[index], 1)
	}
}

// getDistributeIndex 获取分布索引
func (s *Statistics) getDistributeIndex(value int64) int {
	/**
	 * 0-100 1ms 100-buckets total-101-buckets
	 * 101-200 5ms 20-buckets total-121-buckets
	 * 201-500 10ms 30-buckets total-151-buckets
	 * 501-3000 50ms 50-buckets total-201-buckets
	 * 3001-10000 200ms 35-buckets total-236-buckets
	 * 10001-30000 1000ms 20-buckets total-256-buckets
	 * total-256-buckets
	 */
	if value <= 100 {
		return int(value)
	} else if value <= 200 {
		return int(101 + (value-101)/5)
	} else if value <= 500 {
		return int(121 + (value-201)/10)
	} else if value <= 3000 {
		return int(151 + (value-501)/50)
	} else if value <= 10000 {
		return int(201 + (value-3001)/200)
	} else if value <= 30000 {
		return int(236 + (value-10001)/1000)
	} else {
		return 255
	}
}

// GetStatsDataAndReset 获取统计数据并重置
func (s *Statistics) GetStatsDataAndReset() *StatsData {
	count := atomic.SwapInt64(&s.count, 0)
	sum := atomic.SwapInt64(&s.sum, 0)
	max := s.maxValue.GetAndSet(0)

	var avg float64
	if count == 0 {
		avg = 0.0
	} else {
		avg = float64(sum) / float64(count)
	}

	p50Position := int64(float64(count) * 0.5)
	p75Position := int64(float64(count) * 0.75)
	p90Position := int64(float64(count) * 0.90)
	p95Position := int64(float64(count) * 0.95)
	p99Position := int64(float64(count) * 0.99)
	p999Position := int64(float64(count) * 0.999)

	p50, p75, p90, p95, p99, p999 := s.calculatePercentiles(
		count, max, p50Position, p75Position, p90Position, p95Position, p99Position, p999Position, true)

	return &StatsData{
		Count: count,
		Avg:   avg,
		Max:   float64(max),
		P50:   float64(p50),
		P75:   float64(p75),
		P90:   float64(p90),
		P95:   float64(p95),
		P99:   float64(p99),
		P999:  float64(p999),
	}
}

// GetStatsData 获取统计数据（不重置）
func (s *Statistics) GetStatsData() *StatsData {
	count := atomic.LoadInt64(&s.count)
	sum := atomic.LoadInt64(&s.sum)
	max := s.maxValue.Get()

	var avg float64
	if count == 0 {
		avg = 0.0
	} else {
		avg = float64(sum) / float64(count)
	}

	p50Position := int64(float64(count) * 0.5)
	p75Position := int64(float64(count) * 0.75)
	p90Position := int64(float64(count) * 0.90)
	p95Position := int64(float64(count) * 0.95)
	p99Position := int64(float64(count) * 0.99)
	p999Position := int64(float64(count) * 0.999)

	p50, p75, p90, p95, p99, p999 := s.calculatePercentiles(
		count, max, p50Position, p75Position, p90Position, p95Position, p99Position, p999Position, false)

	return &StatsData{
		Count: count,
		Avg:   avg,
		Max:   float64(max),
		P50:   float64(p50),
		P75:   float64(p75),
		P90:   float64(p90),
		P95:   float64(p95),
		P99:   float64(p99),
		P999:  float64(p999),
	}
}

// calculatePercentiles 计算百分位数
func (s *Statistics) calculatePercentiles(count, max, p50Pos, p75Pos, p90Pos, p95Pos, p99Pos, p999Pos int64, reset bool) (p50, p75, p90, p95, p99, p999 int64) {
	p50, p75, p90, p95, p99, p999 = -1, -1, -1, -1, -1, -1

	if s.exact {
		var c int64
		var lastIndexNum int64
		if reset {
			lastIndexNum = atomic.SwapInt64(&s.distribute[len(s.distribute)-1], 0)
		} else {
			lastIndexNum = atomic.LoadInt64(&s.distribute[len(s.distribute)-1])
		}

		for i := 0; i < len(s.distribute); i++ {
			var current int64
			if reset {
				current = atomic.SwapInt64(&s.distribute[i], 0)
			} else {
				current = atomic.LoadInt64(&s.distribute[i])
			}
			c += current

			if p50 == -1 && c >= p50Pos {
				p50 = int64(i)
			}
			if p75 == -1 && c >= p75Pos {
				p75 = int64(i)
			}
			if p90 == -1 && c >= p90Pos {
				p90 = int64(i)
			}
			if p95 == -1 && c >= p95Pos {
				p95 = int64(i)
			}
			if p99 == -1 && c >= p99Pos {
				p99 = int64(i)
			}
			if p999 == -1 && c >= p999Pos {
				p999 = int64(i)
			}
		}

		// 处理超出范围的情况
		maxIndex := int64(len(s.distribute) - 1)
		if p50 == maxIndex {
			p50 = s.quantileExceed(max, maxIndex, lastIndexNum, count, p50Pos)
		}
		if p75 == maxIndex {
			p75 = s.quantileExceed(max, maxIndex, lastIndexNum, count, p75Pos)
		}
		if p90 == maxIndex {
			p90 = s.quantileExceed(max, maxIndex, lastIndexNum, count, p90Pos)
		}
		if p95 == maxIndex {
			p95 = s.quantileExceed(max, maxIndex, lastIndexNum, count, p95Pos)
		}
		if p99 == maxIndex {
			p99 = s.quantileExceed(max, maxIndex, lastIndexNum, count, p99Pos)
		}
		if p999 == maxIndex {
			p999 = s.quantileExceed(max, maxIndex, lastIndexNum, count, p999Pos)
		}
	} else {
		var c int64
		for i := 0; i < len(s.distribute); i++ {
			var current int64
			if reset {
				current = atomic.SwapInt64(&s.distribute[i], 0)
			} else {
				current = atomic.LoadInt64(&s.distribute[i])
			}
			c += current

			if p50 == -1 && c >= p50Pos {
				offset := current - (c - p50Pos)
				p50 = s.index2Real(i, current, offset, max)
			}
			if p75 == -1 && c >= p75Pos {
				offset := current - (c - p75Pos)
				p75 = s.index2Real(i, current, offset, max)
			}
			if p90 == -1 && c >= p90Pos {
				offset := current - (c - p90Pos)
				p90 = s.index2Real(i, current, offset, max)
			}
			if p95 == -1 && c >= p95Pos {
				offset := current - (c - p95Pos)
				p95 = s.index2Real(i, current, offset, max)
			}
			if p99 == -1 && c >= p99Pos {
				offset := current - (c - p99Pos)
				p99 = s.index2Real(i, current, offset, max)
			}
			if p999 == -1 && c >= p999Pos {
				offset := current - (c - p999Pos)
				p999 = s.index2Real(i, current, offset, max)
			}
		}
	}

	return
}

// index2Real 索引转换为真实值
func (s *Statistics) index2Real(index int, current, offset, max int64) int64 {
	if index >= 0 && index <= 100 {
		return int64(index)
	} else if index <= 121 {
		rate := float64(offset) / float64(current)
		return 100 + int64(index-100+1)*5 + int64(5*rate)
	} else if index <= 151 {
		rate := float64(offset) / float64(current)
		return 200 + int64(index-121+1)*10 + int64(10*rate)
	} else if index <= 201 {
		rate := float64(offset) / float64(current)
		return 500 + int64(index-151+1)*50 + int64(50*rate)
	} else if index <= 236 {
		rate := float64(offset) / float64(current)
		return 3000 + int64(index-201+1)*200 + int64(200*rate)
	} else if index < 255 {
		rate := float64(offset) / float64(current)
		return 3000 + int64(index-236+1)*1000 + int64(1000*rate)
	} else if index == 255 {
		return 30000 + int64(float64(offset)/float64(current)*float64(max-30000))
	}
	return 0
}

// quantileExceed 计算超出范围的百分位数
func (s *Statistics) quantileExceed(max, maxIndex, lastIndexNum, count, quantilePosition int64) int64 {
	return int64(math.Round(float64(max) - (float64(max-maxIndex+1)/float64(lastIndexNum))*float64(count-quantilePosition)))
}

// Update 更新最大值
func (m *MaxValue) Update(value int64) {
	for {
		oldValue := atomic.LoadInt64(&m.value)
		if value > oldValue {
			if atomic.CompareAndSwapInt64(&m.value, oldValue, value) {
				break
			}
		} else {
			break
		}
	}
}

// GetAndSet 获取并设置新值
func (m *MaxValue) GetAndSet(newValue int64) int64 {
	return atomic.SwapInt64(&m.value, newValue)
}

// Get 获取当前值
func (m *MaxValue) Get() int64 {
	return atomic.LoadInt64(&m.value)
}
