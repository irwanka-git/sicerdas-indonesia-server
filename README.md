# socerdas-indonesia-server
 Versi Docker Golang API

# Direktori Penting (Migrasi Manual)
+ cover.tar.gz  :   ../storage/cover
+ report.tar.gz :   ../storage/report
+ gambar.tar.gz :   ../public/gambar
+ upload.tar.gz :   ../public/uploads

Server Lama
```
cd /home/scd-id
```
```
cd storage && tar -zcvf cover.tar.gz cover/
```
Server Baru
```
tar -zxvf cover.tar.gz -C /home/docker/php/storage
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