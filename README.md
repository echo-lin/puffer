# puffer
This is a password management tool.

# Overview of Tool Commands

```markdown
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
```


Using this password management tool requires computer administrator privileges, so you will need to enter your password.Otherwise, When executing “puffer add”, the following error occurs：
```markdown
lin@lindeMacBook-Pro puffer % puffer add
Error: This command requires root privileges
```

## Retrieve Account List
```markdown
 sh-3.2# puffer list
+-------+--------+------+----------+
| 归属  | 用户名 | 密码 |   备注   |
+-------+--------+------+----------+
| a.com | lin2   | 123  | test |
| a.com | lin20  | 123  | test |
+-------+--------+------+----------+
```

## Add a account list
```markdown
sh-3.2# puffer add -e 'google accout' -d 'google.com' -p '123456' -u 'google1'
success
```

```markdown
sh-3.2# puffer list
+------------+---------+--------+---------------+
|    归属    | 用户名  |  密码  |     备注      |
+------------+---------+--------+---------------+
| a.com      | lin2    | 123    | test      |
| a.com      | lin20   | 123    | test      |
| google.com | google1 | 123456 | google accout |
+------------+---------+--------+---------------+
```

## Delete a account
```markdown
sh-3.2# puffer del -d 'a.com' -u 'lin20'
success
```

```markdown
sh-3.2# puffer list
+------------+---------+--------+---------------+
|    归属    | 用户名  |  密码  |     备注      |
+------------+---------+--------+---------------+
| a.com      | lin2    | 123    | test      |
| google.com | google1 | 123456 | google accout |
+------------+---------+--------+---------------+
```

