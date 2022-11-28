package zapLogger

import "testing"

func TestZapDemo(t *testing.T) {
    //ZapDemo1()
    //ZapDemo2()
    //ZapDemo3()
    ZapDemo4()
}
//{"level":"error",
//    "ts":1668957704.2193444,
//    "caller":"zapLogger/zapLogger.go:29",
//    "msg":"Error fetching url..",
//    "url":"www.google.com",
//    "error":"Get \"www.google.com\": unsupported protocol scheme \"\"",
//    "stacktrace":"github.com/hpypig/Demos/ZapDemos/zapLogger.simpleHttpGet\n\tD:/sw_study/work/Demos/ZapDemos/zapLogger/zapLogger.go:29\ngithub.com/hpypig/Demos/ZapDemos/zapLogger.ZapDemo\n\tD:/sw_study/work/Demos/ZapDemos/zapLogger/zapLogger.go:15\ngithub.com/hpypig/Demos/ZapDemos/zapLogger.TestZapDemo\n\tD:/sw_study/work/Demos/ZapDemos/zapLogger/zapLogger_test.go:6\ntesting.tRunner\n\tC:/Program Files/Go/src/testing/testing.go:1259"}
