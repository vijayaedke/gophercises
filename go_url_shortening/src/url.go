package main

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

//structure to read YAML file
type artist struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

//default or fallback for other invalid path
func defaultPage() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, reader *http.Request) {
		writer.Write([]byte("Hello Pauper !"))
	})
	return mux
} //defaultPage

func mapHandler(path_to_redirect map[string]string, redirectDefaultPage http.Handler) http.HandlerFunc {
	//return httphandler
	return func(writer http.ResponseWriter, reader *http.Request) {
		//get current URL
		path := reader.URL.Path
		//if match found redirect to corresponding URL
		if redirect_path, path_value := path_to_redirect[path]; path_value {
			http.Redirect(writer, reader, redirect_path, http.StatusFound)
			return
		} //if

		//if not the specified URL then redirect to default page
		redirectDefaultPage.ServeHTTP(writer, reader)
	}
}

func yamlHandler(yamlcontent []byte, redirectDefaultPage http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	//read yaml string
	var artistlist []artist
	err := yaml.Unmarshal(yamlcontent, &artistlist)
	if err != nil {
		return nil, err
	}

	//convert yaml to map string
	path_to_redirect := make(map[string]string)
	for _, paths := range artistlist {
		path_to_redirect[paths.Path] = paths.URL
	} //for
	//invoke the map handler mapping the path to corresponding URL
	return mapHandler(path_to_redirect, redirectDefaultPage), nil
}

func main() {

	//redirect from MapHandler
	// path_to_redirect := map[string]string{
	// 	"/avicii": "https://www.youtube.com/watch?v=96vzFK2wKvg",
	// 	"/miley":  "https://www.youtube.com/watch?v=M11SvDtPBhA",
	// }
	redirectDefaultPage := defaultPage()
	// mapHandlerRedirect := mapHandler(path_to_redirect, redirectDefaultPage)

	//read from yaml file
	yamlPath := `
	- path: /avicii
	  url: https://www.youtube.com/watch?v=96vzFK2wKvg
	- path: /tujamo
	  url: https://www.youtube.com/watch?v=ysNbg3FfsMk
	`

	//redirect from yaml handler
	yamlHandlerRedirect, err := yamlHandler([]byte(yamlPath), redirectDefaultPage)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", yamlHandlerRedirect); err != nil {
		fmt.Println("Failed to connect to server")
		return
	}
}
