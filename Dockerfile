FROM golang:latest as builder
RUN apt-get update && apt-get install -y nocache git ca-certificates && update-ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
COPY .env .env ./
#RUN go env -w GOPROXY="https://goproxy.io,direct"
#RUN go mod tidy
RUN go mod download
COPY . .
RUN go build -o /app/bin/api-service .


FROM golang:latest
RUN apt-get update \
 && apt-get install -y git \
 && apt-get install make

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
RUN git clone https://github.com/klovercloud-ci-cd/klovercloud-ci-jwt-token-generator.git /app/klovercloud-ci-jwt-token-generator
WORKDIR /app/klovercloud-ci-jwt-token-generator
RUN make build
ENV PATH="$PATH:/app/klovercloud-ci-jwt-token-generator"
RUN go install
WORKDIR /app
COPY --from=builder /app/bin /app
EXPOSE 8080
# Run the executable
CMD ["./klovercloud-ci-api-service"]