package main

import (
    "parser"
    "model"
)

func main() {
    events := make([]model.Event, 0, 20)
    parser.ParseEvents("TwoFacesClub", events)
}
