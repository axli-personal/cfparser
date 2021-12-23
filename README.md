# Config File Parser

## Speed

It was Implemented by **binary tree** and only suitable for small project.

## Ignore

Any line starting with specific prefix will be ignored.

## Signal

You can receive signal when config change immediately by channel.

## Example

```go
import "github.com/axli-personal/cfparser"

CFP := cfparser.NewCFParser(file, "#", ' ')

numOfValidLines := CFP.ReadAll()

fmt.Printf("read %v valid lines", numOfValidLines)

if pair := CFP.Get("key"); pair != nil {
    val := pair.String()
}
```
