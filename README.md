# Syncommit

Some developers work on self-hosted instances of gitlab, bitbucket, etc. so they end up with a empty contribution graph on github. This is cli app that inserts two git hooks that will commit an empty commit to a private repo on github whenever you commit to your local repo. This will make your contribution graph look more active.

## Installation

download the latest release on the releases page and extract it to a folder. put the folder on your `$PATH` and run `syncommit init` in the root of your git repo. enter the ssh url of your empty private git repo when prompted. you can also run `syncommit init <url>` to skip the prompt (coming soon...).
after the setup, it should automatically commit an empty commit to your private repo whenever you commit to your local repo. and push all the commits when you push.