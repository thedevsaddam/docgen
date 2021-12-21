package cmd

import (
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/thedevsaddam/docgen/assets_bin"
	"gopkg.in/russross/blackfriday.v2"
)

func getData(a string) string {
	fp, err := assets_bin.AssetFS.Open(a)
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}

func eHTML(s string) string {
	return html.EscapeString(s)
}

func htmlTemplate(v string) template.HTML {
	return template.HTML(v)
}

func cssTemplate(v string) template.CSS {
	return template.CSS(v)
}

func jsTemplate(v string) template.JS {
	return template.JS(v)
}

func snake(v string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9%]+")
	resURI := url.QueryEscape(v)
	if err != nil {

		log.Fatal(err)
	}
	result := reg.ReplaceAllString(resURI, "")
	result = strings.Replace(result, "%", "_", -1)
	return result
}

func trimQueryParams(v string) string {
	if strings.Contains(v, "?") {
		return strings.Split(v, "?")[0]
	}
	return v
}

func addOne(v int) string {
	return strconv.Itoa(v + 1)
}

func trim(v string) string {
	return strings.TrimSpace(v)
}

func lower(v string) string {
	return strings.ToLower(v)
}

func upper(v string) string {
	return strings.ToUpper(v)
}

func githubLink(v string) string {
	v = strings.ToLower(v)

	v = strings.Replace(v, " ", "-", -1)
	nonWordChars := regexp.MustCompile(`[^-\w\d]`)
	v = nonWordChars.ReplaceAllString(v, "")
	return v
}

func githubLinkIncrementer(v string) string {
	k, ok := githubLinkInc[v]
	if ok {
		githubLinkInc[v]++
		return v + "-" + strconv.Itoa((k + 1))
	}
	githubLinkInc[v] = 0
	return v
}

func merge(v1 int, v2 string) string {
	return strconv.Itoa(v1+1) + ". " + v2
}

func markdown(v string) string {
	return string(blackfriday.Run([]byte(v)))
}

func dateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
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

// integer to roman number
func roman(num string) string {
	number, _ := strconv.Atoi(num)
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

// provide env values
func e(key string) string {
	for _, k := range envCollection.Values {
		key = strings.ReplaceAll(key, fmt.Sprintf("{{%s}}", k.Key), k.Value)
	}
	return key
}
