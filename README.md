<h1 data-start="108" data-end="128" class="">📦 Nasabah Service</h1>
<p data-start="130" data-end="325" class="">Sebuah REST API sederhana untuk manajemen nasabah dan rekening berbasis <strong data-start="202" data-end="214">Go Fiber</strong>, <strong data-start="216" data-end="230">PostgreSQL</strong>, dan <strong data-start="236" data-end="246">Docker</strong>. Cocok untuk studi kasus sistem kasir, keuangan, atau backend perbankan mikro.</p>
<hr data-start="327" data-end="330" class="">
<h2 data-start="332" data-end="348" class="">🚀 Tech Stack</h2>
<ul data-start="350" data-end="738">
<li data-start="350" data-end="429" class="">
<p data-start="352" data-end="429" class=""><strong data-start="352" data-end="387"><a data-start="354" data-end="385" rel="noopener" target="_new" class="" href="https://gofiber.io/">Go Fiber</a></strong> – Web framework ringan dan cepat untuk Go</p>
</li>
<li data-start="430" data-end="521" class="">
<p data-start="432" data-end="521" class=""><strong data-start="432" data-end="477"><a data-start="434" data-end="475" rel="noopener" target="_new" class="" href="https://www.postgresql.org/">PostgreSQL</a></strong> – Database relasional open-source yang kuat</p>
</li>
<li data-start="522" data-end="625" class="">
<p data-start="524" data-end="625" class=""><strong data-start="524" data-end="561"><a data-start="526" data-end="559" rel="noopener" target="_new" class="" href="https://www.docker.com/">Docker</a></strong> – Containerization untuk environment development yang konsisten</p>
</li>
<li data-start="626" data-end="738" class="">
<p data-start="628" data-end="738" class=""><strong data-start="628" data-end="682"><a data-start="630" data-end="680" rel="noopener" target="_new" class="">Docker Compose</a></strong> – Orkestrasi service database &amp; app dalam satu perintah</p>
</li>
</ul>
<hr data-start="740" data-end="743" class="">
<h2 data-start="745" data-end="766" class="">📂 Struktur Docker</h2>
<p data-start="768" data-end="845" class="">Project ini menggunakan <code data-start="792" data-end="812">docker-compose.yml</code> untuk menjalankan dua container:</p>
<ol data-start="847" data-end="960">
<li data-start="847" data-end="902" class="">
<p data-start="850" data-end="902" class=""><strong data-start="850" data-end="856">db</strong> – Container PostgreSQL (image: <code data-start="888" data-end="901">postgres:15</code>)</p>
</li>
<li data-start="903" data-end="960" class="">
<p data-start="906" data-end="960" class=""><strong data-start="906" data-end="913">app</strong> – Container aplikasi Go yang menggunakan Fiber</p>
</li>
</ol>
<h3 data-start="962" data-end="976" class="">Dependency</h3>
<p data-start="978" data-end="1146" class="">Container <code data-start="988" data-end="993">app</code> <strong data-start="994" data-end="1011">akan menunggu</strong> hingga container <code data-start="1029" data-end="1033">db</code> berstatus <strong data-start="1044" data-end="1055">healthy</strong> sebelum dijalankan. Ini mencegah aplikasi error saat database belum siap menerima koneksi.</p>
<hr data-start="1148" data-end="1151" class="">
<h2 data-start="1153" data-end="1177" class="">🧬 .env Configuration</h2>
<p data-start="1179" data-end="1273" class="">Berikut isi file <code data-start="1196" data-end="1202">.env</code> yang digunakan untuk konfigurasi koneksi database dalam service <code data-start="1267" data-end="1272">app</code>:</p>
<pre class="overflow-visible!" data-start="1275" data-end="1366"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary"><div class="flex items-center text-token-text-secondary px-4 py-2 text-xs font-sans justify-between h-9 bg-token-sidebar-surface-primary dark:bg-token-main-surface-secondary select-none rounded-t-[5px]">env</div><div class="sticky top-9"><div class="absolute end-0 bottom-0 flex h-9 items-center pe-2"><div class="bg-token-sidebar-surface-primary text-token-text-secondary dark:bg-token-main-surface-secondary flex items-center rounded-sm px-2 font-sans text-xs"><span class="" data-state="closed"><button class="flex gap-1 items-center select-none px-4 py-1" aria-label="Copy"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-xs"><path fill-rule="evenodd" clip-rule="evenodd" d="M7 5C7 3.34315 8.34315 2 10 2H19C20.6569 2 22 3.34315 22 5V14C22 15.6569 20.6569 17 19 17H17V19C17 20.6569 15.6569 22 14 22H5C3.34315 22 2 20.6569 2 19V10C2 8.34315 3.34315 7 5 7H7V5ZM9 7H14C15.6569 7 17 8.34315 17 10V15H19C19.5523 15 20 14.5523 20 14V5C20 4.44772 19.5523 4 19 4H10C9.44772 4 9 4.44772 9 5V7ZM5 9C4.44772 9 4 9.44772 4 10V19C4 19.5523 4.44772 20 5 20H14C14.5523 20 15 19.5523 15 19V10C15 9.44772 14.5523 9 14 9H5Z" fill="currentColor"></path></svg>Copy</button></span><span class="" data-state="closed"><button class="flex items-center gap-1 px-4 py-1 select-none"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-xs"><path d="M2.5 5.5C4.3 5.2 5.2 4 5.5 2.5C5.8 4 6.7 5.2 8.5 5.5C6.7 5.8 5.8 7 5.5 8.5C5.2 7 4.3 5.8 2.5 5.5Z" fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"></path><path d="M5.66282 16.5231L5.18413 19.3952C5.12203 19.7678 5.09098 19.9541 5.14876 20.0888C5.19933 20.2067 5.29328 20.3007 5.41118 20.3512C5.54589 20.409 5.73218 20.378 6.10476 20.3159L8.97693 19.8372C9.72813 19.712 10.1037 19.6494 10.4542 19.521C10.7652 19.407 11.0608 19.2549 11.3343 19.068C11.6425 18.8575 11.9118 18.5882 12.4503 18.0497L20 10.5C21.3807 9.11929 21.3807 6.88071 20 5.5C18.6193 4.11929 16.3807 4.11929 15 5.5L7.45026 13.0497C6.91175 13.5882 6.6425 13.8575 6.43197 14.1657C6.24513 14.4392 6.09299 14.7348 5.97903 15.0458C5.85062 15.3963 5.78802 15.7719 5.66282 16.5231Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path><path d="M14.5 7L18.5 11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path></svg>Edit</button></span></div></div></div><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre! language-env"><span>DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=nasabah_db
DB_PORT=5433
</span></code></div></div></pre>
<ul data-start="1368" data-end="1622">
<li data-start="1368" data-end="1417" class="">
<p data-start="1370" data-end="1417" class=""><code data-start="1370" data-end="1379">DB_HOST</code>: Mengarah ke nama service Docker <code data-start="1413" data-end="1417">db</code></p>
</li>
<li data-start="1418" data-end="1458" class="">
<p data-start="1420" data-end="1458" class=""><code data-start="1420" data-end="1429">DB_USER</code>: Username default PostgreSQL</p>
</li>
<li data-start="1459" data-end="1503" class="">
<p data-start="1461" data-end="1503" class=""><code data-start="1461" data-end="1474">DB_PASSWORD</code>: Password default PostgreSQL</p>
</li>
<li data-start="1504" data-end="1550" class="">
<p data-start="1506" data-end="1550" class=""><code data-start="1506" data-end="1515">DB_NAME</code>: Nama database yang akan digunakan</p>
</li>
<li data-start="1551" data-end="1622" class="">
<p data-start="1553" data-end="1622" class=""><code data-start="1553" data-end="1562">DB_PORT</code>: Port yang di-<em data-start="1577" data-end="1585">expose</em> ke host (5433 → container port 5432)</p>
</li>
</ul>
<hr data-start="1624" data-end="1627" class="">
<h2 data-start="1629" data-end="1653" class="">🐳 docker-compose.yml</h2>
<pre class="overflow-visible!" data-start="1655" data-end="2378"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary"><div class="flex items-center text-token-text-secondary px-4 py-2 text-xs font-sans justify-between h-9 bg-token-sidebar-surface-primary dark:bg-token-main-surface-secondary select-none rounded-t-[5px]">yaml</div><div class="sticky top-9"><div class="absolute end-0 bottom-0 flex h-9 items-center pe-2"><div class="bg-token-sidebar-surface-primary text-token-text-secondary dark:bg-token-main-surface-secondary flex items-center rounded-sm px-2 font-sans text-xs"><span class="" data-state="closed"><button class="flex gap-1 items-center select-none px-4 py-1" aria-label="Copy"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-xs"><path fill-rule="evenodd" clip-rule="evenodd" d="M7 5C7 3.34315 8.34315 2 10 2H19C20.6569 2 22 3.34315 22 5V14C22 15.6569 20.6569 17 19 17H17V19C17 20.6569 15.6569 22 14 22H5C3.34315 22 2 20.6569 2 19V10C2 8.34315 3.34315 7 5 7H7V5ZM9 7H14C15.6569 7 17 8.34315 17 10V15H19C19.5523 15 20 14.5523 20 14V5C20 4.44772 19.5523 4 19 4H10C9.44772 4 9 4.44772 9 5V7ZM5 9C4.44772 9 4 9.44772 4 10V19C4 19.5523 4.44772 20 5 20H14C14.5523 20 15 19.5523 15 19V10C15 9.44772 14.5523 9 14 9H5Z" fill="currentColor"></path></svg>Copy</button></span><span class="" data-state="closed"><button class="flex items-center gap-1 px-4 py-1 select-none"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-xs"><path d="M2.5 5.5C4.3 5.2 5.2 4 5.5 2.5C5.8 4 6.7 5.2 8.5 5.5C6.7 5.8 5.8 7 5.5 8.5C5.2 7 4.3 5.8 2.5 5.5Z" fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"></path><path d="M5.66282 16.5231L5.18413 19.3952C5.12203 19.7678 5.09098 19.9541 5.14876 20.0888C5.19933 20.2067 5.29328 20.3007 5.41118 20.3512C5.54589 20.409 5.73218 20.378 6.10476 20.3159L8.97693 19.8372C9.72813 19.712 10.1037 19.6494 10.4542 19.521C10.7652 19.407 11.0608 19.2549 11.3343 19.068C11.6425 18.8575 11.9118 18.5882 12.4503 18.0497L20 10.5C21.3807 9.11929 21.3807 6.88071 20 5.5C18.6193 4.11929 16.3807 4.11929 15 5.5L7.45026 13.0497C6.91175 13.5882 6.6425 13.8575 6.43197 14.1657C6.24513 14.4392 6.09299 14.7348 5.97903 15.0458C5.85062 15.3963 5.78802 15.7719 5.66282 16.5231Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path><path d="M14.5 7L18.5 11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path></svg>Edit</button></span></div></div></div><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre! language-yaml"><span><span><span class="hljs-attr">version:</span></span><span> </span><span><span class="hljs-string">"3.9"</span></span><span>

