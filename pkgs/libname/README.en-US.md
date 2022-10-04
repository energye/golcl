* [中文](README.md)   
* English  

----


**Custom load liblcl or liblcl at the specified location**

> Custom loading dll or so, you need to operate according to the rules of the go initial package.
 
* example:
```go
package main

import "yourpackage" // Must take effect after being sorted automatically before the lcl package.
import "github.com/energye/golcl/lcl"

```
