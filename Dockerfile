FROM golang:1.22-bookworm AS build
WORKDIR /src
COPY . .
RUN go build -o /out/lab-sample-tracker .
FROM debian:bookworm-slim
COPY --from=build /out/lab-sample-tracker /usr/local/bin/lab-sample-tracker
ENTRYPOINT ["lab-sample-tracker"]
