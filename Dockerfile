FROM golang:1.18-alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o pets-microservice .

#FROM alpine

#COPY --from=builder /build/pets-microservice /

EXPOSE 8080
CMD ["./pets-microservice"]