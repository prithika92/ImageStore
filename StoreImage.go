package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var baseDirectory string

func main() {
	os.Mkdir("." + string(filepath.Separator) + "images",0777)
	baseDirectory = "/go/src/go_docker/images/" // provide the base directory path where the files will be kept
	http.HandleFunc("/", homePage)
	http.HandleFunc("/uploadfile", uploadFile)
	http.HandleFunc("/listfiles", listFiles)
	http.HandleFunc("/deletefile", deleteFile)
	fs := http.FileServer(http.Dir("static/")) //Serving static assets
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, req *http.Request) {
	var options [3]string
	options[0] = "</br><a href = \"/uploadfile\">Click to upload file</a></br>"
	options[1] = "<a href = \"/listfiles\">Click to list files</a></br>"
	options[2] = "<a href = \"/deletefile\">Click to delete file</a></br>"
	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "<h1>%s</h1>, <div>%s</div>", "Home Page\n", options)
}

func uploadFile(w http.ResponseWriter, req *http.Request) {
	//var s string // file coming in
	if req.Method == http.MethodPost {
		f, handler, err := req.FormFile("usrfile") // form file which takes key
		if err != nil {  
			log.Println(err)  // check error
			http.Error(w, "Error uploading file", http.StatusInternalServerError) // http error response writer
			return
		}
		defer f.Close()
		filename := handler.Filename  
		fmt.Println(filename)
		bs, err := ioutil.ReadAll(f)  // read file
		if err != nil {
			log.Println(err)
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		fmt.Println(bs)
		err1 := ioutil.WriteFile(baseDirectory+filename, bs, 0644)
		if err != nil {
			log.Fatal(err1)
		}
		fmt.Println("Success!")
	}

	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8") // write the response
	fmt.Fprintf(w, `<form action="/uploadfile" method="post" enctype="multipart/form-data">
        Upload a file<br>
        <input type="file" name="usrfile"><br>
        <input type="submit">
        </form>
        <br>
        <br>`)

}

func listFiles(w http.ResponseWriter, req *http.Request) {
	files, err := ioutil.ReadDir(baseDirectory)
	if err != nil {
		log.Fatal(err)
	}
	var listOfFiles [20]string
	i := 0
	for _, file := range files {
		listOfFiles[i] = "</br>" + baseDirectory + file.Name() + "</br>"
		fmt.Println(file.Name())
		i = i + 1
	}
	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "<h1>%s</h1>, <div>%s</div>", "List of Files :\n", listOfFiles)
}

func deleteFile(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		fileName := req.FormValue("fileName")
		fmt.Println(fileName)
		fileName = baseDirectory + fileName
		err := os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File deleted successfully !")
	}
	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `<form action="/deletefile" method="post" enctype="multipart/form-data">
        Please provide the name of file to be deleted <br>
        <input type="textbox" name="fileName"><br>
        <input type="submit">
        </form>
        <br>
        <br>`)

}