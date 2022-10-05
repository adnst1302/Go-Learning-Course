package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := openKoneksi()
	defer db.Close()

	ctx := context.Background()
	//kdSupplier := "SP001"

	sqlQuery := "INSERT INTO tbl_supplier(kd_supplier, nama_supplier, no_hp, alamat, created_at, updated_at) VALUES ('SP002', 'CV Subur Karya', '087895336522','Medan Area','2022-10-05 11:07:33','2022-10-05 11:07:33')"

	_, err := db.ExecContext(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}
	fmt.Println("Success")

}

func openKoneksi() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/dbs_fresh_fruit?parseTime=true")
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi berhasil")
	return db
}
