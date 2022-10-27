# Aplikasi Perpustakaan

Ini merupakan aplikasi perpustakaan yang dibangun dengan bahasa pemograman golang

# Proses instalasi

Berikut adalah proses instalasi dari Aplikasi:

## Klonasi dari repositori github

    git clone https://github.com/kasfulk/golang-library.git

## Sebelum melakukan instalasi

Sebelum melakukan instalasi adalah menginstall app go pada PC anda sebelum menjalankan aplikasi tersebut. Untuk instalasi bisa didapatkan pada [Link ini](https://go.dev/dl/)

## Untuk menjalankan aplikasi

Lakukan command pada folder aplikasi seperti ini:

    go mod download
    go run cmd/main.go
anda juga dapat menjalankan dengan mode *watch* yaitu dengan:

    sh watcher.sh
sebelum menjalankan mode *watch* pastikan package nodemon telah terinstall di PC anda.


# Struktur Folder


Berikut struktur folder pada design pattern project ini:

|                |ASCII                          																	|
|----------------|--------------------------------------------------------------------------------------------------|
|`cmd/`          | Pada folder ini merupakan folder yang berisi main program dari project ini.                         |
|`databases/`    | Pada folder ini berisi berbagai macam sub program yang berhubungan dengan databases.                |
|`handlers/`     | Pada folder ini berisi handlers fungsi yang digunakan setelah proses route.                          |
|`routes/`       | Pada folder ini berisi berbagai macam route yang nantinya akan diarahkan kepada folder handlers.     |

Alasan dipilihnya design pattern ini dikarenakan kemudahan dalam memantain apabila aplikasinya telah memiliki banyak route dan function.
