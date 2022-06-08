# use always a version, latest change a lot
FROM golang:latest 

WORKDIR /app

COPY . .

RUN go build ./src/math.go

CMD [ "./math" ]