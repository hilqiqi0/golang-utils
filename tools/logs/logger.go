package logs

/*
有关log打印的封装
*/
import (
	"runtime"

	"github.com/cihub/seelog"
)

func LoggerSetup(c string) {
	//seelog.RegisterCustomFormatter("QuoteMsg", createQuoteMsgFormatter)
	//seelog.RegisterCustomFormatter("QuoteMsg",createQuoteMsgFormatter)
	logger, err := seelog.LoggerFromConfigAsFile(c)
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		_ = seelog.Critical("Important error:", file, ":", line, err)
		panic(err)
		return
	}
	_ = seelog.ReplaceLogger(logger)
}
