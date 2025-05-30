package main

import "fmt"

const (
	maxUsers        = 10
	maxTransactions = 30
	maxNFTs         = 100
)

type NFT struct {
	ID          int
	Nama        string
	Harga       int
	Popularitas int
	Pemilik     string
}

var (
	usernames          [maxUsers]string
	passwords          [maxUsers]string
	balances           [maxUsers]float64
	transactionHistory [maxUsers][maxTransactions]string
	transactionCount   [maxUsers]int
	userCount          int
	loggedInUserIndex  int = -1
	nfts               [maxNFTs]NFT
	nftCount           int
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
	if userCount >= maxUsers {
		fmt.Println("User penuh.")
		return
	}
	var uname, pass string
	fmt.Print("Username baru: ")
	fmt.Scanln(&uname)
	fmt.Print("Password: ")
	fmt.Scanln(&pass)

	for i := 0; i < userCount; i++ {
		if usernames[i] == uname {
			fmt.Println("Username sudah digunakan.")
			return
		}
	}

	usernames[userCount] = uname
	passwords[userCount] = pass
	balances[userCount] = 0
	userCount++

	fmt.Println("Registrasi berhasil!")
}

func login() {
	var uname, pass string
	fmt.Print("Username: ")
	fmt.Scanln(&uname)
	fmt.Print("Password: ")
	fmt.Scanln(&pass)

	for i := 0; i < userCount; i++ {
		if usernames[i] == uname && passwords[i] == pass {
			loggedInUserIndex = i
			fmt.Printf("Selamat datang, %s!\n", uname)
			dashboard()
			return
		}
	}
	fmt.Println("Login gagal.")
}

func dashboard() {
	for {
		fmt.Printf("\n=== DASHBOARD %s ===\n", usernames[loggedInUserIndex])
		fmt.Printf("Saldo: %.2f koin\n", balances[loggedInUserIndex])
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
			loggedInUserIndex = -1
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
		balances[loggedInUserIndex] += jumlah
		catat(fmt.Sprintf("Top up %.2f koin", jumlah))
		fmt.Println("Berhasil top up!")
	} else {
		fmt.Println("Jumlah tidak valid.")
	}
}

func tambahNFT() {
	if nftCount >= maxNFTs {
		fmt.Println("NFT penuh.")
		return
	}
	id := inputInt("ID NFT: ")
	nama := inputStr("Nama: ")
	harga := inputInt("Harga: ")
	pop := inputInt("Popularitas: ")
	nfts[nftCount] = NFT{id, nama, harga, pop, usernames[loggedInUserIndex]}
	nftCount++
	catat(fmt.Sprintf("Menambahkan NFT %s", nama))
	fmt.Println("NFT ditambahkan.")
}

func tampilkanNFT() {
	fmt.Println("\nDAFTAR NFT:")
	for i := 0; i < nftCount; i++ {
		fmt.Printf("[%d] ID: %d | Nama: %s | Harga: %d | Popularitas: %d | Pemilik: %s\n",
			i+1, nfts[i].ID, nfts[i].Nama, nfts[i].Harga, nfts[i].Popularitas, nfts[i].Pemilik)
	}
}

func beliNFT() {
	tampilkanNFT()
	index := inputInt("Pilih nomor NFT yang ingin dibeli: ") - 1
	if index >= 0 && index < nftCount {
		nft := nfts[index]
		if nft.Pemilik == usernames[loggedInUserIndex] {
			fmt.Println("Tidak bisa membeli NFT milik sendiri.")
			return
		}
		if float64(nft.Harga) > balances[loggedInUserIndex] {
			fmt.Println("Saldo tidak cukup.")
			return
		}
		for i := 0; i < userCount; i++ {
			if usernames[i] == nft.Pemilik {
				balances[i] += float64(nft.Harga)
				catatTo(i, fmt.Sprintf("NFT %s terjual ke %s", nft.Nama, usernames[loggedInUserIndex]))
			}
		}
		balances[loggedInUserIndex] -= float64(nft.Harga)
		catat(fmt.Sprintf("Membeli NFT %s", nft.Nama))

		for i := index; i < nftCount-1; i++ {
			nfts[i] = nfts[i+1]
		}
		nftCount--
		fmt.Println("NFT berhasil dibeli!")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func editNFT() {
	idx := inputInt("Nomor NFT yang ingin diedit: ") - 1
	if idx >= 0 && idx < nftCount && nfts[idx].Pemilik == usernames[loggedInUserIndex] {
		nfts[idx].Nama = inputStr("Nama baru: ")
		nfts[idx].Harga = inputInt("Harga baru: ")
		nfts[idx].Popularitas = inputInt("Popularitas baru: ")
		fmt.Println("NFT diperbarui.")
	} else {
		fmt.Println("Index tidak valid atau bukan milik Anda.")
	}
}

func hapusNFT() {
	idx := inputInt("Nomor NFT yang ingin dihapus: ") - 1
	if idx >= 0 && idx < nftCount && nfts[idx].Pemilik == usernames[loggedInUserIndex] {
		for i := idx; i < nftCount-1; i++ {
			nfts[i] = nfts[i+1]
		}
		nftCount--
		fmt.Println("NFT dihapus.")
	} else {
		fmt.Println("Index tidak valid atau bukan milik Anda.")
	}
}

func cariNFT() {
	nama := inputStr("Nama NFT yang dicari: ")
	found := false
	for i := 0; i < nftCount; i++ {
		if nfts[i].Nama == nama {
			fmt.Printf("Ditemukan: %v\n", nfts[i])
			found = true
			break
		}
	}
	if !found {
		fmt.Println("NFT tidak ditemukan.")
	}
}

func urutHargaNaik() {
	for i := 0; i < nftCount-1; i++ {
		minIdx := i
		for j := i + 1; j < nftCount; j++ {
			if nfts[j].Harga < nfts[minIdx].Harga {
				minIdx = j
			}
		}
		nfts[i], nfts[minIdx] = nfts[minIdx], nfts[i]
	}
	fmt.Println("NFT diurutkan berdasarkan harga naik.")
}

func urutPopularitasTurun() {
	for i := 1; i < nftCount; i++ {
		temp := nfts[i]
		j := i - 1
		for j >= 0 && nfts[j].Popularitas < temp.Popularitas {
			nfts[j+1] = nfts[j]
			j--
		}
		nfts[j+1] = temp
	}
	fmt.Println("NFT diurutkan berdasarkan popularitas turun.")
}

func lihatHistori() {
	fmt.Println("\n=== HISTORI TRANSAKSI ===")
	if transactionCount[loggedInUserIndex] == 0 {
		fmt.Println("Tidak ada histori.")
		return
	}
	for i := 0; i < transactionCount[loggedInUserIndex]; i++ {
		fmt.Println("-", transactionHistory[loggedInUserIndex][i])
	}
}

func catat(isi string) {
	if transactionCount[loggedInUserIndex] < maxTransactions {
		transactionHistory[loggedInUserIndex][transactionCount[loggedInUserIndex]] = isi
		transactionCount[loggedInUserIndex]++
	}
}

func catatTo(userIndex int, isi string) {
	if transactionCount[userIndex] < maxTransactions {
		transactionHistory[userIndex][transactionCount[userIndex]] = isi
		transactionCount[userIndex]++
	}
}

func inputInt(prompt string) int {
	var val int
	fmt.Print(prompt)
	fmt.Scanln(&val)
	return val
}

func inputStr(prompt string) string {
	var val string
	fmt.Print(prompt)
	fmt.Scanln(&val)
	return val
}
