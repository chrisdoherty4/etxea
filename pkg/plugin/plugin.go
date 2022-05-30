package plugin

// Stage denotes a stage of execution.
type Stage int

const (
	// PreBootstrap is the very first stage of execution.
	PreBootstrap Stage = 1
	// Bootstrap defines the bootstrapping phase of a provider.
	Bootstrap Stage = 2
	// PostBootstrap is executed after the Bootstrap stage.
	PostBootstrap Stage = 3
)

// StageHandler is bound to a stage and handles the stage.
type StageHandler interface {
	HandleStage() error
}

type StageHandlerFunc func() error

func (fn StageHandlerFunc) HandleStage() error {
	return fn()
}

type BindingAgent interface {
	Bind(Stage, StageHandler) error
}

type Provider interface {
	Init() error
}
