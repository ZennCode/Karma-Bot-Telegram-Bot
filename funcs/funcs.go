package funcs

import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"
	"log"
	errorino "zenncode/tgbot/error"

	_ "github.com/mattn/go-sqlite3"
)

func PlusEins(userID int64, username string, mgsID int64,serverID int64) {
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare(`update "`+strconv.Itoa(int(serverID))+`" set points=? where userID=?`)
	errorino.CheckErr(err)
	points_now := GetPoints(userID, username,mgsID ,serverID) + 1
	res, err := stmt.Exec(points_now, userID)
	errorino.CheckErr(err)
	_, err = res.RowsAffected()
	errorino.CheckErr(err)
	//log.Println("Der Datenbank Eintrag Nr. " + strconv.Itoa(int(affect)) + " wurde aktualisiert")
	log.Println("The user " + username + " has been updated. They now have " + strconv.Itoa(int(points_now)) + " Karma points")
}

func MinusEins(userID int64, username string,mgsID int64, serverID int64) {
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare(`update "`+strconv.Itoa(int(serverID))+`" set points=? where userID=?`)
	errorino.CheckErr(err)
	points_now := GetPoints(userID, username,mgsID,serverID ) - 1
	res, err := stmt.Exec(points_now, userID)
	errorino.CheckErr(err)
	_, err = res.RowsAffected()
	errorino.CheckErr(err)
	//log.Println("Der Datenbank Eintrag Nr. " + strconv.Itoa(int(affect)) + " wurde aktualisiert")
	log.Println("The user " + username + " has been updated. They now have " + strconv.Itoa(int(points_now)) + " Karma points")
}

func GetPoints(userID int64, username string, mgsID int64,serverID int64) (punkte int) {
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.CheckErr(err)
	defer db.Close()
	sqlStatement := `select * from "`+strconv.Itoa(int(serverID))+`" where userID = ?`
	var points int
	var createDate string
	var name string
	row := db.QueryRow(sqlStatement, userID)
	switch err := row.Scan(&userID, &username, &points, &createDate, &name); err {
		case sql.ErrNoRows:
			log.Println("There is no account with the name:", username, "and the ID:", userID)
			AddNewUser(userID, username, mgsID,name,serverID)
		case nil:
		default:
			panic(err)
	}	
	return points
}

func CheckIfUserExists(userID int64, username string, mgsID int64, name string,serverID int64) {
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.CheckErr(err)
	defer db.Close()
	sqlStatement := `select * from "`+strconv.Itoa(int(serverID))+`" where userID = ?`
	var points int64
	var createDate string

	row := db.QueryRow(sqlStatement, userID)
	switch err := row.Scan(&userID, &username, &points, &createDate, &name); err {
		case sql.ErrNoRows:
			log.Println("There is no account with the name:", username, "and the ID:", userID)
			AddNewUser(userID, username, mgsID,name, serverID)
		case nil:
 			layout := "2006-01-02T15:04:05.000Z"
			t, err := time.Parse(layout, createDate)
			if err != nil {
		    	log.Println(err)
			}
			NiceDate := t.Format("02.01.2006 15:04:05")
			log.Println("userID:", userID, "Username:", username, "Points:", points, "Registerd:", NiceDate)
		default:
			panic(err)
	}
}

func AddNewUser(userID int64,username string,mgsID int64, name string,serverID int64){
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.FatalError(err)
	defer db.Close()
	stmt, err := db.Prepare(`INSERT INTO "`+strconv.Itoa(int(serverID))+`"(userID, username, points, createDate, name) values(?,?,?,?,?)`)
	errorino.CheckErr(err)
	res, err := stmt.Exec(userID, username, 0, time.Now().Format("2006-01-02T15:04:05.000Z"),name)
	errorino.CheckErr(err)
	id, err := res.LastInsertId()
	errorino.CheckErr(err)
	log.Println("The database entry for",username,"has been regisered to the ID:",id)
	AddNewUser2(userID, username, mgsID)
}

