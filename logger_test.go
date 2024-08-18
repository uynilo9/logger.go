package logger_test

import (
	"testing"

	"github.com/uynilo9/logger.go"
)

const (
	time    = 10.5
	code    = 213
	line    = 10
	column  = 5
	foo     = "foo"
	url     = "https://uynilo9.is-a.dev"
	id      = 213
	version = "3.14.159"
)

func TestDetail(*testing.T) {
	logger.Detail("- Time: ", time, " - Code: ", code)
}

func TestError(*testing.T) {
	logger.Error("Unexpected character on line ", line, " column ", column)
}

func TestFatal(*testing.T) {
	logger.Fatal("The required data '", foo, "' was missing")
}

func TestTip(*testing.T) {
	logger.Tip("To get more info, visit ", url, "#", id)
}

func TestWarn(*testing.T) {
	logger.Warn("New version ", version, " has been released, please update")
}

func TestDetailln(*testing.T) {
	logger.Detailln("- Time: ", time, " - Code: ", code)
}

func TestErrorln(*testing.T) {
	logger.Errorln("Unexpected character on line ", line, " column ", column)
}

func TestFatalln(*testing.T) {
	logger.Fatalln("The required data '", foo, "' was missing")
}

func TestTipln(*testing.T) {
	logger.Tipln("To get more info, visit ", url, "#", id)
}

func TestWarnln(*testing.T) {
	logger.Warnln("New version ", version, " has been released, please update")
}

func TestDetailf(*testing.T) {
	logger.Detailf("- Time: %g - Code: %d", time, code)
}

func TestErrorf(*testing.T) {
	logger.Errorf("Unexpected character on line %d column %d", line, column)
}

func TestFatalf(*testing.T) {
	logger.Fatalf("The required data '%s' was missing", foo)
}

func TestTipf(*testing.T) {
	logger.Tipf("To get more info, visit %s#%d", url, id)
}

func TestWarnf(*testing.T) {
	logger.Warnf("New version %s has been released, please update", version)
}
