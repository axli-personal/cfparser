# Config File Parser

## Speed

It was Implemented by binary tree and only suitable for small project.

## Ignore

Any line starting with specific prefix will be ignored.

## Signal

You can receive signal when config change immediately by channel.

## Example

```go
parser := NewCFParser(file, "#", "=") // any specific prefix.
validCount := parser.ReadAll()
fmt.Printf("read %v valid config", validCount)
if pair := parser.Get("key"); pair != nil {
	// any api you like.
}
```
