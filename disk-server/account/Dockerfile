FROM golang:1.19

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 18081

CMD ["./main"]

# docker build -t jiapengcheng/disk-account-image:v1 .
# docker push jiapengcheng/disk-account-image:v1
# docker pull jiapengcheng/disk-account-image:v1
# docker run --name disk-account-01 -p 18081:18081 -d disk-account-image:v1