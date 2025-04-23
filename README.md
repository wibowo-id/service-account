<h1>📦 Service Account</h1>
<p>Sebuah REST API sederhana untuk manajemen nasabah dan rekening berbasis <strong>Go Fiber</strong>, <strong>PostgreSQL</strong>, dan <strong>Docker</strong>.</p>
<hr>
<h2>🚀 Tech Stack</h2>
<ul>
<li>
<p><strong><a rel="noopener" target="_new" href="https://gofiber.io/">Go Fiber</a></strong> – Web framework ringan dan cepat untuk Go</p>
</li>
<li>
<p><strong><a rel="noopener" target="_new" href="https://www.postgresql.org/">PostgreSQL</a></strong> – Database relasional open-source yang kuat</p>
</li>
<li>
<p><strong><a rel="noopener" target="_new" href="https://www.docker.com/">Docker</a></strong> – Containerization untuk environment development yang konsisten</p>
</li>
<li>
<p><strong><a rel="noopener" target="_new">Docker Compose</a></strong> – Orkestrasi service database &amp; app dalam satu perintah</p>
</li>
</ul>
<hr>
<h2>📂 Struktur Folder</h2>
  <pre><code>
├── config/
│   └── config.go
├── controller/
│   └── init.go
│   └── rekening_controller.go
├── helper/
│   └── generator.go
│   └── helper.go
├── model/
│   └── mutasi.go
│   └── nasabah.go
│   └── rekening.go
├── route/
│   └── route.go
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yml
├── docker-compose.yml.example
├── Dockerfile
├── go.mod
├── go.sum
└── main.go
</code></pre>
<hr>
<h2>🐳 Struktur Docker</h2>
<p>Project ini menggunakan <code>docker-compose.yml</code> untuk menjalankan dua container:</p>
<ol>
<li>
<p><strong>db</strong> – Container PostgreSQL (image: <code>postgres:15</code>)</p>
  <ul>Menggunakan image postgres:15</ul>
<ul>Database: nasabah_db</ul>
<ul>Port Host: 5433 → Port Container: 5432</ul>
<ul>Healthcheck: Memastikan database siap sebelum service app dijalankan</ul>
</li>
<li>
<p><strong>app</strong> – Container aplikasi Go yang menggunakan Fiber</p>
  <ul>Dibuild dari Dockerfile</ul>
<ul>Menjalankan binary service-account</ul>
<ul>Port: 3000</ul>
<ul>Depends on db: menunggu database siap (service_healthy) sebelum start</ul>
</li>
</ol>
<h3>Dependency</h3>
<p>Container <code>app</code> <strong>akan menunggu</strong> hingga container <code>db</code> berstatus <strong>healthy</strong> sebelum dijalankan. Ini mencegah aplikasi error saat database belum siap menerima koneksi.</p>
<hr>
<h2>🧬 .env Configuration</h2>
<p>Berikut isi file <code>.env</code> yang digunakan untuk konfigurasi koneksi database dalam service <code>app</code>:</p>
<pre class="overflow-visible!"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary"><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre! language-env"><span>DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=nasabah_db
DB_PORT=5433
</span></code></div></div></pre>
<ul>
<li>
<p><code>DB_HOST</code>: Mengarah ke nama service Docker <code>db</code></p>
</li>
<li>
<p><code>DB_USER</code>: Username default PostgreSQL</p>
</li>
<li>
<p><code>DB_PASSWORD</code>: Password default PostgreSQL</p>
</li>
<li>
<p><code>DB_NAME</code>: Nama database yang akan digunakan</p>
</li>
<li>
<p><code>DB_PORT</code>: Port yang di-<em>expose</em> ke host (5433 → container port 5432)</p>
</li>
</ul>
<hr>
<h2>🐳 docker-compose.yml</h2>
<pre class="overflow-visible!"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary"><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre! language-yaml"><span><span><span>version:</span> <span>"3.9"</span>

<span>services:</span>
<span>db:</span>
<span>image:</span> <span>postgres:15</span>
<span>environment:</span>
<span>POSTGRES_USER:</span> <span>postgres</span>
<span>POSTGRES_PASSWORD:</span> <span>postgres</span>
<span>POSTGRES_DB:</span> <span>nasabah_db</span>
<span>volumes:</span>
<span class="hljs-bullet">-</span> <span>dbdata:/var/lib/postgresql/data</span>
<span>ports:</span>
<span class="hljs-bullet">-</span> <span>"5433:5432"</span>
<span>healthcheck:</span>
<span>test:</span> [<span>"CMD-SHELL"</span>, <span>"pg_isready -U postgres"</span>]
<span>interval:</span> <span>5s</span>
<span>timeout:</span> <span>5s</span>
<span>retries:</span> <span class="hljs-number">5</span>

<span>app:</span>
<span>build:</span> <span>.</span>
<span>depends_on:</span>
<span>db:</span>
  <span>condition:</span> <span>service_healthy</span>
<span>ports:</span>
<span class="hljs-bullet">-</span> <span>"3000:3000"</span>
<span>environment:</span>
<span class="hljs-bullet">-</span> <span>DB_HOST=db</span>
<span class="hljs-bullet">-</span> <span>DB_USER=postgres</span>
<span class="hljs-bullet">-</span> <span>DB_PASSWORD=postgres</span>
<span class="hljs-bullet">-</span> <span>DB_NAME=nasabah_db</span>
<span class="hljs-bullet">-</span> <span>DB_PORT=5432</span>
<span>command:</span> [<span>"./service-account"</span>, <span>"--host=0.0.0.0"</span>, <span>"--port=3000"</span>]

<span>volumes:</span>
<span>dbdata:</span>
</span></span></code></div></div></pre>
<h3>Penjelasan</h3>
<ul>
<li>
<p><code>healthcheck</code>: Mengecek apakah PostgreSQL sudah siap melalui <code>pg_isready</code></p>
</li>
<li>
<p><code>depends_on</code>: Menunggu hingga service <code>db</code> "Healthy" sebelum menjalankan <code>app</code></p>
</li>
<li>
<p><code>command</code>: Menjalankan aplikasi Go dengan argumen host dan port</p>
</li>
</ul>
<hr>
<h2>🐳 Dockerfile</h2>
<pre class="overflow-visible!"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary"><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre! language-Dockerfile"><span>FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o service-account .

EXPOSE 3000

CMD ["./service-account"]
</span></code></div></div></pre>

<h3>Penjelasan:</h3>

<ul>
<li>
<p><code>golang:1.23-alpine</code>: base image ringan</p>
</li>
<li>
<p><code>WORKDIR /app</code>: direktori kerja container</p>
</li>
<li>
<p>Menyalin <code>go.mod</code> &amp; <code>go.sum</code> lalu download dependency</p>
</li>
<li>
<p>Menyalin semua source code dan build binary <code>service-account</code></p>
</li>
<li>
<p><code>EXPOSE 3000</code>: dokumentasi port yang dipakai</p>
</li>
<li>
<p><code>CMD</code>: menjalankan binary saat container start</p>
</li>
</ul>
<hr>
<h2>🧪 Menjalankan Project</h2>
<pre class="overflow-visible!"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary"><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre! language-bash"><span><span><span class="hljs-comment"># Jalankan semua service</span>
docker-compose up --build
</span></span></code></div></div></pre>
<p>Akses aplikasi di: <a rel="noopener" target="_new" href="http://localhost:3000">http://localhost:3000</a></p>
