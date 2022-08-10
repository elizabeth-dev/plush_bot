FROM gitpod/workspace-full

USER gitpod

ENV GOPATH="/workspace/go"

# Oh My Zsh
RUN sudo chsh -s $(which zsh) $(whoami)
RUN sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" "--unattended"
## RUN curl -fsSLo $HOME/.oh-my-zsh/custom/themes/elizabeth.zsh-theme https://raw.github.com/elizabeth-dev/dotfiles/main/elizabeth.zsh-theme
## RUN curl -fsSLo $HOME/.zshrc https://raw.github.com/elizabeth-dev/dotfiles/main/.zshrc

# Zsh autosuggestions
RUN echo 'deb http://download.opensuse.org/repositories/shells:/zsh-users:/zsh-autosuggestions/xUbuntu_20.04/ /' | sudo tee /etc/apt/sources.list.d/shells:zsh-users:zsh-autosuggestions.list
RUN curl -fsSL https://download.opensuse.org/repositories/shells:zsh-users:zsh-autosuggestions/xUbuntu_20.04/Release.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/shells_zsh-users_zsh-autosuggestions.gpg > /dev/null

# Install packages
RUN sudo install-packages zsh-syntax-highlighting zsh-autosuggestions

# AWS CLI
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    sudo ./aws/install

# AWS SAM
RUN curl -fsSL "https://github.com/aws/aws-sam-cli/releases/latest/download/aws-sam-cli-linux-x86_64.zip" -o "awssam.zip" && \
    unzip awssam.zip -d sam-installation && \
    sudo ./sam-installation/install
