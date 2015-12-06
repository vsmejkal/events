package main

import (
	"fmt"
	"log"
	"time"
	"github.com/vsmejkal/events/model"
	"github.com/vsmejkal/events/parser"
	"os"
	"github.com/vsmejkal/events/config"
	"path"
)

func getSources() (sources []string) {
    sources = make([]string, 0)

    db := model.GetConnection()
    rows, err := db.Query("SELECT url FROM source;")
    if err != nil {
        log.Println(err)
        return
    }

    var url string
    for rows.Next() {
        if err = rows.Scan(&url); err != nil {
            log.Println(err)
        } else {
            sources = append(sources, url)
        }
    }
    if err = rows.Err(); err != nil {
        log.Println(err)
    }

    return
}

func printUsage() {
	fmt.Printf("Usage: %s config.json\n", path.Base(os.Args[0]))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		printUsage()
		return
	}

	if err := config.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	for _, url := range getSources() {
        fmt.Println("\nParsing", url, "...")

		eventChan := make(chan model.Event, 100)
		errChan := make(chan error, 100)

		go func() {
			parser.ParseEvents(url, eventChan, errChan)
			close(eventChan)
		}()

		loop: for {
			select {
			case event, ok := <-eventChan:
				if !ok {
					break loop
				}

				if event.IsValid() {
					fmt.Println("  ", event.Name, event.Start)
					
					if err := event.Store(); err != nil {
						log.Println(err)
					}
				}

			case err := <-errChan:
				log.Println(err)
			}
		}

        // Primitive load balancing
		time.Sleep(5 * time.Second)
	}
}