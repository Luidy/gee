FROM golang:latest
RUN mkdir /gee
ADD . /gee
WORKDIR /gee
RUN make build
ARG BUILD_PORT=10620
ENV PORT $BUILD_PORT
EXPOSE $BUILD_PORT
ENTRYPOINT ["bin/gee"]
