package database_actions

/*
	Худайбирдин Ильнур Минисламович
	Реализация мини-библиотеки для работы с базой данных SQLite.
	Позволяет сократить количество кода при работе с БД.
*/

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	Db   *sqlx.DB
	Err  error
	Path string
}

func (db *Db) connect() {
	db.Db, db.Err = sqlx.Open("sqlite3", db.Path)
	if db.Err != nil {
		panic(db.Err)
	}
}

func (db *Db) disconnect() {
	db.Db.Close()
}

func (db *Db) GetAllRows(table string, data interface{}) {
	/*
		GetAllRows возвращает все записи из таблицы.
		Parameters:
			table - название таблицы
			data - данные для вставки
	*/
	db.connect()
	defer db.disconnect()
	sqlstr := fmt.Sprintf("select * from %s", table)
	db.Err = db.Db.Select(data, sqlstr)
	if db.Err != nil {
		panic(db.Err)
	}
}

func (db *Db) DeleteAllRows(table string) {
	/*
		DeleteAllRows удаляет все записи из таблицы.
		Parameters:
			table - название таблицы
	*/
	db.connect()
	defer db.disconnect()
	sqlstr := fmt.Sprintf("delete from %s", table)
	_, db.Err = db.Db.Exec(sqlstr)
	if db.Err != nil {
		fmt.Println(db.Err)
	}
}

func (db *Db) InsertRow(table string, ColsNamesSlice []string, data interface{}) {
	/*
		InsertRow добавляет новую строку в таблицу.
		Parameters:
			table - название таблицы
			ColsNamesSlice - названия столбцов
			data - данные для вставки
	*/
	db.connect()
	defer db.disconnect()
	ColsNames := ""
	ValuesNames := ""
	for i := 0; i < len(ColsNamesSlice); i++ {
		ColsNames += ColsNamesSlice[i]
		if i != len(ColsNamesSlice)-1 {
			ColsNames += ", "
		}
		ValuesNames += ":" + ColsNamesSlice[i]
		if i != len(ColsNamesSlice)-1 {
			ValuesNames += ", "
		}
	}
	sqlstr := fmt.Sprintf("INSERT INTO %s (%s) values (%s)", table, ColsNames, ValuesNames)
	_, err := db.Db.NamedExec(sqlstr, data)
	if err != nil {
		fmt.Println(err)
	}
}

func (db *Db) GetRowByCondition(table, cols, condition string, data interface{}) {
	/*
		GetRowByCondition возвращает первую строку по условию.
		Parameters:
			table - название таблицы
			cols - названия столбцов
			condition - условие
			data - данные для вставки
	*/
	db.connect()
	defer db.disconnect()
	db.Err = db.Db.Get(data, "select (?) from ? where ?", cols, table, condition)
	if db.Err != nil {
		fmt.Println(db.Err)
	}
}
