FROM centos
EXPOSE 8080
ENTRYPOINT ["/go-web"]
COPY ./bin/ /
