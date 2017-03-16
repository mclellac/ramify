
FROM scratch
MAINTAINER https://github.com/CentOS/sig-cloud-instance-images
ADD centos-7-docker.tar.xz /

RUN yum update 
RUN yum install -y ca-certificates

ADD . /app
WORKDIR /app

EXPOSE 8080

CMD ["/app/api"]
