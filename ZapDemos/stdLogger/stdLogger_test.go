package stdLogger

import "testing"

func TestSetupLogger(t *testing.T) {
    SetupLogger()
    //SimpleHttpGet("www.google.com")
    //SimpleHttpGet("https://www.google.com") // 这个网有问题，本来应该成功的，空了看下怎么搞
    SimpleHttpGet("https://www.baidu.com")
}
