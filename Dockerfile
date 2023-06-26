FROM ubuntu

RUN apt-get update && apt-get install -y --no-install-recommends \ 
--no-install-suggests wget vim git sudo make gcc g++ \ 
python3 python3-pip golang npm 

WORKDIR /home
