FROM mcr.microsoft.com/azure-cli
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin

RUN ["/bin/bash", "-c", "source /usr/share/bash-completion/bash_completion >> /root/.bashrc"]
RUN ["/bin/bash", "-c", "source <(kubectl completion bash) >> /root/.bashrc"]
RUN /bin/bash