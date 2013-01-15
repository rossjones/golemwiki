package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var config *JsonConfiguration
var config_file string

func init() {
	fmt.Println("Initialising config")
	const (
		default_config = "wiki.cfg"
		config_usage   = "The path to the config file"
	)

	flag.StringVar(&config_file, "config", "wiki.cfg", config_usage)
	flag.StringVar(&config_file, "c", "wiki.cfg", config_usage)
	flag.Parse()

	config = new(JsonConfiguration)
	err := config.LoadFrom(config_file)
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		os.Exit(1)
	}

	init_templates()
	init_database()
}

func main() {

	static_directory := config.Server["static"]
	fmt.Printf("Serving static content from %s\n", static_directory)
	fmt.Printf("Using port %s\n", config.Server["port"])
	http.Handle("/static/",
		http.StripPrefix("/static", http.FileServer(http.Dir(static_directory))))
	http.HandleFunc("/", homepage_view)
	http.HandleFunc("/about", about_view)
	http.HandleFunc("/admin", admin_view)
	http.HandleFunc("/history", history_view)
	http.HandleFunc("/add/group", add_group_view)
	http.HandleFunc("/add/page", add_page_view)

	http.ListenAndServe(":"+config.Server["port"], nil)
}
