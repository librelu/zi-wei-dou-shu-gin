FROM golang:1.15.2-alpine3.12
ENV WORKDIR=/zi-wei-dou-shu-gin
ENV GIN_MODE=release
RUN mkdir ${WORKDIR}
WORKDIR ${WORKDIR}
COPY . ${WORKDIR}
VOLUME .:${WORKDIR}
RUN go install ${WORKDIR}
