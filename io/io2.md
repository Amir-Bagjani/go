Moving data around is one of the most common things you do when programming. That’s why no matter the language you use, you must master handling the data flow.

### Everything is a bunch of bytes
Someone was complaining on Linkedin that it’s confusing that len gives the number of bytes, not the number of characters of a string:

```go
package main

import "fmt"

func main() {
    str := "ana"
    fmt.Println(len(str)) // 3
    str = "世界"
    fmt.Println(len(str)) // 6 not 2
}
```
This may be confusing to someone because of the wrong mental model. In computing, every piece of data is just a series of bytes. If you want its length, then you need the number of bytes, not the number of interpreted things in that data.

If I load the below image in Go and run len on the data, I will not get 1 because there’s 1 cat in the picture, but 293931 because the image has 293931 bytes.

###### Remember: every piece of data you are manipulating, whether that’s a string, an image, or a file is just a bunch of bytes.

Readers and Writers
Data movement comprises two fundamental parts: reading and writing. You read from somewhere, apply some transformations or whatnot, and later on, you write that final data somewhere.

Because reading and writing are fundamental operations, Go offers us two abstractions: io.Reader and io.Writer. From here on, for the sake of readability, I will refer to them as Reader and Writer.

Remember: a Reader is something you can read from, and a Writer is something you can write to.

I emphasized that because when I started with Go, this naming confused me.
When I hear about a writer, I think of someone who writes. But in Go, a Writer is something you can write to.

When working with I/O in Go, you don’t care what something is, but whether you can read from it or write to it.
As long as it implements the Reader interface, you can read from it, and if it implements the Writer interface, you can write to it.

If you have some bytes, you can write them to any Writer. It doesn’t matter if the bytes represent a file, a network connection, an HTTP response, or something else.

The minimal I/O toolkit
Now let’s add the first types and functions to our toolset as promised.

#### bytes.Buffer
As I said above, every piece of data is just a bunch of bytes. Usually, that’s represented as a`[]byte`, but the slice of bytes does not implement Reader and Writer interfaces.

One type I often use is the `bytes.Buffer` because it’s like a `[]byte` that implements both Reader and Writer interfaces simplifying reading and writing to it.

#### io.Copy
Do you have a Reader and want to copy its contents to a Writer? io.Copy does the trick!

Bellow we create a buffer, add a string in it, and then copy that data to `os.Stdout` which is just a file in the end:

```go
package main
import (
    "bytes"
    "io"
    "os"
)
func main() {
    var buf bytes.Buffer
    buf.WriteString("hello world\n")

    io.Copy(os.Stdout, &buf)
}
```

#### fmt.Fprint
Sometimes you just want to print some strings to a Writer. For this, you can use fmt.Fprint% functions:

```go
package main

import (
    "bytes"
    "fmt"
    "os"
)

func main() {
    var buf bytes.Buffer

    // Fprintf can be called to print a string to any io.Writer

    fmt.Fprintf(&buf, "hello world!")         // on bytes.Buffer
    fmt.Println(buf.String())

    fmt.Fprintf(os.Stdout, "hello world!\n")  // on a *os.File
}
```

##### example for server health check

```go
package handlers

import (
    "fmt"
    "net/http"
)

func HealthHandler() http.Handler {
    return http.HandlerFunc(
       func(w http.ResponseWriter, r *http.Request) {
          fmt.Fprintln(w, "OK") // write OK string to response body
       },
    )
}
```

#### io.ReadAll
This method is useful when you want to read all bytes from a certain Reader:

``` go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello World!")

    data, _ := io.ReadAll(r)

    fmt.Println(data)         // [72 101 108 108 111 32 87 111 114 108 100 33]
    fmt.Println(string(data)) // Hello World!
}
```

If you’re dealing with a huge file, it might not always be a good idea to read the full content at once. We will see in future posts what alternatives we have. However, io.ReadAll is a very handy function to have in your toolkit.

##### Conclusion
To recap, here’s the most important information covered in this article:

At the fundamental level, every piece of data is just a bunch of bytes.

Most of what we do when programming is moving data around.

Moving data is composed of two parts: reading and writing. Go offers us io.Reader and io.Writer abstractions for these operations.

A Reader is something you can read from, and a Writer is something you can write to.

As long as something implements io.Reader you can read from it and if it implements io.Writer you can write to it.

We also added a bunch of useful types and functions to our toolkit:

`io.Reader and io.Writer`
`bytes.Buffer`
`io.Copy`
`fmt.Fprintf and fmt.Fprintln`
`io.ReadAll`


##### io.LimitReader

Sometimes, you want to limit the number of bytes you read from a piece of data.

To limit the size of data you read, you can wrap that data source (Reader) with an io.LimitReader function that returns another Reader from which you can read only n bytes:

``` go
package main

import (
    "io"
    "log"
    "os"
    "strings"
)

func main() {
    r := strings.NewReader("Hello!")

    lr := io.LimitReader(r, 4)

    // Ouput: Hell
    if _, err := io.Copy(os.Stdout, lr); err != nil {
        log.Fatal(err)
    }
}
```

##### io.MultiReader

Sometimes, you can have multiple data sources, and you want to treat those sources as one. If you have multiple Readers, you can merge them into one Reader with io.MultiReader:

``` go
package main

import (
    "io"
    "log"
    "os"
    "strings"
)

func main() {
    r1 := strings.NewReader("first reader\n")
    r2 := strings.NewReader("second reader\n")
    r3 := strings.NewReader("third reader\n")

    // merge all 3 readers into one
    r := io.MultiReader(r1, r2, r3)

    if _, err := io.Copy(os.Stdout, r); err != nil {
       log.Fatal(err)
    }

}
```
I recently used io.MultiReader, when I needed to write the bytes coming from different requests into one file in a parallel downloader implementation:

