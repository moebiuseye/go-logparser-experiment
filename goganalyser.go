package main

import "os"
import "fmt"
import "bufio"
import "regexp"

var err error
func routine (haystack string) []string {
    reg := regexp.MustCompile(`^(?P<word1>[^ ]*) (?P<word2>[^ ]*)`)
    
    result := reg.FindStringSubmatch(haystack)
    for i, name := range reg.SubexpNames() {
        fmt.Printf("'%s'\t %d -> %s\n", name, i, result[i])
    }
    return result
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
