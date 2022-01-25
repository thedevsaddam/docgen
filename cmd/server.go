package cmd

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	file      string
	port      int
	serveLive = &cobra.Command{
		Use:   "server",
		Short: "Serve live html from postman collection",
		Long:  `Serve live html from postman collection`,
		Run:   server,
	}
)

func init() {
	serveLive.PersistentFlags().IntVarP(&port, "port", "p", 9000, "port number to listen")
	serveLive.PersistentFlags().StringVarP(&file, "file", "f", "", "postman collection path")
	serveLive.PersistentFlags().BoolVarP(&isMarkdown, "md", "m", false, "display markdown format in preview")
	serveLive.PersistentFlags().BoolVarP(&includeVariable, "var", "v", false, "this flag will include variables in template")
	serveLive.PersistentFlags().BoolVarP(&sorted, "sort", "s", false, "this flag will sort the collection in ascending order")
	serveLive.PersistentFlags().StringVarP(&extraCSS, "css", "c", "", "inject a css file")
	serveLive.PersistentFlags().StringVarP(&env, "env", "e", "", "postman environment variable file path")
}

func server(cmd *cobra.Command, args []string) {
	if file == "" {
		log.Println("You must provide a file name!")
		return
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Println("Invalid file path!")
		return
	}
	http.HandleFunc("/", templateFunc)
	log.Println("Listening on port: ", port)
	log.Printf("Web Server is available at http://localhost:%s/\n", strconv.Itoa(port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		handleErr("Failed to open server", err)
	}
}

func templateFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	var buf *bytes.Buffer
	if !isMarkdown {
		buf = readJSONtoHTML(file)
		_, err := w.Write(buf.Bytes())
		handleErr("Failed to write html in writer", err)
		return
	}
	buf = readJSONtoMarkdownHTML(file)
	_, err := w.Write(buf.Bytes())
	handleErr("Failed to write html in writer", err)
}
