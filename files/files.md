#### File operations in Go
Go's file handling is quite straightforward due to its built-in package os; it provides access to most of the operating system's features, including the file system. It allows you to perform file operations without needing to change the code for it to work with different operating systems.

The packages enable you to perform various file operations in your applications, from writing, reading, and creating files to creating and deleting directories. It also provides helpful error messages whenever it encounters errors while performing file operations.

We’ll explore how to read files in Go in the next section.

#### Reading files in Go
Reading files is probably the most frequent file operation you'll perform in Go, as many use cases require it. In this section, we will explore how to read different types of files in Go.

Let's start with plain text files.

Reading `.txt` files in Go
The os package provides a ReadFile function that makes reading files straightforward in Go. For example, I have a data.txt file in my project folder, and I can read and print it out with the following code:

###### it uses`os.Open` under the hood
``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    filepath := "data.txt"
    data, err := os.ReadFile(filepath)
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println(string(data))
}
```
The code above defines a `filepath` variable with a string value of the file that I'm trying to read, `data.txt`. It then reads the file with the `ReadFile` function and stores the result and error in the `data` and `err` variables. Finally, it checks for errors before printing the result in string format by wrapping it with the `string` helper function. It should return the following output in the terminal:

``` sh
1. fmt
2. net/http
3. os
4. io
5. time
6. encoding/json
7. strings
8. strconv
9. sync
10. testing
11. database/sql
12. encoding/csv
13. regexp
14. math
15. bufio
16. flag
17. path/filepath
18. image
19. html/template
20. crypto
```

And that's it! You've just read your first file in Go! Let's explore how to read a file line-by-line in the next section.

### Reading a `log` file line-by-line in Go
In this section, we will explore a real-life use case of reading files in Go. Let's imagine that you're building an application that helps users deploy Go applications. You need to have a way for users to see the deployment errors so that they have an idea of what to fix before trying to re-deploy their application.

We will write a code that reads a log file line-by-line and prints only the amount of lines requested by the user. This way, users can choose to see only the last five logs and don't have to scroll through all the logs to find the issue with their deployment.

I have a sample `log.txt `file that looks like this:
``` sh
2023-07-11 10:00:00 - Successful: operation completed.
2023-07-11 10:05:12 - Error: Failed to connect to the database.
2023-07-11 10:10:32 - Successful: data retrieval from API.
2023-07-11 10:15:45 - Error: Invalid input received.
2023-07-11 10:20:58 - Successful: file upload.
2023-07-11 10:25:01 - Error: Authorization failed.
2023-07-11 10:30:22 - Successful: record update.
2023-07-11 10:35:37 - Error: Internal server error.
2023-07-11 12:45:59 - Error: Server overloaded.
2023-07-11 12:50:06 - Successful: session created.
2023-07-11 12:55:17 - Error: Invalid input parameters.
2023-07-11 13:00:30 - Successful: software update installed.
2023-07-11 13:05:46 - Error: Access denied.
2023-07-11 13:10:53 - Successful: report generated.
2023-07-11 13:16:01 - Error: Unexpected exception occurred.
2023-07-11 13:20:13 - Successful: user registration.
2023-07-11 13:25:28 - Error: Disk read/write failure.
```
The code for doing this will look like this:

`os.Open` uses `os.OpenFile` under the hood
``` go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func printLastNLines(lines []string, num int) []string {
    var printLastNLines []string
    for i := len(lines) - num; i < len(lines); i++ {
        printLastNLines = append(printLastNLines, lines[i])
    }
    return printLastNLines
}

func main() {
    filepath := "log.txt"
    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println("Error opening file:", err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())

    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

    // print the last 10 lines of the file
    printLastNLines := printLastNLines(lines, 3)
    for _, line := range printLastNLines {
        fmt.Println(line)
        fmt.Println("________")
    }
}
```
The code above uses the `os`' package `Open` function to open the file, defers its `Close` function with the `defer` keyword, defines an empty `lines` slice, and uses the `bufio`'s `NewScanner` function to read the file line-by-line while appending each line to the `lines` array in a text format using the `Text` function. Finally, it uses the `printLastNLines` function to get the last `N` lines of the `lines` array. N is any number of the user's choosing. In this case, it is 3, and the code uses a `for` loop to print each line with an horizontal line between each one.

###### Note: The defer keyword ensures that the file closes if there is an error; if not, it closes the file after the function has finished running.

Using the sample log.txt file, the code above should return the following:

``` go
2023-07-11 13:16:01 - Error: Unexpected exception occurred.
________
2023-07-11 13:20:13 - Successful: user registration.
________
2023-07-11 13:25:28 - Error: Disk read/write failure.
________

