package karyawancontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/kevinapps14/golang-miniproject/entities"
	"github.com/kevinapps14/golang-miniproject/libraries"
	"github.com/kevinapps14/golang-miniproject/models"
)

var validation = libraries.NewValidation()
var KaryawanModel = models.NewKaryawanModel()

func Index(response http.ResponseWriter, request *http.Request) {

	karyawan, _ := KaryawanModel.FindAll()

	data := map[string]interface{}{
		"karyawan": karyawan,
	}

	temp, err := template.ParseFiles("views/karyawan/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

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
		karyawan.NamaLengkap = request.Form.Get("nama_lengkap")
		karyawan.Divisi = request.Form.Get("divisi")
		karyawan.Pekerjaan = request.Form.Get("pekerjaan")
		karyawan.Deadline = request.Form.Get("deadline")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(karyawan)

		if vErrors != nil {
			data["karyawan"] = karyawan
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Berhasil Disimpan"
			KaryawanModel.Create(karyawan)
		}

		temp, _ := template.ParseFiles("views/karyawan/add.html")
		temp.Execute(response, data)
	}

}
func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var karyawan entities.Karyawan
		KaryawanModel.Find(id, &karyawan)

		data := map[string]interface{}{
			"karyawan": karyawan,
		}

		temp, err := template.ParseFiles("views/karyawan/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {

		request.ParseForm()
		var karyawan entities.Karyawan
		karyawan.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		karyawan.NamaLengkap = request.Form.Get("nama_lengkap")
		karyawan.Divisi = request.Form.Get("divisi")
		karyawan.Pekerjaan = request.Form.Get("pekerjaan")
		karyawan.Deadline = request.Form.Get("deadline")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(karyawan)

		if vErrors != nil {
			data["karyawan"] = karyawan
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Berhasil Diperbaharui"
			KaryawanModel.Update(karyawan)
		}

		temp, _ := template.ParseFiles("views/karyawan/edit.html")
		temp.Execute(response, data)
	}

}
func Delete(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	KaryawanModel.Delete(id)

	http.Redirect(response, request, "/pasien", http.StatusSeeOther)
}
