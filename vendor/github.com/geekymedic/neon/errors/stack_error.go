package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

const defDepth = 16

type errChain struct {
	err error
	msg string
}

// String out a chain link error.
// Example: chain1 = errChain { err: errors.New("hello word"), msg: "chain1"}
// chain2 = errChain {err: chain1, msg: "chain2"}
// chain3 = errChain {err: chain2, msg: "chain3"}
// err: "hello word", msg: "chain3, chain2, chain1"
// TODO: Benchmark it
func (errc *errChain) String() string {
	return errc.Error()
}

// err: "hello word",
// msg:
// chain3
//	  chain2
// 		chain1
func (errc *errChain) Error() string {
	if errc == nil {
		return ""
	}
	rawMsg, extMsg := unpackErrChain(errc)
	if len(extMsg) == 0 {
		return fmt.Sprintf("err: %v", rawMsg)
	}
	var extMsgFormat = fmt.Sprintf("msg:\n")
	wrapTable := func(i int) string {
		table := ""
		for ; i > 0; i-- {
			table += "\t"
		}
		return table
	}
	for i := 0; i < len(extMsg); i++ {
		if extMsg[i] == "" {
			continue
		}
		extMsgFormat += fmt.Sprintf("%v\n%s", extMsg[i], wrapTable(i+1))
	}
	return fmt.Sprintf("err: %v\n%v", rawMsg, extMsgFormat)
}

func (errc *errChain) shortTerm() string {
	if errc == nil {
		return ""
	}
	return fmt.Sprintf("%s", errc.msg)
}

func unpackErrChain(err error) (string, []string) {
	var (
		next = err
		raw  = ""
		msg  []string
	)
	for {
		if errc, ok := next.(*errChain); ok {
			next = errc.err
			msg = append(msg, errc.msg)
			continue
		}
		break
	}
	raw = next.Error()
	return raw, msg
}

var OkStackErr = StackError{errChain: nil, stackPtrs: nil}

// StackError out a chain link error.
// Example: chain1 = errChain { err: errors.New("hello word"), msg: "chain1"}
// chain2 = errChain {err: chain1, msg: "chain2"}
// chain3 = errChain {err: chain2, msg: "chain3"}
// stackError = StackError {errChain: chain3, stackPtrs: []uintptr}
// stdout:
// err: "hello word"
// msg:
//	chain1,
//		chain2,
//			chain3
// stacktrace:
//
type StackError struct {
	errChain  error
	stackPtrs []uintptr
}

func NewStackError(txt string) error {
	var (
		ptrs [defDepth]uintptr
		err  = errors.New(txt)
		n    = runtime.Callers(2, ptrs[:])
	)
	return &StackError{
		errChain:  err,
		stackPtrs: ptrs[:n],
	}
}

func newStackErrorWithOther(err error) error {
	var (
		ptrs [defDepth]uintptr
		n    = runtime.Callers(2, ptrs[:])
	)

	return &StackError{
		errChain:  err,
		stackPtrs: ptrs[:n],
	}
}

func (stackErr *StackError) IsNil() bool {
	return stackErr == nil || stackErr.errChain == nil
}

func (stackErr *StackError) ShortTerm() string {
	if stackErr == nil {
		return ""
	}
	errc, ok := stackErr.errChain.(*errChain)
	if ok {
		return errc.shortTerm()
	}
	return ""
}

func (stackErr *StackError) Error() string {
	if stackErr == nil || stackErr.errChain == nil {
		return ""
	}
	var txt = make([]string, 0, len(stackErr.stackPtrs))
	for _, ptr := range stackErr.stackPtrs {
		fn := runtime.FuncForPC(ptr)
		fileName, line := fn.FileLine(ptr)
		funcName := fn.Name()
		output := fmt.Sprintf("%s\n\t %s:%d", funcName, fileName, line)
		txt = append(txt, output)
	}
	frame := strings.Join(txt, "\n")
	chainFrame := stackErr.errChain.Error()
	// Only has top chain and top chain is std error
	if _, ok := stackErr.errChain.(*errChain); !ok {
		return fmt.Sprintf("%s\nstacktrace:\n%s", chainFrame, frame)
	}
	return fmt.Sprintf("%s\nstacktrace:\n%s", chainFrame, frame)
}

func IsNil(err error) bool {
	if err == nil {
		return true
	}
	stackErr, ok := err.(*StackError)
	if ok {
		return stackErr.IsNil()
	}
	return false
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	_, ok := err.(*StackError)
	// wrap
	if ok {
		return err
	}
	return newStackErrorWithOther(err)
}

//WithMessageEx creates a new stack error with err that is formatted by `format v...`
func WithMessage(err error, format string, v ...interface{}) error {
	if err == nil {
		return nil
	}
	e, ok := err.(*StackError)
	if ok {
		return &StackError{
			errChain: &errChain{
				err: e.errChain, msg: fmt.Sprintf(format, v...),
			},
			stackPtrs: e.stackPtrs,
		}
	}
	return &StackError{
		errChain:  &errChain{err: err, msg: fmt.Sprintf(format, v...)},
		stackPtrs: callerStack(),
	}
}

//FormatEx creates a new stack error that is formatted by `format v...`
func Format(format string, v ...interface{}) error {
	return &StackError{
		errChain:  fmt.Errorf(format, v...),
		stackPtrs: callerStack(),
	}
}

func ShortMsg(err error) string {
	if err == nil {
		return ""
	}
	e, ok := err.(*StackError)
	if ok {
		return e.ShortTerm()
	}
	return err.Error()
}

//ByEx creates a new stack error with std error
func By(err error) error {
	if err == nil {
		return nil
	}
	e, ok := err.(*StackError)
	if ok {
		return &StackError{
			errChain:  e.errChain,
			stackPtrs: e.stackPtrs,
		}
	}
	return &StackError{
		errChain:  err,
		stackPtrs: callerStack(),
	}
}

func callerStack() []uintptr {
	var ptrs [defDepth]uintptr
	n := runtime.Callers(2, ptrs[:])
	return ptrs[:n]
}
