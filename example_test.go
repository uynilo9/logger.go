package logger_test

import "github.com/uynilo9/logger.go"

func ExampleDetail() {
	logger.Detail("- Time: ", time, " - Code: ", code)

	// Output: DATAIL
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

	// Output: TIP To get more info, visit https://github.com/uynilo9/logger.go#213
}

func ExampleDetailln() {
	logger.Detailln("- Time: ", time, " - Code: ", code)

	// Output: DATAIL
	// - Time: 10.5 - Code: 213
	// // newline
}

func ExampleErrorln() {
	logger.Errorln("Unexpected character on line ", line, " column ", column)
	
	// Output: ERROR Unexpected character on line 10 column 5
	// // newline
}

func ExampleFatalln() {
	logger.Fatalln("The required data '", foo, "' was missing")

	// Output: FATAL The required data 'foo' was missing
	// // newline
	// // exit
}

func ExampleTipln() {
	logger.Tipln("To get more info, visit ", url, "#", id)

	// Output: TIP To get more info, visit https://github.com/uynilo9/logger.go#213
	// // newline
}

func ExampleDetailf() {
	logger.Detailf("- Time: %g - Code: %d", time, code)

	// Output: DATAIL
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

	// Output: TIP To get more info, visit https://github.com/uynilo9/logger.go#213
}