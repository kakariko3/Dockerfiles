FROM dorowu/ubuntu-desktop-lxde-vnc:bionic-lxqt

# Install ROS Melodic
RUN \
  apt-get update -qq && apt-get install -y --no-install-recommends \
	  gnupg && \
  sh -c 'echo "deb http://packages.ros.org/ros/ubuntu $(lsb_release -sc) main" > /etc/apt/sources.list.d/ros-latest.list' && \
  apt-key adv --keyserver 'hkp://keyserver.ubuntu.com:80' --recv-key C1CF6E31E6BADE8868B172B4F42ED6FBAB17C654 && \
  apt-get update && \
  apt-get install -y \
	  ros-melodic-desktop-full && \
  echo "source /opt/ros/melodic/setup.bash" >> /root/.bashrc && \
  echo "source ~/workspace/devel/setup.bash" >> /root/.bashrc && \
  echo "alias cm='cd ~/workspace && catkin_make'" >> /root/.bashrc && \
  apt-get install -y \
    python-rosdep \
    python-rosinstall \
    python-rosinstall-generator \
    python-wstool \
    build-essential &&\
  rosdep init && \
  rosdep update

# Install other packages
RUN \
  apt-get install -y --no-install-recommends \
    ros-melodic-move-base \
    ros-melodic-slam-gmapping \
    ros-melodic-mouse-teleop \
    ros-melodic-navigation \
    python-pip \
    python3-pip \
    git \
    nano && \
  pip3 install \
    rospkg \
    catkin_pkg

# # Create workspace & catkin_make
# WORKDIR /root/workspace/src
# RUN \
#   /bin/bash -c "source /opt/ros/melodic/setup.bash; catkin_init_workspace" && \
#   git clone https://github.com/haloworld-works/refro_dandy.git
# WORKDIR /root/workspace
# RUN \
#   /bin/bash -c "source /opt/ros/melodic/setup.bash; catkin_make"
# WORKDIR /root
