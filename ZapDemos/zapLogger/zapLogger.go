package zapLogger

import (
    "go.uber.org/zap"
    "net/http"
)

//参考：https://www.liwenzhou.com/posts/Go/zap/

var logger *zap.Logger

func ZapDemo() {
    InitLogger()
    defer logger.Sync()
    simpleHttpGet("www.google.com")
    simpleHttpGet("http://www.google.com")
}


func InitLogger() {
    // 文件跑哪儿去了？貌似要特殊操作才有文件
    // zap.NewProduction()/zap.NewDevelopment()或者zap.Example()
    logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
    resp, err := http.Get(url)
    if err != nil {
        logger.Error(
            "Error fetching url..",
            zap.String("url",url),
            zap.Error(err))
    } else {
        logger.Info("Success..",
            zap.String("statusCode", resp.Status),
            zap.String("url", url))
        resp.Body.Close()
    }
}



