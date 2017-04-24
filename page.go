// https://www.chazzuka.com/2015/03/load-parse-json-file-golang/

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "reflect"
)

type Page struct {
    Name   string `json:"name"`
    Filename string `json:"filename"`
    Url   string `json:"url"`
}

func (p Page) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func getPages() []Page {
    raw, err := ioutil.ReadFile("./chhs.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []Page
    json.Unmarshal(raw, &c)
    return c
}

func main() {

    pages := getPages()

    fmt.Println(len(pages))
    fmt.Println(reflect.TypeOf(pages))

    for _, p := range pages {
        fmt.Println(p.toString())
    }

    fmt.Println(toJson(pages))
}
