# ======================
#  GO FIRST STAGE
# ======================

FROM golang:latest as builder
MAINTAINER DigyLabs Production <work.digy@gmail.com>
USER ${USER}
WORKDIR /app
ENV GO111MODULE="on"
COPY go.mod \
  go.sum ./
RUN go mod download
COPY . ./

# ======================
#  GO FINAL STAGE
# ======================

FROM builder
WORKDIR /app
RUN apt-get update \
  && apt-get install -y \
  make
COPY --from=builder . ./app
RUN make gobuild
EXPOSE 7778
CMD ["./main"]