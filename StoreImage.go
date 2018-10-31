package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var baseDirectory string
var ipaddress string

func main() {
	baseDirectory = "C:\\Files\\" // provide the base directory path where the files will be kept
	ipaddress = "localhost"       // provide the ip address of the webserver
	http.HandleFunc("/", homePage)
	http.HandleFunc("/uploadfile", uploadFile)
	http.HandleFunc("/listfiles", listFiles)
	http.HandleFunc("/deletefile", deleteFile)
	http.ListenAndServe(ipaddress+":8080", nil)
}

func homePage(w http.ResponseWriter, req *http.Request) {
	var options [3]string

	options[0] = "</br><a href = \"http://" + ipaddress + ":8080/uploadfile\">Click to upload file</a></br>"
	options[1] = "<a href = \"http://" + ipaddress + ":8080/listfiles\">Click to list files</a></br>"
	options[2] = "<a href = \"http://" + ipaddress + ":8080/deletefile\">Click to delete file</a></br>"

	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, "<h1>%s</h1>, <div>%s</div>", "Home Page\n", options)
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

func uploadFile(w http.ResponseWriter, req *http.Request) {
	//var s string
	if req.Method == http.MethodPost {
		f, handler, err := req.FormFile("usrfile")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error uploading file", http.StatusInternalServerError)
			return
		}
		defer f.Close()
		filename := handler.Filename
		fmt.Println(filename)
		bs, err := ioutil.ReadAll(f)
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

	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `<form action="/uploadfile" method="post" enctype="multipart/form-data">
        Upload a file<br>
        <input type="file" name="usrfile"><br>
        <input type="submit">
        </form>
        <br>
        <br>`)

}
