FROM golang:1.19.0


RUN mkdir /code && \
    mkdir /result && \
    chmod -R 777 /code && \
    chmod -R 777 /result

COPY /code/server.go /code

WORKDIR /code

# CMD ["go", "run", "/code/server.go"]