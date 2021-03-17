FROM dorowu/ubuntu-desktop-lxde-vnc:bionic-lxqt

RUN apt-get update -qq && apt-get install -y --no-install-recommends \
  git \
  nano \
  gnupg

# Install ROS Melodic
RUN sh -c 'echo "deb http://packages.ros.org/ros/ubuntu $(lsb_release -sc) main" > /etc/apt/sources.list.d/ros-latest.list' && \
  apt-key adv --keyserver 'hkp://keyserver.ubuntu.com:80' --recv-key C1CF6E31E6BADE8868B172B4F42ED6FBAB17C654 && \
  apt-get update -qq && \
  apt-get install -y ros-melodic-desktop-full
RUN ["/bin/bash", "-c", "echo "source /opt/ros/melodic/setup.bash" >> ~/.bashrc"]
RUN apt-get install -y python-rosdep python-rosinstall python-rosinstall-generator python-wstool build-essential && \
  rosdep init && \
  rosdep update

RUN apt-get install -y --no-install-recommends \
  ros-melodic-move-base \
  ros-melodic-slam-gmapping \
  ros-melodic-mouse-teleop \
  ros-melodic-navigation \
  python-pip \
  python3-pip && \
  pip3 install rospkg catkin_pkg

RUN echo "source ~/workspace/devel/setup.bash" >> ~/.bashrc && \
  echo "alias cm='cd ~/workspace && catkin_make'" >> ~/.bashrc

WORKDIR /home/ubuntu/workspace/src
RUN /bin/bash -c "source /opt/ros/melodic/setup.bash; catkin_init_workspace" && \
  git clone https://github.com/kakariko3/refro_test.git
WORKDIR /home/ubuntu/workspace
RUN /bin/bash -c "source /opt/ros/melodic/setup.bash; catkin_make"

# RUN mkdir -p /home/ubuntu/workspace/src && \
# 	cd /home/ubuntu/workspace/src && \
#   git clone https://github.com/kakariko3/refro_test.git && \
#   cd /home/ubuntu/workspace && \
#   /bin/bash -c "source /opt/ros/melodic/setup.bash; catkin_make"
