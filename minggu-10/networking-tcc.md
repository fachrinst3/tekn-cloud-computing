<h1> Latihan TCC-minggu 10  </h1>

<h2> Docker Networking Hands-on Lab </h2>

<h3> A. Networking Basics </h3>

 1. Menggunakan perintah "docker network" untuk melihat command yang dapat digunakan untuk mengelola networking di container

 ![step1](g1.jpg)

 2. Menggunakan perintah "docker network -ls" untuk melihat list network docker yang ada di docker host yang saya gunakan.

 ![step2](g2.jpg)

 3. Menggunakan perintah inspect yang ada di docker network untuk melihat detail konfigurasi dari network yang memiliki nama bridge.

 ![step3](g3.jpg)

 4. Menggunakan perintah "docker info" untuk melihat informasi mengenai docker yang ter-install di vm yang saya gunakan, dari perintah ini dapat dilihat plugin networknya.

 ![step4](g4.jpg) 

<h3> B. Bridge Networking </h3>

 1. Melakukan cek terhadap network container bridge di list menggunakan "docker networking -ls"
 
 ![step5](g2.jpg)

 2. Lalu melakukan update, dan menambahkan command bridge control ke linux vm yang saya gunakan.

 ![step6](g5.jpg)

 ![step7](g6.jpg)

 3. Selanjutnya adalah menghubungkan container, pertama kita melakukan run terhadapa container yang berjalan di background.

 ![step8](g7.jpg)

 ![step9](g8.jpg)

 Lalu menggunakan perintah bridge control untuk melihat bahwa container yang baru di run sudah otomatis menjadi interface dari bridge docker0.

 ![step10](g9.jpg)

 Selanjutnya melakukan tes konektivitas terhadap container baru yang aktif di background melalui ping ke ip yang sebelumnya sudah diketahui menggunkan perintah inspect.

 ![step11](g10.jpg)

 Lalu untuk memastikan bahwa container dapat terhubung ke jaringan luar kita menjalankan shell linux di container yang berjalan dibackground, dan menginstall program untuk melakukan ping ke github.com.

 ![step12](g11.jpg)

 4. Selanjutnya kita akan melakukan konfigurasi NAT untuk koneksi external, pertama kita melakukan run terhadap container yang menggunakan image nginx.
 
 ![step13](g12.jpg)

 ![step14](g13.jpg)
 


 <h3> C. Overlay Networking </h3>

 1. Melakukan inisiasi docker swarm

 ![step15](g14.jpg)

 lalu menyalin output token ke terminal kedua dan bergabung ke node worker
 
 ![step16](g15.jpg)

 Dan melakukan cek apakah keduanya sudah berada pada satu swarm(dicek melalui swarm manager yaitu terminal1/node1).
 
 ![step17](g16.jpg)

 2. Lalu melakukan pembuatan overlay network dengan nama overnet.

 ![step18](g17.jpg)

 menjalankan perintah untuk melihat list network pada node2/terminal2 dan hasilnya overnet tidak kelihatan karena Docker hanya memperluas network overlay ke host ketika diperlukan. Ini biasanya ketika host menjalankan tugas dari layanan yang dibuat di network.
 
 ![step19](g18.jpg)

 3. Selanjutnya membuat service dengan nama myservice, dan memastikannya sudah menjalankan task untuk mereplika berjalan dikedua node.

 ![step20](g19.jpg)

 ![step21](g20.jpg)
 
 Karena sudah berjalan didua node maka overnet akan ditampilkan ketika melihat list network pada node2/terminal2.

 ![step22](g21.jpg)

 4. Lalu untuk melakukan tes bahwa node1, dan node2 terhubung kita menggunakan shell di container myservice2(pada node1), dan melakukan ping ke myservice1(pada node2).
 
 ![step23](g22.jpg)
  
 5. Dan yang terakhir kita akan melakukan tes ping melalui shell ke service myservice yang sudah kita buat.

![step24](g23.jpg)
 

