package parser

import (
    "fmt"
    "net/http"
    "encoding/json"
    "model"
)

const (
    FACEBOOK_API = "https://graph.facebook.com/v2.3/"

    EVENT_FIELDS = "id,name,description,start_time,end_time,is_date_only,place"

    ACCESS_TOKEN = "CAACEdEose0cBABAa6qHGKKA2qPA8i9vHFm9GOgNxHjJq0A9QD4VneL4nY1JNs8P481YrNym2G2QYU1TK7BZC08rC5CVOCWXF0EIoVl1O2aP0GOSCg620CBZAGIkc3tqfMvPFpea0EfhZCaKaivwBYqKd5w0gPZC4bh9JuLv9Gt7ZCpCNDjBUw3toymvZAthuUJAiNUIZAQJxAu19Yc6Qvc5WhBmEhfG32cZD"
)


type fbError struct {
    Message string
    Type string
    Code int
}

type fbPlace struct {
    Id string
    Name string
    Location struct {
        Latitude float64
        Longitude float64
        Street string
        City string
        Zip string
        Country string
    }
}

type fbEvent struct {
    Id string
    Name string
    Description string
    Start_time string
    End_time string
    Is_date_only bool
    Place fbPlace
}

type fbEventList struct {
    Error *fbError
    Data []fbEvent
}


func ParseEvents(pageId string, events []model.Event) (error) {
    url := FACEBOOK_API + pageId + "/events" + "?access_token=" + ACCESS_TOKEN + "&fields=" + EVENT_FIELDS
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    var msg fbEventList
    if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
        return err
    }
    
    fmt.Println(msg.Data)

    return nil
}

