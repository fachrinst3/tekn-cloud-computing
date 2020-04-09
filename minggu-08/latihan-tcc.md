<h1> Latihan TCC-minggu 8 </h1>

1. Pertama saya melakukan pembuatan folder project.

![step1](/step1.jpg)

2. Setelah itu melakukan pembuatan sebuah file python yang isinya menggunakan container redis di jaringan aplikasi.

![step2](/step2.jpg)

3. Sesudah membuat file python, membuat file txt dengan nama requirements, dan membuat file DockerFile yaang menggunakan image python.

![step3](/step3.jpg)

4. Lalu membuat file yaml yang bersikan mengikat wadah dan mesin host ke port 5000. Contoh layanan ini menggunakan port default untuk server web Flask, 5000.

![step4](/step4.jpg)

5. Setelah selesai dengan membuat folder yang dibutuhkan untuk project, kita melakukan run project dengan menggunakan perintah "docker-compose up".

![step5](/dockercomposeup.jpg)

6. Setelah berhasil melakukan run kita membiarkan terminal tetap terbuka, dan membuka url localhost:5000(saya menuliskannya menggunakan ip machine) didalam browser untuk melihat program python yang dijalankan di container.

![step6](/hasilcomposeup.jpg)

7. Di browser ketika kita melakukan refresh terdapat perubahan pada angkanya (bertambah)

![step7](/hasilcomposeup2.jpg)

8. Lalu untuk membuktikan bahwa image sudah ada di dalam pc(direktori local) kita membuka satu terminal lagi dan menggunakan perintah "docker image ls" untuk melihat daftar image yang ada di pc(direktori local).

![step8](/lihatimage.jpg)

9. Selanjutnya kita melakukan perubahan terhadap file yml sebelumnya dengan menambahkan bind mount.

![step9](/edityml.jpg)

10. Selesai dengan penambahan bind mount, kita melakukan build, dam run ulang terhadap project kita menggunakan perintah "docker-compose up".

![step10](/runkedua.jpg)

11. Sama seperti sebelumnya jika perintah sukses dijalankan kita membiarkan terminal tersebut tetap berjalan, dan membuka url localhost:5000 pada browser.

![step11](/hasilcomposeup3.jpg)

12. Karena kode aplikasi sudah kita mount ke container menggunakan volume(perubahan file yml sebelumnya) kita sudah bisa memodifikasi kode file python, dan melihat perubahannya secara instant tanpa harus melakukan build, dan run ulang.

![step12](/editpython.jpg)

13. Hasil dari perubahan kode didalam file python.

![step13](/hasilcomposeuprealtime.jpg)

14. Selanjutnya kita melakukan eksperimen terhadap perintah docker lainnya, kita menggunakan beberapa perintah yang dimana intinya kita melihat container yang sedang berjalan di background.

![step14](/step8eksperimen.jpg)