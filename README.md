## what is it ?

Guardian-shell is a  secure reverse shell tool written in Go. Designed with a focus on responsible use and ethical practices, it facilitates secure communication between authorized parties. Guardian-shell empowers you with the ability to remotely execute commands on target machines, for easy system management and monitoring. 


## usage

to add new user in remote linux server
```
echo ur_sudo_password | sudo -S useradd testuser; echo new_password_123 | passwd testuser
```

to delete user from remote linux server
```
echo ur_sudo_password | sudo -S userdel testuser
```