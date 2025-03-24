package model

type InfoUJN struct {
	KodeQR      string `db:"kodeQRUjian"`
	IDUjian     int    `db:"idUjian"`
	NamaJenjang string `db:"namaJenjang"`
	NamaKelas   string `db:"namaKelas"`
	NamaBab     string `db:"namaBab"`
	NamaMapel   string `db:"namaMapel"`
}

type InfoVID struct {
	KodeQR      string `db:"kodeQR"`
	NamaJenjang string `db:"namaJenjang"`
	NamaKelas   string `db:"namaKelas"`
	NamaMapel   string `db:"namaMapel"`
	NamaBab     string `db:"namaBab"`
	NamaSubBab  string `db:"namaSubBab"`
}
