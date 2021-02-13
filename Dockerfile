FROM golang:1.12 as build
ENV CGO_ENABLED=0
WORKDIR /build
COPY go.mod .
COPY main.go .
RUN go mod download
RUN go build -o /output/api .

FROM scratch
COPY --from=build /output/api /api
ENTRYPOINT [ "/api" ]