package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/k0kubun/pp/v3"
)


type TranslatedText struct{
	OriginalText [] string 
	IPAText [] string
	AudioFile [][] string
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

func GetAudioFiles(ipaText [] string) [][] string{
	var audioFiles [][] string

	for _, ipaWord := range ipaText{
		var wordAudioFiles [] string
		ìpaRunes := []rune(ipaWord)
		for _, ipaRune := range ìpaRunes{
			if string(ipaRune) == "/"{
				continue
			}
			fmt.Println("Rune : ", string(ipaRune))

			filePattern := fmt.Sprintf("./Audio/%v*_*.wav", string(ipaRune))

			log.Println(filePattern)

			fileResult, _ := filepath.Glob(filePattern)
			// pp.Println(fileResult)
			if(len(fileResult) == 0){
				fmt.Println("======")
				fmt.Println(string(ipaRune))
				fmt.Println("======")
				// Couldnt find a char
				// Using the char window should hopefully help here
				// As the Glob pattern only allows first letters
				continue
			}
			pp.Println(fileResult[0])
			wordAudioFiles = append(wordAudioFiles, "./" + fileResult[0])
		}
		audioFiles = append(audioFiles, wordAudioFiles)
	}
	return(audioFiles)
}

func GetIPAText(splitText []string) [] string {
	ipaTranslatedText := []string{}
	t := time.Now()
	for _, word := range splitText{
		fmt.Println(word)
		for i := 0; i < len(engToIPA); i++{
			if word == engToIPA[i][0]{
				ipaTranslatedText = append(ipaTranslatedText, engToIPA[i][1])
				break
			}
		}
	}
	log.Printf("Translation of %v word/s took %v\n", len(splitText), time.Now().Sub(t))
	log.Printf("Translated %v to %v\n", splitText, ipaTranslatedText)
	return ipaTranslatedText
}