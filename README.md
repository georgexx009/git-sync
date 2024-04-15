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

## TO DO
- [x] Reads a .git-sync file, and looks for var ON=true before running the sync
- [ ] in the .git-sync file, add a var DAYS_TO_RESET=7 to set the days to reset the git history
    - [ ] git command to reset the git history
