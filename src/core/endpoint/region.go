package endpoint

// Region 区域定义
type Region int

const (
	CN Region = iota // 中国，默认值
	SG               // 新加坡
)

// String 返回区域的字符串表示
func (r Region) String() string {
	switch r {
	case CN:
		return "CN"
	case SG:
		return "SG"
	default:
		return "UNKNOWN"
	}
}
