package main

import (
	"html/template"
	"log"
	"regexp"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func getData(a string) string {
	bs, err := Asset(a)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func html(v string) template.HTML {
	return template.HTML(v)
}

func css(v string) template.CSS {
	return template.CSS(v)
}

func js(v string) template.JS {
	return template.JS(v)
}

func snake(v string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(v, "")
}

func trimQueryParams(v string) string {
	if strings.Contains(v, "?") {
		return strings.Split(v, "?")[0]
	}
	return v
}

func markdown(v string) string {
	return string(blackfriday.Run([]byte(v)))
}

func color(v string) string {
	switch v {
	case "GET":
		return "info"
	case "POST":
		return "success"
	case "PATCH":
		return "warning"
	case "PUT":
		return "warning"
	case "DELETE":
		return "danger"
	default:
		return "info"

	}
}
