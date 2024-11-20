# Gormatter

## Introduction
Gormatter in a small lib to help making colored and stylized outputs in the terminal.

## How to use
Import the package :

```go
import (
	gmt "github.com/halsten-dev/gormatter"
```

Then there is choice between multiple function :

```go
fmt.Println(gmt.Bold(gmt.Red, gmt.None, "Text in red and bold"))
fmt.Println(gmt.BoldItalic(gmt.Blue, gmt.Yellow, "Text in blue, bold and italic with yellow background"))
fmt.Println(gmt.Italic(gmt.Yellow, gmt.Cyan, "Text in yellow in italic with cyan background"))
```

Available colors :

```go
Black
Red
Green
Yellow
Blue
Magenta
Cyan
White

HighBlack
HighRed
HighGreen
HighYellow
HighBlue
HighMagenta
HighCyan
HighWhite
```

Available functions :

```go
func Normal(foregroundColor int, backgroundColor int, text string) string

func Bold(foregroundColor int, backgroundColor int, text string) string

func Italic(foregroundColor int, backgroundColor int, text string) string

func Underline(foregroundColor int, backgroundColor int, text string) string

func BoldItalic(foregroundColor int, backgroundColor int, text string) string

func BoldUnderline(foregroundColor int, backgroundColor int, text string) string

func BoldItalicUnderline(foregroundColor int, backgroundColor int, text string) string

func ItalicUnderline(foregroundColor int, backgroundColor int, text string) string
```

## License
MIT License

Copyright (c) 2024 Lionel Leeser

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

