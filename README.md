# Golang Observability Stack

Project ini adalah demonstrasi implementasi Full Observability Stack (Metrics, Logs, Traces) menggunakan Go, Prometheus, Loki, Tempo, Alloy, dan Grafana.

## Arsitektur
- **Aplikasi (Go):** Menghasilkan metrics (Prometheus), logs (JSON format), dan traces (OpenTelemetry).
- **Grafana Alloy:** Kolektor yang mengambil logs dari container Docker dan mengirimkannya ke Loki.
- **Prometheus:** Menyimpan metrics dari aplikasi.
- **Loki:** Menyimpan logs dari aplikasi.
- **Tempo:** Menyimpan distributed traces (OTLP).
- **Grafana:** Visualisasi dan korelasi antar data (Metrics -> Logs -> Traces).

## Prasyarat
- Docker dan Docker Compose installed.

## Cara Menjalankan
1. Clone repository ini.
2. Jalankan stack menggunakan Docker Compose:
   ```bash
   docker compose up -d --build
   ```
3. Tunggu hingga semua container berstatus `healthy` (terutama `go-app`).

## Akses Layanan
- **Grafana:** [http://localhost:3000](http://localhost:3000) (Login otomatis sebagai Admin)
- **Prometheus:** [http://localhost:9090](http://localhost:9090)
- **App API:** [http://localhost:8080](http://localhost:8080)

---

## Panduan Testing & Verifikasi

### 1. Generate Traffic (Data)
Lakukan beberapa request ke aplikasi untuk menghasilkan data observability:
```bash
# Klik link ini beberapa kali atau gunakan curl
curl http://localhost:8080/
curl http://localhost:8080/health
```

### 2. Verifikasi di Grafana
Buka Grafana ([http://localhost:3000](http://localhost:3000)) dan navigasi ke Dashboard **"HOME - Observability"**.

#### A. Verifikasi Metrics
- Cek panel **"HTTP Requests Total"**. Anda akan melihat grafik jumlah request berdasarkan path (`/` atau `/health`).
- Cek panel **"HTTP Request Duration"** untuk melihat latensi aplikasi.

#### B. Verifikasi Logs (Loki)
- Cek panel **"Application Logs"**. Anda akan melihat log dalam format JSON.
- Perhatikan field `trace_id` di dalam log. Klik pada baris log, dan jika konfigurasi benar, akan muncul tombol/link **"TraceID"** (Tempo) di sebelah ID trace tersebut.

#### C. Verifikasi Traces (Tempo)
- Klik link **"TraceID"** dari panel log, atau buka menu **Explore** -> Pilih datasource **Tempo**.
- Cari traces yang baru saja dihasilkan. Anda akan melihat timeline eksekusi fungsi `home-request` beserta atribut seperti `http.method` dan `user_agent`.

### 3. Verifikasi Korelasi (The "Power" of this stack)
1. Pergi ke menu **Explore**.
2. Pilih **Loki** dan jalankan query `{job="go-app"}`.
3. Cari log dengan pesan `"request completed"`.
4. Klik log tersebut, lalu klik link **Tempo** di field `trace_id`.
5. Grafana akan membuka split view yang menampilkan detail tracing dari log tersebut secara otomatis.

## Load Testing dengan Hey

Untuk melihat bagaimana stack ini bekerja di bawah beban (load), Anda bisa menggunakan tool [hey](https://github.com/rakyll/hey).

1. **Install hey** (jika belum ada):
   ```bash
   # Linux/macOS
   go install github.com/rakyll/hey@latest
   # Atau menggunakan brew
   brew install hey
   ```

2. **Jalankan Load Test**:
   Anda bisa menggunakan script yang sudah disediakan atau menjalankan perintah langsung:
   ```bash
   ./load-test.sh
   # ATAU
   hey -z 10s -c 10 http://localhost:8080/
   ```

3. **Amati Hasilnya**:
   - **Grafana Dashboard:** Lihat lonjakan pada grafik "HTTP Requests Total".
   - **Tempo:** Lihat distribusi latensi yang lebih bervariasi.
   - **Loki:** Lihat ribuan log JSON yang terstruktur dengan korelasi `trace_id`.

## Troubleshooting
- Jika dashboard tidak muncul: Pastikan folder `grafana/provisioning` memiliki izin baca.
- Jika logs tidak muncul: Pastikan Alloy memiliki akses ke `/var/run/docker.sock`.
- Jika traces tidak muncul: Cek apakah container `tempo` sudah berjalan normal dan port `4317` dapat diakses oleh container `app`.
