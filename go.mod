module github.com/hpypig/Demos

go 1.17

require go.uber.org/zap v1.23.0
// indirect ?????啥呀？我就引入了一个zap，怎么下面还多了俩？
require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
