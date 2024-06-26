FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o main .

FROM alpine
RUN adduser -S -D -H -h /app justme
USER justme
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]