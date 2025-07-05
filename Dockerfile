FROM golang:1.24 AS build

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -tags lambda.norpc -o main ./main.go

FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /app/main ./main
ENTRYPOINT [ "./main" ]