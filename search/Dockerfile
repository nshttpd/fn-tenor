FROM fnproject/go:dev as build-stage
WORKDIR /function
ADD . /go/src/github.com/nshttpd/fn-tenor/search
RUN cd /go/src/github.com/nshttpd/fn-tenor/search && \
    go build -o func .

FROM fnproject/go
WORKDIR /function
COPY --from=build-stage /go/src/github.com/nshttpd/fn-tenor/search/func /function/
ENTRYPOINT [ "/function/func" ]