package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func prepareBase(board string, thread string) map[string]io.Reader {
	var baseReader map[string]io.Reader
	var comment string
	var caption string
	var name string
	comment = ""
	caption = ""
	name = fmt.Sprintf("%v%v", jsonPayload.Person, tripcode)
	// If instagram story

	switch jsonPayload.From {
	case "vk":
		log.Println("Got post from VK")
		switch jsonPayload.Type {
		case "status":
			comment = fmt.Sprintf("✏️[b]VK статус изменился:[/b] %v\n\n[sup]%v ⛓[/sup]", jsonPayload.Caption, jsonPayload.Source)
			log.Printf("Caption: \"%v\"", comment)
		case "post":
			comment = fmt.Sprintf("[sup]%v ⛓[/sup]\n\n", jsonPayload.Source)
			if jsonPayload.Caption != "" {
				caption = strings.Replace(jsonPayload.Caption, "\n", "\n> ", -1)
				caption = strings.Replace(caption, "\n> ⠀\n> ", "\n\n> ", -1)
				comment = fmt.Sprintf("%v> %v", comment, caption)
			}
			log.Printf("Caption: \"%v\"", comment)
		case "public":
			comment = fmt.Sprintf("[sup]%v ⛓[/sup]\n\n", jsonPayload.Source)
			if jsonPayload.Caption != "" {
				caption = strings.Replace(jsonPayload.Caption, "\n", "\n> ", -1)
				caption = strings.Replace(caption, "\n> ⠀\n> ", "\n\n> ", -1)
				comment = fmt.Sprintf("%v> %v", comment, caption)
			}
			log.Printf("Caption: \"%v\"", comment)
		}
	case "ig":
		log.Println("Got post from IG")
		switch jsonPayload.Type {
		case "story":
			comment = fmt.Sprintf("[sup]%v ⛓[/sup]\n\n", jsonPayload.Source)
			if jsonPayload.Caption != "" {
				caption = strings.Replace(jsonPayload.Caption, "\n", "\n> ", -1)
				caption = strings.Replace(caption, "\n> ⠀\n> ", "\n\n> ", -1)
				comment = fmt.Sprintf("%v> %v", comment, caption)
			}
			log.Printf("Caption: \"%v\"", comment)
		case "post":
			comment = fmt.Sprintf("[sup]%v ⛓[/sup]\n\n", jsonPayload.Source)
			if jsonPayload.Caption != "" {
				caption = strings.Replace(jsonPayload.Caption, "\n", "\n> ", -1)
				caption = strings.Replace(caption, "\n> ⠀\n> ", "\n\n> ", -1)
				comment = fmt.Sprintf("%v> %v", comment, caption)
			}
			log.Printf("Caption: \"%v\"", comment)
		}
	case "twitch":
		log.Println("Got post from twitch")
		switch jsonPayload.Type {
		case "live":
			comment = fmt.Sprintf("[sup]Стрим запустился! %v ⛓[/sup]\n\n", jsonPayload.Source)
			if jsonPayload.Caption != "" {
				caption = strings.Replace(jsonPayload.Caption, "\n", "\n> ", -1)
				caption = strings.Replace(caption, "\n> ⠀\n> ", "\n\n> ", -1)
				comment = fmt.Sprintf("%vЗаголовок стрима: \n> \"%v\"", comment, caption)
			}
			log.Printf("Caption: \"%v\"", comment)
		}
	}

	baseReader = map[string]io.Reader{
		"task": strings.NewReader("post"),
		//"board":  strings.NewReader(json["2ch_board"].(string)),  // https://2ch.hk/test/
		"board":  strings.NewReader(board),  // https://2ch.hk/test/
		"thread": strings.NewReader(thread), // https://2ch.hk/test/res/28394.html
		"name":   strings.NewReader(name),   // Tripcode for attention whore
		//"email": strings.NewReader(""), // R u fucking kidding me?
		//"subject": strings.NewReader(jsonPayload.Person), // Oldfags never use it
		"comment": strings.NewReader(comment), // Post text

		//"comment": strings.NewReader(caption), // Post text
	}
	return baseReader
}

func prepareFiles() map[string]io.Reader {
	var filesReader map[string]io.Reader
	files := jsonPayload.Files
	url := []string{}
	for k, v := range files {
		fmt.Println(k, "is:", v)
		url = append(url, v)
	}
	if len(url) == 0 {
		return filesReader
	}
	count := len(url)
	// I know, I know. But it works...
	switch count {
	case 1:
		fmt.Println("One file")
		for k, v := range files {
			fmt.Println(k, "is:", v)
			resp1, e := http.Get(v)
			if e != nil {
				fmt.Println("http.Get error:", e)
				reportTg(e)
			}
			//defer resp.Body.Close()
			filesReader = map[string]io.Reader{
				`files1`: resp1.Body,
			}
		}
	case 2:
		fmt.Println("Two files")
		resp1, e := http.Get(url[0])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		resp2, e := http.Get(url[1])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		//defer resp.Body.Close()
		filesReader = map[string]io.Reader{
			`files1`: resp1.Body,
			`files2`: resp2.Body,
		}
	case 3:
		fmt.Println("Three files")
		resp1, e := http.Get(url[0])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		resp2, e := http.Get(url[1])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		resp3, e := http.Get(url[2])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		//defer resp.Body.Close()
		filesReader = map[string]io.Reader{
			`files1`: resp1.Body,
			`files2`: resp2.Body,
			`files3`: resp3.Body,
		}
	default:
		fmt.Println(len(url))
		resp1, e := http.Get(url[0])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		resp2, e := http.Get(url[1])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		resp3, e := http.Get(url[2])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		resp4, e := http.Get(url[3])
		if e != nil {
			fmt.Println("http.Get error:", e)
			reportTg(e)
		}
		//defer resp.Body.Close()
		filesReader = map[string]io.Reader{
			`files1`: resp1.Body,
			`files2`: resp2.Body,
			`files3`: resp3.Body,
			`files4`: resp4.Body,
		}
	}
	return filesReader
}
