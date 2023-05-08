#!/bin/bash

if [[ $(git config --get remote.origin.url) == *"github"* ]]; then
  echo "skipping sync commit push since remote is github."
  exit 0
fi

syncommit push

exit 0