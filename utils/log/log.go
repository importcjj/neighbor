package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

// level
const (
	DEBUG int = iota
	INFO
	WARN
	ERROR
)

var levelNames = [4]string{"DEBUG", "INFO", "WARN", "ERROR"}

var colors = map[string]int{
	"black":   0,
	"red":     1,
	"green":   2,
	"yellow":  3,
	"blue":    4,
	"magenta": 5,
	"cyan":    6,
	"white":   7,
}
var levelColors = map[int]string{
	DEBUG: "blue",
	INFO:  "green",
	WARN:  "yellow",
	ERROR: "red",
}

var (
	enabled           = true
	level             = INFO
	colored           = false
	w       io.Writer = os.Stderr
	mutex   sync.Mutex
)

// SetLevel 设置日志而级别
func SetLevel(l int) {
	level = l % len(levelNames)
}

// SetColored 是否开启颜色
func SetColored(b bool) {
	colored = b
}

// SetWriter 设置writer
func SetWriter(writer io.Writer) {
	w = writer
}

// Disable 关闭日志
func Disable() {
	enabled = false
}

// Enable 开启日志
func Enable() {
	enabled = true
}

// Colored 为字符串加上颜色
func Colored(color string, text string) string {
	return fmt.Sprintf("\033[3%dm%s\033[0m", colors[color], text)
}

// header 构造日志输出格式的Header
func header(time string, level int, filepath string, line int) string {
	levelName := fmt.Sprintf("%s", levelNames[level])
	levelColor := levelColors[level]

	if colored {
		levelName = Colored(levelColor, levelName)
	}

	return fmt.Sprintf("%s [%s] [%s:%d]", time, levelName, filepath, line)
}

func log(l int, msg string) error {
	if enabled && l >= level {
		_, filename, line, _ := runtime.Caller(2)
		pkgName := path.Base(path.Dir(filename))
		filepath := path.Join(pkgName, path.Base(filename))
		now := time.Now().Format("2006/02/02 15:04:05")

		header := header(now, l, filepath, line)
		mutex.Lock()
		defer mutex.Unlock()
		_, err := fmt.Fprintf(w, "%s %s\n", header, msg)
		return err
	}
	return nil
}

// Debugf 调试
func Debugf(format string, a ...interface{}) error {
	return log(DEBUG, fmt.Sprintf(format, a...))
}

// Infof 普通
func Infof(format string, a ...interface{}) error {
	return log(INFO, fmt.Sprintf(format, a...))
}

// Warnf 警告
func Warnf(format string, a ...interface{}) error {
	return log(WARN, fmt.Sprintf(format, a...))
}

// Errorf 错误
func Errorf(format string, a ...interface{}) error {
	return log(ERROR, fmt.Sprintf(format, a...))
}

// Fatalf 错误并退出
func Fatalf(format string, a ...interface{}) {
	log(ERROR, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Debug 调试
func Debug(a ...interface{}) error {
	return log(DEBUG, fmt.Sprint(a...))
}

// Info 普通
func Info(a ...interface{}) error {
	return log(INFO, fmt.Sprint(a...))
}

// Warn 普通
func Warn(a ...interface{}) error {
	return log(WARN, fmt.Sprint(a...))
}

// Error 错误
func Error(a ...interface{}) error {
	return log(ERROR, fmt.Sprint(a...))
}

// Fatal 错误并退出
func Fatal(a ...interface{}) {
	log(ERROR, fmt.Sprint(a...))
	os.Exit(1)
}
