
mkdir -p .git/myhooks
curl  -L -o ./.git/myhooks/gommit_windows_amd64.zip https://github.com/klemens-galus/gommit/releases/download/1.0.2/gommit_windows_amd64.zip
tar -xf .git/myhooks/gommit_windows_amd64.zip -C .git/myhooks/
mv .git/myhooks/gommit_windows_amd64.exe .git/myhooks/prepare-commit-msg.exe
rm -f .git/myhooks/*.zip
git config --local core.hooksPath .git/myhooks




# chmod +x .git/myhooks/prepare-commit-msg
