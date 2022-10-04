FROM golang:1.19-alpine as builder
COPY go.mod go.sum /go/src/github.com/madz0k/demo-crm-backend/
WORKDIR /go/src/github.com/madz0k/demo-crm-backend
RUN go mod download
COPY . /go/src/github.com/madz0k/demo-crm-backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/demo-crm-backend github.com/madz0k/demo-crm-backend

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/madz0k/demo-crm-backend/build/demo-crm-backend /usr/bin/crm
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/crm"]