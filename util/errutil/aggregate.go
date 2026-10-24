package errutil

import (
	"strings"
	"sync"
)

// RTTI xBpAmFj.k70VYqq，原方法 @ 0xb26b20。没有 Unwrap 方法。
type k70VYqq []error

func (errs k70VYqq) Error() string {
	if len(errs) == 1 && errs[0] != nil {
		return errs[0].Error()
	}
	var text strings.Builder
	for i, err := range errs {
		if err != nil {
			text.WriteString(err.Error())
			if i < len(errs)-1 {
				text.WriteByte('\n')
			}
		}
	}
	return text.String()
}

// T2i1ItMaa @ 0xb26d20。过滤 nil 后仍返回命名切片，即使只剩一个错误。
func T2i1ItMaa(errs ...error) error {
	var nonNil k70VYqq
	for _, err := range errs {
		if err != nil {
			nonNil = append(nonNil, err)
		}
	}
	if len(nonNil) == 0 {
		return nil
	}
	return nonNil
}

// D5aDHlQDLp @ 0xb26e60；各任务并发执行，按完成顺序收集错误。
func D5aDHlQDLp(tasks ...func() error) error {
	if len(tasks) == 0 {
		return nil
	}
	var waiting sync.WaitGroup
	waiting.Add(len(tasks))
	errors := make(chan error, len(tasks))
	for _, task := range tasks {
		go func(task func() error) {
			if err := task(); err != nil {
				errors <- err
			}
			waiting.Done()
		}(task)
	}
	waiting.Wait()
	var nonNil k70VYqq
	for {
		select {
		case err, ok := <-errors:
			if !ok {
				if len(nonNil) == 0 {
					return nil
				}
				return nonNil
			}
			if err != nil {
				nonNil = append(nonNil, err)
			}
		default:
			if len(nonNil) == 0 {
				return nil
			}
			return nonNil
		}
	}
}
