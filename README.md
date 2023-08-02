# backend-website-prokerin

data dikirim saat like ( id user bisa dari token.. cukup id proker atau id
comment dari url langsung )

<!-- cara penggunaan -->

file rouute.go itu routing dari api nya

folder schema itu untuk mengetahui param apa aja yang dipake saat melakukan request(post dan put)

di bagian routing jika ada ":id" dsb yang startnya titik2 itu maksudnya adalah untuk menerima parameter dari url
example : api/project/:id --- ditulisnya saat request api/project/123
maka nilai id = 123

<!-- cara pemakaian docker -->

instal docker desktop lalu buka apknya (ada di google tinggal download)
buka cmd arahin ke dir project.. jalanin docker compose up (atay docker compose build --> docker compose up)
klo misal error coba beberapa kali.. klo misal tetep error langsung chat aja