``` go 
results := make([]io.Reader, n)
// ...
if err := writeToFile(destinationFileName, io.MultiReader(results...)err != nil {
 return fmt.Errorf("could not write to file: %w", err)
}
```

##### io.MultiWriter

Similar to io.MultiReader, we have an io.MultiWriter function, which creates a writer that reproduces its writes to all the provided writers:
``` go 
package main

import (
    "bytes"
    "fmt"
    "io"
    "strings"
)

func main() {
    var (
       buf1 bytes.Buffer
       buf2 bytes.Buffer
    )

    // writing into mw will write to both buf1 and buf2
    mw := io.MultiWriter(&buf1, &buf2)

    // r is the source of data(Reader)
    r := strings.NewReader("some io.Reader stream to be read")

    // write to mw from r
    io.Copy(mw, r)

    fmt.Println("data inside buffer1 :", buf1.String())
    fmt.Println("data inside buffer2 :", buf2.String())

}
```

I like to use `io.MultiWriter` when I'm trying to debug what was written in a certain Writer.
If a function writes to a Writer and, for some reason, it is too hard to get those contents, I connect a `bytes.Buffer` to it, and then I check the contents of the buffer, which will be the same as the contents of my inaccessible Writer:

``` go
package main

import (
    "bytes"
    "fmt"
    "io"
)

func main() {
    buf := new(bytes.Buffer)
    weirdWriter := new(bytes.Buffer)

    debug := io.MultiWriter(buf, weirdWriter) // attach buf to weirdWriter

    complicatedFunctionWithAWriter(debug) // a function what normally used weirdWriter

    // The contents of the buffer will be the same as in weirdWriter
    fmt.Println(buf.String())
    fmt.Println(buf.Bytes())

}

func complicatedFunctionWithAWriter(w io.Writer) {
    fmt.Fprintf(w, "i'm writing something")
}
```

##### io.TeeReader

Imagine reading and writing data in a place, and you want to write the same data somewhere else.

In the above image, we read from R to W and write to an extra Logs Writer at the same time.

Let’s see how that looks in code:
``` go 
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    logs := new(bytes.Buffer)

    data := strings.NewReader("Hello World!\n")

    teeReader := io.TeeReader(data, logs)

    // logs will also receives contents from teeReader
    io.Copy(os.Stdout, teeReader)

    fmt.Println("Content of logs:", logs.String())
}
```

##### io.Pipe

Speaking of Linux and plumbing, you might be familiar with the pipe operator, which combines two or more commands so that the output of one becomes the input of the other.

``` sh
echo hello | tr l y
```
This outputs:

``` sh 
heyyo
```
I like the pipe's universality. It doesn't matter what programs you connect as long as one writes and the other reads. I could have just as easily used `cat` to fetch the contents of a file instead of `echo`:

``` sh
cat file | tr l y
```
In Go, we achieve this behavior with io.Pipe that can be used to connect code expecting an io.Reader with code expecting an io.Writer.

Let's try to replicate the same `echo hello | tr l y` example to see how that works:

``` go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    pipeReader, pipeWriter := io.Pipe()

    echo(pipeWriter, "hello")
    tr(pipeReader, "e", "i")
}

func echo(w io.Writer, s string) {
    fmt.Fprint(w, s)
}

func tr(r io.Reader, old string, new string) {
    data, _ := io.ReadAll(r)
    res := strings.Replace(string(data), old, new, -1)
    fmt.Println(res)
}
```
Running this program, we get an error:

``` sh
fatal error: all goroutines are asleep - deadlock!
```

To understand why, we need to read the documentation of the io.Pipe function:

``` sh
// Pipe creates a synchronous in-memory pipe.
// That is, each Write to the [PipeWriter] blocks until it has satisfied
// one or more Reads from the [PipeReader] that fully consume
// the written data.
```

Our code is not synchronous but sequential: we call `echo` and then `tr`. We get a deadlock when we try to write because no reading is happening.

Let's fix that:

``` go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    pipeReader, pipeWriter := io.Pipe()

    // Run echo concurrently with tr in a separate goroutine
    go echo(pipeWriter, "hello")
    tr(pipeReader, "e", "i")
}

func echo(w io.Writer, s string) {
    fmt.Fprint(w, s)
}

func tr(r io.Reader, old string, new string) {
    data, _ := io.ReadAll(r)
    res := strings.Replace(string(data), old, new, -1)
    fmt.Println(res)
}
```

Running the modified program gives us the same error:

``` go
fatal error: all goroutines are asleep - deadlock!
```
Why is that? Let's check the `PipeReader` documentation on the Read method:

``` go 
// Read implements the standard Read interface:
// it reads data from the pipe, blocking until a writer
// arrives or the write end is closed.
// If the write end is closed with an error, that error is
// returned as err; otherwise err is EOF.
func (r *PipeReader) Read(data []byte) (n int, err error) {
    return r.pipe.read(data)
}
```

The second part of this sentence is key: blocking until a writer arrives or the write-end is closed.

The read is still blocked because the writer is not closed after writing our data. Let's fix that by calling `Close()` on the `pipeWriter`:

``` go
package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    pipeReader, pipeWriter := io.Pipe()

    go func() {
       echo(pipeWriter, "hello")
       // we close the writer so we unblock the reader
       pipeWriter.Close()
    }()
    tr(pipeReader, "e", "i")
}

func echo(w io.Writer, s string) {
    fmt.Fprint(w, s)
}

func tr(r io.Reader, old string, new string) {
    data, _ := io.ReadAll(r)
    res := strings.Replace(string(data), old, new, -1)
    fmt.Println(res)
}
```