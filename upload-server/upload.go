package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"flag"
	"path/filepath"
	"io/ioutil"

	"github.com/golang/glog"
)

var (
	flagPort = flag.Int("p", 80, "port to bind")
	flagDir  = flag.String("d", "/", "directory to store uploaded files")
)

func main() {
	flag.Parse()
	defer glog.Flush()
	if _, err := ioutil.ReadDir(*flagDir); err != nil {
		glog.Fatal(err)
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	glog.Infof("listening %d", *flagPort)
	glog.Info(http.ListenAndServe(fmt.Sprintf(":%d", *flagPort), nil))
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		glog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	defer file.Close()
	f, err := os.OpenFile(filepath.Join(*flagDir, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		glog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}
	defer f.Close()
	io.Copy(f, file)
	glog.Infof("received file %s", filepath.Join(*flagDir, handler.Filename))
	fmt.Fprintln(w, fmt.Sprintf("receive file %s !", handler.Filename))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tpl))
}

const tpl = `<html>
<head>
<title>upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
 <input type="file" name="uploadfile" />
 <input type="submit" value="upload" />
</form>
</body>
</html>`
