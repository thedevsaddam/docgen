package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	serveLive.PersistentFlags().StringVarP(&file, "file", "f", "", "postman collection file relative path")
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
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func templateFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	buf := readJSONtoHTML(file)
	w.Write(buf.Bytes())
}