```
Next, let's explore how to read `.json` files in Go.

#### Reading `.json` files in Go
Reading and using data from `.json` files is also a popular use case in programming, so let's learn how to do that. For example, you have a `.json` configuration file that looks like the following:

``` json
{
    "database_host": "localhost",
    "database_port": 5432,
    "database_username": "myuser",
    "database_password": "mypassword",
    "server_port": 8080,
    "server_debug": true,
    "server_timeout": 30
}
```
You also need to read for your application to work correctly, which can be done by reading and decoding the `.json` file:

``` go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Config struct {
    DBHost        string `json:"database_host"`
    DBPort        int    `json:"database_port"`
    DBUsername    string `json:"database_username"`
    DBPassword    string `json:"database_password"`
    ServerPort    int    `json:"server_port"`
    ServerDebug   bool   `json:"server_debug"`
    ServerTimeout int    `json:"server_timeout"`
}

func main() {
    filepath := "config.json"
    var config Config

    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(config)
}
```
The code above uses the `os`' package `Open` function to open the file, defers its `Close` function with the defer keyword, and uses the `json`'s `NewDecoder` function to read the file. It then uses the `Decoder` function to decode it into an object and checks for errors before printing the `config` details in the terminal:

```
{localhost 5432 myuser mypassword 8080 true 30}
```
That's it! You can now split the `config` object up and use it however you like in your application.

Next, let's look at how to read a `.csv` file.

#### Reading `.csv` files in Go
Comma-separated values (CSV) is one of the most popular file formats. Let's explore how to read the following .csv file in this section:

``` sh
Name,Email,Phone,Address
John Doe,johndoe@example.com,555-1234,123 Main St
Jane Smith,janesmith@example.com,555-5678,456 Elm St
Bob Johnson,bobjohnson@example.com,555-9876,789 Oak St
```
Go provides an `encoding/csv` package that can be used to read `.csv `files:

``` go
package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    filepath := "data.csv"

    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(records)
}
```
The code above uses the `os`' package `Open` function to open the file, defer its `Close` function with the defer keyword, and create a `reader` variable that stores the result of the `NewReader` function. It then uses the `ReadAll` function to read the file, check for errors, and print the result in the terminal. The output should look like this:

``` sh
[[Name Email Phone Address] [John Doe johndoe@example.com 555-1234 123 Main St] [Jane Smith janesmith@example.com 555-5678 456 Elm St] [Bob Johnson bobjohnson@example.com 555-9876 789 Oak St]]
```

Let's explore how to read bytes from files in the next section.

#### Reading bytes from files in Go

In some cases, you might want to read a specific number of bytes from files in Go:

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    filepath := "data.txt"

    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    data := make([]byte, 10)
    _, err = file.Read(data)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(data)
    fmt.Println(string(data))
}
```
The code above uses the `os`' package `Open` function to open the file, defer its `Close` function with the `defer` keyword, and create a `data` variable that has the value of 10 bytes using the `make` function. It then checks for errors before printing the `data` out in bytes and string formats:

``` sh
[49 46 32 102 109 116 10 50 46 32]
1. fmt
2. 
```
Now that we've explored how to perform various read operations on different file formats in Go, let's explore how to handle write operations.

#### Writing and manipulating files in Go
Reading and writing files is relatively straightforward in Go, as it provides developers with a lot of functions to perform such tasks without needing to download a third-party library. In this section, we will explore different Go writing tasks and how to handle them properly.

#### Creating a file in Go
Go's os package provides a Create function that creates files with any extension. For example, suppose you have an application that helps users to deploy applications. In this case, you might want to create an empty log.txt file containing the logs of the application immediately after the user creates a new project:

Create uses `OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)` under the hood

``` go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Create("log.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fmt.Print("File created successfully")
}
```

The code above uses the `os`'s Create function to create an empty `log.txt` file, checks for any error, and prints a success message to the user. The user should see the following output in the terminal:

``` sh
File created successfully
```
Before continuing with writing files in Go, let's explore the different file opening flags in Go, as we will use some of them in the later in the article.

#### File-opening flags in Go
Go provides file-opening flags represented by constants defined in the os package. These flags determine the behavior of file operations, such as opening, creating, and truncating files. The following is a list of the flags and what they do.

1. os.O_RDONLY: Opens the file as read-only. The file must exist.

2. os.O_WRONLY: Opens the file as write-only. If the file exists, its contents are truncated. If it doesn't exist, a new file is created.

