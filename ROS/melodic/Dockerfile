FROM dorowu/ubuntu-desktop-lxde-vnc:bionic-lxqt

RUN apt-get update -q && \
    apt-get upgrade -yq && \
    apt-get install -yq wget curl git build-essential vim sudo lsb-release locales bash-completion tzdata gosu && \
    rm -rf /var/lib/apt/lists/*
RUN useradd --create-home --home-dir /home/ubuntu --shell /bin/bash --user-group --groups adm,sudo ubuntu && \
    echo ubuntu:ubuntu | chpasswd && \
    echo "ubuntu ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
RUN git clone https://github.com/Tiryoh/ros_setup_scripts_ubuntu.git /tmp/ros_setup_scripts_ubuntu && \
    gosu ubuntu /tmp/ros_setup_scripts_ubuntu/ros-melodic-desktop.sh && \
    rm -rf /var/lib/apt/lists/*
ENV USER ubuntu

RUN sudo apt-get update && sudo apt-get upgrade -y && \
  sudo apt-get install -y --no-install-recommends \
  ros-melodic-move-base \
  ros-melodic-slam-gmapping \
  ros-melodic-mouse-teleop \
  ros-melodic-navigation \
  python-pip \
  python3-pip \
  nano && \
  pip3 install rospkg catkin_pkg

RUN echo "source ~/workspace/devel/setup.bash" >> ~/.bashrc && \
  echo "alias cm='cd ~/workspace && catkin_make'" >> ~/.bashrc
