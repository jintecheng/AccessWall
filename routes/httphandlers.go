package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	sessions "github.com/goincremental/negroni-sessions"
)

// HandlerFunc for "index" is main
func index(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query()["dir"]
	if len(value) > 0 {
		gpath = strings.Trim(value[0], "/")
		re := strings.LastIndex(gpath, "..")
		if re != -1 {
			gpath = gpath[:re-3]
			gpath = gpath[:strings.LastIndex(gpath, "/")]
		}
	} else {
		gpath = userDir
	}
	//사용자가 자기 디렉토리 이외에 접근 하려 할때 제어
	if strings.Index(gpath, userDir) < 0 {
		gpath = userDir
	}
	if gpath == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	data, err := getFileLists(gpath)
	if err != nil {
		http.Redirect(w, r, "/?dir="+userDir, http.StatusFound)
		return
	}
	t, _ := template.ParseFiles("index.html", "./templates/filemanager.html", "./templates/upload.html")
	t.ExecuteTemplate(w, "index", data)
}

func mailPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("send_mail.html", "./templates/_module.html")
	t.ExecuteTemplate(w, "mail", nil)
}

// HandlerFunc for login
func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	if err != nil || email == "" && password == "" {
		t, _ := template.ParseFiles("./templates/login.html")
		t.Execute(w, nil)
	}
	if r.Method != "POST" {
		return
	}

	//http.PostForm 은 form 데이터를 만들어 보낼 때 사용 한다
	resp, err := http.PostForm("http://localhost:8001/api/user/login", url.Values{"name": {""}, "email": {email}, "password": {password}})

	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		//str := string(respBody)
	}

	loginResult := &LoginResult{}
	err = json.Unmarshal(respBody, loginResult)
	if err != nil {
		fmt.Println(err)
	}

	//sotauth로 로그인 성공 했을때 움직이는 것
	if loginResult.Status == true {
		lo := loginResult.Account
		sotAuthUser := &User{Uid: lo.Token, Name: lo.Name, Email: lo.Email, Expired: lo.Expired, Token: lo.Token}
		gUser = sotAuthUser
		SetCurrentUser(r, sotAuthUser)
		s := sessions.GetSession(r)
		s.Get("sotauth")
		userDir = rootDir + "/" + strings.Split(sotAuthUser.Email, "@")[0]
		gpath = userDir
		http.Redirect(w, r, "/?dir="+userDir, http.StatusFound)
		//여기서 부터는 좀 쉬고 해야할 일은 이제 권한을 획득 하는 함수를 만드는 것
	}
	defer resp.Body.Close()
	t, _ := template.ParseFiles("./templates/login.html")
	t.Execute(w, nil)
}

// Handler for logout
func logout(w http.ResponseWriter, r *http.Request) {
	sessions.GetSession(r).Delete(currentUserKey)
	userDir = ""
	gUser = nil
	linkQuery = ""
	http.Redirect(w, r, "/login", http.StatusFound)
}

// HandlerFunc for post request 'delete'
func delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println(w, "Method error")
		return
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(w, err)
		return
	}
	client, _ := connDB()
	defer disConnDB(client)
	for _, filePath := range r.Form["dirs[]"] {
		err = deleteFile(filePath)
		//시스템에서 파일을 삭제하는데 에러가 없고 파일의 가장 뒤가 "/" 가 아니면
		//데이터베이스에서 해당 파일 명을 삭제 한다.
		tmpFile := strings.Replace(filePath, userDir, "", -1)
		tmpFile = strings.Replace(tmpFile, "/", "", -1)
		deleteFileDB(client, &FileStruct{Name: tmpFile, User: gUser.Email, Guest: nil})
	}

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, "Success")
}

func searchFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, "Method error")
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	filename := r.PostFormValue("name")
	client, _ := connDB()
	defer disConnDB(client)
	results, _ := findFileDB(client, &FileStruct{Name: filename, User: gUser.Email, Guest: nil})
	var result string
	if len(results) > 0 {
		result = strings.Join(results[0].Guest, ";")
	} else {
		result = ""
	}

	fmt.Fprint(w, result)
}

