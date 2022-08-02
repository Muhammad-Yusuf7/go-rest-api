package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  float64 `json:"price"`
}

var sourse = User + ":" + Password + "@/" + Db_name
var Db, err = sql.Open("mysql", sourse)

func Db_init() {
	if err != nil {
		panic(err)
	}

	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS albums (id INT NOT NULL AUTO_INCREMENT, title VARCHAR(250) NOT NULL, artist VARCHAR(250) NOT NULL, price FLOAT NOT NULL, PRIMARY KEY(id) )")
	if err != nil {
		panic(err)
	}
}

func Insert_to_db(db_args *Album) (r sql.Result, err error) {
	r,err = Db.Exec("INSERT INTO albums (title, artist, price) VALUES (?,?,?)",db_args.Title, db_args.Artist, db_args.Price)
	return
}

func Get_albums_from_db() []Album {
	var album []Album
	rows,err := Db.Query("SELECT * FROM albums")
	if err != nil {
		panic(err)
	}
	for rows.Next(){
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil{
			panic(err)
		}
		album = append(album, alb)
	}
	return album
}

func Get_albums_from_db_width_id(id int) []Album {
	var album []Album
	rows,err := Db.Query("SELECT * FROM albums WHERE id = ?",id)
	if err != nil {
		panic(err)
	}
	for rows.Next(){
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil{
			panic(err)
		}
		album = append(album, alb)
	}
	return album
}