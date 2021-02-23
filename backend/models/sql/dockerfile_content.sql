INSERT INTO `dockerfile_tree`.`dockerfile`(`id`, `dockerfile`, `level_id`) VALUES (1, 'FROM hank997/dockerfile-tree:base-scripts as base-scripts
FROM nvidia/cuda:10.0-base-ubuntu16.04

ENV CUDA_MAJOR_VERSION 10
ENV CUDA_MINOR_VERSION 0

RUN 
    sed -i \'s/archive.ubuntu.com/mirrors.aliyun.com/g\' /etc/apt/sources.list && 
    sed -i \'s/security.ubuntu.com/mirrors.aliyun.com/g\' /etc/apt/sources.list

# 包含cuda编译环境
ENV NCCL_VERSION 2.5.6

RUN mkdir -p /usr/local/src
COPY --from=base-scripts /clean-layer.sh  /usr/bin/clean-layer.sh
COPY --from=base-scripts /remove-cuda-static.sh /usr/local/src/
RUN chmod a+rwx /usr/bin/clean-layer.sh

RUN apt-get update && apt-get install -y --no-install-recommends 
        cuda-libraries-${CUDA_PKG_VERSION} 
        cuda-nvtx-${CUDA_PKG_VERSION} 
        libnccl2=${NCCL_VERSION}-1+cuda${CUDA_MAJOR_VERSION}.${CUDA_MINOR_VERSION} 
    && apt-mark hold libnccl2 
    && apt-get update && apt-get install -y --no-install-recommends 
        cuda-nvml-dev-${CUDA_PKG_VERSION} 
        cuda-command-line-tools-${CUDA_PKG_VERSION} 
        cuda-libraries-dev-${CUDA_PKG_VERSION} 
        cuda-minimal-build-${CUDA_PKG_VERSION} 
        libnccl-dev=${NCCL_VERSION}-1+cuda${CUDA_MAJOR_VERSION}.${CUDA_MINOR_VERSION} 
    && bash /usr/local/src/remove-cuda-static.sh 
    && clean-layer.sh

ENV LIBRARY_PATH /usr/local/cuda/lib64/stubs

# Downgrade cudnn to 7.6.5 for tensorflow 2.0 compatibility
ENV CUDNN_VERSION 7.6.5.32
RUN 
    mkdir -p /usr/local/src && 
    apt-get update && 
    apt-get install -y --no-install-recommends wget && 
    # apt-get install -y --no-install-recommends apt-utils && 
    # apt-mark unhold libcudnn7 && 
    apt-get -y --allow-downgrades install libcudnn7=${CUDNN_VERSION}-1+cuda${CUDA_MAJOR_VERSION}.${CUDA_MINOR_VERSION} && 
    apt-mark hold libcudnn7 && 
    ln -s /usr/lib/x86_64-linux-gnu/libcudnn.so.7 /usr/lib/x86_64-linux-gnu/libcudnn.so && 
    # Install cudnn header
    cd /usr/local/src && wget https://developer.download.nvidia.cn/compute/machine-learning/repos/ubuntu1604/x86_64/libcudnn7-dev_${CUDNN_VERSION}-1+cuda${CUDA_MAJOR_VERSION}.${CUDA_MINOR_VERSION}_amd64.deb && dpkg -x libcudnn7-dev_${CUDNN_VERSION}-1+cuda${CUDA_MAJOR_VERSION}.${CUDA_MINOR_VERSION}_amd64.deb ./ && mv usr/include/x86_64-linux-gnu/cudnn_v7.h /usr/local/cuda/include/cudnn.h && 
    clean-layer.sh
ENV LD_LIBRARY_PATH=/usr/local/cuda/lib64:/usr/local/cuda/lib64/stubs:$LD_LIBRARY_PATH
ENV PKG_CONFIG_PATH=${PKG_CONFIG_PATH}:/usr/local/lib/pkgconfig
ENV DEBIAN_FRONTEND noninteractive

# 配置语言环境
RUN apt-get update 
    && apt-get -y install locales 
    && locale-gen en_US.UTF-8 
    && mv /etc/localtime /etc/localtime.bak 
    && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 
    && clean-layer.sh
ENV LANG=en_US.UTF-8 LANGUAGE=en_US.UTF-8

# 基础软件包
RUN 
    apt-get update && 
    apt-get -y install build-essential wget gcc make unzip zip g++ vim openssh-server libgoogle-glog-dev && 
    # Install latest cmake
    cd /usr/local/src && wget -q https://github.com/Kitware/CMake/releases/download/v3.15.6/cmake-3.15.6-Linux-x86_64.sh && bash cmake-3.15.6-Linux-x86_64.sh --prefix=/usr/local/ --skip-license && 
    clean-layer.sh && 
    rm -f /etc/apt/sources.list.d/cuda.list /etc/apt/sources.list.d/nvidia-ml.list

ENV IMAGE_TAG_INFO ""', 13);

INSERT INTO `dockerfile_tree`.`dockerfile`(`id`, `dockerfile`, `level_id`) VALUES (2, '## Python 3.6
FROM cuda-image
COPY get-pip.py /usr/local/src/get-pip.py
RUN \
    sed -i \'s/mirrors.aliyun.com/mirrors.cloud.tencent.com/g\' /etc/apt/sources.list && \
    sed -i \'s/mirrors.aliyun.com/mirrors.cloud.tencent.com/g\' /etc/apt/sources.list
RUN \
    apt-get update && \
    apt-get install -y software-properties-common && \
    add-apt-repository -y  ppa:deadsnakes/ppa && \
    apt-get update && \
    apt-get install -y python3.6 python3.6-dev && \
    # 设置默认python版本链接
    update-alternatives --install /usr/bin/python python /usr/bin/python3.6 2 && \
    update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.6 2 && \
    update-alternatives --set python /usr/bin/python3.6 && \
    update-alternatives --set python3 /usr/bin/python3.6 && \
    # pip, wheel, setuptools
    python3.6 /usr/local/src/get-pip.py && \
    clean-layer.sh',14);