package main

import "testing"

// di golang file untuk keperluan testing dipisahkan dengan file utama
// dan nama file untuk testing harus berakhiran _test.go
// unit test di golang dituliskan dalam bentuk function dengan nama function harus diawali dengan kata Test

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	// method t.Logf digunakan untuk menampilkan log,
	// method ini equivalen dengan fmt.Printf
	t.Logf("Volume : %.2f", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		// method t.Errorf sama dengan t.Logf namun diikuti dengan fail pada saat testing
		t.Errorf("SALAH harusnya %.2f", volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("SALAH harusnya %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH harusnya %.2f", kelilingSeharusnya)
	}
}

// cara eksekusi testing dengan cara menggunakan command `go test` diikuti file testing,
// karena file testing dan file yang akan di test berada di file yang berbeda, maka kedua file harus diikutsertakan sebagai argumen pada saat menjalankan command go test
// argument -v atau verbose digunakan untuk menampilkan semua output log pada saat pengujian

// jika tidak tau file apa yang akan di test, bisa juga dengan menggunakan command
// go test nama_module
// example : go test github.com/fanialfi/go-unit-test -v itu akan menjalankan semua unit test yang ada di module

// jika hanya ingin menjalankan 1 function unit test bisa denga menggunakan command
// go test -run ^TestFunction$ namaModule -v
// example : go test -run ^TestHitungLuas$ github.com/fanialfi/go-unit-test -v

// benchmarking
// package testing selain digunakan untuk melakukan testing, juga bisa digunakan untuk melakukan benchmarking, cara pembuatannya sama seperti cara pembuatan function untuk testing
// bedanya nama awal function diawali dengan kata Benchmark dan memiliki parameter pada function bertipe *testing.B

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}

// untuk menjalankan function benchmark diikuti dengan argument -bench=.
