FROM golang

WORKDIR /app

COPY . .

RUN go build .

# Not sure if there is better way
# to trigger execution. 
ENTRYPOINT [ "go","run","." ]