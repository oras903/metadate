FROM ubuntu
MAINTAINER ftpw <352965446@qq.com>
RUN apt-get update
RUN apt-get install openssh-server supervisor haproxy -y
ADD supervisord.conf /etc/supervisor/supervisord.conf

ADD https://pan.baidu.com/s/1c3VAsM0 /tmp/confd
RUN cd /usr/sbin && cp /tmp/confd /usr/sbin/confd  && chmod +x /usr/sbin/confd && rm /tmp/confd
RUN mkdir -p /etc/confd/conf.d
RUN mkdir -p /etc/confd/templates
ADD haproxy.toml /etc/confd/conf.d/
ADD haproxy.cfg.tmpl /etc/confd/templates/

EXPOSE 8888 22


RUN /etc/init.d/ssh stop
RUN /etc/init.d/haproxy stop
CMD ["/usr/bin/supervisord","-c","/etc/supervisor/supervisord.conf"]
