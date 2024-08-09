markdown
Copy code
# Golang Fiber PostgreSQL CRUD API

Ini adalah proyek sederhana yang menggunakan [Golang](https://golang.org/), [Fiber](https://gofiber.io/), dan [PostgreSQL](https://www.postgresql.org/) untuk membuat RESTful API sederhana dengan operasi CRUD (Create, Read, Update, Delete) untuk buku.

## Fitur

- **Create Book**: Menambahkan data buku baru.
- **Get All Books**: Mendapatkan semua data buku yang ada.
- **Get Book By ID**: Mendapatkan data buku berdasarkan ID.
- **Delete Book**: Menghapus data buku berdasarkan ID.

## Struktur Proyek

- `main.go`: File utama yang berisi pengaturan routing dan logika utama aplikasi.
- `models/`: Folder untuk mendefinisikan model database.
- `.env`: File untuk menyimpan variabel lingkungan seperti informasi koneksi database.

## Prasyarat

- Go 1.16 atau versi lebih baru
- PostgreSQL yang terinstal dan dikonfigurasi

## Instalasi

1. Clone repositori ini:

   ```bash
   git clone https://github.com/septianhari/go-fiber-postgres
Instal dependensi:

bash
Copy code
go mod tidy
Buat file .env di direktori proyek Anda dan tambahkan konfigurasi database PostgreSQL:

env
Copy code
DB_HOST=localhost
DB_USER=your_username
DB_PASS=your_password
DB_NAME=your_database_name
DB_PORT=5432
Migrasikan database:

Model Books diharapkan sudah didefinisikan dalam package models. Fungsi MigrateBooks di models akan digunakan untuk memigrasi database.

go
Copy code
func MigrateBooks(db *gorm.DB) error {
    // Implementasikan migrasi model di sini
    return db.AutoMigrate(&Books{})
}
Menjalankan Aplikasi
Jalankan aplikasi dengan perintah berikut:

bash
Copy code
go run main.go
Aplikasi akan berjalan di http://localhost:8080.

API Endpoint
POST /api/create_book

Menambahkan buku baru.

Contoh Request Body:

json
Copy code
{
  "author": "Author Name",
  "title": "Book Title",
  "publisher": "Publisher Name"
}
GET /api/get_books

Mendapatkan semua data buku.
GET /api/get_book/

Mendapatkan buku berdasarkan ID.
DELETE /api/delete_book/

Menghapus buku berdasarkan ID.
