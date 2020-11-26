package routes

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net/http"
	"net/mail"
	"strconv"
	"strings"

	"github.com/taknb2nch/go-pop3"
)

type popMail struct {
	Title   string
	From    string
	To      string
	Uid     string
	Date    string
	Content string
}

func pop(w http.ResponseWriter, r *http.Request) {
	addr := "pop.naver.com:995"

	client, err := pop3.Dial(addr) //pop3.go의 net.Dial -> tls.Dial (,nil)

	if err != nil {
		log.Printf("Error1: %v\n", err)
	}

	defer func() {
		client.Quit()
	}()

	if err = client.User("in3166@naver.com"); err != nil {
		log.Printf("Error2: %v\n", err)
		return
	}
	if err = client.Pass("dlsgh12"); err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	var count int
	var size uint64

	if count, size, err = client.Stat(); err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	log.Printf("Count: %d, Size: %d\n", count, size)

	// if count, size, err = client.List(2); err != nil {
	// 	log.Printf("Error: %v\n", err)
	// 	return
	// }
	// log.Printf("Number: %d, Size: %d\n", count, size)

	//var mis []pop3.MessageInfo

	// if mis, err = client.ListAll(); err != nil {
	// 	log.Printf("Error: %v\n", err)
	// 	return
	// }

	// for _, mi := range mis {
	// 	log.Printf("Number: %d, Size: %d, Uid: %q\n", mi.Number, mi.Size, mi.Uid)
	// }

	// var number int
	// var uid string

	// if number, uid, err = client.Uidl(2); err != nil {
	// 	log.Printf("Error: %v\n", err)
	// 	return
	// }

	// var bufs1 bytes.Buffer
	// wr := transform.NewWriter(&bufs1, korean.EUCKR.NewEncoder())
	// wr.Write([]byte(uid))
	// wr.Close()

	// log.Printf("Number: %d, Uid: %s\n \n", number, bufs1.String())

	// if mis, err = client.UidlAll(); err != nil {
	// 	log.Printf("Error: %v\n", err)
	// 	return
	// }

	// i := 0

	// for _, mi := range mis {
	// 	i++
	// 	log.Printf("Number: %d, Uid: %s\n", mi.Number, mi.Uid)
	// }

	var content string
	var data1 []popMail
	for i := count; i > 180; i-- {

		if number, uid, err := client.Uidl(i); err != nil {
			log.Printf("Uidl Error: %v\n", err)
			return
		} else {
			log.Printf("i: %d, Number: %d, Uid: %s\n \n", i, number, uid)

			if content, err = client.Retr(i); err != nil {
				log.Printf("Error: %v\n", err)
				return
			} else {
				msg, err := mail.ReadMessage(bytes.NewBufferString(content))
				if err != nil {
					log.Fatal("Cannot parse myMessage.")
				}

				fmt.Println("\n--------2body--------------")
				date1, _ := msg.Header.Date()
				fmt.Println("1DATE: ", date1.Format("2006-01-02 15:04:05"))
				fmt.Println("1To: ", msg.Header.Get("To"))
				fmt.Println("1From: ", decodeRFC2047(msg.Header.Get("From")))
				fmt.Println("1Subject: ", decodeRFC2047(msg.Header.Get("Subject")))
				fmt.Println("msg.Header.Get(Content-Type): ", msg.Header.Get("Content-Type"))
				fmt.Println("msg.Header.Get(Content-Transfer-encoding): ", msg.Header.Get("Content-Transfer-encoding"))

				//multipart
				ss := multipartFunc(w, content)
				if err != nil {
					log.Fatal(err)
				}
				//fmt.Println(ss)
				fmt.Println("\n--------2end--------------")

				data := popMail{
					Title:   decodeRFC2047(msg.Header.Get("Subject")),
					From:    decodeRFC2047(msg.Header.Get("From")),
					To:      msg.Header.Get("To"),
					Date:    date1.Format("2006-01-02 15:04:05"),
					Uid:     uid,
					Content: ss,
				}

				data1 = append(data1, popMail{
					Title:   decodeRFC2047(msg.Header.Get("Subject")),
					From:    decodeRFC2047(msg.Header.Get("From")),
					To:      msg.Header.Get("To"),
					Date:    date1.Format("2006-01-02 15:04:05"),
					Uid:     uid,
					Content: ss,
				})

				fmt.Println(i, ": --Last-- data")

				//fmt.Println(data)
				file, _ := json.MarshalIndent(data, "", " ")
				t := strconv.Itoa(i)
				var url string = "public/mailboxes/in3166@naver.com/inbox/" + string(t) + ".json"
				err1 := ioutil.WriteFile(url, file, 0777)
				if err1 != nil {
					log.Println("저장 오류", err1)
				}
			}
		}
	}

	//log.Println("Content:\n%s\n", content)

	// r1 := strings.NewReader(content)
	// m, err := mail.ReadMessage(r1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// header := m.Header

	// date, _ := header.Date()
	// fmt.Println("DATE1: ", date)
	// fmt.Println("DATE2: ", header.Get("Date"))
	// fmt.Println("From: ", header.Get("From"))
	// fmt.Println("To: ", header.Get("To"))
	// //fmt.Println("Subject: ", header.Get("subject"))
	// fmt.Println("Content-Type: ", header.Get("Content-Type"))

	// subject1 := header.Get("subject")
	// dec := new(mime.WordDecoder)
	// subject, _ := dec.DecodeHeader(subject1)
	// fmt.Println("Subject: ", subject)
	// from2, _ := dec.DecodeHeader(header.Get("From"))
	// fmt.Println("From2: ", from2)

	// fmt.Println("\n--------body--------------")
	// //body, err := ioutil.ReadAll(m.Body)
	// //fmt.Println(string(body))
	// //fmt.Println("\n--------2body--------------")

	// msg, err := mail.ReadMessage(bytes.NewBufferString(content))
	// if err != nil {
	// 	log.Fatal("Cannot parse myMessage.")
	// }

	// fmt.Println("\n--------2body--------------")
	// date1, _ := msg.Header.Date()
	// fmt.Println("1DATE: ", date1.Format("2006-01-02 15:04:05"))
	// fmt.Println("1To: ", msg.Header.Get("To"))
	// fmt.Println("1From: ", decodeRFC2047(msg.Header.Get("From")))
	// fmt.Println("1Subject: ", decodeRFC2047(msg.Header.Get("Subject")))
	// fmt.Println("msg.Header.Get(Content-Type): ", msg.Header.Get("Content-Type"))
	// fmt.Println("msg.Header.Get(Content-Transfer-encoding): ", msg.Header.Get("Content-Transfer-encoding"))

	// //multipart
	// ss := multipartFunc(w, content)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(ss)
	// fmt.Println("\n--------2end--------------")

	// if err = client.Dele(6); err != nil {
	// 	log.Printf("Error: %v\n", err)
	// 	return
	// }

	if err = client.Noop(); err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	if err = client.Rset(); err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	// data := popMail{
	// 	Title:   decodeRFC2047(msg.Header.Get("Subject")),
	// 	From:    decodeRFC2047(msg.Header.Get("From")),
	// 	To:      msg.Header.Get("To"),
	// 	Date:    date1.Format("2006-01-02 15:04:05"),
	// 	Content: ss,
	// }

	log.Printf("완료")
	//log.Println("합쳤냐?", data1)
	//	file, _ := json.MarshalIndent(data, "", " ")

	//	err1 := ioutil.WriteFile("public/mailboxes/in3166@naver.com/inbox/o.json", file, 0777)

	//if err1 != nil {
	//	log.Println("저장 오류", err1)
	//}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	b, err := json.MarshalIndent(data1, "", "  ")
	if err != nil {
		panic(err)
	}

	w.Write(b)
}

