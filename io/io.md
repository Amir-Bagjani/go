# 1 Write to standard output
The most common example every Go programming tutorial teaches you:

``` go
package main
import "fmt"
func main() {
	fmt.Println("Hello Medium")
}
```
But, no one tells you the above program is a simplified version of this one:

```go
package main
import (
	"fmt"
	"os"
)
func main() {
	fmt.Fprintln(os.Stdout, "Hello Medium")
}
```
The Fprintln method takes an `io.Writer` type and a string to write into a writer. The `os.Stdout` satisfies `io.Writer` interface.

This example is excellent and can extend to any writer besides `os.Stdout`.

# 2 Write to a custom writer
Let’s create a custom writer and store some information there. One can do it by initializing an empty buffer and writing content to it.

```go
package main
import (
	"bytes"
	"fmt"
)
func main() {
	// Empty buffer (implements io.Writer)
	var b bytes.Buffer
	fmt.Fprintln(&b, "Hello Medium") // Don't forget &

	// Optional: Check the contents stored
	fmt.Println(b.String()) // Prints `Hello Medium`
}
```

# 3 Write to multiple writers at once
Sometimes, one needs to write a string into multiple writers. We can easily do that using the MultiWriter method from the io package.

```go
package main
import (
	"bytes"
	"fmt"
	"io"
)
func main() {
	// Two empty buffers
	var foo, bar bytes.Buffer

	// Create a multi writer
	mw := io.MultiWriter(&foo, &bar)

	// Write message into multi writer
	fmt.Fprintln(mw, "Hello Medium")

	// Optional: verfiy data stored in buffers
	fmt.Println(foo.String())
	fmt.Println(bar.String())
}
```

# 4 Create a simple reader
An I/O reader helps hold information that can retrieve with API

Go provides `io.Reader` interface to implement an IO reader. A reader does not read but provides data for others. It is a temporary store for information with many methods like WriteTo, Seek, etc.

```go
package main
import (
	"fmt"
	"io"
	"strings"
)
func main() {
	// Create a new reader (Readonly)
	r := strings.NewReader("Hello Medium")

	// Read all content from reader
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	// Optional: verify data
	fmt.Println(string(b))
}
```
Note: `os.Stdin` is a commonly used reader for collecting standard input.

# 5 Read from multiple readers at once

Similar to io.MultiWriter, we can also create an io.MultiReader to read data from multiple readers. The data is collected sequentially in the order of readers passed to `io.MultiReader`. It is like gathering information from various data stores at once but in the given order.

```go
package main
import (
	"fmt"
	"io"
	"strings"
)
func main() {
	// Create two readers
	foo := strings.NewReader("Hello Foo\n")
	bar := strings.NewReader("Hello Bar")

	// Create a multi reader
	mr := io.MultiReader(foo, bar)

	// Read data from multi reader
	b, err := io.ReadAll(mr)

	if err != nil {
		panic(err)
	}

	// Optional: Verify data
	fmt.Println(string(b))
}
```
##### Note: Don’t use `io.ReadAll` for big buffers, as they can choke memory.


### Copying data from a reader to writer


Reader: From whom I can copy data

Writer: To whom I can write data to

These definitions make it easy to figure out that we need to load data from a reader (string reader) and dump it into a writer (like os.Stdout or a buffer). This copy process can happen in two ways:

1. Reader pushes data to a writer
2. The Writer pulls data from a reader


# 6 Reader pushes data to a writer (copy variant 1)

This part explains the first copy variation, i.e., the reader pushes data into a writer. It uses the `reader.WriteTo(writer)` API.

```go
package main
import (
	"bytes"
	"fmt"
	"strings"
)
func main() {
	// Create a reader
	r := strings.NewReader("Hello Medium")

	// Create a writer
	var b bytes.Buffer

	// Push data
	r.WriteTo(&b) // Don't forget &

	// Optional: verify data
	fmt.Println(b.String())
}
```

# 7 Writer pulls data from a reader (copy variant 2)
The method `writer.ReadFrom(reader)` is used by a writer to pull data from a given reader. Let’s see an example:

```go
package main
import (
	"bytes"
	"fmt"
	"strings"
)
func main() {
	// Create a reader
	r := strings.NewReader("Hello Medium")

	// Create a writer
	var b bytes.Buffer

	// Pull data
	b.ReadFrom(r)

	// Optional: verify data
	fmt.Println(b.String())
}
```
# 8 Copy data from a reader to writer (copy variant 3, io.Copy)
The `io.Copy` is a utility function that allows one to move data from a reader to a writer without worrying about which variant to pick from above.

```go
package main
import (
	"bytes"
	"fmt"
	"io"
	"strings"
)
func main() {
	// Create a reader
	r := strings.NewReader("Hello Medium")

	// Create a writer
	var b bytes.Buffer

	// Copy data
	_, err := io.Copy(&b, r) // Don't forget &

	if err != nil {
		panic(err)
	}
	// Optional: verify data
	fmt.Println(b.String())
}
```

# 9 Create a data tunnel with io.Pipe

The `io.Pipe` returns a reader and a writer pair, where writing data into a writer automatically allows programs to consume data from the Reader. It is like a Unix pipe.

###### Pipe creates a synchronous in-memory pipe. It can be used to connect code expecting an io.Reader with code expecting an io.Writer.

You must put writing logic into a separate go-routine (from the main go-routine) because the pipe blocks the Writer until the data is read from the Reader, and the Reader is also blocked until the Writer is closed.

```go
package main
import (
	"fmt"
	"io"
)
func main() {
	pr, pw := io.Pipe()

	// Writing data to writer should be in a go-routine
	// because pipe is synchronous.
	go func() {
		defer pw.Close() // Important! To notify writing is done
		fmt.Fprintln(pw, "Hello Medium")
	}()

	// Code is blocked until someone writes to writer and closes it
	b, err := io.ReadAll(pr)

	if err != nil {
		panic(err)
	}
	// Optional: verify data
	fmt.Println(string(b))
}
```

# 10 Capture stdout of a function into a variable with io.Pipe, io.Copy and io.MultiWriter

Let’s say we are building a CLI application. As part of that process, we should tap into the standard output generated by a function(to console) and capture the same information into a variable. How can we do that? We can use the techniques discussed above to create a solution.

```go
package main
import (
	"bytes"
	"fmt"
	"io"
	"os"
)
// Your function
func foo(w *io.PipeWriter) {
	defer w.Close()
	// Write a message to pipe writer
	fmt.Fprintln(w, "Hello Medium")
}

func main() {
	// Create a pipe
	pr, pw := io.Pipe()

	// Pass writer to function
	go foo(pw)

	// Variable to get standard output of function
	var b bytes.Buffer

	// Create a multi writer that is a combination of
	// os.Stdout and our variable byte buffer
	mw := io.MultiWriter(os.Stdout, &b)
	// Copies reader content to standard output
	_, err := io.Copy(mw, pr)

	if err != nil {
		panic(err)
	}

	// Optional: verify data
	fmt.Println(b.String())
}
```

Using the iopackage in Go, we could manipulate data as we have seen in these patterns.