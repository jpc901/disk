FROM golang:1.19

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]


# docker build -t jiapengcheng/disk-master-image:v1 .
# docker buildx build --platform linux/amd64 -t jiapengcheng/disk-master-image:v3-amd64 
# docker push jiapengcheng/disk-master-image:v3-amd64
# docker pull jiapengcheng/disk-master-image:v1
# docker run --name disk-master-01 -p 8081:8081 -d disk-master-image:v1