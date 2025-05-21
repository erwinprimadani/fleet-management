# Fleet Management - README

## Menjalankan Program dengan Docker dan SQL Script

Dokumentasi ini menjelaskan cara menjalankan aplikasi Fleet Management menggunakan Docker, termasuk menjalankan script SQL untuk setup database.

---

## Prasyarat

- Docker dan Docker Compose sudah terinstall di mesin Anda
- Port 8080 (atau port lain yang Anda gunakan) tersedia untuk akses API dari luar Docker

---

## Struktur Proyek Terkait Docker

- `docker-compose.yml` - konfigurasi layanan Docker (Postgres, RabbitMQ, Mosquitto, API)
- Folder `script/database/` berisi script SQL untuk setup database
- Folder `cmd/fleettracking/` berisi simulasi fleet tracking dengan geofencing
---

## Langkah Menjalankan

### 1. Jalankan Docker Compose

Di root folder proyek, jalankan perintah berikut untuk membangun dan menjalankan semua container:

```bash
step 1. docker-compose build
step 2. docker-compose up postgres rabbitmq mosquitto -d
step 3. docker-compose up api -d atau menggunakan run and debug
```

Perintah ini akan:

- Menarik image yang diperlukan (Postgres, RabbitMQ, Mosquitto, dll)
- Membangun image aplikasi Anda dari Dockerfile
- Menjalankan container secara background (`-d`)

### 2. Menjalankan Script SQL setelah step

step 2.1. Lakukan ke database progres sesuai config yang berada di env.
step 2.2 Untuk data geofence insert data ke table location_landmark sesuai yang diinginkan. contoh: -6.2430015, 106.8246234 untuk halte mampang prapatan
step 2.3 anda dapat melanjukan ke step 3

### 3. Cek Status Container

Pastikan semua container berjalan dengan baik:

```bash
docker-compose ps
```

### 4. Akses API

Jika port binding sudah diatur (misal `3000:3000`), Anda bisa mengakses API dari host dengan:

```bash
curl http://localhost:8080/healthcheck
```

atau endpoint lain sesuai dokumentasi API.

### 5. Monitoring RabbingMQ
Anda cukup membuka dashboar rabbitMQ di http://localhost:15672/ 

### 6. Simulasi Fleet Tracking
Cukup mudah jika anda ingin mencoba simulasi geofencing dengan fleet tracking.
step 1. cd cmd/fleettracking
step 2. go run main.go
step 3. pantau log sampai menujukan 'GEOFENCE ALERT'
Selamat anda sudah melakukan simulasi geofencing


## Penutup

Dengan mengikuti langkah di atas, Anda dapat menjalankan aplikasi Fleet Management lengkap dengan database dan messaging queue menggunakan Docker secara mudah dan konsisten berserta simulasi fleet tracking.

Jika ada pertanyaan lebih lanjut, silakan hubungi ke kontak yang sudah ada.

---# fleet-management