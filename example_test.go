package logger_test

import "github.com/uynilo9/logger.go"

func ExampleDetail() {
	logger.Detail("- Time: ", time, " - Code: ", code)
	// Output: DETAIL
	// - Time: 10.5 - Code: 213
}

func ExampleError() {
	logger.Error("Unexpected character on line ", line, " column ", column)
	// Output: ERROR Unexpected character on line 10 column 5
}

func ExampleFatal() {
	logger.Fatal("The required data '", foo, "' was missing")
	// Output: FATAL The required data 'foo' was missing
	// // exit
}

func ExampleTip() {
	logger.Tip("To get more info, visit ", url, "#", id)
	// Output: TIP To get more info, visit https://uynilo9.is-a.dev#213
}

func ExampleWarn() {
	logger.Warn("New version ", version, " has been released, please update")
	// Output: WARN New version 3.14.159 has been released, please update
}

func ExampleDetailln() {
	logger.Detailln("- Time: ", time, " - Code: ", code)
	// Output: DETAIL
	// - Time: 10.5 - Code: 213
	// // newline
}

func ExampleErrorln() {
	logger.Errorln("Unexpected character on line ", line, " column ", column)
	// Output: ERROR Unexpected character on line  10  column  5
	// // newline
}

func ExampleFatalln() {
	logger.Fatalln("The required data '", foo, "' was missing")
	// Output: The required data ' foo ' was missing
	// // newline // exit
}

func ExampleTipln() {
	logger.Tipln("To get more info, visit ", url, "#", id)
	// Output: TIP To get more info, visit  https://uynilo9.is-a.dev # 213
	// // newline
}

func ExampleWarnln() {
	logger.Warnln("New version ", version, " has been released, please update")
	// Output: WARN New version 3.14.159 has been released, please update
	// // newline
}

func ExampleDetailf() {
	logger.Detailf("- Time: %g - Code: %d", time, code)
	// Output: DETAIL
	// - Time: 10.5 - Code: 213
}

func ExampleErrorf() {
	logger.Errorf("Unexpected character on line %d column %d", line, column)
	// Output: ERROR Unexpected character on line 10 column 5
}

func ExampleFatalf() {
	logger.Fatalf("The required data '%s' was missing", foo)
	// Output: FATAL The required data 'foo' was missing
	// // exit
}

func ExampleTipf() {
	logger.Tipf("To get more info, visit %s#%d", url, id)
	// Output: TIP To get more info, visit https://uynilo9.is-a.dev#213
}

func ExampleWarnf() {
	logger.Warnf("New version %s has been released, please update", version)
	// Output: WARN New version 3.14.159 has been released, please update
}
