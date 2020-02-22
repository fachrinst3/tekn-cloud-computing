<h1> Pembahasan Praktikum TCC pertemuan-3 </h1>
 
Sebelum menggunakan Heroku terlebih dahulu kita harus membuat akun, dan mengkonfirmasinya melalui email yang kita gunakan untuk masuk kedalam akun yang sudah dibuat.

![confirmationEmail](/img/image-01..jpg)

Setelah selesai masuk ke link getting started on heroku with Python. Ketika kita akan melakukan pendownload-an CLI heroku agar dapat mengikuti proses berikutnya. Ketika selesai melakukan download kita akan menggunakan perintah "Heroku login" untuk mengkoneksikan bash yang kita gunakan kepada akun heroku yang sudah dibuat sebelumnya.
![logintoCLI](/img/image-02.jpg)

Lalu melakukan clone terhadap aplikasi dummy yang disediakan heroku untuk melakukan proses pengenalan penggunaan heroku ini.
![cloningdummyAPP](/img/image-03.jpg)

Setelah itu kita menggunakan perintah "heroku create" untuk membuat repo aplikasi heroku didalam akun kita, ketika melakukannya kita akan diberikan nama acak untuk (repo) aplikasinya serta dibuatkan remote otomatis sesuai dengan nama yang diberikan heroku, Perlu diingat jika ingin mengubah nama aplikasi kita juga harus mengkonfigurasikan remote yang sudah ada menjadi nama aplikasi yang kita ubah.

Selanjutnya menggunakan perintah untuk melakukan deploy terhadap aplikasi yang sudah kita clone sebelumnya dengan menggunakan perintah yang mirip dengan cara untuk push ke github. Untuk menambahkan perubahan file kita menggunakan "git add .", lalu menggunakan "git commit -m" untuk memberikan keterangan perubahan yang dilakukan(dalam proses pengenalan heroku kita menggunakan keterangan "DEMO"), dan menggunakan perintah "git push heroku master" untuk melakukan deploy aplikasi lokal ke server heroku. Ketika sudah selesai dalam proses deploy kita dapat membuka aplikasi tersebut menggunakan perintah "heroku open" didalam bash yang akan otomatis membuka aplikasi kita didalam web.

Jika tidak ingin langsung mendeploy perubahan yang kita lakukan terhadap aplikasi kita dapat menggunakan localhost menggunakan perintah "heroku local web -f Procfile.windows" (sebelumnya kita sudah melakukan setting agar procfile menyediakan localhost:5000 untuk aplikasi kita), dan langsung membuka localhost:5000 didalam web kita untuk melihat aplikasinya versi localhost
![commandLocalApp](/img/image-05.jpg)

Lalu untuk langkah selanjutnya kita diminta untuk melakukan modifikasi terhadap file views.py dimana kita akan menambahkan module requests, dan function request, lalu melakukan save
![teapotMAN](/img/image-06.jpg) dan menggunakan perintah heroku untuk membuka aplikasi kita didalam localhost, hasilnya menjadi seperti gambar dibawah.
![teapotMANversiLocalhost](/img/image-07.jpg)
Setelah hasil sesuai dengan perubahan yang diinginkan kita dapat melakukan deploy ke repo aplikasi heroku di akun kita menggunakan perintah yang sama ketika melakukan deploy aplikasi ini pertama kali. Ketika selelasi melakukan deploy aplikasinya kita dapat membukanya melalui bash menggunakan perintah "heroku open", dan web browser akan membuka aplikasi kita yang versi host diheroku. 
![appWebVersion](/img/image-08.jpg)

<h3>KESIMPULAN PRAKTIKUM</h3>

Dari keseluruhan praktik dapat saya simpulkan cara penggunaan heroku sebagai berikut :


1. Buat repo untuk aplikasi yang ingin dibuat didalam heroku.
2. Lakukan push terhadap aplikasi kita ke repo heroku.
3. Jika ingin melakukan perubahan cukup lakukan di folder lokal(pc) lalu preview di localhost.
4. Ketika perubahan yang dilakukan sudah sesuai dengan keinginan, dan tidak error didalam localhost kita dapat mendeploynya, dan selesai.

<h3>TAMBAHAN PROS & CON menurut saya</h3>
Pros :

1. Sangat mempermudah developer untuk melakukan pengolahan, dan storing aplikasi terutama yang sudah terbiasa menggunakan layanan cloud(tidak perlu takut data di pc hilang).

Cons :

1. Karena berbayar kita hanya akan dapat menggunakan service host selama 30 menit sejak repo aplikasi di create(heroku create), hal ini menyebabkan ketika aplikasi dibuka di host heroku pada waktu setelah lebih dari 30 menit perlu waktu untuk berjalan normal(heroku menyebutnya dynos akan tidur ketika lebih dari 30 menit, dan jika bangun perlu waktu untuk menerima perintah setelah itu berjalan normal).

2. Fitur add-ons tidak dapat digunakan didalam akun versi pro (padahal untuk keerluan pembelajaran :( ).

