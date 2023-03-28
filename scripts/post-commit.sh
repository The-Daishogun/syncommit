#!/bin/bash

commit_message=$(git log -n 1 HEAD --pretty=format:%s)
branch_name=$(git branch --show-current)
repo_name=$(basename -s .git `git config --get remote.origin.url`)

echo $commit_message
echo $branch_name
echo $repo_name

syncmessage="commit: $commit_message on branch $branch_name on repo: $repo_name"
syncommit commit -m "$(echo $syncmessage)"
exit 0