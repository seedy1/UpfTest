FROM golang:1.18

# Set destination for COPY
WORKDIR /app

# Download Go modules
# COPY go.mod go.sum ./
COPY . .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./
# COPY config.json ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /upf-agg
EXPOSE 8080
# Run
CMD ["/upf-agg"]
