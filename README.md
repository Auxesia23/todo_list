# Todo List API

Aplikasi Todo List berbasis REST API yang dibangun menggunakan Go, GORM, dan PostgreSQL. Aplikasi ini menyediakan manajemen tugas (todo) dengan fitur autentikasi pengguna.

## Fitur Utama

- Manajemen Todo (CRUD)
  - Membuat todo baru
  - Melihat daftar todo
  - Mengupdate todo
  - Menghapus todo
  - Menandai todo sebagai selesai
- Manajemen Pengguna
  - Registrasi pengguna
  - Login pengguna
  - Autentikasi menggunakan JWT
- Relasi antara Todo dan Pengguna
- Pencatatan waktu tenggat (due date) untuk setiap todo

## Teknologi yang Digunakan

- [Go](https://golang.org/) - Bahasa pemrograman
- [GORM](https://gorm.io/) - ORM (Object Relational Mapping)
- [Chi Router](https://github.com/go-chi/chi) - HTTP router
- [PostgreSQL](https://www.postgresql.org/) - Database
- [JWT](https://github.com/golang-jwt/jwt) - JSON Web Token untuk autentikasi
- [godotenv](https://github.com/joho/godotenv) - Manajemen environment variables

## Struktur Data

### Todo
```go
type Todo struct {
    Title       string    // Judul todo
    Description string    // Deskripsi todo
    DueDate     time.Time // Waktu tenggat
    Completed   bool      // Status penyelesaian
    UserEmail   string    // Email pengguna pemilik todo
}
```

## Cara Instalasi

1. Clone repositori ini
```bash
git clone https://github.com/Auxesia23/todo_list.git
```

2. Masuk ke direktori proyek
```bash
cd todo_list
```

3. Salin file .env.example ke .env dan sesuaikan konfigurasi
```bash
cp .env.example .env
```

4. Install dependensi
```bash
go mod download
```

5. Jalankan aplikasi
```bash
go run cmd/api.go
```

## Penggunaan API

### Autentikasi

- **Register**: `POST /api/register`
- **Login**: `POST /api/login`

### Todo

- **Membuat Todo**: `POST /api/todos`
- **Mengambil Semua Todo**: `GET /api/todos`
- **Mengambil Todo by ID**: `GET /api/todos/{id}`
- **Update Todo**: `PUT /api/todos/{id}`
- **Hapus Todo**: `DELETE /api/todos/{id}`
- **Menandai Todo Selesai**: `PATCH /api/todos/{id}/complete`

## Kontribusi

Kontribusi selalu diterima. Silakan buat pull request untuk memberikan perbaikan atau penambahan fitur.
