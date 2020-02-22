<h1> Pengerjaan deployement ke PaaS di Heroku untuk PHP. </h1>


1. Melakukan cloning aplikasi dummy dari heroku.
![appdummy](/minggu-03/appDummyByHeroku.jpg)

2. Lalu menggunakan git-bash untuk melakukan deploy ke heroku, dan membuka aplikasi yang sudah dideploy.
![openApp](/minggu-03/hasilDeploy.jpg)

3. Ketika melakukan perintah "heroku create" heroku akan otomatis membuatkan repo aplikasi di akun kita dengan nama acak serta remote untuk mengelolanya melalui bash local. Jika ingin melakukan perubahan nama aplikasi dapat melalui setting pada website heroku, akan tetapi ketika melakukan perubahan setelah selesai harus melakukan perubahan nama remote yang ada menjadi nama aplikasi yang baru.

4. Perlu diingat untuk melakukan pengaturan scale menjadi 1 melalui gitbash (dynos untuk versi free adalah 1), jika tidak aplikasi tidak akan bisa dibuka melalui web ataupun bash.

5. Melalui bash local kita dapat menggunakan bash khusus untuk aplikasi yang sudah dideploy ke heroku(salah satu fitur heroku).
![bashApp](/minggu-03/bashAppheroku.jpg)

6. Lalu melakukan modifikasi terhadap file index, melakukan deploy ke heroku, dan menampilkannya.
![pushLocalChange](/minggu-03/hasiltambahcomposer.jpg)

7. Untuk fitur add-ons diabaikan karena untuk mengaksesnya kita perlu untuk memverifikasi akun(langsung masuk ke payment method >_<).

8. Dan yang terakhir karena kita menggunakan versi free dynos yang kita gunakan akan sleep(tidak aktif) terhitung dari 1 jam setelah perintah "heroku create" digunakan. Aplikasi dapat diakses hanya saja akan terdapat penundaan beberapa detik untuk permintaan pertama saat aktif kembali. Permintaan selanjutnya akan berjalan normal. Dyno gratis juga menggunakan kuota bulanan gratis per akun selama kuota tidak habis, aplikasi gratis dapat terus berjalan.