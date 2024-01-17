package basic

import (
	"github.com/IvanKyrylov/cloud-native-go-learn/pkg/logger"
	"github.com/IvanKyrylov/cloud-native-go-learn/pkg/runner"
)

var _ runner.Runner = &service{}

var functions = []func(){
	logicValue,
	complexValue,
	stringValue,
	variable,
	constValue,
	arrayValue,
	sliceValue,
	subSliceValue,
	stringSliceValue,
	mapValue,
	pointerValue,
	controlStructures,
	errorProcessing,
	function,
	structure,
	methods,
	interfaces,
	goroutines,
	channels,
	selectControl,

	generic,
}

type service struct {
	logger    logger.Logger
	functions []func()
}

func New(logger logger.Logger) runner.Runner {
	return &service{
		logger:    logger,
		functions: functions,
	}
}
