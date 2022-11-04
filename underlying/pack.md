

> go语言的包(package)是多个Go源码的集合，go语言有很多内置包，比如fmt，os，io等。我们也可以自定义包。

> 在一个go语言程序中使用其它包的对象或者函数时，首先要通过 import 引入它



1. 引入包的路径
   第一种方式相对路径：

```
import "./module" // 引入的包在当前文件同一目录的 module 目录，不建议使用此方式。

```

第二种方式绝对路径：


```azure
import “LearnGo/init” // 引入的包在 gopath/src/LearnGo/init 目录。


```


2. 引入包的特殊方式
   下面展示一些特殊的 import 方式。

1) 点操作
   我们有时候会看到如下的方式导入包
```
import . “fmt”

```
这个点操作的含义就是这个包导入之后，在调用这个包的函数时，可以省略前缀的包名。

例如：fmt.Println("hello world") 可以省略的写成 Println("hello world")。

2) 别名操作
   别名操作就是可以把包命名成另一个容易记忆的名字。

import f "fmt"
别名操作的话调用包函数时前缀变成了我们的前缀，即 f.Println("hello world")。

3) _操作
   _操作是一个让很多人费解的操作符，例如：

import _ "github.com/go-sql-driver/mysql"
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的 init 函数。