func filesPerm(w http.ResponseWriter, r *http.Request) {
	filenames := []string{}
	emails := []string{}
	files := []FileStruct{}
	if r.Method != "POST" {
		fmt.Fprint(w, "Method error")
		return
	}

	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fnStr := r.PostFormValue("edit_filenames")
	emStr := r.PostFormValue("edit_emails")

	fnStr = strings.Replace(fnStr, userDir, "", -1)
	fnStr = strings.Replace(fnStr, "/", "", -1)
	filenames = strings.Split(fnStr, ";")

	emails = strings.Split(emStr, ";")

	fmt.Println(filenames, gUser.Email, emails)

	for _, file := range filenames {
		files = append(files, FileStruct{file, gUser.Email, emails})
	}

	client, _ := connDB()
	defer disConnDB(client)
	//insert 중에 중복 파일의 경우 업데이트를 한다.
	for _, fs := range files {
		findresult, _ := findFileDB(client, &fs)
		fmt.Println(len(findresult))
		if len(findresult) > 0 {
			err = updateFileDB(client, &fs)
		} else {
			err = insertFileDBOne(client, &fs)
		}
	}
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
}

func makeDir(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println(w, "Method error")
		return
	}

	dirname := r.PostFormValue("dirname")

	if dirname == "" {
		fmt.Println(w, "Failed make new folder!")
		return
	}

	dirname = gpath + "/" + dirname

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		err = os.MkdirAll(dirname, 0755)
		if err != nil {
			fmt.Println(w, err)
			return
		}
		fmt.Fprint(w, "Success make new folder!!")
	}
}

// HandlerFunc for post request 'upload'
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println(w, "Method error")
		return
	}

	var file multipart.File
	var fileHeader *multipart.FileHeader
	var err error
	var errMsg string
	var uploadFileName string

	file, fileHeader, err = r.FormFile("upfile")
	if err != nil {
		if fileHeader == nil {
			errMsg = "None upload file-header"
		}
		fmt.Fprintln(w, "%s : %s", err, errMsg)
		return
	}

	//uploadFileName = fileHeader.Filename
	var saveImage *os.File
	if _, err := os.Stat(gpath + "/" + fileHeader.Filename); err == nil {
		uploadFileName = plusFileNumber(fileHeader.Filename)
	} else if os.IsNotExist(err) {
		uploadFileName = fileHeader.Filename
	}

	saveImage, err = os.Create(gpath + "/" + uploadFileName)
	if err != nil {
		fmt.Fprintln(w, "Failed save file : %s", err)
		return
	}

	defer saveImage.Close()
	defer file.Close()
	_, err = io.Copy(saveImage, file)
	if err != nil {
		fmt.Fprintln(w, "Failed upload : %s ", err)
		return
	}

	http.Redirect(w, r, "/?dir="+gpath, http.StatusFound)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "relative/path/to/favicon.ico")
	http.Redirect(w, r, "/?dir="+gpath, http.StatusFound)
}

func linkLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/login.html")
	t.Execute(w, nil)
}

func linkFile(w http.ResponseWriter, r *http.Request) {
	if gUser == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	value := r.URL.Query()["file"]
	tmps := strings.Split(value[0], "/")
	filename := tmps[len(tmps)-1]
	client, _ := connDB()
	guest := gUser.Email
	guests := []string{guest}
	defer disConnDB(client)
	result, _ := findLinkFileDB(client, &FileStruct{Name: filename, User: "", Guest: guests})
	if len(result) == 0 {
		fmt.Fprint(w, "Permission error : ")
	} else {
		sendFile(w, r, filename)
	}
}

func sendFile(w http.ResponseWriter, r *http.Request, filename string) {
	filepath := r.URL.Query()["file"][0]
	filepath = "./Users" + filepath

	streamBytes, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Fprint(w, err)
	}

	b := bytes.NewBuffer(streamBytes)
	w.Header().Set("Content-Dispostion", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/pdf")
	if _, err := b.WriteTo(w); err != nil {
		fmt.Fprint(w, err)
	}

}

/*
func sendFile(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Query()["file"][0]
	filepath = "./Users" + filepath

	openFile, err := os.Open(filepath)
	if err != nil {
		fmt.Fprint(w, "File download error : "+err.Error())
	}
	defer openFile.Close()

	fileHeader := make([]byte, 512)
	openFile.Read(fileHeader)
	fileContentType := http.DetectContentType(fileHeader)
	fileStat, _ := openFile.Stat()
	fileSize := strconv.FormatInt(fileStat.Size(), 10)
	w.Header().Set("Content-Dispostion", "attachment; filename="+filepath)
	w.Header().Set("Content-Type", fileContentType)
	w.Header().Set("Content-Length", fileSize)

	openFile.Seek(0, 0)
	io.Copy(w, openFile)
	return
}
*/
