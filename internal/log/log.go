package log

import (
	"log"
	"os"
	"sync"
)

const IsErrorLog = true
const IsInfoLog = true //false

//自己定义一个关于日志的简单框架
var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Llongfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Llongfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log methods
var (
	Errorf = func(err ...interface{}) {
		if IsErrorLog {
			errorLog.Println(err)
		}
	}

	Info = func(err ...interface{}) {
		if IsInfoLog {
			infoLog.Println(err)
		}
	}
)
