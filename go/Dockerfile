FROM golang:1.21
LABEL maintainer="IRWAN KURNIAWAN <irwanka.emaiil@gmail.com>"
WORKDIR /app   
# COPY go/go.mod  ./
# COPY go/go.mod  ./
COPY go/ ./
RUN chmod +x /app/entrypoint.sh

ENV TZ=Asia/Jakarta
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy
RUN air init 
CMD ["/app/entrypoint.sh"]
 