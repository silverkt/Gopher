package main
import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
)


/**
create table student(
	id int primary key auto_increment,
	name varchar(20),
	age int,
	created varchar(10)
	) DEFAULT CHARSET=utf8;
*/
func main() {
    db := opendb("root:12345678@tcp(localhost:3306)/gosqlx_db?charset=utf8mb4")
    // id:=insert(db)
	query(db)
	//fmt.Println(id);
    // update(db,id)

}

//打开数据库连接
func opendb(dbstr string) ( * sql.DB) {
//dsn: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]
    db, err := sql.Open("mysql", dbstr)
    prerr(err)
    return db
}

//插入数据
func insert(db  * sql.DB) int64 {

    stmt, err := db.Prepare("INSERT INTO student SET name=?,age=?,created=?")
    prerr(err)

    res, err := stmt.Exec("abloz4", 32, "2013-8-21")
    prerr(err)

    id, err := res.LastInsertId()
    prerr(err)

    fmt.Println(id)
    return id

}
//更新数据
func update(db  *sql.DB,id int64) {
    stmt, err := db.Prepare("update student set name=? where id=?")
    prerr(err)

    res, err := stmt.Exec("silver", id)
    prerr(err)

    affect, err := res.RowsAffected()
    prerr(err)

    fmt.Println(affect)
}

//查询数据
func query(db  * sql.DB) {

    rows, err := db.Query("SELECT * FROM student")
    prerr(err)

    for rows.Next() {
        var id int
        var name sql.NullString
        var age sql.NullInt64
        var created sql.NullString
        err = rows.Scan(&id, &name, &age, &created)
        prerr(err)
        fmt.Println(id, name.String, age.Int64, created.String)
	 
    }
}

//删除数据
func del(db  * sql.DB, id int64) {
    stmt, err := db.Prepare("delete from student where id=?")
    prerr(err)

    res, err := stmt.Exec(id)
    prerr(err)

    affect, err := res.RowsAffected()
    prerr(err)

    fmt.Println(affect)
}
func prerr(err error) {
    if err != nil {
        panic(err)
    }
}