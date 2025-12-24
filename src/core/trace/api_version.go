package trace

// ApiVersion API版本
type ApiVersion string

const (
	// V1 API版本1
	V1 ApiVersion = "v1"
	// V2 API版本2
	V2 ApiVersion = "v2"
)

// String 返回字符串表示
func (a ApiVersion) String() string {
	return string(a)
}
