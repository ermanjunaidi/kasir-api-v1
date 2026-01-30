# Kasir API v1

API sederhana untuk sistem kasir berbasis Go yang terintegrasi dengan database PostgreSQL.

## Fitur
- Manajemen Produk (CRUD)
- Manajemen Kategori
- Integrasi PostgreSQL menggunakan `pgx`
- Konfigurasi menggunakan Environment Variables

## Persyaratan
- [Go](https://golang.org/dl/) (versi 1.20 atau terbaru)
- [PostgreSQL](https://www.postgresql.org/download/)

## Instalasi & Persiapan

1. **Clone repository ini:**
   ```bash
   git clone <repository-url>
   cd kasir-api-v1
   ```

2. **Setup Environment Variables:**
   Buat file `.env` di direktori root dan sesuaikan konfigurasinya:
   ```env
   PORT=8080
   DB_CONN=postgres://username:password@localhost:5432/database_name?sslmode=disable
   ```

3. **Install Dependensi:**
   ```bash
   go mod tidy
   ```

## Cara Menjalankan

Jalankan aplikasi dengan perintah:
```bash
go run main.go
```
Aplikasi akan berjalan di `http://localhost:8080`.

## Dokumentasi API

Anda dapat menggunakan file `data.rest` jika menggunakan ekstensi REST Client di VS Code.

### Produk
- **GET** `/api/produk` - Mengambil semua produk
- **GET** `/api/produk/{id}` - Mengambil detail produk berdasarkan ID
- **POST** `/api/produk` - Membuat produk baru
- **PUT** `/api/produk/{id}` - Memperbarui produk yang ada
- **DELETE** `/api/produk/{id}` - Menghapus produk

### Health Check
- **GET** `/health` - Mengecek status API

## Struktur Proyek
- `main.go`: Entry point aplikasi dan setup router.
- `handlers/`: Logika untuk menangani request HTTP.
- `services/`: Logika bisnis.
- `repositories/`: Interaksi langsung dengan database.
- `models/`: Definisi struktur data (struct).
- `database/`: Inisialisasi koneksi database.
- `data.rest`: File testing untuk REST Client.
