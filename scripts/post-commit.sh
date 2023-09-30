#!/bin/bash

if [[ $(git config --get remote.origin.url) == *"github"* ]]; then
  echo "skipping sync commit since remote is github."
  exit 0
fi

commit_hash=$(git log -1 --format=%h)
commit_message=$(git log -n 1 HEAD --pretty=format:%s)
branch_name=$(git branch --show-current)
repo_name=$(basename -s .git `git config --get remote.origin.url`)

sync_message="hash: $commit_hash $commit_message on branch $branch_name on repo: $repo_name"
syncommit commit -m "$(echo $sync_message)"
exit 0