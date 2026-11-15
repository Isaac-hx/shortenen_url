Fitur intinya mengikuti tantangan roadmap.sh: buat, ambil, ubah, dan hapus short URL, serta lacak statistik akses. Tantangan itu juga menyatakan tidak perlu autentikasi atau otorisasi. Jadi, di project ini jangan buat login dulu. Itu bagian Project 2.
go
roadmap

Bagian yang saya tambahkan sendiri, di luar tantangan roadmap.sh:

Redirect langsung lewat GET /{code}.
Cache Redis dan rate limiting.
Custom alias. 2. Endpoint API
Method Path Fungsi Status sukses
POST /api/v1/urls Buat short URL 201
GET /{code} Redirect ke URL asli 302
GET /api/v1/urls/{code} Detail URL 200
PUT /api/v1/urls/{code} Ubah URL tujuan atau masa berlaku 200
DELETE /api/v1/urls/{code} Hapus 204
GET /api/v1/urls/{code}/stats Statistik akses 200
GET /health Cek kesehatan service 200

Contoh request dan response POST /api/v1/urls:

json
// request
{ "url": "https://example.com/artikel/panjang?id=123", "custom_alias": "artikel1", "expires_at": "2026-12-31T23:59:59Z" }

// response 201
{ "code": "artikel1", "short_url": "http://localhost:8080/artikel1", "url": "https://example.com/artikel/panjang?id=123", "created_at": "...", "expires_at": "..." }

custom_alias dan expires_at bersifat opsional.

Format error yang konsisten:

json
{ "error": { "code": "invalid_url", "message": "URL harus diawali http:// atau https://" } }

Status error yang dipakai:

Status Kapan dipakai
400 Input tidak valid
404 Kode tidak ditemukan
409 Alias sudah dipakai
410 Link sudah kedaluwarsa
429 Terkena rate limit 3. Aturan bisnis
Validasi URL. Hanya skema http dan https, panjang maksimal 2048 karakter, dan format harus valid.
Kode pendek. Panjang 7 karakter base62 (a-z, A-Z, 0-9), dibuat acak dengan generator yang aman secara kriptografi (crypto/rand di Go, secrets di Python). Ruang kodenya 62⁷ ≈ 3,5 triliun kombinasi. Sumber desain menyebut bahwa link pendek sebaiknya tidak bisa ditebak. Karena itu, jangan pakai auto-increment ID yang di-encode ke base62, sebab urutannya mudah ditebak.
designgurus
Tabrakan kode. Kolom code wajib UNIQUE. Kalau insert gagal karena duplikat, coba lagi dengan kode baru (maksimal 5 kali).
Custom alias. 3-32 karakter, hanya a-z A-Z 0-9 - \_. Tolak kata yang bentrok dengan rute sistem, misalnya api dan health.
Masa berlaku. Kalau expires_at kosong, link tidak kedaluwarsa. Sumber lain menyarankan masa berlaku default yang bisa diatur sendiri oleh pengguna. Kamu bebas memilih salah satu. Yang penting, dokumentasikan pilihanmu.
educative
Jenis redirect. Pakai 302, bukan 301. Browser menyimpan 301 secara permanen, sehingga klik berikutnya tidak lewat servermu dan statistik jadi kurang akurat. Ini pengetahuan umum tentang semantik HTTP. Kamu bisa memverifikasinya di RFC 9110 bagian 15.4.
