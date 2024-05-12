Re-use


We can re-use the same bufio.NewWriterSize for different writers using the reset() method:

``` go
writerOne := new(Writer)
bw := bufio.NewWriterSize(writerOne,2) 
writerTwo := new(Writer)
bw.Reset(writerTwo) 
```


Check available space

We can check available space with the Available() method.


#### Reading with bufio


bufio allows us to read in batches with bufio.Reader. After a read, data is released, as required, from the buffer to the consumer. In the example below, we will look at:

1. Peek
2. ReadSlice
3. ReadLine
4. ReadByte
5. Scanner


##### Peek
The Peek method lets us see the first 𝑛 bytes (referred to as peek value) in the buffer without consuming them. The method operates in the following way.

If the peek value is less than buffer capacity, the characters equal to the peek value are returned.
If the peek value is greater than buffer capacity, bufio.ErrBufferFull is returned.
If the peek value includes EOF and is less than buffer capacity, EOF is returned.


##### ReadSlice
ReadSlice has a signature of:

``` go 
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
```

It returns a slice of the string including the delimiter. For example, if the input is 1, 2, 3 and we use commas as delimiters, the output will be:

```
1,
2,
3
```

If the delimiter cannot be found, and EOF has been reached, then io.EOF is returned. If the delimiter is not reached and readSlice has exceeded buffer capacity, then io.ErrBufferFull is returned


##### ReadLine
ReadLine is defined as:

``` go
ReadLine() (line []byte, isPrefix bool, err error)
```
ReadLine uses ReadSlice under the hood. However, it removes new-line characters (\n or \r\n) from the returned slice.

Note that its signature is different because it returns the isPrefix flag as well. This flag returns true when the delimiter has not been found and the internal buffer is full.

Readline does not handle lines longer than the internal buffer. We can call it multiple times to finish reading.


##### ReadByte
ReadByte has a signature of:

``` go
func (b *Reader) ReadBytes(delim byte) ([]byte, error)
```

Similar to ReadSlice, ReadBytes returns slices before and including the delimiter. In fact, ReadByte works over ReadSlice, which acts as the underlying low-level function. However, ReadByte can call multiple instances of ReadSlice to accumulate return data; therefore, circumventing buffer size limitations. Additionally, since ReadByte returns a new slice of byte, it is safer to use because consequent read operations will not overwrite the data.



##### Scanner
Scanner breaks a stream of data by splitting it into tokens. Scanning stops at EOF, at first IO error, or if a token is too large to fit into the buffer. If more control over error handling is required, use bufio.Reader. Scanner has a signature of:

``` go
func NewScanner(r io.Reader) *Scanner
```

This is the split function used for dividing the text into token defaults to ScanLines; however, you change it if need be​.