3. os.O_RDWR: Opens the file for reading and writing. If the file exists, its contents are truncated. If it doesn't exist, a new file is created.

4. os.O_APPEND: Appends data to the file when writing. Writes occur at the end of the file.

5. os.O_CREATE: Creates a new file if it doesn't exist.

6. os.O_EXCL: Used with O_CREATE, it ensures that the file is created exclusively, preventing creation if it already exists.

7. os.O_SYNC: Open the file for synchronous I/O operations. Write operations are completed before the call returns.

8. os.O_TRUNC: If the file exists and is successfully opened, its contents are truncated to zero length.

9. os.O_NONBLOCK: Opens the file in non-blocking mode. Operations like read or write may return immediately with an error if no data is available or the operation would block.

These flags can be combined using the bitwise OR (|) operator. For example, `os.O_WRONLY|os.O_CREATE` would open the
file for writing, creating it if it doesn't exist.

When using these flags, it's important to check for errors returned by file operations to handle cases where the file cannot be opened or created as expected.

Let's look at how to write text to files in the next section.

#### Writing text to files in Go
The `os` package also provides a `WriteString` function that helps you write strings to files. For example, you want to update the `log.txt` file with a log message:
``` go
package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    data := "2023-07-11 10:05:12 - Error: Failed to connect to the database. _________________"
    _, err = file.WriteString(data)
    if err != nil {
        log.Fatal(err)
    }

}
```
The code above uses the `OpenFile` function to open the `log.txt` file in write-only mode and creates it if it doesn't exist. It then creates a `data` variable containing a string and uses the `WriteString` function to write string data to the file.

#### Appending to a file in Go
The code in the previous section deletes the data inside the file before writing the new data every time the code is run, which is acceptable in some cases. However, for a log file, you want it to retain all the previous logs so that the user can refer to them as many times as needed to, for example, perform analytics.

You can open a file in append mode like this:
``` go 
package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    data := "\n 2023-07-11 10:05:12 - Error: Failed to connect to the database.\n __________________ \n"
    _, err = file.WriteString(data)
    if err != nil {
        log.Fatal(err)
    }

}
```
The code above uses the `os.O_APPEND` to open the file in append mode and will retain all the existing data before adding new data to the `log.txt` file. You should get an updated file each time you run the code instead of a new file.

#### Writing bytes to files in Go
Go allows you to write bytes to files as strings with the `Write` function. For example, if you are streaming data from a server and it is returning bytes, you can write the bytes to a file to be readable

``` go
package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("data.bin", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21, 0x0A}
    _, err = file.Write(data)
    if err != nil {
        log.Fatal(err)
    }

}
```
The code above opens the `data.bin` file in write-only and append mode and creates it if it doesn't already exist. The code above should return a `data.bin` file containing the following:

``` sh 
Hello, World!
```

Next, let's explore how to write formatted data to a file section.

#### Writing formatted data to a file in Go
This is one of the most common file-writing tasks when building software applications. For example, if you are building an e-commerce website, you will need to build order confirmation receipts for each buyer, which will contain the details of the user's order. Here is how you can do this in Go:

``` go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    username, orderNumber := "Adams_adebayo", "ORD6543234"
    file, err := os.Create(username + orderNumber + ".pdf")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    item1, item2, item3 := "shoe", "bag", "shirt"
    price1, price2, price3 := 1000, 2000, 3000

    _, err = fmt.Fprintf(file, "Username: %s\nOrder Number: %s\nItem 1: %s\nPrice 1: %d\nItem 2: %s\nPrice 2: %d\nItem 3: %s\nPrice 3: %d\n", username, orderNumber, item1, price1, item2, price2, item3, price3)
    if err != nil {
        log.Fatal(err)
    }

}
```

The code above defines two variables, `username` and `orderNumber`, creates a `.pdf` based on the variables, checks for errors, and defers the `Close` function with the `defer` keyword. It then defines three variables, `item1`, `item2`, and `item3`, formats a message with the `fmt`'s `Fprintf` all the variables, and writes it to the `.pdf` file.

The code above then creates an `Adams_adebayoORD6543234.pdf` file with the following contents:

``` sh
Username: Adams_adebayo
Order Number: ORD6543234
Item 1: shoe
Price 1: 1000
Item 2: bag
Price 2: 2000
Item 3: shirt
Price 3: 3000
```

#### Writing to `.csv` files in Go
With the help of the `encoding/csv` package, you can write data to `.csv` files easily with Go. For example, you want to store new users' profile information in a `.csv` file after they sign up:

