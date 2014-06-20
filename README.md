**NAME**

base64 - Base64 encode/decode in Go

**SYNOPSIS**

base64 [-d] [-s string]

**DESCRIPTION**

Base64 encode or decode string, or standard input, to standard output.

-d      Decode data.
	
-s string
        Encode string.

**EXIT STATUS**

The base64 utility exit 0 on success and 1 if there is an error.

**EXAMPLES**

```
$ cat /bin/echo | base64 | base64 -d > echo.new
$ chmod a+x echo.new
$ ./echo.new test
test
```

**AUTHOR**

Andy Kosela

**SEE ALSO**

http://en.wikipedia.org/wiki/Base64
