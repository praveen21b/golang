FROM golang:1.15 as dev
# dev as alias for this container

WORKDIR /work 
# present inside the container

## Build stage
FROM golang:1.15 as Build

WORKDIR /app

COPY ./app/* /app/

RUN go build -o app
# creates binary file named as app

### Runtime app
FROM alpine as running
COPY --from=Build /app/app /
# copy from build stage: copies app binary from app folder and save it in / or in root dir
CMD [ "./app" ]