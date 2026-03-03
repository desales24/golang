Baik 👍 karena kamu **belum pernah install Golang sebelumnya** dan pakai **zsh di WSL Ubuntu**, kita mulai dari nol dengan cara yang paling aman dan direkomendasikan (install versi resmi dari situs Go).

---

# 🚀 Instalasi Golang di WSL Ubuntu (zsh)

## ✅ 1. Update Ubuntu dulu

```bash
sudo apt update && sudo apt upgrade -y
```

---

## ✅ 2. Download Go versi terbaru

Download dari situs resmi:
👉 [https://go.dev/dl/](https://go.dev/dl/)

Atau langsung via terminal (contoh versi 1.22.0 — kalau ada versi lebih baru, bisa ganti angkanya):

```bash
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
```

---

## ✅ 3. Install Go ke sistem

```bash
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```

---

## ✅ 4. Tambahkan Go ke PATH (karena kamu pakai zsh)

Edit file `.zshrc`:

```bash
nano ~/.zshrc
```

Tambahkan baris ini di paling bawah:

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Simpan lalu keluar.

---

## ✅ 5. Reload zsh

```bash
source ~/.zshrc
```

---

## ✅ 6. Cek apakah berhasil

```bash
go version
```

Kalau berhasil, akan muncul seperti:

```
go version go1.22.0 linux/amd64
```

---

# 🎯 Tes Program Go Pertama

Buat file:

```bash
nano hello.go
```

Isi dengan:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go di WSL!")
}
```

Jalankan:

```bash
go run hello.go
```

Kalau keluar tulisan `Hello, Go di WSL!` berarti sukses 🎉

---

Kalau mau, saya bisa bantu lanjut ke:

* Setup Go Modules (`go mod init`)
* Setup struktur project yang benar
* Integrasi Go dengan VS Code di WSL
* Atau cara install via apt (lebih simpel tapi biasanya versi lama)

Kamu pakai Ubuntu versi berapa di WSL? (ketik `lsb_release -a`)
