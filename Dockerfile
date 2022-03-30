FROM golang:1.17

WORKDIR /go/src/food

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=1

ENV FOOD_SQLITE_PATH=/tmp/food/food.sqlite
ENV FOOD_SESSION_STORE_PATH=/tmp/food/sessions
ENV FOOD_UPLOAD_FOLDER=/tmp/food/uploads

RUN mkdir -p ${FOOD_SESSION_STORE_PATH} ${FOOD_UPLOAD_FOLDER}
RUN apt-get update && apt-get install -y sqlite3

COPY . .

RUN go build -o food
RUN cp ./food /usr/local/bin/food

CMD /usr/local/bin/food
