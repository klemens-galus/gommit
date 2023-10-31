os=$(uname -a | grep -e Linux -e MINGW64)

echo $os

if [ -z "${os}" ]; then
    echo "OS is not available"
    exit 1
fi

if [ -z "$(uname -a | grep -e Linux)" ]; then
    mkdir -p .git/myhooks
    curl  -L -o ./.git/myhooks/prepare-commit-msg.zip https://github.com/klemens-galus/gommit/releases/download/1.0.2/gommit_linux_amd64.zip
    tar -xf .git/myhooks/gommit_windows_amd64.zip -C .git/myhooks/
    chmod +x .git/myhooks/prepare-commit-msg
    git config --local core.hooksPath .git/myhooks
fi 

if [ -z "$(uname -a | grep -e MINGW64)" ]; then
    mkdir -p .git/myhooks
    curl  -L -o ./.git/myhooks/gommit_windows_amd64.zip https://github.com/klemens-galus/gommit/releases/download/1.0.2/gommit_windows_amd64.zip
    tar -xf .git/myhooks/gommit_windows_amd64.zip -C .git/myhooks/
    mv .git/myhooks/gommit_windows_amd64.exe .git/myhooks/prepare-commit-msg.exe
    rm -f .git/myhooks/*.zip
    git config --local core.hooksPath .git/myhooks
fi 



# chmod +x .git/myhooks/prepare-commit-msg
