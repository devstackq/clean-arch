#pull golang base  image
FROM golang:latest

RUN mkdir -p app
# ENV GOPATH=/
WORKDIR /app 
# add all in folder app
ADD . /app
COPY . .
# ADD go.mod go.sum ./
# downlaod dependency
# RUN go mod download
RUN go get
# RUN go build
# copy main file to curent image
# COPY cmd/main.go ./
# copy source code
# COPY . .
# run app  in os linux then build app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./cmd/main .
# RUN  go build -o /projectx

# create volume
# VOLUME [ "/app/shared" ]
# set entrypoint
# ENTRYPOINT [ "./cmd/main" ]
#open out port for OS, inside container app start port - 6969
EXPOSE 6969
# command with arg - projectx; when docker run
CMD ["./projectx"]