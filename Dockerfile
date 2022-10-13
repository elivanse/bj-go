# imagen base de dist de go
FROM golang:1.18.1-buster as builder

# directorio de la app
WORKDIR /home/ivanse/go/src/bj-go

# copiamos todo a la raiz del dir al dir /app
COPY go.mod .

# instalamos dependencias
RUN go mod download

# copiamos archivos.go al dir de la ./app
COPY *.go .

# buildeamos la app con config opcional
RUN CGO_ENABLED=0 GOOS=linuix go build -o /opt/go-docker-multistage

FROM alpine:latest
COPY --from=builer /opt/go-docker-multistage /opt/go-docker-multistage
EXPOSE 8080
ENTRYPOINT [ "executable" ]

# desde script e
# docker build --rm -t go-docker-multistage:alpha .
# docker image ls

# docker run -d -p 8080:8081 
