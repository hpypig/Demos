package stdLogger

import (
    "log"
    "net/http"
    "os"
)

func SetupLogger() {
    logFileLocation, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744) // 0 代表八进制
    log.SetOutput(logFileLocation) // 给一个 io.Writer 即一个可写的地方，一个输出流
}
func SimpleHttpGet(url string) {
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Error fetching url %s: %s", url, err.Error())
    } else {
        log.Printf("Status Code for %s : %s", url, resp.Status)
        resp.Body.Close()
    }


}
