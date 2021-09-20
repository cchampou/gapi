package logger

import "fmt"

type Logger struct {
}

func (l *Logger) Log(str string) (int, error) {
	return fmt.Println(str)
}
