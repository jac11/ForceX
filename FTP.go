package main

import (
    "bufio"
    "log"
    "net"
    "os"
    "github.com/jlaffaye/ftp"
)

func (ArgVar *ArgVar) FTPConnect() (string, string) {
    var UserName []string
    var PassWords []string

    if ArgVar.User != "" {
        UserName = append(UserName, ArgVar.User)
    } else if ArgVar.UserList != "" {
        file, err := os.Open(ArgVar.UserList)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            UserName = append(UserName, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }

    if ArgVar.Pass != "" {
        PassWords = append(PassWords, ArgVar.Pass)
    } else if ArgVar.PassList != "" {
        file, err := os.Open(ArgVar.PassList)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            PassWords = append(PassWords, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }
    var user, pass string
    DomainNet := net.JoinHostPort(ArgVar.Address, ArgVar.Port)

    //argsuser-argspass
    if ArgVar.PassList  == "" && ArgVar.UserList == ""{
        connect, err := ftp.Dial(DomainNet)
        err = connect.Login(ArgVar.User,ArgVar.Pass)
        if err == nil {
            return ArgVar.User ,ArgVar.Pass
        }

    //username -passwordlist    
    }else if ArgVar.User !="" && ArgVar.PassList !=""{
        for _, pass = range PassWords {
            connect, err := ftp.Dial(DomainNet)
            if err != nil {
                continue
            }
            defer connect.Quit()
            err = connect.Login(ArgVar.User, pass)
            if err == nil {
                return ArgVar.User, pass
            }
        }

    //userlist-argpass  
    }else if ArgVar.Pass !="" && ArgVar.UserList!=""{
        for _, user = range UserName {
            connect, err := ftp.Dial(DomainNet)
            if err != nil {
                continue
            }
            defer connect.Quit()
            err = connect.Login(user,ArgVar.Pass)
            if err == nil {
                return user, ArgVar.Pass
            }
        }
    //userlist-passwordlist
    } else if ArgVar.PassList != "" && ArgVar.UserList != "" {
        for _, user := range UserName {
            for _, pass := range PassWords {
                connect, err := ftp.Dial(DomainNet)
                if err != nil {
                    continue
                }
                
                defer connect.Quit()
                err = connect.Login(user, pass)
                if err == nil {
                    return user, pass
                }
            
            }
        }
    }

    return "", ""
}