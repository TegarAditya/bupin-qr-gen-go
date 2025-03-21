package model

import (
	"database/sql"
)

type InfoUJN struct {
	KodeQRUjian string `db:"kodeQR"`
	IDUjian     int    `db:"idUjian"`
	IDJenjang   int    `db:"idJenjang"`
	IDKelas     int    `db:"idKelas"`
	IDMapel     int    `db:"idMapel"`
	IDBab       int    `db:"idBab"`
	NamaJenjang string `db:"namaJenjang"`
	NamaKelas   string `db:"namaKelas"`
	NamaBab     string `db:"namaBab"`
	NamaMapel   string `db:"namaMapel"`
}

type InfoVID struct {
	KodeQR      string         `db:"kodeQR"`
	NamaJenjang string         `db:"namaJenjang"`
	NamaKelas   string         `db:"namaKelas"`
	NamaMapel   string         `db:"namaMapel"`
	NamaBab     string         `db:"namaBab"`
	NamaSubBab  string         `db:"namaSubBab"`
	IDKelas     int            `db:"idKelas"`
	IDMapel     int            `db:"idMapel"`
	IDBab       int            `db:"idBab"`
	IDSubBab    int            `db:"idSubBab"`
	LinkVideo   string         `db:"linkVideo"`
	Ytid        string         `db:"ytid"`
	LinkDMP     sql.NullString `db:"linkDmp"`
	YtidDMP     sql.NullString `db:"ytidDmp"`
	TP          string         `db:"tp"`
}
