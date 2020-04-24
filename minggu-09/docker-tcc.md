<h1> Latihan TCC-minggu 9 </h1>

<h2> A. Run a single task in an Alpine Linux container </h2>

1. Melakukan Clone tweet app dari repo di Github.

![step1](g1.jpg)

2. Melakukan run container linux alpine secara single task

![step2](g2.jpg)

3. Melakukan cek terhadap container yang dijalankan sebelumnya dengan nama "hostname".

![step3](g3.jpg)

<h2> B. Run an interactive Ubuntu container</h2>

1. Menjalankan container interaktif ubuntu dan menggunakan perintah dasar linux didalamnya. Karena menjalankannya menggunakan perintah --rm maka ketika dicek pada list container ubuntu tidak akan terdaftar (langsung dihapus ketika diberhentikan).

![step4](g4.jpg)

2. Lalu kita melakukan pengecekan versi dari virtual machinenya.

![step5](g5.jpg)

<h2> C. Run a background MySQL container</h2>

1. Penggunaan perintah untuk menjalankan container Mysql di background.

![step6](g6.jpg)

2. Melakukan pengecekan terhadap container yang berjalan di background menggunakan perintah docker untuk melihat list container.

![step7](g7.jpg)

3. Menggunakan perintah docker untuk melihat proses yang berjalan di container.

![step8](g8.jpg)

4. Menggunakan perintah docker untuk membuka shell didalam container yang berjalan dibackground untuk melakukan cek versi mysql-nya.

![step9](g9.jpg)

<h2> D. Package and run a custom app using Docker</h2>

1. Masuk ke direktori linux tweet-app yang sebelumnya di clone, dan membuat DockerFile yang menggunakan image nginx.

![step10](g10.jpg)

2. Melakukan export environment variabel yang isinya id dockerhub kita.

![step11](g11.jpg)

3. Melakukan build image yang sebelumnya sudah disiapkan.

![step12](g12.jpg)

4. Melakukan run terhadap image, dan melihat hasilnya dibrowser.

![step13](g13.jpg)

![step14](g14.jpg)

5. Memberhentikan container-nya menggunakan perintah force, dan melakukan cek apakah berhasil.

![step15](g15.jpg)

![step16](g16.jpg)

<h2> E. Modify a running website</h2>

1. Pertama kita menjalankan container pada task sebelumnya (tweet app), akan tetapi kali ini container di mount untuk proses terhadap task ini.

![step17](g17.jpg)

![step18](g14.jpg)

2. Lalu kita menambahlan file html didalam direktori tweet-app, dan melakukan cek terhadap perubahan di website tweet-app yang berjalan.

![step19](g18.jpg)

![step20](g19.jpg)

3. Meskipun kita telah memodifikasi index.html sistem file lokal dan melihatnya berubah didalam container yang sedang berjalan, kita belum benar-benar mengubah image Docker dicontainer tersebut (karena menggunakan mount), maka kita melakukan pemberhentian container, dan melakukan start ulang, dan hasilnya image tidak berubah sama sekali.

![step21](g20.jpg)

![step22](g21.jpg)

4. Kita akan melakukan modifikasi(update) terhadap image tweet-app dan menggunakan tag 2.0 didalam proses build ulangnya. Dan melihat hasilnya di list container.

![step23](g22.jpg)

![step24](g23.jpg)

5. Lalu kita melakukan run (yang didalam direktorinya kita sudah menggunakan file html yang di update sebelumnya), dan melihat hasilnya.

![step25](g24.jpg)

![step26](g25.jpg)

6. Lalu untuk perbandingan terhadap image baru, dan lama kita akan menjalankan container lain dengan port berbeda, yaitu 80-8080 (karena yang baru sudah berjalan di port 80-80), dan nama containernya adalah old linux tweet app. Lalu melihat image yang lama di port tersebut.

![step27](g26.jpg)

![step28](g27.jpg)

7. Lalu kita akan melakukan push terhadap 2 image yang sudah kita buat didalam proses tutorial ini, yaitu tweet app 1.0, dan tweet app 2.0 .

![step29](g28.jpg)

8. Sebelum push kita harus melakukan login, karena sebelumnya sudah login saya langsung menadapatkan feedback ini.

![step30](g30.jpg)

9. Lalu melakukan push ke repo dockerhub di akun saya.

![step31](g31.jpg)

10. Dan berikut adalah hasil dari push ke repo dockerhub saya. 

![step32](g32.jpg)