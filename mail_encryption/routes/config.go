package routes

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type accountInfo struct {
	smtp_add    string
	smtp_port   string
	imap_add    string
	imap_port   string
	pop3_add    string
	pop3_port   string
	mail_passwd string
}

const (
	host     = "13.124.175.231"
	port     = 5432
	user     = "jhkim"
	password = "qwer1234"
	dbname   = "emailtest"
)

// DB에 연결
func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func disconnectDB(db *sql.DB) error {
	err := db.Close()

	if err != nil {
		log.Fatal()
	}
	return err
}

func getAccountInfo(db *sql.DB, mailadd string) accountInfo {

	var accountinfo accountInfo

	db.QueryRow("SELECT smtp_add FROM mail_info WHERE mail=$1", mailadd).Scan(&accountinfo.smtp_add)
	db.QueryRow("SELECT smtp_port FROM mail_info WHERE mail=$1", mailadd).Scan(&accountinfo.smtp_port)
	db.QueryRow("SELECT imap_add FROM mail_info WHERE mail=$1", mailadd).Scan(&accountinfo.imap_add)
	db.QueryRow("SELECT imap_port FROM mail_info WHERE mail=$1", mailadd).Scan(&accountinfo.imap_port)
	db.QueryRow("SELECT pop3_add FROM mail_info WHERE mail=$1", mailadd).Scan(&accountinfo.pop3_add)
	db.QueryRow("SELECT pop3_port FROM mail_info WHERE mail=$1", mailadd).Scan(&accountinfo.pop3_port)
	db.QueryRow("SELECT convert_from(decrypt(decode(mail_passwd, 'hex'), 'epjtwihnsasdc', 'aes'),'utf-8') from mail_info WHERE mail=$1", mailadd).Scan(&accountinfo.mail_passwd)

	return accountinfo
}

func getPresetInfo(db *sql.DB, mailadd string) accountInfo {

	var accountinfo accountInfo

	db.QueryRow("SELECT smtp_add FROM mail_preset WHERE mail_add=$1", mailadd).Scan(&accountinfo.smtp_add)
	db.QueryRow("SELECT smtp_port FROM mail_preset WHERE mail_add=$1", mailadd).Scan(&accountinfo.smtp_port)
	db.QueryRow("SELECT imap_add FROM mail_preset WHERE mail_add=$1", mailadd).Scan(&accountinfo.imap_add)
	db.QueryRow("SELECT imap_port FROM mail_preset WHERE mail_add=$1", mailadd).Scan(&accountinfo.imap_port)
	db.QueryRow("SELECT pop3_add FROM mail_preset WHERE mail_add=$1", mailadd).Scan(&accountinfo.pop3_add)
	db.QueryRow("SELECT pop3_port FROM mail_preset WHERE mail_add=$1", mailadd).Scan(&accountinfo.pop3_port)
	//db.QueryRow("SELECT convert_from(decrypt(decode(mail_passwd, 'hex'), 'epjtwihnsasdc', 'aes'),'utf-8') from mail_info WHERE mail=$1", mailadd).Scan(&accountinfo.mail_passwd)

	return accountinfo
}
