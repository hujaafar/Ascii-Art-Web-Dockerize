FROM golang:1.19

WORKDIR /app

LABEL description="Generate ASCII-ART from text using different banners"

LABEL version="1.0"

LABEL release-date="2024-07-16"


COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /asciiArtWeb

EXPOSE 2004

CMD ["/asciiArtWeb"]