``` go
package main

import (
    "encoding/csv"
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("users.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
    data := []string{"Adams Adebayo", "30", "Lagos"}
    err = writer.Write(data)
    if err != nil {
        log.Fatal(err)
    }

}
```
The code above opens the `users.csv` file in write-only and append mode and creates it if it doesn't already exist. It will then use the `NewWriter` function to create a `writer` variable, defer the `Flush` function, create a `data` variable with the `string` slice, and write the data to the file with the `Write` function.

The code above will then return a `users.csv` file with the following contents:

``` sh
Adams Adebayo,30,Lagos
```
#### Writing JSON data to a file in Go
Writing JSON data to `.json` files is a common use case in software development. For example, you are building a small application and want to use a simple `.json` file to store your application data:

``` go
package main

import (
    "encoding/json"
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("users.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    data := map[string]interface{}{
        "username": "olodocoder",
        "twitter":  "@olodocoder",
        "email":    "hello@olodocoder.com",
        "website":  "https://dev.to/olodocoder",
        "location": "Lagos, Nigeria",
    }

    encoder := json.NewEncoder(file)
    err = encoder.Encode(data)
    if err != nil {
        log.Fatal(err)
    }

}
```
The code above opens the `users.csv` file in write-only and append mode and creates it if it doesn't already exist, defers the `Close` function, and defines a `data` variable containing the user data. It then creates an `encoder` variable with the `NewEncoder` function and encodes it with the `Encoder` function.

The code above then returns a `users.json` file containing the following:
``` json
{"email":"hello@olodocoder.com","location":"Lagos, Nigeria","twitter":"@olodocoder","username":"olodocoder","website":"https://dev.to/olodocoder"}
```
#### Writing XML data to files in Go
You can also write XML data to files in Go using the `encoding/xml` package:
``` go
package main

import (
    "encoding/xml"
    "log"
    "os"
)

func main() {
    type Person struct {
        Name string `xml:"name"`
        Age  int    `xml:"age"`
        City string `xml:"city"`
    }

    file, err := os.OpenFile("users.xml", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    data := Person{
        Name: "John Doe",
        Age:  30,
        City: "New York",
    }

    encoder := xml.NewEncoder(file)
    err = encoder.Encode(data)
    if err != nil {
        log.Fatal(err)
    }

}
```
The code above defines a `Person` struct with three fields, opens the `users.xml` file in write-only and append mode and creates it if it doesn't already exist, defers the `Close` function, and defines a `data` variable that contains the user data. It then creates an encoder variable with the `NewEncoder` function and encodes it with the Encoder function.

The code above should return a `user.xml` file that contains the following contents:

```xml
<Person><name>John Doe</name><age>30</age><city>New York</city></Person>
```

#### Renaming files in Go
Go enables you to rename files from your code using the `Rename` function:

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Rename("users.xml", "data.xml")
    if err != nil {
        fmt.Println(err)
    }

}
```
The code above renames the `users.xml` file created in the previous section to `data.xml`.

#### Deleting files in Go
Go enables you to delete files with the `Remove` function:

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Remove("data.bin")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("File deleted")
}
```
The code above deletes the `data.bin` file from the specified path.

Now that you understand how to write and manipulate different types of files in Go, let's explore how to work with directories.

#### Working with directories in Go
In addition to files, Go also provides functions that you can use to perform different tasks in applications. We will explore some of these tasks in the following sections.

#### Creating a directory
Go provides a `Mkdir` function that you can use to create an empty directory:

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Mkdir("users", 0755)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Directory Created Successfully")
}
```

The code above creates a `users` folder in the current working directory.

#### Creating multiple directories in Go
You can create multiple directories in Go using the `MkdirAll` function:

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.MkdirAll("data/json_data", 0755)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Directory Created Successfully")
}
```
The code above will create a `data` directory and a `json_data` directory inside it.

Note: If a data directory already exists, the code will only add a `json_data` directory inside it.

#### Checking if a directory exists in Go
To avoid errors, checking if a directory exists before creating a file or directory inside is good practice. You can use the `Stat` function and the `IsNotExist` function to do a quick check:

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    if _, err := os.Stat("data/csv_data"); os.IsNotExist(err) {
        fmt.Println("Directory does not exist")
    } else {
        fmt.Println("Directory exists")
    }

}
```
The code above returns a message based on the results of the check. In my case, it will return the following:

``` shell
Directory exists
```
#### Renaming directories in Go
You can also use the `Rename` function to rename directories:

``` go 
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Rename("data/csv_data", "data/xml_data")
    if err != nil {
        fmt.Println(err)
    }

}
```
The code above renames the `data/csv_data` directory to `data/xml_data`.

#### Deleting an empty directory in Go
You can use the `Remove` function to delete folders in your applications:

``` go 
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.Remove("data/json_data")
    if err != nil {
        fmt.Println(err)
    }

}
```
The code above removes the `json_data` directory from the `data` directory.

#### Deleting a directory with all its content in Go
Go provides a `RemoveAll` function that allows you to remove all the directories and everything inside them, including files and folders:
``` go 
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.RemoveAll("users")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("users directory and all it's content has been removed")
}
```
The code above deletes the `users` directory and everything inside it.

Note: It's good practice to check if the directory exists before attempting to delete it.

#### Get a list of files and directories in a directory in Go
You can retrieve a list of all the files and directories in a directory using the `ReadDir` function:
``` go 
package main

