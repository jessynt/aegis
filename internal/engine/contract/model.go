package contract

// 风控模型
type Model interface {
	GetModelName() string
	GetModelId() int64
	GetModelGUID() string
}