func multipartFunc(w http.ResponseWriter, content string) string {

	fmt.Println("@@multipart In!")
	//fmt.Println("content in Func: \n", content)
	log.Println("1")
	strings.ReplaceAll(content, "\n", "\r\n")
	msg, err := mail.ReadMessage(bytes.NewBufferString(content))
	if err != nil {
		log.Println("1 error: ", err)
		return ""
	}

	fmt.Println("$$$ msg.Header.Get(Content-Type): ", msg.Header.Get("Content-Type"))

	var uDec []byte // 본문

	mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	fmt.Println("$$$$$ params boundary: ", params["boundary"])
	fmt.Println("$$$$$ mediaType: ", mediaType)

	if err != nil {
		log.Fatal("2 error: ", err)
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(msg.Body, params["boundary"])
		i := 1
		for {

			fmt.Println(i, ": !!!!!!!!!!!!!!!!!!")
			i++
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("$$$$$ p.Header.Get(Content-Type): ", p.Header.Get("Content-Type"))
			slurp, err := ioutil.ReadAll(p) //본문
			if err != nil {
				log.Fatal(err)
			}
			encoding := p.Header.Get("Content-Transfer-Encoding\n")
			if encoding == "" {
				encoding = "7bit"
			}
			if strings.Contains(p.Header.Get("Content-Type"), "multipart") {
				newContent := "Mime-Version: 1.0\n" + "Message-ID: " + randomString(12) + "\nContent-Type: " + p.Header.Get("Content-Type") + "\nContent-Transfer-Encoding: " + encoding + ";\n\n" + string(slurp) // 헤더와 본문 사이의 띄움

				ss := multipartFunc(w, newContent)
				log.Println("현재 컨텐츠: ", ss)
				return ss
			}

			if !(strings.Contains(strings.ToUpper(p.Header.Get("Content-Type")), "UTF-8") || strings.Contains(strings.ToUpper(p.Header.Get("Content-Type")), "EUC-KR") || strings.Contains(strings.ToUpper(p.Header.Get("Content-Type")), "US")) {
				fmt.Println("break: ", p.Header.Get("Content-Type"))
				break
			}

			//fmt.Println("사전 검사 본문: ", string(slurp))
			//fmt.Println("사전 검사 Encoding: ", p.Header.Get("Content-Transfer-Encoding"))

			if strings.Contains(strings.ToUpper(p.Header.Get("Content-Transfer-Encoding")), "BASE64") {
				uDec, _ = base64.StdEncoding.DecodeString(string(slurp))
			} else if strings.Contains(strings.ToUpper(p.Header.Get("Content-Transfer-Encoding")), "QUOTED-PRINTABLE") {
				uDec, _ = ioutil.ReadAll(quotedprintable.NewReader(strings.NewReader(string(slurp))))
			} else {
				uDec = []byte(string(slurp))
			}

			//fmt.Printf("%s %v\n", b, err)
			//fmt.Printf("^^결과: Content-type: %s / Content-Transfer-encoding:  \n --------\n %s\n", p.Header.Get("Content-Type"), uDec)
		}
	} else {
		body, _ := ioutil.ReadAll(msg.Body)

		fmt.Println("********************")
		//fmt.Println("사전 검사 본문: ", string(body))
		fmt.Println("사전 검사: ", msg.Header.Get("Content-Transfer-Encoding"))
		//var uDec []byte

		if strings.Contains(strings.ToUpper(msg.Header.Get("Content-Transfer-encoding")), "BASE64") {
			fmt.Println("1")
			uDec, _ = base64.StdEncoding.DecodeString(string(body))
			//fmt.Println("1: ", string(uDec))
		} else if strings.Contains(strings.ToUpper(msg.Header.Get("Content-Transfer-encoding")), "QUOTED-PRINTABLE") {
			fmt.Println("2")
			uDec, _ = ioutil.ReadAll(quotedprintable.NewReader(strings.NewReader(string(body))))
		} else {
			uDec = body
		}
		//fmt.Printf("%s %v\n", b, err)
		//fmt.Printf("  <본문 내용> \n %s\n", uDec)
		fmt.Println("********************")
	}

	//fmt.Printf("  <본문 내용> \n %s\n", uDec)
	//fmt.Fprintf(w, string(uDec))
	return string(uDec)
}

// decodeRFC2047 ... UTF8
func decodeRFC2047(s string) string {
	// GO 1.5 does not decode headers, but this may change in future releases...
	decoded, err := (&mime.WordDecoder{}).DecodeHeader(s)
	if err != nil || len(decoded) == 0 {
		return s
	}
	return decoded
}

// adapted from mime.randomBoundary
func randomString(length int) string {
	buf := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf[:])
}
