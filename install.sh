git config --local core.hooksPath .git/myhooks
rm -fr .git/myhooks/
mkdir -p .git/myhooks
cp hooks/* .git/myhooks/

chmod +x .git/myhooks/prepare-commit-msg
