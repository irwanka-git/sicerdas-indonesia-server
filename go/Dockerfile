FROM golang:1.21
LABEL maintainer="IRWAN KURNIAWAN <irwanka.emaiil@gmail.com>"
RUN apt-get update

#PDFUNITE
RUN apt-get install -y poppler-utils

#WKHTMLTOPDF
RUN apt-get install -y openssl build-essential libssl-dev libxrender-dev git-core libx11-dev libxext-dev libfontconfig1-dev libfreetype6-dev fontconfig
RUN wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
RUN tar xvJf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
RUN cp wkhtmltox/bin/wkhtmlto* /usr/bin/

WORKDIR /app   
# COPY go/go.mod  ./
# COPY go/go.mod  ./
COPY go/ ./
RUN chmod +x /app/entrypoint.sh
RUN chmod +x /app/export.sh

ENV TZ=Asia/Jakarta
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy
RUN air init 
CMD ["/app/entrypoint.sh"]
 