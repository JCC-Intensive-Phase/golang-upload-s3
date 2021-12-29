# golang-upload-s3

# go-aws

Repo untuk belajar AWS

## Requirement
1. Akun di AWS
2. Bucket di AWS S3

## Setup
1. isi value dari variables di `.env`
2. set variables di `.env` ke environment variables dengan command
```
set -a
source .env
```
3. `go run main.go`

## Explanation
Saat ini hanya ada 1 endpoint yang bisa dicoba yaitu `POST /upload`.
Endpoint tersebut akan menerima request body dalam bentuk `multipart/form-data` berupa file dan mengupload file tersebut ke S3 bucket yang sudah ditentukan sebelumnya.

Contoh pengiriman request
```
curl --location --request POST 'localhost:8080/upload' \
--form 'file=@"/D:/test.txt"'
```

Contoh balikan dari server
```
{
    "file_location": "https://jabar-coding-camp.s3.ap-southeast-1.amazonaws.com/test.txt",
    "message": "file uploaded"
}
```
