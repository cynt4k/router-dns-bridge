package herror

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
)

type InternalError struct {
	Err    error
	Stack  []byte
	Fields []map[string]interface{}
	Panic  bool
}

func (i *InternalError) Error() string {
	if i.Panic {
		return fmt.Sprintf("[Panic] %s\n%s", i.Err.Error(), i.Stack)
	}
	return fmt.Sprintf("%s\n%s", i.Err.Error(), i.Stack)
}

func InternalServerError(err error) error {
	context := errorReport(runtime.Caller(1))
	return &InternalError{
		Err:   err,
		Stack: debug.Stack(),
		Fields: []map[string]interface{}{{
			"context": context,
		}},
		// Fields: []zap.Field{zapdriver.ErrorReport(runtime.Caller(1)), zap.Error(err)},
		Panic: false,
	}
}

func Panic(err error) error {
	context := errorReport(runtime.Caller(1))
	return &InternalError{
		Err:   err,
		Stack: debug.Stack(),
		Fields: []map[string]interface{}{{
			"context": context,
		}},
		Panic: true,
	}
}

// reportLocation is the source code location information associated with the log entry
// for the purpose of reporting an error,
// if any.
type reportLocation struct {
	File     string `json:"filePath"`
	Line     string `json:"lineNumber"`
	Function string `json:"functionName"`
}

// reportContext is the context information attached to a log for reporting errors
type reportContext struct {
	ReportLocation reportLocation `json:"reportLocation"`
}

func errorReport(pc uintptr, file string, line int, ok bool) *reportContext {
	if !ok {
		return nil
	}

	var function string
	if fn := runtime.FuncForPC(pc); fn != nil {
		function = fn.Name()
	}
	return &reportContext{
		ReportLocation: reportLocation{
			File:     file,
			Line:     strconv.Itoa(line),
			Function: function,
		},
	}
}
