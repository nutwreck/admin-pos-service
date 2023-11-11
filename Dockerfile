# ======================
#  GO FIRST STAGE
# ======================

FROM golang:latest as builder
MAINTAINER DigyLabs Production <work.digy@gmail.com>
USER ${USER}
WORKDIR /app
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod \
  go.sum ./
RUN go mod download
COPY . ./

# ======================
#  GO FINAL STAGE
# ======================

FROM builder
WORKDIR /app
# RUN apt-get update \
#   && apt-get install -y \
#   make
# COPY --from=builder . ./app
# RUN make gobuild
RUN go build -o main .
EXPOSE 8881
CMD ["./main"]