FROM ubuntu:latest
RUN apt update && apt install  openssh-server sudo -y
RUN groupadd leonatone && useradd -ms /bin/bash -g leonatone leonatone -u 1024
RUN mkdir /var/jfs
RUN mkdir -p /home/leonatone/.ssh
COPY id_rsa.pub /home/leonatone/.ssh/authorized_keys
RUN chown leonatone:leonatone /home/leonatone/.ssh/authorized_keys && chmod 600 /home/leonatone/.ssh/authorized_keys
RUN service ssh start
EXPOSE 22
CMD ["/usr/sbin/sshd","-D"]