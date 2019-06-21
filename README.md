# Mark some test as hardcore, So that it won't be run by default

```go

func TestIsEmailValidForHost(t *testing.T) {
	hardcoretest.Mark(t)
	for _, c := range isEmailValidForHostCases {
		isValid := helpers.IsEmailValidForHost(c.email)
		assert.Equal(t, c.expectedIsValid, isValid, c.email)
	}
}


```

This test can only be run by:

```
HARDCORE=1 go test -v .
```
