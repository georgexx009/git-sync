# git sync

This is a golang program to sync a git repository git a command.
Every N days, it will push force to reset the git history.

WARNING: This is not a tool for applications. It's a tool for notes.

## Install
Run the next shell
```
# first time
chmod +x install.sh

./install.sh
```

Add this to your `.zshrc`
```
# git-sync
alias gs='~/.config/bin/git-sync'
```

