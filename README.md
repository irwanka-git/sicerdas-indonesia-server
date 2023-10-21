# socerdas-indonesia-server
 Versi Docker Golang API

# Direktori Penting (Migrasi Manual)
+ cover.tar.gz  :   ../storage/cover
+ report.tar.gz :   ../storage/report
+ gambar.tar.gz :   ../public/gambar
+ upload.tar.gz :   ../public/uploads

*** Create Backup tar.gz dari server lama
```
tar -zcvf /home/scd-id/storage_cover.tar.gz /home/scd-id/storage/cover/
```
```
tar -zcvf /home/scd-id/storage_report.tar.gz /home/scd-id/storage/report/
```
```
tar -zcvf /home/scd-id/public_gambar.tar.gz /home/scd-id/public/gambar/
```
```
tar -zcvf /home/scd-id/public_uploads.tar.gz /home/scd-id/public/uploads/
```

*** Extract Backup to server di server baru