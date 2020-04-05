package routes

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type File struct {
	Dir  bool
	Name string
	Path string
	Date time.Time
}

type FileInfos struct {
	Files    []File
	PathName string
}

// getFileLists 파일 리스트를 가져온다.
// FileInfos 는 특정 패스내의 파일 배열 구조체
// error 에러 메세지
func getFileLists(p string) (FileInfos, error) {
	fs := []File{}
	result := FileInfos{}
	fis, err := ioutil.ReadDir(p)
	if err != nil {
		log.Printf("filehandlers.getFileLists() error = %s %s\n", p, err)
		return result, err
	}
	fs = append(fs, File{Dir: true, Name: "..", Path: p + "/" + ".."})
	for _, f := range fis {
		if f.IsDir() {
			fs = append(fs, File{Dir: true, Name: f.Name() + "/", Path: p + "/" + f.Name(), Date: f.ModTime()})
		}
	}

	for _, f := range fis {
		if !f.IsDir() {
			fs = append(fs, File{Dir: false, Name: f.Name(), Path: p + "/" + f.Name(), Date: f.ModTime()})
		}
	}
	result.Files = fs
	result.PathName = p

	return result, err
}

//사용자 ID로 디렉토리 만들기
func createDir(userName string) error {
	dirname := rootDir + "/" + userName
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		err = os.MkdirAll(dirname, 0755)
		if err != nil {
			return err
		}
	}
	userDir = dirname
	return nil
}

//Delect file with file path
func deleteFile(filePath string) error {
	err := os.Remove(filePath)

	if err != nil {
		return err
	}
	return nil
}

//동일한 파일이 있을 경우 추가 되는 파일에 시간 데이터를 합친다.
func plusFileNumber(filename string) string {
	v := strings.LastIndex(filename, ".")
	t := time.Now()
	ts := t.Format("060102150405")
	temp := filename[:v] + "_" + ts
	return temp + filename[v:]
}
