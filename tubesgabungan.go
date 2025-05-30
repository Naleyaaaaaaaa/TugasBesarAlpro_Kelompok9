package main

import "fmt"

const (
	maksUser        = 10
	maksTransaksi   = 30
	maksNFT         = 100
)

type NFT struct {
	ID          int
	Nama        string
	Harga       int
	Popularitas int
	Pemilik     string
}

var (
	daftarUsername       [maksUser]string
	daftarPassword       [maksUser]string
	saldo               [maksUser]float64
	historiTransaksi    [maksUser][maksTransaksi]string
	jumlahTransaksi     [maksUser]int
	jumlahUser          int
	indeksUserLogin     int = -1
	daftarNFT           [maksNFT]NFT
	jumlahNFT           int
)

func main() {
	for {
		fmt.Println("\n=== SISTEM INVESTASI NFT ===")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			registrasi()
		case 2:
			login()
		case 3:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func registrasi() {
	if jumlahUser >= maksUser {
		fmt.Println("User penuh.")
		return
	}
	var uname, pass string
	fmt.Print("Username baru: ")
	fmt.Scanln(&uname)
	fmt.Print("Password: ")
	fmt.Scanln(&pass)

	for i := 0; i < jumlahUser; i++ {
		if daftarUsername[i] == uname {
			fmt.Println("Username sudah digunakan.")
			return
		}
	}

	daftarUsername[jumlahUser] = uname
	daftarPassword[jumlahUser] = pass
	saldo[jumlahUser] = 0
	jumlahUser++

	fmt.Println("Registrasi berhasil!")
}

func login() {
	var uname, pass string
	fmt.Print("Username: ")
	fmt.Scanln(&uname)
	fmt.Print("Password: ")
	fmt.Scanln(&pass)

	for i := 0; i < jumlahUser; i++ {
		if daftarUsername[i] == uname && daftarPassword[i] == pass {
			indeksUserLogin = i
			fmt.Printf("Selamat datang, %s!\n", uname)
			dashboard()
			return
		}
	}
	fmt.Println("Login gagal.")
}

func dashboard() {
	for {
		fmt.Printf("\n=== DASHBOARD %s ===\n", daftarUsername[indeksUserLogin])
		fmt.Printf("Saldo: %.2f koin\n", saldo[indeksUserLogin])
		fmt.Println("1. Top Up")
		fmt.Println("2. Tambah NFT")
		fmt.Println("3. Tampilkan NFT")
		fmt.Println("4. Beli NFT")
		fmt.Println("5. Edit NFT")
		fmt.Println("6. Hapus NFT")
		fmt.Println("7. Cari NFT (Nama)")
		fmt.Println("8. Urutkan Harga (Naik)")
		fmt.Println("9. Urutkan Popularitas (Turun)")
		fmt.Println("10. Histori Transaksi")
		fmt.Println("11. Logout")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			topUp()
		case 2:
			tambahNFT()
		case 3:
			tampilkanNFT()
		case 4:
			beliNFT()
		case 5:
			editNFT()
		case 6:
			hapusNFT()
		case 7:
			cariNFT()
		case 8:
			urutHargaNaik()
		case 9:
			urutPopularitasTurun()
		case 10:
			lihatHistori()
		case 11:
			indeksUserLogin = -1
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func topUp() {
	var jumlah float64
	fmt.Print("Jumlah top up: ")
	fmt.Scanln(&jumlah)
	if jumlah > 0 {
		saldo[indeksUserLogin] += jumlah
		catat(fmt.Sprintf("Top up %.2f koin", jumlah))
		fmt.Println("Berhasil top up!")
	} else {
		fmt.Println("Jumlah tidak valid.")
	}
}

func tambahNFT() {
	if jumlahNFT >= maksNFT {
		fmt.Println("NFT penuh.")
		return
	}
	id := inputInt("ID NFT: ")
	nama := inputStr("Nama: ")
	harga := inputInt("Harga: ")
	pop := inputInt("Popularitas: ")
	daftarNFT[jumlahNFT] = NFT{id, nama, harga, pop, daftarUsername[indeksUserLogin]}
	jumlahNFT++
	catat(fmt.Sprintf("Menambahkan NFT %s", nama))
	fmt.Println("NFT ditambahkan.")
}

func tampilkanNFT() {
	fmt.Println("\nDAFTAR NFT:")
	for i := 0; i < jumlahNFT; i++ {
		fmt.Printf("[%d] ID: %d | Nama: %s | Harga: %d | Popularitas: %d | Pemilik: %s\n",
			i+1, daftarNFT[i].ID, daftarNFT[i].Nama, daftarNFT[i].Harga, daftarNFT[i].Popularitas, daftarNFT[i].Pemilik)
	}
}

func beliNFT() {
	tampilkanNFT()
	indeks := inputInt("Pilih nomor NFT yang ingin dibeli: ") - 1
	if indeks >= 0 && indeks < jumlahNFT {
		nft := &daftarNFT[indeks]
		if nft.Pemilik == daftarUsername[indeksUserLogin] {
			fmt.Println("Tidak bisa membeli NFT milik sendiri.")
			return
		}
		if float64(nft.Harga) > saldo[indeksUserLogin] {
			fmt.Println("Saldo tidak cukup.")
			return
		}

		for i := 0; i < jumlahUser; i++ {
			if daftarUsername[i] == nft.Pemilik {
				saldo[i] += float64(nft.Harga)
				catatTo(i, fmt.Sprintf("NFT %s terjual ke %s", nft.Nama, daftarUsername[indeksUserLogin]))
				break
			}
		}
		saldo[indeksUserLogin] -= float64(nft.Harga)
		catat(fmt.Sprintf("Membeli NFT %s dari %s", nft.Nama, nft.Pemilik))
		nft.Pemilik = daftarUsername[indeksUserLogin]
		fmt.Println("NFT berhasil dibeli dan kini menjadi milik Anda!")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func editNFT() {
	indeks := inputInt("Nomor NFT yang ingin diedit: ") - 1
	if indeks >= 0 && indeks < jumlahNFT && daftarNFT[indeks].Pemilik == daftarUsername[indeksUserLogin] {
		daftarNFT[indeks].Nama = inputStr("Nama baru: ")
		daftarNFT[indeks].Harga = inputInt("Harga baru: ")
		daftarNFT[indeks].Popularitas = inputInt("Popularitas baru: ")
		fmt.Println("NFT diperbarui.")
	} else {
		fmt.Println("Index tidak valid atau bukan milik Anda.")
	}
}

func hapusNFT() {
	indeks := inputInt("Nomor NFT yang ingin dihapus: ") - 1
	if indeks >= 0 && indeks < jumlahNFT && daftarNFT[indeks].Pemilik == daftarUsername[indeksUserLogin] {
		for i := indeks; i < jumlahNFT-1; i++ {
			daftarNFT[i] = daftarNFT[i+1]
		}
		jumlahNFT--
		fmt.Println("NFT dihapus.")
	} else {
		fmt.Println("Index tidak valid atau bukan milik Anda.")
	}
}

func cariNFT() {
	nama := inputStr("Nama NFT yang dicari: ")
	ditemukan := false
	for i := 0; i < jumlahNFT; i++ {
		if daftarNFT[i].Nama == nama {
			fmt.Printf("Ditemukan: %v\n", daftarNFT[i])
			ditemukan = true
			break
		}
	}
	if !ditemukan {
		fmt.Println("NFT tidak ditemukan.")
	}
}

func urutHargaNaik() {
	for i := 0; i < jumlahNFT-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahNFT; j++ {
			if daftarNFT[j].Harga < daftarNFT[minIdx].Harga {
				minIdx = j
			}
		}
		daftarNFT[i], daftarNFT[minIdx] = daftarNFT[minIdx], daftarNFT[i]
	}
	fmt.Println("NFT diurutkan berdasarkan harga naik.")
}

func urutPopularitasTurun() {
	for i := 1; i < jumlahNFT; i++ {
		temp := daftarNFT[i]
		j := i - 1
		for j >= 0 && daftarNFT[j].Popularitas < temp.Popularitas {
			daftarNFT[j+1] = daftarNFT[j]
			j--
		}
		daftarNFT[j+1] = temp
	}
	fmt.Println("NFT diurutkan berdasarkan popularitas turun.")
}

func lihatHistori() {
	fmt.Println("\n=== HISTORI TRANSAKSI ===")
	if jumlahTransaksi[indeksUserLogin] == 0 {
		fmt.Println("Tidak ada histori.")
		return
	}
	for i := 0; i < jumlahTransaksi[indeksUserLogin]; i++ {
		fmt.Println("-", historiTransaksi[indeksUserLogin][i])
	}
}

func catat(isi string) {
	if jumlahTransaksi[indeksUserLogin] < maksTransaksi {
		historiTransaksi[indeksUserLogin][jumlahTransaksi[indeksUserLogin]] = isi
		jumlahTransaksi[indeksUserLogin]++
	}
}

func catatTo(indeks int, isi string) {
	if jumlahTransaksi[indeks] < maksTransaksi {
		historiTransaksi[indeks][jumlahTransaksi[indeks]] = isi
		jumlahTransaksi[indeks]++
	}
}

func inputInt(pesan string) int {
	var val int
	fmt.Print(pesan)
	fmt.Scanln(&val)
	return val
}

func inputStr(pesan string) string {
	var val string
	fmt.Print(pesan)
	fmt.Scanln(&val)
	return val
}
