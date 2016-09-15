package main

import (
	"fmt"
    "io"
    "log"
    "net/http"
    "os"
    "encoding/xml"
)

type Query struct {
    //Series Show
    // Have to specify where to find episodes since this
    // doesn't match the xml tags of the data that needs to go into it
    AddonSiteList []AddonSite `xml:"sdk:addon-site>"`
}

type AddonSite struct {
    sdk:name string
    sdk:url string
}

//        <sdk:name>Google Inc.</sdk:name>
//        <sdk:url>addon.xml</sdk:url>

func main() {
	addonListUrl := "https://dl.google.com/android/repository/addons_list-2.xml"

	if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
        os.Exit(1)
    }

    response, err := http.Get(addonListUrl)
        if err != nil {
                log.Fatal(err)
        } else {
                defer response.Body.Close()
                
                var q Query
                xmlFile := response.Body
                xml.Unmarshal(xmlFile, &q)
                
                _, err := io.Copy(os.Stdout, response.Body)
                if err != nil {
                        log.Fatal(err)
                }
        }
}

func Hello() string {
	return string("hello, world\n")
}