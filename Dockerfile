FROM golang:1.18

# Set destination for COPY
WORKDIR /app
COPY . .
RUN go mod download
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /upf-agg
EXPOSE 8080
# Run
CMD ["/upf-agg"]
