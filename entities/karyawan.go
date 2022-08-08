package entities

type Karyawan struct {
	Id          int64
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	Divisi      string `validate:"required"`
	Pekerjaan   string `validate:"required"`
	Deadline    string `validate:"required"`
}
