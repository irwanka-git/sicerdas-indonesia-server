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
cd /home/docker/sicerdas-indonesia-server && docker compose build
```
```
cd /home/docker/sicerdas-indonesia-server && docker compose up -d
```
```
cd /home/docker/sicerdas-indonesia-server && docker compose down
```
```
cd /home/docker/sicerdas-indonesia-server && docker compose down
```

- shell php container:
```
cd /home/docker/sicerdas-indonesia-server && docker compose exec php bash
```
```
composer install
```
```
.env for PHP
```
```
APP_NAME=Laravel
APP_KEY=base64:2yPF8wd8dc3SoLzkSQk6i2AXxeXJJNdLgP90byo/r8U=
APP_DEBUG=true
APP_LOG_LEVEL=debug
APP_URL=https://sicerdas.web.id
GO_API_URL=http://host.docker.internal:8045

DB_CONNECTION=pgsql
DB_HOST=host.docker.internal
DB_PORT=5432
DB_DATABASE=pgstore
DB_USERNAME=postgres
DB_PASSWORD=@Scd2022*


BROADCAST_DRIVER=log
CACHE_DRIVER=file
SESSION_DRIVER=file
QUEUE_DRIVER=sync

REDIS_HOST=127.0.0.1
REDIS_PASSWORD=null
REDIS_PORT=6379

MAIL_DRIVER=smtp
MAIL_HOST=smtp.mailtrap.io
MAIL_PORT=2525
MAIL_USERNAME=null
MAIL_PASSWORD=null
MAIL_ENCRYPTION=null

PUSHER_APP_ID=
PUSHER_APP_KEY=
PUSHER_APP_SECRET=


PATH_REPORT=/var/www/html/storage/report/
RENDER_REPORT=http://host.docker.internal:8042
WKHTML=wkhtmltopdf
BASE_URL_API=http://host.docker.internal:8042/api

PATH_COVER=/var/www/html/storage/cover/
PYTHON_SCRIPT=/home/pyhton/cover.py
PYTHON_CMD=python3

NOCAPTCHA_SECRET=6LctlYUiAAAAABGE9CnClJJ561leUpSBi5uvyf6n
NOCAPTCHA_SITEKEY=6LctlYUiAAAAAKdVThRI5nRAOn3exy5Wxywq--xW
```
```
chmod -R a+rw storage
chmod -R a+rw storage/cover
chmod -R a+rw storage/report
chmod -R a+rw bootstrap/cache
chmod -R a+rw public/gambar
chmod -R a+rw public/uploads
chmod -R a+rw public/report 
```
-=-==-08
shell go container:
```
cd /home/docker/sicerdas-indonesia-server && docker compose exec go bash
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

.env for golang (go)
```
PORT=":5000"
EMAIL_FIREBASE="sicerdas.service@gmail.com"
FIREBASE_BUCKET="sicerdas-indonesia-service-repository"

TIMEZONE="Asia/Jakarta"
DSN_SICERDAS="host=host.docker.internal user=postgres password=@Scd2022* dbname=pgstore port=5432 sslmode=disable TimeZone=Asia/Jakarta"
JWT_SIGN_KEY = "6YW1pZ29zLXNzby5teWxvZ2luLmNvbSIs"
URL_ISS="https://sicerdas.web.id"
REF_MOBILE_LOGIN="cLwx789acDEol95erTY73"
PATH_JSON_SOAL="/app/storage/json-soal"
URL_FIREBASE_STORAGE_PREFIX="https://storage.googleapis.com/%v/%v"
URL_FIREBASE_PREFIX="https://storage.googleapis.com/%v/%v/%v"
URL_SICERDAS="https://sicerdas.web.id"
```
