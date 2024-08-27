## Strings в Golang

### What is the strings in Go

String in go this is array of bytes
```
type _string struct {
	elements *byte // underlying bytes
	len      int   // number of bytes
}
```

### How do I determine the number of characters for a string?" or "What are the nuances when iterating over a string?

String this is array of bytes.
And in standard encoding we can use: `len(string)`
But sometimes one symbol this is more than one byte and in this case we need cast it to runes:
`len([]rune(string))`
So, in the biggest part of cases we can use standard cod

utf8.RuneCountInString(str)