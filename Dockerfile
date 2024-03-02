FROM golang:1.21.5 AS builder

WORKDIR /app
COPY go.* ./
COPY . .

COPY .env .

RUN chmod a+r .env

RUN go build -o main .

FROM alpine:3.19

WORKDIR /app

RUN apk --no-cache add libc6-compat

ENV GIN_MODE=release

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]


#FROM golang:1.21.5
#
## Set the working directory inside the container
#WORKDIR /app
#
## Copy the local code to the container
#COPY . .
#
## Build the Go application
#RUN go build -o main .
#
## Expose the port the app runs on
#EXPOSE 8080
#
## Command to run the executable
#CMD ["./main"]

