FROM golang:1.20
ARG STAGE
ENV STAGE="prod"
WORKDIR /app
COPY . .

RUN go build -o main .

CMD ["./main"]
