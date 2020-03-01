FROM gcc:5.4.0

RUN apt-get update -y && apt-get upgrade -y

RUN curl -Lo boost.tar.gz http://sourceforge.net/projects/boost/files/boost/1.60.0/boost_1_60_0.tar.gz
RUN mkdir -p /opt/boost/
RUN tar xvzf boost.tar.gz -C /opt/boost

CMD ["bash"]
