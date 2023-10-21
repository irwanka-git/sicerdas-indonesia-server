FROM devilbox/php-fpm:7.4-prod
LABEL maintainer="IRWAN KURNIAWAN <irwanka.emaiil@gmail.com>"
RUN apt-get update
 

#SUPPORT
RUN apt-get install -y libpq-dev
RUN apt-get install -y wget
RUN apt-get install -y zlib1g-dev libzip-dev libpng-dev
RUN apt-get install -y libfreetype6-dev libjpeg62-turbo-dev libgd-dev
RUN apt-get install -y poppler-utils
RUN apt-get install -y cron
RUN apt-get install -y vim
RUN apt-get install -y xvfb
RUN apt-get install -y libxrender1
RUN apt-get install -y libfontconfig
RUN apt-get install -y libxtst6
RUN apt-get install -y xz-utils


#COMPOSER
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
RUN php -r "if (hash_file('sha384', 'composer-setup.php') === 'e21205b207c3ff031906575712edab6f13eb0b361f2085f1f1237b7126d785e826a450292b6cfd1d64d92e6563bbde02') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;"
RUN php composer-setup.php
RUN php -r "unlink('composer-setup.php');"
RUN mv composer.phar /usr/local/bin/composer

#WKHTMLTOPDF
RUN apt-get install -y openssl build-essential libssl-dev libxrender-dev git-core libx11-dev libxext-dev libfontconfig1-dev libfreetype6-dev fontconfig
RUN wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
RUN tar xvJf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
RUN cp wkhtmltox/bin/wkhtmlto* /usr/bin/
