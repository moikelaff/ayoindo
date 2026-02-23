# Jawaban Soal 2 & 3 - Technical Test Ayo Indonesia

## Soal 2: Analisis Produk (Website & Aplikasi)

### 1. Review Website (https://ayo.co.id)
Berdasarkan observasi pada website Ayo Indonesia, berikut adalah beberapa hal yang dapat ditingkatkan:
*   **Performa & Waktu Muat (Load Time):** Optimasi ukuran gambar dan *caching* dapat ditingkatkan agar halaman utama (landing page) dapat dimuat lebih cepat, terutama pada koneksi internet yang kurang stabil.
*   **Responsivitas Mobile (Mobile UI/UX):** Meskipun sudah responsif, beberapa elemen UI (seperti tabel klasemen atau jadwal) terkadang masih sulit dibaca di layar *smartphone* berukuran kecil. Penggunaan *card layout* untuk data tabular di versi mobile akan sangat membantu.
*   **SEO & Meta Tags:** Struktur URL dan *meta description* pada halaman spesifik (seperti halaman profil tim atau detail pertandingan) dapat dioptimalkan lebih lanjut untuk meningkatkan visibilitas di mesin pencari (Google).
*   **Aksesibilitas (Accessibility):** Menambahkan atribut *alt-text* pada semua gambar dan memastikan kontras warna teks memenuhi standar WCAG agar website lebih ramah bagi pengguna dengan keterbatasan visual.

### 2. Review Aplikasi Mobile (Playstore/Appstore)
Berdasarkan penggunaan aplikasi Ayo Indonesia, berikut adalah area yang dapat dikembangkan:
*   **Onboarding & Tutorial Pengguna Baru:** Proses *onboarding* bisa dibuat lebih interaktif. Pengguna baru (terutama kapten tim amatir) mungkin membutuhkan panduan singkat (tooltips) tentang cara membuat tim, mengundang pemain, dan mendaftar *sparring*.
*   **Sistem Notifikasi (Push Notifications):** Notifikasi terkadang *delay* atau kurang spesifik. Akan lebih baik jika ada kustomisasi notifikasi (misal: hanya notifikasi jadwal tanding, hasil pertandingan, atau undangan tim).
*   **Fitur Pencarian & Filter (Sparring/Lawan):** Filter pencarian lawan *sparring* bisa dibuat lebih detail, misalnya berdasarkan radius jarak (km) dari lokasi pengguna, level kemampuan tim (beginner/intermediate), atau ketersediaan lapangan.
*   **Optimasi Penggunaan Baterai & Data:** Aplikasi terkadang memakan cukup banyak *resource* saat memuat banyak gambar/logo tim. Implementasi *lazy loading* dan *image compression* di sisi *client* akan sangat membantu performa aplikasi di *device* *low-end*.

---

## Soal 3: Pertanyaan Personal & HR

### 1. Alasan menjadi kandidat terbaik dan kontribusi untuk perusahaan
Saya merupakan kandidat terbaik karena memiliki latar belakang pengalaman yang terstruktur: **2 tahun sebagai Software Quality Assurance (SQA)** dan **1 tahun sebagai Software Engineer**, ditambah keahlian dalam pengembangan *backend* menggunakan Go/Golang, arsitektur REST API, dan manajemen *database* relasional. Kombinasi ini menjadikan saya unik — saya tidak hanya bisa membangun fitur, tetapi juga memahami betul bagaimana sebuah sistem harus diuji, divalidasi, dan dijaga kualitasnya sejak tahap *development*. Kontribusi yang dapat saya berikan adalah membangun sistem *backend* yang *scalable*, aman, dan mudah di-*maintain* sekaligus memastikan kualitasnya sejak awal (*shift-left testing*), sehingga mengurangi *bug* di *production* dan mempercepat siklus pengembangan tim Ayo Indonesia.

