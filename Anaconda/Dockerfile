FROM ubuntu:latest
RUN apt update && apt install -y \
  sudo \
  wget \
  vim \
  nano
WORKDIR /opt
RUN wget https://repo.anaconda.com/archive/Anaconda3-2020.11-Linux-x86_64.sh && \
  sh /opt/Anaconda3-2020.11-Linux-x86_64.sh -b -p /opt/anaconda3 && \
  rm -f Anaconda3-2020.11-Linux-x86_64.sh
ENV PATH /opt/anaconda3/bin:$PATH

WORKDIR /
RUN pip install --upgrade pip
