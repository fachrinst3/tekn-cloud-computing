<h1> Latihan TCC Minggu-06 </h1>

<h3>A. Membuat program Go untuk koneksi dan membaca data dari MySQL </h3>

1. Sebelum melakukan pembuatan program kita diwajibkan untuk menginstall Mysql Community Edition, compiler Go, dan Mysql driver untuk Go.

2. Setelah proses instalasi selesai kita langsung membuat sebuah database yang akan digunakan untuk melakukan koneksi, dan pembacaan data menggunakann Go.
<div style="text-align:center">

![databaseMysql](/minggu-06/mysql.jpg)

</div>

3. Lalu membuat program menggunakan bahasa Go untuk melakukan koneksi, dan membaca data yang sudah dibuat sebelumnya di database mysql. [Program Go Mysql](/minggu-06/featuring-mysql.go)

4. Hasil dari run program Go yang dibuat :
<div style="text-align:center">

![bacamysql](/minggu-06/bacaMysql.jpg)

</div>

<h3>B. Membuat program Go untuk koneksi dan membaca data dari MongoDB</h3>

1. Sebelum melakukan pembuatan program kita diwajibkan untuk menginstall MongoDB (atau gunakan versi cloudnya MongoDB atlas), compiler Go, dan Mongo driver untuk Go.

2. Setelah proses instalasi selesai kita langsung membuat sebuah database yang akan digunakan untuk melakukan koneksi, dan pembacaan data menggunakann Go.

3. Karena saya menggunakan MongoDB Atlas maka harus membuat akun untuk bisa membuat cluster sendiri, disini nama cluster saya di MongoDB Atlas adalah nim saya sendiri.

![buatCluster](/minggu-06/UsingMDBatlas.jpg)

4. Lalu mendapatkan URI string dari MongoDB Atlas agar dapat mengakses cluster dan databasenya. Jika tidak menggunakan MongoDB Atlas cukup isi URI string dengan port localhost, dan nama databasenya.

![URIstring](/minggu-06/setup-fromatlas.jpg)

5. Setelah mendapatkan URI string kita sudah bisa melakukan proses coding dengan Go untuk melakukan koneksi, dan membaca data dari database yang ada di MongoDB. [Program Go MongoDB](/minggu-06/featuringmongo.go)

6. Sebelum melakukan run kita akan mengisikan data didalam collection.

![isicollection](/minggu-06/datadiatlas.jpg)

7. Hasil dari run program Go :

![hasilMongo](/minggu-06/mongocomplete.jpg)</br>

Untuk output yang disorot warna hijau output yang dikeluarkan jika sukses melakukan koneksi, dan output yang disorot warna kuning adalah isi collection yang ada di database.
