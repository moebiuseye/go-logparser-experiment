package main
import (
    "github.com/vjeantet/grok"
    "os"
    "fmt"
    "bufio"
)

var err error

func routine (haystack string) map[string]string {
    g, _ := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
    // Example
    // 127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207
    values, _ := g.Parse("%{COMMONAPACHELOG}", haystack)
    return values
}

func main () {
    reader := bufio.NewReader(os.Stdin)
    for err == nil {
        str,err:= reader.ReadString('\n')
        fmt.Printf("%v -- %v\n", str, err)
        result := routine(str)
        fmt.Printf("= %#v\n", result)
    }
}
