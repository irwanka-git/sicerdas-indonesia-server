# socerdas-indonesia-server
 Versi Docker Golang API

# Direktori Penting (Migrasi Manual)
+ cover.tar.gz  :   ../storage/cover
+ report.tar.gz :   ../storage/report
+ gambar.tar.gz :   ../public/gambar
+ upload.tar.gz :   ../public/uploads

Server Lama
```
cd /home/scd-id/storage && tar -zcvf cover.tar.gz cover/
```
```
cd /home/scd-id/storage && tar -zcvf report.tar.gz report/
```
```
cd /home/scd-id/public && tar -zcvf gambar.tar.gz gambar/
```
```
cd /home/scd-id/public && tar -zcvf uploads.tar.gz uploads/
```

Server Baru
```
tar -zxvf cover.tar.gz -C /home/docker/php/storage
```
```
tar -zxvf report.tar.gz -C /home/docker/php/storage
```
```
tar -zxvf gambar.tar.gz -C /home/docker/php/public
```
```
tar -zxvf uploads.tar.gz -C /home/docker/php/public
```
 