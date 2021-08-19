package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

//Logger helps us with logging
type Logger struct {
}

//the number of characters that will be part of the code path
const sourceChars = 50

//the format of the timestamp within the log
const timeStampFormat string = "2006-01-02T15:04:05.000Z"

// The singleton instance
var instance *Logger

//To log or not to log
var enabled bool = true

// Little magic to call a function exactly once
var once sync.Once

//GetInstance returns a threadsafe singleton of the logger
func GetInstance() *Logger {
	once.Do(func() {
		instance = new(Logger)
	})
	return instance
}

//TurnOffLogs turns off all logging for the purpose of running tests
func (l *Logger) TurnOffLogs() {
	enabled = false
}

//Info logs at the info level to the stdout
func (l *Logger) Info(msg string) {
	if !enabled {
		return
	}
	t := time.Now()
	ts := t.Format(timeStampFormat)
	fmt.Fprintf(os.Stdout, "%s | INFO  | %s | %s\n", ts, l.formatTxtColumn(l.getCallStackLine(), sourceChars), msg)
}

//Warn logs at the warn level to the stdout
func (l *Logger) Warn(msg string) {
	if !enabled {
		return
	}
	t := time.Now()
	ts := t.Format(timeStampFormat)
	fmt.Fprintf(os.Stdout, "%s | WARN  | %s | %s\n", ts, l.formatTxtColumn(l.getCallStackLine(), sourceChars), msg)
}

//Error logs at the error level to the stderr as well as a stack trace.
func (l *Logger) Error(msg string) {
	if !enabled {
		return
	}
	t := time.Now()
	ts := t.Format(timeStampFormat)
	sTrace := string(debug.Stack())
	sTraceArr := strings.Split(sTrace, "\n")[5:]
	sTrace = strings.Join(sTraceArr, "\n")
	fmt.Fprintf(os.Stderr, "%s | ERROR | %s | %s\n%s", ts, l.formatTxtColumn(l.getCallStackLine(), sourceChars), msg, sTrace)
}

//Fatal logs at the fatal level to the stderr as well as a stack trace along with exiting the process with the given exitCode
func (l *Logger) Fatal(msg string, exitCode int) {
	if !enabled {
		os.Exit(exitCode)
	}
	t := time.Now()
	ts := t.Format(timeStampFormat)
	sTrace := string(debug.Stack())
	sTraceArr := strings.Split(sTrace, "\n")[5:]
	sTrace = strings.Join(sTraceArr, "\n")
	fmt.Fprintf(os.Stderr, "%s | FATAL | %s | %s\n%s", ts, l.formatTxtColumn(l.getCallStackLine(), sourceChars), msg, sTrace)
	os.Exit(exitCode)
}

func (l *Logger) getCallStackLine() string {
	sTrace := string(debug.Stack())
	sTraceArr := strings.Split(sTrace, "\n")
	line := strings.Replace(sTraceArr[8], "\t", "", -1)
	line = strings.Split(line, " ")[0]
	return line
}

//formats the text to fit in the number of col provided
func (l *Logger) formatTxtColumn(txt string, col int) string {
	if txt == "" {
		return txt
	}
	txtLen := len(txt)
	if txtLen == col {
		txt = fmt.Sprintf("   %s", txt) //prefix left by 3 spaces to match the dotdotdot below
	} else if txtLen > col {
		diff := txtLen - col
		txt = fmt.Sprintf("...%s", txt[diff:]) //prefix with dotdotdot
	} else if txtLen < col {

		diff := MaxInt32(col-txtLen+3, 3)
		for i := 0; i < diff; i++ {
			txt = fmt.Sprintf(" %s", txt)
		}
	}
	return txt
}
