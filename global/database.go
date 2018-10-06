package global

import (
	"bufio"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

const(
	Salt = "onlinecodes"

	QueryCodesvalue = "select codesvalue from share where hashkey = ?"
	InsetSnippet = "insert into share (hashkey, codesvalue) values(?,?) on duplicate key update codesvalue=?"
	ShareIsExists = "show tables like \"share\""
)

type store interface{
	GetSnippet(ctx context.Context, snippet *Snippet)(interface{},error)
	PutSnippet(ctx context.Context, snippet *Snippet) (string,error)
}

type Snippet struct{
	ID string	`json:"id"`
	Src []byte	`json:"src"`
}

type DatabaseController struct{
	db *sql.DB
}

func (c *DatabaseController) GetSnippet(ctx context.Context, snippet *Snippet)(interface{},error){
	query, err := c.db.Prepare(QueryCodesvalue)
	if err != nil{
		logrus.Error(err)
		return nil,nil
	}
	var res string
	if err := query.QueryRow(snippet.ID).Scan(&res); err != nil{
		logrus.Error(err)
		return nil, nil
	}
	return res,nil
}

func (c *DatabaseController)PutSnippet(ctx context.Context, snippet *Snippet)(string, error){
	id := snippet.getID()
	if _, err := c.db.ExecContext(ctx,InsetSnippet,id,snippet.Src,snippet.Src); err != nil{
		logrus.Debug(err)
		return "",err
	}
	return id , nil
}

func (s *Snippet)getID()string{
	h := sha256.New()
	io.WriteString(h,Salt)
	h.Write(s.Src)
	return hex.EncodeToString(h.Sum(nil))[:32]
}

func NewController() *DatabaseController{
	dir := "./"
	c := DecodeConfigJson(dir)

	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	authURL := c["database_username"] + ":" + c["database_password"] + "@tcp(" + c["database_ip"] + ")/"
	initDatabase(authURL,dir)
	db, err := sql.Open("mysql",authURL + "onlinecode")

	if err != nil{
		panic(err)
	}
	return &DatabaseController{
		db:db,
	}
}

func initDatabase(authURL, dir string){
	initOnlineCodeDatabaseSQL := filepath.Join(dir,"sql/init-onlinecode-database.sql")
	db, err := sql.Open("mysql",authURL)
	if err != nil{
		panic(err)
	}

	reader, err :=os.Open(initOnlineCodeDatabaseSQL)
	if err != nil{
		logrus.Errorf("Error when initDatabase, cannot optn initdatabase.sql file [%v]",err)
		panic(err)
	}
	bufReader := bufio.NewReader(reader)
	for{
		line, _, err := bufReader.ReadLine()
		if err != nil{
			break;
		}
		db.Exec(string(line))
	}
}