</span><span><span class="hljs-attr">services:</span></span><span>
  </span><span><span class="hljs-attr">db:</span></span><span>
    </span><span><span class="hljs-attr">image:</span></span><span> </span><span><span class="hljs-string">postgres:15</span></span><span>
    </span><span><span class="hljs-attr">environment:</span></span><span>
      </span><span><span class="hljs-attr">POSTGRES_USER:</span></span><span> </span><span><span class="hljs-string">postgres</span></span><span>
      </span><span><span class="hljs-attr">POSTGRES_PASSWORD:</span></span><span> </span><span><span class="hljs-string">postgres</span></span><span>
      </span><span><span class="hljs-attr">POSTGRES_DB:</span></span><span> </span><span><span class="hljs-string">nasabah_db</span></span><span>
    </span><span><span class="hljs-attr">volumes:</span></span><span>
      </span><span><span class="hljs-bullet">-</span></span><span> </span><span><span class="hljs-string">dbdata:/var/lib/postgresql/data</span></span><span>
    </span><span><span class="hljs-attr">ports:</span></span><span>
      </span><span><span class="hljs-bullet">-</span></span><span> </span><span><span class="hljs-string">"5433:5432"</span></span><span>
    </span><span><span class="hljs-attr">healthcheck:</span></span><span>
      </span><span><span class="hljs-attr">test:</span></span><span> [</span><span><span class="hljs-string">"CMD-SHELL"</span></span><span>, </span><span><span class="hljs-string">"pg_isready -U postgres"</span></span><span>]
      </span><span><span class="hljs-attr">interval:</span></span><span> </span><span><span class="hljs-string">5s</span></span><span>
      </span><span><span class="hljs-attr">timeout:</span></span><span> </span><span><span class="hljs-string">5s</span></span><span>
      </span><span><span class="hljs-attr">retries:</span></span><span> </span><span><span class="hljs-number">5</span></span><span>

  </span><span><span class="hljs-attr">app:</span></span><span>
    </span><span><span class="hljs-attr">build:</span></span><span> </span><span><span class="hljs-string">.</span></span><span>
    </span><span><span class="hljs-attr">depends_on:</span></span><span>
      </span><span><span class="hljs-attr">db:</span></span><span>
        </span><span><span class="hljs-attr">condition:</span></span><span> </span><span><span class="hljs-string">service_healthy</span></span><span>
    </span><span><span class="hljs-attr">ports:</span></span><span>
      </span><span><span class="hljs-bullet">-</span></span><span> </span><span><span class="hljs-string">"3000:3000"</span></span><span>
    </span><span><span class="hljs-attr">environment:</span></span><span>
      </span><span><span class="hljs-bullet">-</span></span><span> </span><span><span class="hljs-string">DB_HOST=db</span></span><span>
      </span><span><span class="hljs-bullet">-</span></span><span> </span><span><span class="hljs-string">DB_USER=postgres</span></span><span>
      </span><span><span class="hljs-bullet">-</span></span><span> </span><span><span class="hljs-string">DB_PASSWORD=postgres</span></span><span>
      </span><span><span class="hljs-bullet">-</span></span><span> </span><span><span class="hljs-string">DB_NAME=nasabah_db</span></span><span>
      </span><span><span class="hljs-bullet">-</span></span><span> </span><span><span class="hljs-string">DB_PORT=5432</span></span><span>
    </span><span><span class="hljs-attr">command:</span></span><span> [</span><span><span class="hljs-string">"./service-account"</span></span><span>, </span><span><span class="hljs-string">"--host=0.0.0.0"</span></span><span>, </span><span><span class="hljs-string">"--port=3000"</span></span><span>]

