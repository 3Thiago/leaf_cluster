package log

import "fmt"

type tagLog struct {
	tag interface{}
}

func (logger *tagLog) doPrintf(level int, printLevel string, format string, a ...interface{}) {
	fmt.Println(logger.tag, level, printLevel, format, a)
}

func (logger *tagLog) Debug(format string, a ...interface{}) {
	if !logger.matchLog(debugLevel) {
		return
	}
	logger.doPrintf(debugLevel, printDebugLevel, format, a...)
}

func (logger *tagLog) Release(format string, a ...interface{}) {
	if !logger.matchLog(releaseLevel) {
		return
	}
	logger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
}

func (logger *tagLog) Error(format string, a ...interface{}) {
	if !logger.matchLog(errorLevel) {
		return
	}
	logger.doPrintf(errorLevel, printErrorLevel, format, a...)
}

func (logger *tagLog) Fatal(format string, a ...interface{}) {
	if !logger.matchLog(fatalLevel) {
		return
	}
	logger.doPrintf(fatalLevel, printFatalLevel, format, a...)
}

func (Logger *tagLog) matchLog(level int) bool{
	return true
}

func TagLog(tag interface{}) tagLog {
	return tagLog{tag:tag}
}


