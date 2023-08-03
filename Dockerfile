FROM golang:latest 

WORKDIR /usr/src/app 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 3000

CMD [ "go", "run", "main.go" ]