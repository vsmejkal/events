package parser

import (
    "fmt"
    "io/ioutil"
    "net/http"
    // "encoding/json"
    // "model"
)

const (
    FACEBOOK_API = "https://graph.facebook.com/v2.3/"

    ACCESS_TOKEN = "CAACEdEose0cBAIxcA7xKECXlINSAWZCYWu2WfbQwQgFDLtvVRnfCtn5DJQx5D26AJ7Tr9ovSVgArUgBRx0sunuhl7GKJ46rR8hVpUq530bdaDuXhYJh1ZAbcYPeRgoVrOVkqbWmvUApDEF0ZB3F6cJN9vcbkRU8y13T4dwQImWBYnCjVjs4AL1O6FCZBWe6sR8R3lu7yBStDYFZBPk5lBHagBFkHJkgIZD"
)

func Parse(node string) (error) {
    url := FACEBOOK_API + node + "/events" + "?access_token=" + ACCESS_TOKEN

    res, _ := http.Get(url)
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(body))

    return nil;
}

