# socerdas-indonesia-server
 Versi Docker Laravel, Nginx, Postgres Golang API

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
tar -zxvf cover.tar.gz -C /home/docker/sicerdas-indonesia-server/php/storage
```
```
tar -zxvf report.tar.gz -C /home/docker/sicerdas-indonesia-server/php/storage
```
```
tar -zxvf gambar.tar.gz -C /home/docker/sicerdas-indonesia-server/php/public
```
```
tar -zxvf uploads.tar.gz -C /home/docker/sicerdas-indonesia-server/php/public
```
 

# COMMAND SETUP SERVER
```
cd /home/docker
```
```
git clone git@github.com:irwanka-git/sicerdas-indonesia-server.git
```
```
cd /home/docker/sicerdas-indonesia-server && git pull
```
```
cd /home/docker/sicerdas-indonesia-server && docker composer build
```
```
cd /home/docker/sicerdas-indonesia-server && docker composer up -d
```
```
cd /home/docker/sicerdas-indonesia-server && docker composer down
```

setup nginx proxy 
- add include proxy path in server (/etc/nginx/conf.d)
```
/home/docker/sicerdas-indonesia-server/nginx/proxy.conf
```
```
http{
  ....
  include /home/docker/sicerdas-indonesia-server/nginx/proxy.conf;  
}
```

