package logger

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

func Detail(a ...any) (int, error) {
	a = append([]any{color.Gray + "DETAIL" + color.Reset + "\n"}, a...)
	n, err := fmt.Print(a...)
	return n, err
}

func Error(a ...any) (int, error) {
	a = append([]any{color.Red + "ERROR" + color.Reset + " "}, a...)
	n, err := fmt.Print(a...)
	return n, err
}

func Fatal(a ...any) {
	a = append([]any{color.Red + "FATAL" + color.Reset + " "}, a...)
	fmt.Print(a...)
	os.Exit(1)
}

func Tip(a ...any) (int, error) {
	a = append([]any{color.Blue + "TIP" + color.Reset + " "}, a...)
	n, err := fmt.Print(a...)
	return n, err
}

func Detailln(a ...any) (int, error) {
	a = append([]any{color.Gray + "DETAIL" + color.Reset + "\n"}, a...)
	n, err := fmt.Println(a...)
	return n, err
}

func Errorln(a ...any) (int, error) {
	a = append([]any{color.Red + "ERROR" + color.Reset}, a...)
	n, err := fmt.Println(a...)
	return n, err
}

func Fatalln(a ...any) {
	a = append([]any{color.Red + "FATAL" + color.Reset}, a...)
	fmt.Println(a...)
	os.Exit(1)
}

func Tipln(a ...any) (int, error) {
	a = append([]any{color.Blue + "TIP" + color.Reset}, a...)
	n, err := fmt.Println(a...)
	return n, err
}

func Detailf(format string, a ...any) (int, error) {
	n, err := fmt.Printf(color.Gray + "DETAIL" + color.Reset + "\n" + format, a...)
	return n, err
}

func Errorf(format string, a ...any) (int, error) {
	n, err := fmt.Printf(color.Red + "ERROR" + color.Reset + " " + format, a...)
	return n, err
}

func Fatalf(format string, a ...any) {
	fmt.Printf(color.Red + "FATAL" + color.Reset + " " + format, a...)
	os.Exit(1)
}

func Tipf(format string, a ...any) (int, error) {
	n, err := fmt.Printf(color.Blue + "TIP" + color.Reset + " " + format, a...)
	return n, err
}