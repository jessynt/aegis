package contract

type Activation interface {
	Id() int64
	GetActivationName() string
	GetModelId() int64
	GetWarningScore() int64
	GetBlockScore() int64
}
