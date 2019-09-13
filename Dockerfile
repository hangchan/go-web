FROM centos
EXPOSE 8080
ENTRYPOINT ["/go-web"]
COPY ./bin/ /
COPY ./count/count.txt /tmp/count.txt