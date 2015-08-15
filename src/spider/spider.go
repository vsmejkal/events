package main

import (
    "fmt"
    "log"
    "time"
    "model"
    "parser"
)

func main() {
    sources := []string {
        // Kluby, bary, kavarny
        "https://www.facebook.com/fledaclub",
        "https://www.facebook.com/kabinetmuz",
        "https://www.facebook.com/Melodka.cz",
        "https://www.facebook.com/favalcz",
        "https://www.facebook.com/pages/Mersey-klub-Brno/149025198548565",
        "https://www.facebook.com/Klub.Alterna",
        "https://www.facebook.com/starapekarna",
        "https://www.facebook.com/desertbrno",
        "https://www.facebook.com/pages/Veselá-vačice/161305307286151",
        "https://www.facebook.com/pages/Klub-ŠELEPKA/207146942634999",
        "https://www.facebook.com/studentpartySemilasso",
        "https://www.facebook.com/ReisyClub",
        "https://www.facebook.com/SonoCentrum",
        "https://www.facebook.com/karaburgut",
        "https://www.facebook.com/Sklenick",
        "https://www.facebook.com/leitnerka",
        "https://www.facebook.com/perpetuumklub",
        "https://www.facebook.com/Industrabrno",
        "https://www.facebook.com/mala.amerika.brno",
        "https://www.facebook.com/Livingstone.club",
        "https://www.facebook.com/cafeprah",
        "https://www.facebook.com/vegalitebrno",
        "https://www.facebook.com/musilkaksomega",
        "https://www.facebook.com/elevenclub",
        "https://www.facebook.com/Bastilamusic",
        "https://www.facebook.com/pages/Jazz-bar-U-kouřícího-králíka-official-site/130274893694606",
        "https://www.facebook.com/vyhlidkacafe",
        "https://www.facebook.com/TwoFacesBar",
        "https://www.facebook.com/TwoFacesClub",
        "https://www.facebook.com/CLUBWASH",
        "https://www.facebook.com/YachtMusicPub",
        "https://www.facebook.com/klub.m13",
        "https://www.facebook.com/7nebe",
        "https://www.facebook.com/staratkalcovna.cz",
        "https://www.facebook.com/TabarinClub",
        "https://www.facebook.com/pages/Andel-Cafe-Brno/161767873865667",
        "https://www.facebook.com/musiclabbrno",
        "https://www.facebook.com/pages/RUSTY-NAIL-Music-Club/551585758206006",
        "https://www.facebook.com/abajoclub",
        "https://www.facebook.com/OldiesMusicClubFicak",
        "https://www.facebook.com/mandarinclubbrno",
        "https://www.facebook.com/pages/RATEJNA-music-restaurant/278880934245",
        "https://www.facebook.com/schrottbrno",
        "https://www.facebook.com/nadraze",
        "https://www.facebook.com/pages/VIBE-club/1595145777437797",
        "https://www.facebook.com/donlagarto",
        "https://www.facebook.com/ahoybrno",
        "https://www.facebook.com/caribicbrno",
        "https://www.facebook.com/BrnoTheEight",
        "https://www.facebook.com/DejaVuBrno",
        "https://www.facebook.com/upalecka",
        "https://www.facebook.com/kabaretspacek",
        "https://www.facebook.com/netopyrbrno",
        "https://www.facebook.com/charlieshat",
        "https://www.facebook.com/KavarnaTrojka",
        "https://www.facebook.com/sklepniscena",
        "https://www.facebook.com/tresgallos2",

        // Kina
        "https://www.facebook.com/kinoscalabrno",
        "https://www.facebook.com/artbrno",
        "https://www.facebook.com/LetniKinoSpilberk",
        "https://www.facebook.com/pages/Letní-kino-Brno-na-Dvoře-MdB/464222550388191",
        "https://www.facebook.com/Kino.Lucerna.Brno",
        "https://www.facebook.com/pages/Kinokavárna/355040474608176",

        // Divadla
        "https://www.facebook.com/husanaprovazku",
        "https://www.facebook.com/operadiversa",
        "https://www.facebook.com/janacekopera",
        "https://www.facebook.com/cinohrandb",
        "https://www.facebook.com/DivadloNaOrli",
        "https://www.facebook.com/bezbarierovedivadlobarka",
        "https://www.facebook.com/divadlobolkapolivky",
        "https://www.facebook.com/rubinbrno",
        "https://www.facebook.com/divadlokorab",
        "https://www.facebook.com/divadlopolarka",
        "https://www.facebook.com/hadivadlo",
        "https://www.facebook.com/dsparadox",
        "https://www.facebook.com/mestske.divadlo.brno",
        "https://www.facebook.com/studioMarta",

        // Galerie
        "https://www.facebook.com/pages/FAIT-Gallery/198731196900551",
        "https://www.facebook.com/ragallery",

        // Ostatní
        "https://www.facebook.com/hvezdarna.brno",
        "https://www.facebook.com/tmbrno",
        "https://www.facebook.com/pages/Koupaliště-Kraví-hora/44095453644",
        "https://www.facebook.com/mzk.cz",
        "https://www.facebook.com/klubcestovatelubrno",
        "https://www.facebook.com/masarykova.univerzita",
        "https://www.facebook.com/phil.muni.cz",
        "https://www.facebook.com/filharmoniebrno",
        "https://www.facebook.com/pdfmu",
        "https://www.facebook.com/PrfMUni",
        "https://www.facebook.com/MUNIEBrno",
        "https://www.facebook.com/galerievankovka",
        "https://www.facebook.com/vidabrno",
        "https://www.facebook.com/ticbrno",
        "https://www.facebook.com/FactoryFashionMarket",
        "https://www.facebook.com/pages/Spirit-Bar/442458112553968",
        "https://www.facebook.com/expeditionclubbrno",
        "https://www.facebook.com/pages/Kulturní-Centrum-Líšeň/472564170044",
        "https://www.facebook.com/slevarna",
        "https://www.facebook.com/lasershowhallbrno",
        "https://www.facebook.com/bbakaly",
        "https://www.facebook.com/tanecnak",

    }

    for _, url := range sources {
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
                    event.Store()
                }
                fmt.Println("NEW:", event.Name)

            case err := <-errChan:
                log.Println(err)
            }
        }

        time.Sleep(1000 * time.Millisecond)
    }
}