package utils

import (
	"bupin-qr-gen-go/database"
	"bupin-qr-gen-go/model"
	"fmt"
	"regexp"
	"strings"
)

func GetInfoUJN(id string) (*model.InfoUJN, error) {
	var info model.InfoUJN

	q := `
		SELECT
			qrujian.kodeQR,
			qrujian.id_ujian AS idUjian,
			qrujian.idJenjang,
			qrujian.idKelas,
			qrujian.idMapel,
			qrujian.idBab,
			jenjang.namaJenjang,
			kelas.namaKelas,
			bab.namaBab,
			mapel.namaMapel
		FROM
			qrujian
			INNER JOIN kelas ON qrujian.idKelas = kelas.idKelas
			INNER JOIN mapel ON qrujian.idMapel = mapel.idMapel
			INNER JOIN bab ON qrujian.idBab = bab.idBab
			INNER JOIN jenjang ON qrujian.idJenjang = jenjang.idJenjang 
		WHERE
			qrujian.kodeQRUjian = ?
	`

	err := database.DB.Get(&info, q, id)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func GetInfoVID(id string) (*model.InfoVID, error) {
	var info model.InfoVID

	q := `
		SELECT
			qrvap.kodeQR,
			jenjang.namaJenjang,
			kelas.namaKelas,
			mapel.namaMapel,
			bab.namaBab,
			subbab.namaSubBab,
			qrvap.idKelas,
			qrvap.idMapel,
			qrvap.idBab,
			qrvap.idSubBab,
			submateri.linkVideo,
			submateri.ytid,
			videodmp.linkDmp,
			videodmp.ytidDmp,
			qrvap.tp
		FROM
			qrvap
			LEFT JOIN jenjang ON jenjang.idJenjang = qrvap.idJenjang
			LEFT JOIN kelas ON kelas.idKelas = qrvap.idKelas
			LEFT JOIN mapel ON mapel.idMapel = qrvap.idMapel
			LEFT JOIN bab ON bab.idBab = qrvap.idBab
			LEFT JOIN subbab ON subbab.idSubBab = qrvap.idSubBab
			LEFT JOIN submateri ON submateri.idSubBab = subbab.idSubBab
			LEFT JOIN videodmp ON subbab.idSubBab = videodmp.idSubBab
		WHERE
		qrvap.kodeQR = ?
	`

	err := database.DB.Get(&info, q, id)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func GetJenjangKelas(input string) string {
	j := []string{"SD", "SMP", "SMA", "MI", "MTS", "SMK", "MA"}
	p := strings.Join(j, "|")

	rp := fmt.Sprintf(`\b(%s)(?:-)?\s*(\d+|VII|VIII|IX|X|XI|XII)\b`, p)
	r := regexp.MustCompile(`(?i)` + rp)

	m := r.FindStringSubmatch(input)

	if m == nil {
		return input
	}

	jenjang := m[1]
	kelas := m[2]

	return fmt.Sprintf("%s %s", jenjang, kelas)
}

func GetKurikulum(input string) string {
	s := strings.ToLower(input)

	switch {
	case strings.Contains(s, "merdeka"):
		return "KURMER"
	case strings.Contains(s, "kma 143"), strings.Contains(s, "kma 183"), strings.Contains(s, "kma 347"):
		return "KMA 143"
	case strings.Contains(s, "btq"):
		return "BTQ"
	case strings.Contains(s, "2013"):
		return "K13"
	default:
		return "UNKNOWN"
	}
}

func GetBab(input string) string {
	r := regexp.MustCompile(`\b(?:bab|chapter|subtema|wulangan|unit)\s+(\d+)`)
	m := r.FindStringSubmatch(strings.ToLower(input))

	if m == nil {
		return input
	}

	var t string
	if strings.Contains(strings.ToLower(input), "subtema") {
		t = "SUBTEMA"
	} else {
		t = "BAB"
	}

	return fmt.Sprintf("%s %s", t, m[1])
}

func GetSubBab(input string) string {
	re := regexp.MustCompile(`[A-Z]\.`)

	switch {
	case strings.Contains(input, "AKM"):
		return "AKM"
	case strings.Contains(input, "P3"):
		return "P3"
	}

	m := re.FindStringSubmatch(strings.ToUpper(input))

	if m == nil {
		return input
	}

	return fmt.Sprintf("SUBBAB %s", strings.Replace(m[0], ".", "", -1))
}

func GetFileName(v *model.InfoVID) string {
	var sb strings.Builder

	sb.WriteString(GetJenjangKelas(v.NamaKelas))
	sb.WriteString(" - ")
	sb.WriteString(GetKurikulum(v.NamaKelas))
	sb.WriteString(" - ")
	sb.WriteString(v.NamaMapel)
	sb.WriteString(" - ")
	sb.WriteString(GetBab(v.NamaBab))
	sb.WriteString(" - ")
	sb.WriteString(GetSubBab(v.NamaSubBab))
	sb.WriteString(" - ")
	sb.WriteString(v.KodeQR)

	return sb.String()
}
