#ARG APP_NAME="shorty"

# Build stage
FROM golang:1.18-alpine as builder
WORKDIR /app/src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shorty .
#RUN go get -d -v ./... \
#    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${APP_NAME} .

# Production stage \
FROM scratch
WORKDIR /app/bin
COPY --from=builder /app/src .
EXPOSE 80
CMD ["./shorty"]