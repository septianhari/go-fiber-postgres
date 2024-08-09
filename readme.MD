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

   git clone https://github.com/username/repo-name.git
   cd repo-name

2. Instal dependensi:
   
    go mod tidy

3. Buat file .env di direktori proyek Anda dan tambahkan konfigurasi database PostgreSQL: 

    DB_HOST=localhost
    DB_USER=your_username
    DB_PASS=your_password
    DB_NAME=your_database_name
    DB_PORT=5432

4. Migrasikan database: 
Model Books diharapkan sudah didefinisikan dalam package models. Fungsi MigrateBooks di models akan digunakan untuk memigrasi database.

func MigrateBooks(db *gorm.DB) error {
    // Implementasikan migrasi model di sini
    return db.AutoMigrate(&Books{})
}


## Menjalankan Aplikasi

1. Jalankan aplikasi dengan perintah berikut:

go run main.go

2. Aplikasi akan berjalan di http://localhost:8080.

## API Endpoint

- POST /api/create_book
- Menambahkan buku baru.
- Contoh Request Body:

{
  "author": "Author Name",
  "title": "Book Title",
  "publisher": "Publisher Name"
}


- GET /api/get_books
    - Mendapatkan semua data buku.

- GET /api/get_book/
    - Mendapatkan buku berdasarkan ID.

- DELETE /api/delete_book/
    - Menghapus buku berdasarkan ID.






