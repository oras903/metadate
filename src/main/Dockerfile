FROM golang:latest
MAINTAINER ftpw <352965446@qq.com>
RUN apt-get update
RUN apt-get install openssh-server supervisor haproxy -y
RUN apt-get install psmisc -y
ADD supervisord.conf /etc/supervisor/supervisord.conf


#RUN wget https://pan.baidu.com/s/1o9AXm90  -O /tmp/go1.9.2.linux-amd64.tar.gz
#RUN tar -xzvf /tmp/go1.9.2.linux-amd64.tar.gz -C /usr/local
#RUN mkdir -p  /home/go

#ENV GOPATH /home/go  
#ENV GOROOT $PATH:/usr/local/go
#ENV PATH $GOROOT/bin:$PATH

RUN mkdir -p $GOPATH/src/github.com/kelseyhightower
RUN git clone https://github.com/kelseyhightower/confd.git $GOPATH/src/github.com/kelseyhightower/confd
WORKDIR $GOPATH/src/github.com/kelseyhightower/confd
RUN apt-get update
RUN apt-get install gcc
RUN apt-get install make
RUN make
RUN chmod +x $GOPATH/src/github.com/kelseyhightower/confd/bin/confd
RUN mkdir -p /etc/confd/conf.d
RUN mkdir -p /etc/confd/templates
ADD haproxy.toml /etc/confd/conf.d/haproxy.toml
ADD haproxy.cfg.tmpl /etc/confd/templates/haproxy.cfg.tmpl

EXPOSE 8888 22


RUN /etc/init.d/ssh stop
#RUN killall haproxy
#RUN rm /var/run/haproxy.pid
CMD ["/usr/bin/supervisord","-c","/etc/supervisor/supervisord.conf"]
