FROM centos:7

RUN yum -y install epel-release && yum -y install git && yum -y install golang
RUN mkdir /usr/local/mars
COPY bulo.fun* /usr/local/mars
RUN cd /var/local \
   && git clone https://github.com/YingshuLu/mars.git \
   && cd mars && go build && mv mars* /usr/local/mars \
   && cd /usr/local/mars && rm -fr /var/local/mars

EXPOSE 443

STOPSIGNAL SIGQUIT
WORKDIR "/usr/local/mars"
CMD ["./mars"]