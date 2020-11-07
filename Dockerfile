FROM docker-ci.cyber.sas.com/cid-golang-build AS talongo

RUN mkdir -p /opt/edsw 

RUN mkdir -p /opt/edsw/go/bin /opt/edsw/go/src /opt/edsw/go/pkg /opt/edsw/go/src/talon-kafka \
    /opt/edsw/go/src/sas.com/new-kafka/talon

WORKDIR /opt/edsw

#COPY MockIt-1.0-SNAPSHOT.tar /opt/edsw

#RUN ls -l /opt/edsw
#RUN tar xvf MockIt-1.0-SNAPSHOT.tar
#RUN ls -lR /opt/edsw
#RUN java -version
#EXPOSE 5555
#ENTRYPOINT ["/opt/edsw/MockIt-1.0-SNAPSHOT/bin/MockIt", "udp", "5555", "kafka:9092", "edwin.test""]

ENV PATH="/usr/local/go/bin:/opt/edsw/go/bin:${PATH}"

RUN ls -ltr /opt/edsw/go

RUN find . -name sarama -print

COPY protobuf.tgz protobuf.tgz

RUN ls -l /opt/edsw

WORKDIR /opt/edsw

RUN uname -a

RUN cat /etc/centos-release

RUN yum group install -y "Development Tools" perl-core libtemplate-perl zlib-devel
RUN yum clean all

WORKDIR /opt/edsw

RUN     git clone https://github.com/protocolbuffers/protobuf.git && \
    cd protobuf && \
    git submodule update --init --recursive && \
    ./autogen.sh && \
     ./configure && \
     make -j 4 && \
     make install && \
     ldconfig 

RUN /usr/local/go/bin/go version

RUN GOPATH=/opt/edsw/go /usr/local/go/bin/go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN echo "Hi, Ed"

RUN find / -name protoc-gen-go -print

RUN GOPATH=/opt/edsw/go  /usr/local/go/bin/go get github.com/Shopify/sarama
RUN GOPATH=/opt/edsw/go  /usr/local/go/bin/go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN GOPATH=/opt/edsw/go  go get github.com/prometheus/client_golang/prometheus
RUN GOPATH=/opt/edsw/go  go get github.com/prometheus/client_golang/prometheus/promauto
RUN GOPATH=/opt/edsw/go  go get github.com/prometheus/client_golang/prometheus/promhttp

WORKDIR /opt/edsw/go/src/sas.com/new-kafka

COPY main.go ./
COPY proto proto/

RUN GOPATH=/opt/edsw/go /usr/local/bin/protoc -I=proto --go_out=talon   ./proto/*.proto

RUN GOPATH=/opt/edsw/go go build

# RUN /opt/edsw/go/src/sas.com/new-kafka/new-kafka -brokers d-crd-esp08.dev.cyber.sas.com:9092 -topics sas.cyber.NetflowParsed -group edtest

ENTRYPOINT [ "/opt/edsw/go/src/sas.com/new-kafka/new-kafka", "-brokers", "d-crd-esp08.dev.cyber.sas.com:9092", "-topics", "sas.cyber.NetflowParsed", "-group", "edtest" ]
