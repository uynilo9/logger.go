package logger_test

import (
	"testing"

	"github.com/uynilo9/logger.go"
)

const (
	time = 10.5
	code = 213
	line = 10
	column = 5
	foo = "foo"
	url = "https://uynilo9.is-a.dev"
	id = 213
)

func TestDetail(t *testing.T) {
	logger.Detail("- Time: ", time, " - Code: ", code)
}

func TestError(t *testing.T) {
	logger.Error("Unexpected character on line ", line, " column ", column)
}

func TestFatal(t *testing.T) {
	logger.Fatal("The required data '", foo, "' was missing")
}

func TestTip(t *testing.T) {
	logger.Tip("To get more info, visit ", url, "#", id)
}

func TestDetailln(t *testing.T) {
	logger.Detailln("- Time: ", time, " - Code: ", code)
}

func TestErrorln(t *testing.T) {
	logger.Errorln("Unexpected character on line ", line, " column ", column)
}

func TestFatalln(t *testing.T) {
	logger.Fatalln("The required data '", foo, "' was missing")
}

func TestTipln(t *testing.T) {
	logger.Tipln("To get more info, visit ", url, "#", id)
}

func TestDetailf(t *testing.T) {
	logger.Detailf("- Time: %g - Code: %d", time, code)
}

func TestErrorf(t *testing.T) {
	logger.Errorf("Unexpected character on line %d column %d", line, column)
}

func TestFatalf(t *testing.T) {
	logger.Fatalf("The required data '%s' was missing", foo)
}

func TestTipf(t *testing.T) {
	logger.Tipf("To get more info, visit %s#%d", url, id)
}