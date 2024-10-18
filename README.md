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

Create Table audit_log

```
CREATE TABLE `audit_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `module` varchar(255) NOT NULL,
  `action_type` varchar(50) NOT NULL,
  `search_key` varchar(255) NOT NULL,
  `before_data` text NOT NULL,
  `after_data` text NOT NULL,
  `action_by` varchar(100) NOT NULL,
  `action_time` datetime NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5737 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
```

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
