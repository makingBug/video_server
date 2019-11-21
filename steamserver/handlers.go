package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"html/template"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")

	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video,err := os.Open(vl)
	if err != nil{
		log.Printf("Error when try to open file %v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}

	//强制提醒,浏览器会自动将他作为mp4来解析
	w.Header().Set("Content-Type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	//先校验文件大小是否超出限制
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	//拿到文件,将数据写到文件
	file, _, err := r.FormFile("file")//<form name="file">
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}


	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}

