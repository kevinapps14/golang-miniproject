package karyawancontroller

import (
	"fmt"
	"html/template"
	"net/http"

	"git.hub/kevinapps14/golang-miniproject/entities"
)

func Index(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/karyawan/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/karyawan/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()
		var karyawan entities.Karyawan
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.Divisi = request.Form.Get("divisi")
		pasien.Pekerjaan = request.Form.Get("pekerjaan")
		pasien.Deadline = request.Form.Get("deadline")

		fmt.Println(pasien)
	}

}
func Edit(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Its Work!! Edit")
}
func Delete(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Its Work!! Delete")
}
