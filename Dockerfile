FROM golang:1.16-alpine AS build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV DOCKER_BUILDKIT=1
WORKDIR /src
COPY . .
RUN go build  -o /out/diamonds cmd/main.go
FROM alpine AS base
COPY --from=build /out/diamonds /
RUN addgroup -g 2000 diamonds && \
    adduser -H -D -s /bin/sh -u 2000 -G diamonds diamonds
USER diamonds
CMD /diamonds
