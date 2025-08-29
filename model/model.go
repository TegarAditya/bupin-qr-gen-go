package model

type InfoUJN struct {
	KodeQR      string `db:"kodeQRUjian"`
	NamaJenjang string `db:"namaJenjang"`
	NamaKelas   string `db:"namaKelas"`
	NamaMapel   string `db:"namaMapel"`
	NamaBab     string `db:"namaBab"`
	IDUjian     int    `db:"idUjian"`
}

type InfoVID struct {
	KodeQR      string `db:"kodeQR"`
	NamaJenjang string `db:"namaJenjang"`
	NamaKelas   string `db:"namaKelas"`
	NamaMapel   string `db:"namaMapel"`
	NamaBab     string `db:"namaBab"`
	NamaSubBab  string `db:"namaSubBab"`
}
