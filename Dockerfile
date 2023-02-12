FROM golang:1.20-alpine3.17 AS builder
LABEL stage=builder
WORKDIR /desafio-sword 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a

FROM scratch AS final
ARG ENV
COPY --from=builder /desafio-sword .
CMD [ "./desafio-sword", "app:server" ]