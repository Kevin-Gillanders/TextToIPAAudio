package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)


type TranslatedText struct{
	OriginalText [] string 
	IPAText [] string
	AudioFile [] string
}

func CreateTranslatedText(originalText string) TranslatedText{

	splitText := strings.Fields(originalText)
	ipaText := GetIPAText(splitText)
	audioFiles := GetAudioFiles(ipaText)
	return TranslatedText{
		OriginalText: splitText,
		IPAText: ipaText,
		AudioFile: audioFiles,
	}
}

func GetAudioFiles(ipaText [] string) [] string{
	panic("unimplemented")
}

func GetIPAText(splitText []string) [] string {
	ipaTranslatedText := []string{}
	t := time.Now()
	for _, word := range splitText{
		fmt.Println(word)
		for i := 0; i < len(engToIPA); i++{
			if word == engToIPA[i][0]{
				log.Println("under l : ", engToIPA[i])
				ipaTranslatedText = append(ipaTranslatedText, engToIPA[i][1])
				break
			}
		}
	}
	log.Printf("Translation of %v word/s took %v\n", len(splitText), time.Now().Sub(t))
	log.Printf("Translated %v to %v\n", splitText, ipaTranslatedText)
	return ipaTranslatedText
}