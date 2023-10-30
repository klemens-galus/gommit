os=$(uname -a | grep -e Linux -e MINGW64)

echo $os

if [ -z "${os}" ]; then
    echo "OS is not available"
    exit 1
fi

if [ -z "$(uname -a | grep -e Linux)" ]; then
    mkdir -p .git/myhooks
    curl https://github.com/klemens-galus/gommit/releases/download/1.0.0/gommit_linux_amd64 -o .git/myhooks/prepare-commit-msg
    chmod +x .git/myhooks/prepare-commit-msg
fi 

if [ -z "$(uname -a | grep -e MINGW64)" ]; then
    mkdir -p .git/myhooks
    curl https://github.com/klemens-galus/gommit/releases/download/1.0.0/gommit_windows_amd64 -o .git/myhooks/prepare-commit-msg
    chmod +x .git/myhooks/prepare-commit-msg
fi 


# chmod +x .git/myhooks/prepare-commit-msg
