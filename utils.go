package main

import "runtime"

func getFilePathFuncNameAndLine() (file, funcName string, line int) {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	_, file, line, _ = runtime.Caller(1)
	return file, f.Name(), line
}
