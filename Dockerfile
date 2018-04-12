FROM node:8

#
# very simple way to run `les` without installing it
#
# $ docker build . -t les
#
# $ docker run -v $(pwd):/work les les convert
# $ docker run -v $(pwd):/work les les-node -b
#

WORKDIR /work

RUN curl -L https://github.com/Adaptech/letseventsource/blob/master/releases/les/0.10.0/les-Linux-x86_64?raw=true -o /usr/local/bin/les
RUN chmod +x /usr/local/bin/les

RUN curl -L https://github.com/Adaptech/letseventsource/blob/master/releases/les-node/0.10.0/les-node-Linux-x86_64?raw=true -o /usr/local/bin/les-node
RUN chmod +x /usr/local/bin/les-node

