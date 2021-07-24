package main

import "fmt"

const englishHelloPrefix = "hello"
const spanish = "Spanish"
const helloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + " " + name
}

/**
在我们的函数签名中，我们使用了 命名返回值（prefix string）。
这将在你的函数中创建一个名为 prefix 的变量。
它将被分配「零」值。这取决于类型，例如 int 是 0，对于字符串它是 ""。
你只需调用 return 而不是 return prefix 即可返回所设置的值。
这将显示在 Go Doc 中，所以它使你的代码更加清晰。
如果没有其他 case 语句匹配，将会执行 default 分支。
函数名称以小写字母开头。在 Go 中，公共函数以大写字母开始，私有函数以小写字母开头。我们不希望我们算法的内部结构暴露给外部，所以我们将这个功能私有化。
*/
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println("test")
}
