FROM golang:1.17

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /envoy-als
EXPOSE 5000
CMD [ "/envoy-als"]