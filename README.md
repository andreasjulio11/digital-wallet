Digital Wallet API
Aplikasi RESTful API untuk sistem dompet digital menggunakan Golang dan Gin Framework. Proyek ini dirancang dengan fitur manajemen saldo yang aman menggunakan database transaction dan sistem autentikasi JWT

Fitur :
1. User Management: Registrasi dan Login menggunakan JWT.
2. Wallet System: Fitur Top-up dan Withdraw saldo.
3. Database Transaction: Menjamin keamanan data saldo
4. Auto Migration: Skema database akan otomatis menyesuaikan dengan struct Go setiap kali aplikasi dijalankan.

Tech Stack "
1. Language: Go (Golang) 25.5
2. Framework: Gin Gonic
3. ORM: GORM
4. Database: PostgreSQL / MySQL
5. Auth: JWT

Persiapan & Instalasi
1. Clone Repository & Install Dependencies:
2. Konfigurasi Database: Pastikan koneksi database, rename .env.example sesuaikan dengan konfigurasi database lokal