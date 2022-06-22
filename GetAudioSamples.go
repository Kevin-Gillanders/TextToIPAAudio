package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)


func GetAudioSamples(){
	endPoint := "https://jbdowse.com/ipa/"
	response, err := http.Get(endPoint)

	if err != nil{
		log.Fatal("Error http get against endPoint : ", err)
	}


	// content, err := ioutil.ReadAll(response.Body)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println(string(content))

    doc, err := goquery.NewDocumentFromReader(response.Body)
    defer response.Body.Close()


    if err != nil{
    	log.Fatal("Error passing data to goquery : ", err)
    }

    fmt.Println(doc)

    r, _ := regexp.Compile(".*\\/(.*)\\.wav")

    doc.Find("table").Each(func(i int, table *goquery.Selection) {
		// For each item found, get the title
		table.Find("td").Each(func(i int, td *goquery.Selection){
			

			represenation := td.Find("p").Text()
			if strings.TrimSpace(represenation) == ""{
				h, err := td.Html()
				if err != nil{
					log.Fatal("Error getting represenation in fail safe : ", err)
				}
				represenation = strings.TrimSpace(strings.Split(h, "<")[0])
			}
			// runeToCheck := "ɠ"


            // <div></div>
            // <div></div>
            // <p>ɠ̥</p>
            // <a href="s/e90.wav" target="b"></a>
            // <a href="s/e91.wav" target="b"></a>
            // <a href="s/e92.wav" target="b"></a>
            // <a href="s/e93.wav" target="b"></a>
			// if strings.Contains(represenation, runeToCheck) {
				// fmt.Printf("%v\n", represenation)
				// con, _ := td.Html()
				// fmt.Printf("%v\n", con)

				td.Find("a").EachWithBreak(func(i int, a *goquery.Selection) bool {
					file, ok := a.Attr("href")
					if ok{
						// if strings.Contains(file, "0a0"){
							fmt.Println("Saving : " + file)
							fileNameOnSite := r.FindStringSubmatch(file)[1]
							fmt.Println("Saving : " + file)
							fmt.Println("fileNameOnSite : " + fileNameOnSite)
							fmt.Println("represenation : " + represenation)
							h, _ := td.Html()
							fmt.Println("HMTL : ", h)
							out, err := os.Create("./Audio/" + represenation + "_" + fileNameOnSite + ".wav")
							defer out.Close()
	
							if err != nil{
								log.Fatal("Error creating file : ", err)
							}
	
							resp, err := http.Get(endPoint + file)
							if err != nil{
								log.Fatal("Error downloading file : ", err)
							}
							defer resp.Body.Close()
	
	
							_, err = io.Copy(out, resp.Body)
	
							if err != nil{
								log.Fatal("Error saving file : ", represenation, file, err)
							}
	
							time.Sleep(time.Second * 2)
							fmt.Println("Saved : ", file)
							fmt.Println("=======")
							return false
						}
					// }
					return true
				})

			// }
		})
	})
}