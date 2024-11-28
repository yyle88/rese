# rese

**rese** 简化了 Go 中的错误处理和多值函数返回值提取。它将错误检查和结果提取合并为一个操作。

**rese** 代表 **res**（结果）+ **err**（错误）。

## English README

[English documentation](README.md)

## 安装

```bash
go get github.com/yyle88/rese
```

---

## 函数

| 函数                                                        | 作用                                               | 返回值                                 |
|-----------------------------------------------------------|--------------------------------------------------|-------------------------------------|
| `V0(err error)`                                           | 检查错误，如果错误不为 `nil` 则触发 panic。没有返回值。               | 无                                   |
| `V1[T1 any](v1 T1, err error) T1`                         | 检查错误，如果没有错误，返回 `v1`。                             | 返回类型为 `T1` 的 `v1`                   |
| `V2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2)`        | 检查错误，如果没有错误，返回 `v1` 和 `v2`。                      | 返回类型为 `T1` 的 `v1` 和 类型为 `T2` 的 `v2` |
| `P0(err error)`                                           | 检查错误，如果错误不为 `nil` 则触发 panic。没有返回值。               | 无                                   |
| `P1[T1 any](v1 *T1, err error) *T1`                       | 检查错误，检查 `v1` 是否为非 `nil`，并返回 `v1`。                | 返回类型为 `T1` 的 `v1` 的指针               |
| `P2[T1, T2 any](v1 *T1, v2 *T2, err error) (*T1, *T2)`    | 检查错误，检查 `v1` 和 `v2` 是否为非 `nil`，并返回 `v1` 和 `v2`。  | 返回类型为 `T1` 和 `T2` 的 `v1` 和 `v2` 的指针 |
| `C0(err error)`                                           | 检查错误，如果错误不为 `nil` 则触发 panic。没有返回值。               | 无                                   |
| `C1[T1 comparable](v1 T1, err error) T1`                  | 检查错误，检查 `v1` 是否为零值，如果不是零值，返回 `v1`。               | 返回类型为 `T1` 的 `v1`                   |
| `C2[T1, T2 comparable](v1 T1, v2 T2, err error) (T1, T2)` | 检查错误，检查 `v1` 和 `v2` 是否为零值，如果不是零值，返回 `v1` 和 `v2`。 | 返回类型为 `T1` 的 `v1` 和 类型为 `T2` 的 `v2` |

## 示例

### 示例 1: 简单的错误和结果检查，使用 `V1`

```go
package main

import (
	"fmt"
	"github.com/yyle88/rese"
)

func run() (string, error) {
	return "Hello, World!", nil
}

func main() {
	// 使用 V1 来检查错误并获取结果
	result := rese.V1(run()) 
	fmt.Println(result) // 输出: Hello, World!
}
```

### 示例 2: 确保指针非 `nil`，使用 `P1`

```go
package main

import (
	"fmt"
	"github.com/yyle88/rese"
)

func getSomething() (*int64, error) {
	v := int64(42)
	return &v, nil
}

func main() {
	// 使用 P1 来检查错误并确保指针非 `nil`
	ptr := rese.P1(getSomething()) 
	fmt.Println(*ptr) // 输出: 42
}
```

### 示例 3: 检查零值，使用 `C1`

```go
package main

import (
	"fmt"
	"github.com/yyle88/rese"
)

func getInt() (int, error) {
	return 20, nil
}

func main() {
	// 使用 C1 来检查错误并确保非零值
	num := rese.C1(getInt())
	fmt.Println("Received:", num) // 输出: 20
}
```

## 许可协议

此项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE) 文件。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉
