FROM gliderlabs/alpine
MAINTAINER Massimiliano Dessi (@desmax74)
WORKDIR /app
ADD src/org.desmax/gdgsardegna/gdgsardegna /app/gdgsardegna
EXPOSE 8080
ENTRYPOINT ["/app/gdgsardegna"]