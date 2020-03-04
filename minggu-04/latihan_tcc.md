<h1> Melakukan install DevStack di image Ubuntu 18.04 </h1>

<h3> Proses :</h3>

1. Masuk kedalam virtual box dan start machine Ubuntu 18.04

![virtualbox](/minggu-04/Virtual)

2. Melakukan instalasi git di dalam virtual machine dengan menggunakan perintah "sudo apt install git." pada terminal, karena kita akan melakukan cloning terhadap repo devstack. Repo devstack berisi skrip yang menginstal OpenStack dan templat untuk file konfigurasi.

3. Sebelum melakukan clone, kita disarankan untuk membuat user dengan nama stack , dan memberikannya sudo privileges agar tidak perlu menggunakan password ketika akan menjalankan suatu perintah. 

4. Setelah itu kita akan langsung melakukan clone terhadap repo devstack ke repo local.
![commandclone](/minggu-04/cloning_dev.jpg)

5. Berikut hasil dari cloning dari repo devstack.
![local](/minggu-04/devstack_local.jpg)

6. Lalu melakukan instalasi devstack, karena mengalami banyak kendala seperti file yang tidak bisa diakses walaupun user stack sudah diberi privileges root, saya tidak bisa mengakses dashboard devstack sampai laporan ini dibuat.