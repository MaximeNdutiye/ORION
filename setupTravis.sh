curl https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip > terraform_0.11.7_linux_amd64.zip && \
unzip terraform_0.11.7_linux_amd64.zip -d . && \
export PATH="$PATH:$HOME/gopath/src/github.com/MaximeNdutiye/ORION/terraform" && \
rm -f terraform_0.11.7_linux_amd64.zip && echo $PATH