### 2. Tiga keunggulan utama
1.  **Keahlian Teknis yang Relevan:** Menguasai Golang (termasuk *framework* Gin dan GORM), PostgreSQL, serta implementasi keamanan API (JWT, *password hashing*), didukung oleh 1 tahun pengalaman langsung di bidang pengembangan perangkat lunak.
2.  **Quality-First Mindset dari Latar Belakang SQA:** Dengan 2 tahun pengalaman sebagai SQA, saya terbiasa berpikir dari sudut pandang *edge case*, validasi data, dan ketahanan sistem. Hal ini membuat kode yang saya tulis sebagai *software engineer* cenderung lebih robust, minim *bug*, dan memiliki *error handling* yang matang.
3.  **Jembatan antara Engineering & Quality:** Kemampuan saya bergerak di dua peran (SQA dan *Software Engineer*) memudahkan kolaborasi lintas tim. Saya dapat membantu mendefinisikan *test case* untuk fitur baru, menulis API yang mudah di-*test*, serta mempercepat proses *QA* karena saya memahami kebutuhan kedua belah pihak.

### 3. Contoh konkret pencapaian/pengalaman relevan
Selama **2 tahun sebagai SQA**, saya bertanggung jawab merancang *test plan*, melakukan *functional testing*, *regression testing*, dan *API testing* (menggunakan Postman dan tools otomasi) untuk aplikasi berbasis web dan mobile. Pengalaman ini memberikan saya pemahaman mendalam tentang bagaimana sebuah sistem bisa gagal, yang kini saya terapkan langsung saat menulis kode.

Kemudian selama **1 tahun sebagai Software Engineer**, saya merancang dan mengembangkan REST API dari nol untuk sistem manajemen data, mengimplementasikan arsitektur *database* yang efisien, sistem autentikasi berbasis JWT, dan fitur *soft-delete* untuk menjaga integritas data historis — persis seperti yang dibutuhkan pada posisi ini.

Gabungan keduanya sangat relevan bagi Ayo Indonesia: saya dapat membangun fitur *backend* yang andal sekaligus memastikan API tersebut sudah memiliki validasi dan *error handling* yang cukup sebelum diserahkan ke tim QA, sehingga memperpendek siklus *review* dan mempercepat *time-to-production*.

### 4. Motivasi melamar dan rencana dampak positif
Motivasi utama saya adalah visi Ayo Indonesia yang mendigitalisasi dan memajukan ekosistem sepak bola amatir di Indonesia. Sebagai seseorang yang menyukai teknologi dan sepak bola, ini adalah kesempatan untuk menggabungkan dua *passion* saya. Selain itu, latar belakang saya di SQA membuat saya menyadari betapa pentingnya kualitas produk terhadap kepercayaan pengguna — sesuatu yang sangat krusial untuk aplikasi dengan basis komunitas seperti Ayo. Jika terpilih, saya berencana memberikan dampak positif dengan membangun fitur *backend* yang tidak hanya fungsional tetapi juga *reliable*, mengoptimalkan *query database*, serta berkolaborasi aktif dengan tim QA untuk memangkas *bug* di *production* berdasarkan pengalaman saya di kedua sisi pengembangan perangkat lunak.

### 5. Langkah-langkah dalam 3 bulan pertama (30-60-90 Days Plan)
*   **Bulan 1 (Learning & Onboarding):** Fokus memahami arsitektur sistem yang ada, *codebase*, standar *coding* perusahaan, dan alur *deployment* (CI/CD). Berbekal pengalaman SQA, saya juga akan langsung mempelajari alur *testing* yang berlaku — *test case*, *bug reporting*, dan *release process* — agar saya bisa menulis kode yang selaras dengan standar QA tim sejak hari pertama.
*   **Bulan 2 (Contributing & Optimizing):** Mulai mengambil tanggung jawab pada fitur yang lebih besar secara mandiri. Saya akan melakukan *code review*, mengidentifikasi *bottleneck* pada API yang ada, dan mengusulkan optimasi (*query database*, *caching*, atau perbaikan validasi). Pengalaman SQA saya akan saya manfaatkan untuk menulis *unit test* dan memastikan setiap fitur yang saya bangun sudah ter-*cover* sebelum masuk ke proses QA formal.
*   **Bulan 3 (Leading & Innovating):** Sepenuhnya mandiri dalam menangani *epic/feature* dari awal hingga *deployment*. Saya juga akan mulai mendokumentasikan API dan proses teknis yang mungkin belum tertulis, serta berperan sebagai jembatan antara tim *engineering* dan QA — mengusulkan praktik *shift-left testing* agar kualitas produk Ayo Indonesia semakin meningkat dari waktu ke waktu.
