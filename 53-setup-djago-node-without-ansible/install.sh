#!/usr/bin/env bash


export DEBIAN_FRONTEND=noninteractive


function update_apt_cache() {
    apt-get update
}


function install_common_packages() {
    apt-get install --yes rsync mc vim git curl tmux htop iotop make strace sudo gettext
}


function install_infra_packages() {
    apt-get install --yes postgresql postgresql-contrib nginx
}


function update_bash_rc() {

cat > ~/.bashrc <<'EOF'
    export EDITOR=vim
    export GREP_OPTIONS="--color=auto"

    export LS_OPTIONS='--color=auto'
    eval "`dircolors`"
    alias ls='ls $LS_OPTIONS'
    alias ll='ls $LS_OPTIONS -l'
    alias l='ls $LS_OPTIONS -lA'

    alias rm='rm -i'
    alias cp='cp -i'
    alias mv='mv -i'

    export PS1="\\[\\033[31;1m\\]\\h\\[\\033[0;32m\\] \\w\\[\\033[00m\\]: "

    export PATH="/root/.pyenv/bin:$PATH"
    eval "$(pyenv init -)"
    eval "$(pyenv virtualenv-init -)"
EOF
    source ~/.bashrc
}


function install_python_packages() {
    apt-get install --yes python3 python3-pip

    apt-get install --yes libpq-dev libsqlite3-dev libbz2-dev libreadline-dev libjpeg-dev

    if [[ ! -d "~/.pyenv" ]]; then
        curl --silent --fail -L https://raw.githubusercontent.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash

        update_bash_rc
    fi

    if [[ ! -d "~/.pyenv/versions/3.5.1" ]]; then
        pyenv install 3.5.1
        pyenv global 3.5.1
    fi

    pip3 install --upgrade pip
    pip3 install --upgrade tox
}


function configure_git_deploy_keys() {
    mkdir -pv ~/.ssh

    cat > ~/.ssh/deploy-key <<'EOF'
-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----
EOF
    chmod 0600 ~/.ssh/deploy-key

    cat > ~/.ssh/config <<'EOF'
Host *
IdentitiesOnly yes

HostName bitbucket.org
IdentityFile ~/.ssh/deploy-key
StrictHostKeyChecking no
EOF
}


function fetch_project_sources() {
    git_repo='git@github.com:agalitsyn/play.git'
    git clone --depth=1 $git_repo $PROJECT_FOLDER || git -C $PROJECT_FOLDER pull
}


function install_project_dependencies() {
    pip3 install -r $APP_FOLDER/requirements.txt
    pip3 install -r $APP_FOLDER/dev-requirements.txt

    sudo -u postgres psql -c "CREATE DATABASE test;"
    sudo -u postgres psql -c "CREATE USER test WITH PASSWORD 'test'; GRANT ALL PRIVILEGES ON DATABASE test TO test;"
    sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE test TO test;"

    make -C $APP_FOLDER setup-app
}


function create_project_unit_file() {
    cp $APP_FOLDER/etc/test.service /etc/systemd/system/test.service
    mkdir -pv /var/log/test
}


function create_nginx_config() {
    rm /etc/nginx/sites-available/default

    cp $APP_FOLDER/etc/nginx.conf /etc/nginx/sites-available/test
    ln -s /etc/nginx/sites-available/test /etc/nginx/sites-enabled

    nginx -t
}


function reload_unit_files() {
    systemctl daemon-reload
}


function enable_and_start_unit() {
    local service=${1:?"usage: enable_and_start_unit <service>"}

    systemctl enable "$service"
    systemctl restart "$service"
    systemctl status "$service"
}


# Constants

PROJECT_FOLDER="/opt/test"
APP_FOLDER="$PROJECT_FOLDER/src/app"


# Main logic

update_apt_cache

install_common_packages
install_infra_packages
install_python_packages

configure_git_deploy_keys
fetch_project_sources
install_project_dependencies

create_project_unit_file
create_nginx_config

reload_unit_files

enable_and_start_unit "postgresql"
enable_and_start_unit "test"
enable_and_start_unit "nginx"
