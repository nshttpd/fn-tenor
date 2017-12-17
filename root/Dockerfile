FROM fnproject/go:dev as build-stage
WORKDIR /function
ADD . /go/src/github.com/nshttpd/fn-tenor
RUN cd /go/src/github.com/nshttpd/fn-tenor && \
    go build -o func .

FROM fnproject/go
WORKDIR /function
COPY --from=build-stage /go/src/github.com/nshttpd/fn-tenor/func /function/
ENTRYPOINT [ "/function/func" ]