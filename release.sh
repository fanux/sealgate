git add .
git commit -m $1
git tag $2
git push && git push --tags
