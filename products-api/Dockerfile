# Hangi docker image'ı kullanılacak
FROM golang:1.16-alpine

# app isminde bir dizin oluşturuyoruz.
WORKDIR /app

# Projemizdeki mod ve sum dosyalarını klasöre taşıyoruz
COPY go.mod ./
COPY go.sum ./

# uygulamada kullandığımız bağımlılıkları indiriyoruz echo gibi
RUN go mod download

# main.go gibi go uzantılı kaynak dosyalarımızı taşıyoruz
COPY *.go ./

# Projeyi build ediyporuz
RUN go build -o /products-api

# hangi portta host edeceğimizi söylüyoruz.
EXPOSE 8000

CMD [ "/products-api" ]