# Single-binary container. ~12 MB final image.
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download || true
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/cloister-bell ./

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/cloister-bell /cloister-bell
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/cloister-bell"]
