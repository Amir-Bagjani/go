#### File operations in Go
Go's file handling is quite straightforward due to its built-in package os; it provides access to most of the operating system's features, including the file system. It allows you to perform file operations without needing to change the code for it to work with different operating systems.

The packages enable you to perform various file operations in your applications, from writing, reading, and creating files to creating and deleting directories. It also provides helpful error messages whenever it encounters errors while performing file operations.

We’ll explore how to read files in Go in the next section.

#### Reading files in Go
Reading files is probably the most frequent file operation you'll perform in Go, as many use cases require it. In this section, we will explore how to read different types of files in Go.

Let's start with plain text files.

Reading `.txt` files in Go
The os package provides a ReadFile function that makes reading files straightforward in Go. For example, I have a data.txt file in my project folder, and I can read and print it out with the following code:

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


