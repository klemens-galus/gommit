mkdir -p .git/myhooks
curl  -L -o ./.git/myhooks/gommit_darwin_arm64.tar https://github.com/klemens-galus/gommit/releases/download/1.0.2/gommit_darwin_arm64.tar
tar -xf .git/myhooks/gommit_darwin_arm64.tar -C .git/myhooks/
mv .git/myhooks/gommit_darwin_arm64 .git/myhooks/prepare-commit-msg
git config --local core.hooksPath .git/myhooks




# chmod +x .git/myhooks/prepare-commit-msg
