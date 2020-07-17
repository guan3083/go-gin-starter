FROM golang:1.14 as build-stage
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOARM=6
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download && go mod verify && go build -ldflags="-s -w" -o quickpass

FROM alpine:3.11
WORKDIR /app
RUN mkdir conf
COPY --from=build-stage /app/quickpass .
COPY --from=build-stage /app/conf/* /app/conf/
COPY --from=build-stage /app/script/* /app/script/
RUN rm /app/conf/app.ini && mv /app/conf/app_release.ini /app/conf/app.ini
EXPOSE 8000
CMD ["./quickpass"]
