FROM fullstorydev/grpcurl:latest as builder

FROM alpine 
WORKDIR /test

COPY --from=builder /bin/grpcurl /bin/grpcurl
COPY test.sh .
RUN chmod +x test.sh

ENTRYPOINT [ "/bin/sh", "test.sh" ]