package main

import (
	"fmt"
)

type Student struct {
	ID         int
	Name       string
	Department string
	TestScore  int
	Accepted   bool
}

type Department struct {
	ID   int
	Name string
}

const maxStudents = 10
const maxDepartments = 10

var students [maxStudents]Student
var departments [maxDepartments]Department
var studentCount, departmentCount int

func addStudent(id int, name, department string) {
	if studentCount < maxStudents {
		students[studentCount] = Student{ID: id, Name: name, Department: department}
		studentCount++
		fmt.Println("Mahasiswa berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas mahasiswa penuh.")
	}
}

func addDepartment(id int, name string) {
	if departmentCount < maxDepartments {
		departments[departmentCount] = Department{ID: id, Name: name}
		departmentCount++
		fmt.Println("Jurusan berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas jurusan penuh.")
	}
}

func deleteStudent(id int) {
	for i := 0; i < studentCount; i++ {
		if students[i].ID == id {
			students[i] = students[studentCount-1]
			studentCount--
			fmt.Println("Mahasiswa berhasil dihapus.")
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}

func deleteDepartment(id int) {
	for i := 0; i < departmentCount; i++ {
		if departments[i].ID == id {
			departments[i] = departments[departmentCount-1]
			departmentCount--
			fmt.Println("Jurusan berhasil dihapus.")
			return
		}
	}
	fmt.Println("Jurusan tidak ditemukan.")
}

func addTestScore(studentID int, score int) {
	for i := 0; i < studentCount; i++ {
		if students[i].ID == studentID {
			if students[i].TestScore == 0 {
				students[i].TestScore = score
				students[i].Accepted = score > 70
				fmt.Println("Nilai tes berhasil ditambahkan.")
			} else {
				fmt.Println("Mahasiswa sudah memiliki nilai tes. Gunakan fungsi ubah nilai tes untuk mengubahnya.")
			}
			return
		}
	}
	fmt.Println("Mahasiswa dengan ID tersebut tidak ditemukan.")
}

func editTestScore(studentID int, score int) {
	for i := 0; i < studentCount; i++ {
		if students[i].ID == studentID {
			students[i].TestScore = score
			students[i].Accepted = score > 70
			fmt.Println("Nilai tes berhasil diubah.")
			return
		}
	}
	fmt.Println("Mahasiswa dengan ID tersebut tidak ditemukan.")
}

func deleteTestScore(studentID int) {
	for i := 0; i < studentCount; i++ {
		if students[i].ID == studentID {
			students[i].TestScore = 0
			students[i].Accepted = false
			fmt.Println("Nilai tes berhasil dihapus.")
			return
		}
	}
	fmt.Println("Mahasiswa dengan ID tersebut tidak ditemukan.")
}

func insertionSortByTestScore() {
	for i := 1; i < studentCount; i++ {
		key := students[i]
		j := i - 1

		for j >= 0 && students[j].TestScore < key.TestScore {
			students[j+1] = students[j]
			j = j - 1
		}
		students[j+1] = key
	}

	fmt.Println("Mahasiswa berdasarkan nilai tes (diurutkan dengan insertion sort):")
	for i := 0; i < studentCount; i++ {
		fmt.Printf("Nama: %s, Jurusan: %s, Nilai Tes: %d\n", students[i].Name, students[i].Department, students[i].TestScore)
	}
}

func selectionSortByName() {
	for i := 0; i < studentCount-1; i++ {
		minIndex := i
		for j := i + 1; j < studentCount; j++ {
			if students[j].Name < students[minIndex].Name {
				minIndex = j
			}
		}
		students[i], students[minIndex] = students[minIndex], students[i]
	}

	fmt.Println("Mahasiswa berdasarkan nama (diurutkan dengan selection sort):")
	for i := 0; i < studentCount; i++ {
		fmt.Printf("Nama: %s, Jurusan: %s, Nilai Tes: %d\n", students[i].Name, students[i].Department, students[i].TestScore)
	}
}

func displayStudentsByDepartment() {
	for i := 0; i < studentCount-1; i++ {
		for j := 0; j < studentCount-i-1; j++ {
			if students[j].Department > students[j+1].Department {
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
	fmt.Println("Mahasiswa berdasarkan jurusan:")
	for i := 0; i < studentCount; i++ {
		fmt.Printf("Nama: %s, Jurusan: %s, Nilai Tes: %d\n", students[i].Name, students[i].Department, students[i].TestScore)
	}
}

func sequentialSearchStudentByName(name string) []Student {
	var result []Student
	for i := 0; i < studentCount; i++ {
		if students[i].Name == name {
			result = append(result, students[i])
		}
	}
	return result
}

func displayStudentsByAdmissionStatus(accepted bool) {
	status := "diterima"
	if !accepted {
		status = "ditolak"
	}
	fmt.Println("Mahasiswa yang", status, ":")
	for i := 0; i < studentCount; i++ {
		if students[i].Accepted == accepted {
			fmt.Printf("Nama: %s, Jurusan: %s, Nilai Tes: %d\n", students[i].Name, students[i].Department, students[i].TestScore)
		}
	}
}

func binarySearchByID(id int) *Student {
	low := 0
	high := studentCount - 1

	for low <= high {
		mid := (low + high) / 2
		if students[mid].ID == id {
			return &students[mid]
		} else if students[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func sortStudentsByID() {
	for i := 0; i < studentCount-1; i++ {
		for j := 0; j < studentCount-i-1; j++ {
			if students[j].ID > students[j+1].ID {
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
}

func main() {
	for {
		var userType int
        fmt.Println("==============================================")
		fmt.Println("Selamat datang di sistem penerimaan mahasiswa.")
        fmt.Println("==============================================")
		fmt.Println("Pilih jenis pengguna:")
		fmt.Println("1. Admin")
		fmt.Println("2. Mahasiswa")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&userType)

		if userType == 1 {
			adminMenu()
		} else if userType == 2 {
			studentMenu()
		} else if userType == 3 {
			fmt.Println("Keluar dari program.")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}

}

func adminMenu() {
	for {
		fmt.Println("Menu Admin:")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tambah Jurusan")
		fmt.Println("3. Tambah Nilai Tes")
		fmt.Println("4. Ubah Nilai Tes")
		fmt.Println("5. Hapus Mahasiswa")
		fmt.Println("6. Hapus Jurusan")
		fmt.Println("7. Tampilkan Mahasiswa berdasarkan Nilai Tes")
		fmt.Println("8. Tampilkan Mahasiswa berdasarkan Nama")
		fmt.Println("9. Cari Mahasiswa berdasarkan Nama")
		fmt.Println("10. Tampilkan Mahasiswa berdasarkan Jurusan")
		fmt.Println("11. Tampilkan Mahasiswa yang Diterima")
		fmt.Println("12. Tampilkan Mahasiswa yang Ditolak")
		fmt.Println("13. Hapus Nilai Tes Mahasiswa")
		fmt.Println("14. Cari Mahasiswa berdasarkan ID")
		fmt.Println("15. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id int
			var name, department string
			fmt.Print("Masukkan ID mahasiswa: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&name)
			fmt.Print("Masukkan jurusan mahasiswa: ")
			fmt.Scanln(&department)
			addStudent(id, name, department)
		case 2:
			var id int
			var name string
			fmt.Print("Masukkan ID jurusan: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan nama jurusan: ")
			fmt.Scanln(&name)
			addDepartment(id, name)
		case 3:
			var studentID, score int
			fmt.Print("Masukkan ID mahasiswa: ")
			fmt.Scanln(&studentID)
			fmt.Print("Masukkan nilai tes: ")
			fmt.Scanln(&score)
			addTestScore(studentID, score)
		case 4:
			var studentID, score int
			fmt.Print("Masukkan ID mahasiswa: ")
			fmt.Scanln(&studentID)
			fmt.Print("Masukkan nilai tes baru: ")
			fmt.Scanln(&score)
			editTestScore(studentID, score)
		case 5:
			var studentID int
			fmt.Print("Masukkan ID mahasiswa yang ingin dihapus: ")
			fmt.Scanln(&studentID)
			deleteStudent(studentID)
		case 6:
			var departmentID int
			fmt.Print("Masukkan ID jurusan yang ingin dihapus: ")
			fmt.Scanln(&departmentID)
			deleteDepartment(departmentID)
		case 7:
			insertionSortByTestScore()
		case 8:
			selectionSortByName()
		case 9:
			var name string
			fmt.Print("Masukkan nama mahasiswa yang ingin dicari: ")
			fmt.Scanln(&name)
			fmt.Println("Mencari mahasiswa dengan nama '", name, "':")
			result := sequentialSearchStudentByName(name)
			if len(result) == 0 {
				fmt.Println("Mahasiswa tidak ditemukan.")
			} else {
				for _, student := range result {
					fmt.Printf("Nama: %s, Jurusan: %s, Nilai Tes: %d\n", student.Name, student.Department, student.TestScore)
				}
			}
		case 10:
			displayStudentsByDepartment()
		case 11:
			displayStudentsByAdmissionStatus(true)
		case 12:
			displayStudentsByAdmissionStatus(false)
		case 13:
			var studentID int
			fmt.Print("Masukkan ID mahasiswa: ")
			fmt.Scanln(&studentID)
			deleteTestScore(studentID)
		case 14:
			var studentID int
			fmt.Print("Masukkan ID mahasiswa yang ingin dicari: ")
			fmt.Scanln(&studentID)
			sortStudentsByID()
			student := binarySearchByID(studentID)
			if student != nil {
				fmt.Printf("Nama: %s, Jurusan: %s, Nilai Tes: %d\n", student.Name, student.Department, student.TestScore)
			} else {
				fmt.Println("Mahasiswa dengan ID tersebut tidak ditemukan.")
			}
		case 15:
			fmt.Println("Keluar dari menu admin.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func studentMenu() {
	for {
		fmt.Println("Menu Mahasiswa:")
		fmt.Println("1. Tampilkan semua mahasiswa")
		fmt.Println("2. Tampilkan mahasiswa berdasarkan jurusan")
		fmt.Println("3. Tampilkan mahasiswa berdasarkan nama")
		fmt.Println("4. Tampilkan mahasiswa berdasarkan nilai tes")
		fmt.Println("5. Kembali ke menu utama")
		fmt.Println("6. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displayAllStudents()
		case 2:
			displayStudentsByDepartment()
		case 3:
			selectionSortByName()
		case 4:
			insertionSortByTestScore()
		case 5:
			return
		case 6:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func displayAllStudents() {
	fmt.Println("Semua Mahasiswa:")
	for i := 0; i < studentCount; i++ {
		fmt.Printf("Nama: %s, Jurusan: %s, Nilai Tes: %d\n", students[i].Name, students[i].Department, students[i].TestScore)
	}
}