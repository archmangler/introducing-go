#!/usr/bin/bash
#Steps to install golang

wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
sha256sum go1.10.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_profile
source ~/.bash_profile
mkdir -p $HOME/go
ls -l $HOME/go

