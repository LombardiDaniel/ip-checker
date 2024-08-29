#### - BUILDER - ####
FROM golang:1.22.3 AS builder

WORKDIR /cmd

COPY cmd/go.mod go.mod
RUN go mod download

COPY cmd/ ./

RUN go build -o /bin/main main.go


#### - SERVER - ####
FROM alpine:3.19.1 as server

RUN apk add --no-cache gcompat=1.1.0-r4 libstdc++=13.2.1_git20231014-r0
# RUN apk add --no-cache gcompat libstdc++

WORKDIR /cmd

COPY --from=builder /bin/main ./main

# RUN adduser --system --no-create-home nonroot
# USER nonroot

CMD ["./main"]
