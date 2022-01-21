FROM golang:1.14 AS compiler
WORKDIR /src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ./contact.out .

FROM gcr.io/distroless/static
COPY --from=compiler /src/app/contact.out /server
ENTRYPOINT ["/server"]