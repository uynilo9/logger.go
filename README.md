# 🌴 - A simple logger for Go

## ✨ Intro

A simple logger written in Go, which is used to print errors, tips, etc. messages, was born while I was working on another project [github.com/uynilo9/bento](https://github.com/uynilo9/bento) of mine

Btw it used to be in [pkg/logger](https://github.com/uynilo9/bento/tree/main/pkg/logger) under [github.com/uynilo9/bento](https://github.com/uynilo9/bento) originally. So far it still exists there but will no longer get updates there (if it's supposed to get any here lol).

## 📦 Installation

1. Install the module from GitHub url:

```sh
go get github.com/uynilo9/logger.go
```

2. Import the module in your Go code:

```go
import "github.com/uynilo9/logger.go"
```

## ⚙️ Functionality

#### 1. `logger.Detail(a ...any)`

##### (a) Argument(s):

- `a` (`...any`): The detailed message(s) to be printed. You can pass any number of arguments of any type.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Detail("- Time: ", time, " - Code: " code)

_, err := logger.Detail("- Time: ", time, " - Code: " code)
if err != nil {
    // ...
```

#### 2. `logger.Error(a ...any)`

##### (a) Argument(s):

- `a` (`...any`): The error message(s) to be printed. You can pass any number of arguments of any type.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Error("Unexpected character on line ", line, " column ", column)

_, err := logger.Error("Unexpected character on line ", line, " column ", column)
if err != nil {
    // ...
```

#### 3. `logger.Fatal(a ...any)`

##### (a) Argument(s):

- `a` (`...any`): The fatal message(s) to be printed. You can pass any number of arguments of any type.

##### (b) Return:

> [!WARNING]
> This function returns NOTHING because it directly EXITS after printing the message(s).

##### (c) Example:

```go
logger.Fatal("The required data \'", data, "\' was missing")
```

#### 4. `logger.Tip(a ...any)`

##### (a) Argument(s):

- `a` (`...any`): The tip message(s) to be printed. You can pass any number of arguments of any type.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Tip("To get more info, visit ", url, "#", id)

_, err := logger.Tip("To get more info, visit ", url, "#", id)
if err != nil {
    // ...
```

#### 5. `logger.Detailln(a ...any)`

##### (a) Argument(s):

- `a` (`...any`): The detailed message(s) to be printed. You can pass any number of arguments of any type. A newline is always appended.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Detailln("- Time: ", time, " - Code: " code)

_, err := logger.Detailln("- Time: ", time, " - Code: " code)
if err != nil {
    // ...
```

#### 6. `logger.Errorln(a ...any)`

##### (a) Argument(s):

- `a` (`...any`): The error message(s) to be printed. You can pass any number of arguments of any type. A newline is always appended.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Error("Unexpected character on line ", line, " column ", column)

_, err := logger.Errorln("Unexpected character on line ", line, " column ", column)
if err != nil {
    // ...
```

#### 7. `logger.Fatalln(a ...any)`

##### (a) Argument(s):

- `a` (`...any`): The fatal message(s) to be printed. You can pass any number of arguments of any type. A newline is always appended.

##### (b) Return:

> [!WARNING]
> This function returns NOTHING because it directly EXITS after printing the message(s).

##### (c) Example:

```go
logger.Fatalln("The required data \'", data, "\' was missing")
```

#### 8. `logger.Tipln(a ...any)`

##### (a) Argument(s):

- `a` (`...any`): The tip message(s) to be printed. You can pass any number of arguments of any type. A newline is always appended.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Tipln("To get more info, visit ", url, "#", id)

_, err := logger.Tipln("To get more info, visit ", url, "#", id)
if err != nil {
    // ...
```

#### 9. `logger.Detailf(format string, a ...any)`

##### (a) Argument(s):

- `format` (`string`): The format string for the detailed message. It follows the same syntax as `fmt.Printf`, allowing for formatted printing.
- `a` (`...anyThe values to be formatted and printed according to the format string.type.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Detailf("- Time: %g - Code: %d", time, code)

_, err := logger.Detailf("- Time: %g - Code: %d", time, code)
if err != nil {
    // ...
```

#### 10. `logger.Errorf(format string, a ...any)`

##### (a) Argument(s):

- `format` (`string`): The format string for the error message. It follows the same syntax as `fmt.Printf`, allowing for formatted printing.
- `a` (`...any`): The values to be formatted and printed according to the format string.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Errorf("Unexpected character on line %d column %d", line, column)

_, err := logger.Errorf("Unexpected character on line %d column %d", line, column)
if err != nil {
    // ...
```

#### 11. `logger.Fatalf(format string, a ...any)`

##### (a) Argument(s):

- `format` (`string`): The format string for the fatal message. It follows the same syntax as `fmt.Printf`, allowing for formatted printing.
- `a` (`...any`): The values to be formatted and printed according to the format string.

##### (b) Return:

> [!WARNING]
> This function returns NOTHING because it directly EXITS after printing the message(s).

##### (c) Example:

```go
logger.Fatalf("The required data \'%s\' was missing", data)
```

#### 12. `logger.Tipf(format string, a ...any)`

##### (a) Argument(s):

- `format` (`string`): The format string for the tip message. It follows the same syntax as `fmt.Printf`, allowing for formatted printing.
- `a` (`...any`): The values to be formatted and printed according to the format string.

##### (b) Return:

- `n` (`int`): The number of bytes written.
- `err` (`error`): Any write error encountered.

##### (c) Example:

```go
logger.Tipf("To get more info, visit %s#%d", url, id)

_, err := logger.Tipf("To get more info, visit %s#%d", url, id)
if err != nil {
    // ...
```

## 🔖 Version

0.0.1-dev.2

## 📜 License

MIT License