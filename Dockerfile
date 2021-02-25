FROM golang:1.15
WORKDIR /app/src/
COPY . .
RUN make build
RUN ["chmod", "+x", "/app/src/wait-for-it.sh"]
CMD ["/app/src/main"]