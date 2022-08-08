package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kevinapps14/golang-miniproject/config"
	"github.com/kevinapps14/golang-miniproject/entities"
)

type KaryawanModel struct {
	conn *sql.DB
}

func NewKaryawanModel() *KaryawanModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &KaryawanModel{
		conn: conn,
	}
}

func (p *KaryawanModel) FindAll() ([]entities.Karyawan, error) {

	rows, err := p.conn.Query("select * from karyawan")
	if err != nil {
		return []entities.Karyawan{}, err
	}
	defer rows.Close()

	var dataKaryawan []entities.Karyawan
	for rows.Next() {
		var karyawan entities.Karyawan
		rows.Scan(&karyawan.Id,
			&karyawan.NamaLengkap,
			&karyawan.Divisi,
			&karyawan.Pekerjaan,
			&karyawan.Deadline)
		//mengubah format yyyy-mm-dd menjadi format indo
		deadline, _ := time.Parse("2006-01-02", karyawan.Deadline)
		//diubah menjadi dd-mm-yyyy
		karyawan.Deadline = deadline.Format("02-01-2006")

		dataKaryawan = append(dataKaryawan, karyawan)
	}

	return dataKaryawan, nil

}

func (p *KaryawanModel) Create(karyawan entities.Karyawan) bool {

	result, err := p.conn.Exec("insert into Karyawan (nama_lengkap, divisi, pekerjaan, deadline) values(?,?,?,?)",
		karyawan.NamaLengkap, karyawan.Divisi, karyawan.Pekerjaan, karyawan.Deadline)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *KaryawanModel) Find(id int64, karyawan *entities.Karyawan) error {

	return p.conn.QueryRow("select * from karyawan where id = ?", id).Scan(&karyawan.Id,
		&karyawan.NamaLengkap,
		&karyawan.Divisi,
		&karyawan.Pekerjaan,
		&karyawan.Deadline)
}

func (p *KaryawanModel) Update(karyawan entities.Karyawan) error {

	_, err := p.conn.Exec(
		"udpate karyawan set nama_lengkap = ?, divisi = ?, pekerjaan = ?, deadline = ? where id = ?",
		karyawan.NamaLengkap, karyawan.Divisi, karyawan.Pekerjaan, karyawan.Deadline, karyawan.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *KaryawanModel) Delete(id int64) {
	p.conn.Exec("delete from karyawan where id = ?", id)
}
