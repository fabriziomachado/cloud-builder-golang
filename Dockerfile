FROM golang:1.8-alpine

WORKDIR /go/src/calculator
ADD ./src/calculator /go/src/calculator

RUN go install .

#CMD ["go", "test", "-v"]
#CMD ["go", "run", "add.go"]
CMD ["calculator"]
