# Getting Started

Selamat datang di panduan instalasi Golang menggunakan pendekatan Go Modules. Go Modules telah menjadi standar di Golang sejak versi 1.13, yang memungkinkan struktur proyek lebih fleksibel.

## Instalasi Go di Ubuntu (WSL atau Native)

### 1. Update Paket Ubuntu
Pertama-tama, pastikan sistem paket Anda diperbarui:

```bash
sudo apt update
sudo apt upgrade
```

### 2. Unduh dan Instal Go
Kunjungi [halaman unduhan resmi Go](https://go.dev/dl/) untuk mendapatkan versi terbaru. Gunakan perintah di bawah ini sebagai contoh untuk mengunduh Go versi 1.19:  

```
wget https://go.dev/dl/go1.19.6.linux-amd64.tar.gz
```

Ekstrak Go ke direktori `/usr/local`:
```
sudo tar -C /usr/local -xzf go1.19.6.linux-amd64.tar.gz
```

### 3. Tambahkan Go ke `PATH`
Edit file .bashrc atau .zshrc (tergantung shell yang Anda gunakan) untuk menambahkan Go ke variabel PATH:

```
nano ~/.zshrc
```

Tambahkan baris berikut di bagian bawah file:
```
export PATH=$PATH:/usr/local/go/bin
```

Aktifkan perubahan:
```
source ~/.zshrc
```

### 4. Verifikasi Instalasi
Pastikan Go terinstal dengan benar dengan menjalankan perintah berikut:

```
go version
```

Anda akan melihat output versi Go yang terinstal, misalnya:

```
go version go1.19.6 linux/amd64
```


## Membuat Proyek Go Baru dengan Go Modules
Dengan Go Modules, Anda bisa membuat proyek di mana saja, tanpa harus berada di dalam direktori GOPATH.

### 1. Buat Direktori Proyek
Pilih direktori di mana Anda ingin membuat proyek Go. Misalnya:

```
mkdir Tes-Golang
cd Tes-Golang
```

### 2. Inisialisasi Go Modules
Jalankan perintah berikut untuk menginisialisasi **Go Modules** (Contoh dengan direktori github saya). Ini akan membuat file `go.mod` di dalam direktori proyek Anda:
```
go mod init github.com/Rafly1818/Project-GoLang
```

### 3. Tulis Kode Go Pertama Anda
Buat file baru bernama `main.go` di dalam direktori proyek dan tulis kode berikut:
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go Modules!")
}
```

### 4. Jalankan Program
```
go run main.go
```

Anda akan melihat output:
```
Hello, Go Modules!
```

## Menambahkan Dependensi dengan Go Modules (Jika Sudah Pro)
Jika proyek Anda memerlukan dependensi dari luar, Go Modules akan mengelola versi library secara otomatis. Misalnya, untuk menambahkan library `gorilla/mux` sebagai dependency:

```
go get github.com/gorilla/mux
```

Ini akan menambahkan dependensi tersebut ke dalam file `go.mod` dan `go.sum`.

## Build Proyek Go
Setelah selesai menulis kode, Anda dapat membuild proyek Go Anda menjadi file binary. Jalankan perintah berikut untuk membuild proyek:

```
go build
```

Ini akan menghasilkan file binary bernama Project-GoLang (atau sesuai nama modul Anda), yang bisa dieksekusi.


## Struktur Direktori Proyek
Dengan pendekatan Go Modules, struktur direktori proyek Anda akan terlihat seperti ini:

```
Project-GoLang/
  ├── go.mod
  ├── go.sum
  └── main.go
```
- go.mod: File yang mendefinisikan modul dan dependensi proyek Anda.
- go.sum: File yang menyimpan versi spesifik dari dependensi proyek.
- main.go: File sumber kode utama.

## Menjalankan Tes
Go memiliki dukungan testing bawaan. Untuk menulis tes, buat file dengan nama yang diakhiri dengan `_test.go`, lalu tulis tes Anda menggunakan paket `testing`.  
Misalnya, buat file `main_test.go`:

```
package main

import "testing"

func TestMain(t *testing.T) {
    expected := "Hello, Go Modules!"
    result := "Hello, Go Modules!"

    if result != expected {
        t.Errorf("Hasil tidak sesuai: got %v, want %v", result, expected)
    }
}
```

Untuk menjalankan tes, cukup jalankan perintah:
```
go test
```

---

<h3 align="left">Socials</h3>
<p align="center"> 
  <a href="https://www.github.com/Rafly1818" target="_blank" rel="noreferrer"> 
    <picture> 
      <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/danielcranney/readme-generator/main/public/icons/socials/github-dark.svg" /> 
      <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/danielcranney/readme-generator/main/public/icons/socials/github.svg" /> 
      <img src="https://raw.githubusercontent.com/danielcranney/readme-generator/main/public/icons/socials/github.svg" width="32" height="32" alt="GitHub" /> 
    </picture> 
  </a> 
  <a href="http://www.instagram.com/flyyr_" target="_blank" rel="noreferrer"> 
    <picture> 
      <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/danielcranney/readme-generator/main/public/icons/socials/instagram-dark.svg" /> 
      <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/danielcranney/readme-generator/main/public/icons/socials/instagram.svg" /> 
      <img src="https://raw.githubusercontent.com/danielcranney/readme-generator/main/public/icons/socials/instagram.svg" width="32" height="32" alt="Instagram" /> 
    </picture> 
  </a>
</p>