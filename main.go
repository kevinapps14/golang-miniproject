package main

import (
	"net/http"

	"github.com/kevinapps14/golang-miniproject/controllers/karyawancontroller"
)

func main() {

	http.HandleFunc("/", karyawancontroller.Index)
	http.HandleFunc("/karyawan", karyawancontroller.Index)
	http.HandleFunc("/karyawan/index", karyawancontroller.Index)
	http.HandleFunc("/karyawan/add", karyawancontroller.Add)
	http.HandleFunc("/karyawan/edit", karyawancontroller.Edit)
	http.HandleFunc("/karyawan/delete", karyawancontroller.Delete)
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("assets/img/"))))

	http.ListenAndServe(":9000", nil)

}
