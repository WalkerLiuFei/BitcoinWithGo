package utils

import (
	"go.uber.org/zap"
	"fmt"
	"runtime"
	"reflect"
)

var sugarLogger *zap.SugaredLogger

func GetSugarLogger() *zap.SugaredLogger {
	if sugarLogger != nil {
		return sugarLogger
	}
	logger, err := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	if err != nil {
		fmt.Println(err)
	}
	return logger.Sugar()
}

//在打印函数名称的时候调用
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
