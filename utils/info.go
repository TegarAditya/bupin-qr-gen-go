package utils

import (
	"bupin-qr-gen-go/database"
	"bupin-qr-gen-go/model"
	"fmt"
	"regexp"
	"strings"
)

func GetInfoUJN(id string) (*model.InfoUJN, error) {
	var infoUJN model.InfoUJN

	queryString := `
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

	err := database.DB.Get(&infoUJN, queryString, id)
	if err != nil {
		return nil, err
	}

	return &infoUJN, nil
}

func GetInfoVID(id string) (*model.InfoVID, error) {
	var infoVID model.InfoVID

	queryString := `
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
		qrvap.kodeQR = ?`

	err := database.DB.Get(&infoVID, queryString, id)
	if err != nil {
		return nil, err
	}

	return &infoVID, nil
}

func GetJenjangKelas(inputString string) string {
	jenjangs := []string{"SD", "SMP", "SMA", "MI", "MTS", "SMK", "MA"}
	jenjangsPattern := strings.Join(jenjangs, "|")

	regexPattern := fmt.Sprintf(`\b(%s)(?:-)?\s*(\d+|VII|VIII|IX|X|XI|XII)\b`, jenjangsPattern)
	regex := regexp.MustCompile(`(?i)` + regexPattern)

	match := regex.FindStringSubmatch(inputString)

	if match == nil {
		return inputString
	}

	jenjang := match[1]
	kelas := match[2]

	return fmt.Sprintf("%s %s", jenjang, kelas)
}

func GetKurikulum(inputString string) string {
	stringLower := strings.ToLower(inputString)

	switch {
	case strings.Contains(stringLower, "merdeka"):
		return "KURMER"
	case strings.Contains(stringLower, "kma 143"), strings.Contains(stringLower, "kma 183"), strings.Contains(stringLower, "kma 347"):
		return "KMA 143"
	case strings.Contains(stringLower, "btq"):
		return "BTQ"
	case strings.Contains(stringLower, "2013"):
		return "K13"
	default:
		return "UNKNOWN"
	}
}

func GetBab(inputString string) string {
	regex := regexp.MustCompile(`\b(?:bab|chapter|subtema|wulangan|unit)\s+(\d+)`)
	match := regex.FindStringSubmatch(strings.ToLower(inputString))

	if match == nil {
		return inputString
	}

	var typeStr string
	if strings.Contains(strings.ToLower(inputString), "subtema") {
		typeStr = "SUBTEMA"
	} else {
		typeStr = "BAB"
	}

	return fmt.Sprintf("%s %s", typeStr, match[1])
}

func GetSubBab(inputString string) string {
	regex := regexp.MustCompile(`[A-Z]\.`)

	switch {
	case strings.Contains(inputString, "AKM"):
		return "AKM"
	case strings.Contains(inputString, "P3"):
		return "P3"
	}

	match := regex.FindStringSubmatch(strings.ToUpper(inputString))

	if match == nil {
		return inputString
	}

	return fmt.Sprintf("SUBBAB %s", strings.Replace(match[0], ".", "", -1))
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