</span><span><span class="hljs-attr">volumes:</span></span><span>
  </span><span><span class="hljs-attr">dbdata:</span></span><span>
</span></span></code></div></div></pre>
<h3 data-start="2380" data-end="2402" class="">Penjelasan Penting</h3>
<ul data-start="2404" data-end="2622">
<li data-start="2404" data-end="2479" class="">
<p data-start="2406" data-end="2479" class=""><code data-start="2406" data-end="2419">healthcheck</code>: Mengecek apakah PostgreSQL sudah siap melalui <code data-start="2467" data-end="2479">pg_isready</code></p>
</li>
<li data-start="2480" data-end="2556" class="">
<p data-start="2482" data-end="2556" class=""><code data-start="2482" data-end="2494">depends_on</code>: Menunggu hingga service <code data-start="2520" data-end="2524">db</code> "Healthy" sebelum menjalankan <code data-start="2551" data-end="2556">app</code></p>
</li>
<li data-start="2557" data-end="2622" class="">
<p data-start="2559" data-end="2622" class=""><code data-start="2559" data-end="2568">command</code>: Menjalankan aplikasi Go dengan argumen host dan port</p>
</li>
</ul>
<hr data-start="2624" data-end="2627" class="">
<h2 data-start="2629" data-end="2654" class="">🧪 Menjalankan Project</h2>
<pre class="overflow-visible!" data-start="2656" data-end="2718"><div class="contain-inline-size rounded-md border-[0.5px] border-token-border-medium relative bg-token-sidebar-surface-primary"><div class="sticky top-9"><div class="absolute end-0 bottom-0 flex h-9 items-center pe-2"><div class="bg-token-sidebar-surface-primary text-token-text-secondary dark:bg-token-main-surface-secondary flex items-center rounded-sm px-2 font-sans text-xs"><span class="" data-state="closed"><button class="flex items-center gap-1 px-4 py-1 select-none"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-xs"><path d="M2.5 5.5C4.3 5.2 5.2 4 5.5 2.5C5.8 4 6.7 5.2 8.5 5.5C6.7 5.8 5.8 7 5.5 8.5C5.2 7 4.3 5.8 2.5 5.5Z" fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"></path><path d="M5.66282 16.5231L5.18413 19.3952C5.12203 19.7678 5.09098 19.9541 5.14876 20.0888C5.19933 20.2067 5.29328 20.3007 5.41118 20.3512C5.54589 20.409 5.73218 20.378 6.10476 20.3159L8.97693 19.8372C9.72813 19.712 10.1037 19.6494 10.4542 19.521C10.7652 19.407 11.0608 19.2549 11.3343 19.068C11.6425 18.8575 11.9118 18.5882 12.4503 18.0497L20 10.5C21.3807 9.11929 21.3807 6.88071 20 5.5C18.6193 4.11929 16.3807 4.11929 15 5.5L7.45026 13.0497C6.91175 13.5882 6.6425 13.8575 6.43197 14.1657C6.24513 14.4392 6.09299 14.7348 5.97903 15.0458C5.85062 15.3963 5.78802 15.7719 5.66282 16.5231Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path><path d="M14.5 7L18.5 11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path></svg>Edit</button></span></div></div></div><div class="overflow-y-auto p-4" dir="ltr"><code class="whitespace-pre! language-bash"><span><span><span class="hljs-comment"># Jalankan semua service</span></span><span>
docker-compose up --build
</span></span></code></div></div></pre>
<p data-start="2720" data-end="2785" class="">Akses aplikasi di: <a data-start="2739" data-end="2785" rel="noopener" target="_new" class="" href="http://localhost:3000">http://localhost:3000</a></p>
