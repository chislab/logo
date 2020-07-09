打印chislab的logo
```golang
package main
import (
    "time"
    "github.com/chislab/logo"
)

func main() {
    //每行延迟100毫秒
    logo.Print(100 * time.Millisecond)
}
```