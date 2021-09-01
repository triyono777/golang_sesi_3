package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	namaProgram := os.Args
	idSiswa, _ := strconv.Atoi(namaProgram[1])
	getDataKelas(idSiswa)
	//fmt.Println(namaProgram)

}

type namaSiswa struct {
	absensi   int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var dataKelas []namaSiswa

func getDataKelas(noabsen int) {
	dataKelas = []namaSiswa{
		{absensi: 1, nama: "MUHAMMAD ZUNAN ALFIKRI", alamat: "ALAMAT MUHAMMAD", pekerjaan: "MAHASISWA", alasan: "ALASAN MUHAMMAD"},
		{absensi: 2, nama: "ARFAN JADULHAQ", alamat: "ALAMAT ARFAN", pekerjaan: "MAHASISWA", alasan: "ALASAN ARFAN"},
		{absensi: 3, nama: "TRIYONO", alamat: "ALAMAT TRIYONO", pekerjaan: "MAHASISWA", alasan: "ALASAN TRIYONO"},
	}
	for _, siswa := range dataKelas {

		if siswa.absensi == noabsen {
			fmt.Println(siswa.nama)
			fmt.Println(siswa.alamat)
			fmt.Println(siswa.pekerjaan)
			fmt.Println(siswa.alasan)

		}

	}

}
