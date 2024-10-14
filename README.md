# Audit Log (API)
 
App Audit Log API

Aplikasi ini adalah sebuah API untuk mencatat log audit dalam aplikasi. API ini menyediakan berbagai endpoint untuk mengelola log audit dan memberikan cara yang terstruktur untuk melacak aktivitas dalam aplikasi.
Fitur
- Pencatatan Log Audit: Mencatat berbagai jenis aktivitas seperti pembuatan, pembaruan, dan penghapusan data.
- Endpoint CRUD: Mendukung operasi CRUD (Create, Read, Update, Delete) untuk log audit.
- Keamanan: Menggunakan mekanisme autentikasi dan otorisasi untuk melindungi endpoint.

Instalasi
Clone Repository:
    
    git clone https://github.com/irfandiricon/app-audit-log-api.git
    cd app-audit-log-api
    go mod tidy

Konfigurasi

Buat file .env di root direktori dan tambahkan konfigurasi berikut:

    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=username
    DB_PASS=password
    DB_NAME=audit_log

Menjalankan dalam Mode Pengembangan:

    go run main.go

Struktur Project

Proyek ini memiliki struktur dasar sebagai berikut:

```
├───src
│   ├───common
│   │   ├───http
│   │   └───utils
│   │       ├───httpresponse
│   │       └───middlewares
│   ├───domain
│   │   ├───models
│   │   └───repositories
│   ├───infrastructure
│   │   ├───api
│   │   ├───cache
│   │   ├───database
│   │   │   └───mysql
│   │   └───rabbitmq
│   ├───interfaces
│   │   ├───kafka
│   │   ├───mq
│   │   ├───rest
│   │   │   └───controller
│   │   └───rpc
│   ├───payload
│   │   ├───request
│   │   └───response
│   └───usecases
└───tmp
```