import (
    "fmt"
    "os"
)

func main() {
    dirEntries, err := os.ReadDir("data")
    if err != nil {
        fmt.Println(err)
    }

    for _, entry := range dirEntries {
        fmt.Println(entry.Name())
    }
}
```
The code above returns a list of all the directories and files inside the `data` folder.

Now that you know how to work with directories in Go applications, let's explore some of the advanced file operations in the next section.

#### Advanced file operations in Go
In this section, we will explore some of the advanced file operations you might encounter in Go applications.

#### Writing compressed data to a file in Go
Working with compressed files is uncommon, but here's how to create a `.txt` file inside a compressed file using the `compress/gzip` package:

``` go 
package main

import (
    "compress/gzip"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("data.txt.gz", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    gzipWriter := gzip.NewWriter(file)
    defer gzipWriter.Close()

    data := "Data to compress"
    _, err = gzipWriter.Write([]byte(data))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File compressed successfully")
}
```
The code above creates a `data.txt.gz`, which contains a `data.txt` file in the working directory.

#### Writing encrypted data to a file in Go
When building applications that require secure files, you can create an encrypted file with Go's `crypto/aes` and `crypto/cipher` packages:

``` go 
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "log"
    "os"
)

func main() {
    // file, err := os.OpenFile("encrypted.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    file, err := os.Create("encrypted.txt")
    if err != nil {
        log.Fatal(err)
        fmt.Println("Error")
    }
    defer file.Close()

    key := []byte("cacf2ebb8cf3402964356547f20cced5")
    plaintext := []byte("This is a secret! Don't tell anyone!🤫")

    block, err := aes.NewCipher(key)
    if err != nil {
        log.Fatal(err)
        fmt.Println("Error")
    }

    ciphertext := make([]byte, len(plaintext))
    stream := cipher.NewCTR(block, make([]byte, aes.BlockSize))
    stream.XORKeyStream(ciphertext, plaintext)

    _, err = file.Write(ciphertext)
    if err != nil {
        log.Fatal(err)
        fmt.Println("Error")
    }
    fmt.Println("Encrypted file created successfully")
}
```
The code above creates an `encrypted.txt` file containing an encrypted version of the plaintext string:

``` sh 
?Э_g?L_.?^_?,_?_;?S???{?Lؚ?W4r
W?8~?
```
#### Copying a file to another directory in Go
Copying existing files to different locations is something we all do frequently. Here's how to do it in Go:

``` go 
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    srcFile, err := os.Open("data/json.go")
    if err != nil {
        fmt.Println(err)
    }
    defer srcFile.Close()

    destFile, err := os.Create("./json.go")
    if err != nil {
        fmt.Println(err)
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, srcFile)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Copy done!")
}
```
The code above copies the `json.go` file in the data directory and its contents and then creates another `json.go` with the same in the root directory.

#### Get file properties in Go
Go allows you to get the properties of a file with the `Stat` function:

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    fileInfo, err := os.Stat("config.json")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("File name:", fileInfo.Name())
    fmt.Println("Size in bytes:", fileInfo.Size())
    fmt.Println("Permissions:", fileInfo.Mode())
    fmt.Println("Last modified:", fileInfo.ModTime())

    fmt.Println("File properties retrieved successfully")
}
```
The code above returns the name, size, permissions, and last modified date of the `config.json` file:

``` sh 
File name: config.json
Size in bytes: 237
Permissions: -rw-r--r--
Last modified: 2023-07-11 22:46:59.705875417 +0100 WAT
File properties retrieved successfully
```
#### Get the current working directory path in Go
You can get the current working directory of your application in Go:

``` go 
package main

import (
    "fmt"
    "os"
)

func main() {
    wd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Current working directory:", wd)

}
```
The code above will return the full path of my current working directory:

###### Current working directory: /Users/user12/Documents/gos/go-files