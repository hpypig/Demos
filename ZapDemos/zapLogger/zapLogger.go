package zapLogger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "io"
    "net/http"
    "os"
)

//参考：https://www.liwenzhou.com/posts/Go/zap/
/*
使用步骤：
1）NewProduction 生成 *zap.Logger 对象，全局的
   （可进一步 ——> Sugar() 生成*zap.SugaredLogger）
2）打印
   Logger: Error()  Info() Debug() Warn() Panic()
   SugaredLogger: Errorf() Infof()
----
定制 Logger
1) 设置 core
2) zap.New(core) 生成 Logger
3) 生成 SugaredLogger

好处： 设置日志格式、级别
-------
问题：
    代码简洁问题：encoder 的获取。通过参数指定不同 encoder，而不是改代码？

*/
var logger *zap.Logger
var sugaredLogger *zap.SugaredLogger

func ZapDemo1() {
    InitLogger()
    defer logger.Sync()
    simpleHttpGet("www.baidu.com")
    simpleHttpGet("https://www.baidu.com")
    simpleHttpGet("https://www.google.com")
}
func ZapDemo2() {
    InitLogger1() // 普通
    defer logger.Sync()
    simpleHttpGetWithSugar("www.baidu.com")
    simpleHttpGetWithSugar("https://www.baidu.com")
    simpleHttpGetWithSugar("https://www.google.com")
}
func ZapDemo3() {
    InitLogger2() // JSON
    defer sugaredLogger.Sync()
    simpleHttpGetWithSugar("www.baidu.com")
    simpleHttpGetWithSugar("https://www.baidu.com")
    simpleHttpGetWithSugar("https://www.google.com")
}
func ZapDemo4() {
    InitLogger3() // console
    defer sugaredLogger.Sync()
    simpleHttpGetWithSugar("www.baidu.com")
    simpleHttpGetWithSugar("https://www.baidu.com")
    simpleHttpGetWithSugar("https://www.google.com")
}



//-------------

func InitLogger() {
    // 文件跑哪儿去了？貌似要特殊操作才有文件
    // zap.NewProduction()/zap.NewDevelopment()或者zap.Example()
    logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
    resp, err := http.Get(url)
    if err != nil {
        //参数是 string, ...Field 固定了格式
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

//--------------

func InitLogger1() {
    logger, _ = zap.NewProduction()
    sugaredLogger = logger.Sugar()
}
func simpleHttpGetWithSugar(url string) {
    resp, err := http.Get(url)
    if err != nil {
        // 多了 Errorf，可自己定义格式
        sugaredLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
    } else {
        sugaredLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
        resp.Body.Close()
    }
}

//----定制 Logger

func InitLogger2() {
    // 这三个函数就是写定制逻辑的，不过现在先用现成的？
    encoder := getJSONEncoder()         // 日志格式
    writeSyncer := getLogWriter()       // 日志位置
    levelEnabler := zapcore.DebugLevel  // 日志级别

    core := zapcore.NewCore(encoder, writeSyncer, levelEnabler)

    logger = zap.New(core) // 常规日志
    sugaredLogger = logger.Sugar()
}

func getJSONEncoder() zapcore.Encoder {
    // 日志设置为 json 格式，传入参数为配置（格式？）
    return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
    file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)

    // 可以多端输出
    ws := io.MultiWriter(file, os.Stdout)

    // 设置用于同步日志的 Writer
    //return zapcore.AddSync(file)
    return zapcore.AddSync(ws)
}



//---------

func InitLogger3() {
    // 这三个函数就是写定制逻辑的，不过现在先用现成的？
    encoder := getConsoleEncoder()         // 日志格式
    writeSyncer := getLogWriter()       // 日志位置
    levelEnabler := zapcore.DebugLevel  // 日志级别

    core := zapcore.NewCore(encoder, writeSyncer, levelEnabler)
    logger = zap.New(core, zap.AddCaller()) // 多加一个功能，日志增加函数调用信息
    //logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 跳过调用？？
    sugaredLogger = logger.Sugar()
}
// 控制台格式的日志
func getConsoleEncoder() zapcore.Encoder {
    encoderConfig := zap.NewProductionEncoderConfig()
    return zapcore.NewConsoleEncoder(encoderConfig)
}

func getConsoleEncoder2() zapcore.Encoder {
    encoderConfig := zap.NewProductionEncoderConfig()

    // 修改时间编码器
    // 在日志文件中使用大写字母记录日志级别
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

    return zapcore.NewConsoleEncoder(encoderConfig)
}

//----------

// 多级别日志输出

func InitLogger4() {
    encoder := getJSONEncoder()
    // test.log记录全量日志
    logF, _ := os.Create("./test4.log")
    c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
    // test.err.log记录ERROR级别的日志
    errF, _ := os.Create("./test4.err.log")
    c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
    // 使用NewTee将c1和c2合并到core
    core := zapcore.NewTee(c1, c2)
    logger = zap.New(core, zap.AddCaller())
}

// 按大小切割文件
func getLogWriter2() zapcore.WriteSyncer {
    lumberJackLogger := &lumberjack.Logger{
        Filename:   "./test.log",
        MaxSize:    10, // MB  超过以后会把原文件加时间戳重命名，然后新建一个log
        MaxBackups: 5, // 不知道是不是切割数的意思？？？？？
        MaxAge:     30,
        Compress:   false,
    }
    return zapcore.AddSync(lumberJackLogger)
}