func AddNewUser2(userID int64,username string, mgsID int64,){
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.FatalError(err)
	defer db.Close()
	stmt, err := db.Prepare("INSERT or REPLACE INTO Mention(UserID, lastMention) values(?,?)")
	errorino.CheckErr(err)
	res, err := stmt.Exec(userID, 0)
	errorino.CheckErr(err)
	id, err := res.LastInsertId()
	errorino.CheckErr(err)
	log.Println("The database entry for the last mention has been registered to the ID:",id)
}

func UpdateLastMention(userID int64,MsgID int64, username string){
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.FatalError(err)
	defer db.Close()
	stmt, err := db.Prepare("update Mention set lastMention=? where UserID=?")
	errorino.CheckErr(err)
	res, err := stmt.Exec(MsgID, userID)
	errorino.CheckErr(err)
	_, err = res.RowsAffected()
	errorino.CheckErr(err)
	log.Println("The last Mention Message ID for",username, "has been updated to", MsgID)
}

func UpdateUserName(userID int64,MsgID int64, username string,serverID int64){
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.FatalError(err)
	defer db.Close()
	stmt, err := db.Prepare(`update "`+strconv.Itoa(int(serverID))+`" set username=? where userID=?`)
	errorino.CheckErr(err)
	res, err := stmt.Exec(username, userID)
	errorino.CheckErr(err)
	_, err = res.RowsAffected()
	errorino.CheckErr(err)
	log.Println("The username for userID:",userID,"has been updated to", username)
}

func UpdateName(userID int64,MsgID int64, username string, name string,serverID int64){
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.FatalError(err)
	defer db.Close()
	stmt, err := db.Prepare(`update "`+strconv.Itoa(int(serverID))+`" set name=? where userID=?`)
	errorino.CheckErr(err)
	res, err := stmt.Exec(name, userID)
	errorino.CheckErr(err)
	_, err = res.RowsAffected()
	errorino.CheckErr(err)
	log.Println("The name for userID:",userID,"has been updated to", name)
}

func CheckMention(UserID int64, username string, mgsID int64) int64 {
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.CheckErr(err)
	defer db.Close()
	sqlStatement := "select * from Mention where UserID = ?"
	var lastMention int64
	row := db.QueryRow(sqlStatement, UserID)
	switch err := row.Scan(&UserID, &lastMention); err {
		case sql.ErrNoRows:
			log.Println("DB Mention Error", UserID)
			AddNewUser2(UserID, username, mgsID)
		case nil:
			log.Println("userID:", UserID, "UserName:", username, "MessageID:", mgsID)
		default:
			panic(err)
	}
	return lastMention
}

func LeaderBoard(userID int64, username string,serverID int64) string {
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.CheckErr(err)
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM "`+strconv.Itoa(int(serverID))+`" ORDER BY points DESC LIMIT 10`)
	errorino.CheckErr(err)
	var points int64
	var createDate string
	var name string
	i := 1
	var i2 string
	var i3 string
	for rows.Next()  {
		err = rows.Scan(&userID, &username, &points, &createDate,&name)
		errorino.CheckErr(err)
		// log.Println(i, userID, username, points, createDate)
		i2 += strconv.Itoa(int(userID))+ username+ strconv.Itoa(int(points))+ createDate+"\n"
		i3 += strconv.Itoa(i)+". place, with " + strconv.Itoa(int(points)) + " Karma "+ name +"\n"
		i += 1
	}
	return i3
}



func NewTable(serverID int64){
	//dbname:= strconv.Itoa(int(serverID))
	db, err := sql.Open("sqlite3", "./foo.db")
	errorino.CheckErr(err)
	defer db.Close()
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS "`+strconv.Itoa(int(serverID))+`" ("userID" NUMERIC, "username" TEXT, "points" INTEGER NOT NULL, "createDate" TEXT NOT NULL, "name" TEXT NOT NULL, UNIQUE("userID")
	)
	`
	_, err = db.Exec(sqlStmt)
	errorino.CheckErr(err)

}

func RandomHallo() string {
	var resu [6]string
	resu[0] = "Hey there"
	resu[1] = "Hi"
	resu[2] = "Howdy"
	resu[3] = "Greetings"
	resu[4] = "What's up"
	resu[5] = "Yo"
	return resu[rand.Intn(6)]
}

