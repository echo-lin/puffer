# puffer
This is a password management tool.

# Overview of Tool Commands：

Usage:
  puffer [command]

Available Commands:
  add         Add a account to system
  completion  Generate the autocompletion script for the specified shell
  del         Delete a account
  help        Help about any command
  list        Show accounts list

Flags:
      --config string   config file (default is $HOME/puffer.yaml)
  -h, --help            help for puffer
  -t, --toggle          Help message for toggle

Use "puffer [command] --help" for more information about a command.


Using this password management tool requires computer administrator privileges, so you will need to enter your password.Otherwise, When executing “puffer add”, the following error occurs：
Error: This command requires root privileges


