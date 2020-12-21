# hank997/dockerfile-tree:base-scripts

FROM busybox 
WORKDIR /
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py 
COPY clean-layer.sh remove-cuda-static.sh /
    