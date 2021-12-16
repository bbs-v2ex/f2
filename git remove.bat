https://blog.csdn.net/lizhijian21/article/details/94026307
git rev-list --all --objects | grep "$(git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -n 3 | awk -F ' '  '{print $1}')"


git filter-branch --force --index-filter "git rm -rf --cached --ignore-unmatch KKIP/0002-1.wav" --prune-empty --tag-name-filter cat -- --all


git stash

git push origin master --force

rm -rf .git/refs/original/
git reflog expire --expire=now --all
git gc --prune=now