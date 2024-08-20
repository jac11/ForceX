package main

import (
    "bufio"
    "log"
    "net"
    "os"
    "github.com/jlaffaye/ftp"
    "fmt"
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

    if ArgVar.PassList  == "" && ArgVar.UserList == ""{
        connect, err := ftp.Dial(DomainNet)
        err = connect.Login(ArgVar.User,ArgVar.Pass)
        if err == nil {
            fmt.Print("\033[G\033[K")
            return ArgVar.User ,ArgVar.Pass
        }else{
            fmt.Print("\033[G\033[K")
            fmt.Printf("❌ User Can not Login   USERNAME %s  PASSWORD  %s",  ArgVar.User ,ArgVar.Pass)
        }
    
    }else if ArgVar.User !="" && ArgVar.PassList !=""{
        for _, pass = range PassWords {
            connect, err := ftp.Dial(DomainNet)
            defer connect.Quit()
            err = connect.Login(ArgVar.User, pass)
            if err == nil {
                fmt.Print("\033[G\033[K")
                return ArgVar.User, pass
            }else{
                fmt.Print("\033[G\033[K")
                fmt.Printf("❌ User Can not Login   USERNAME %s  PASSWORD  %s", ArgVar.User, pass)
            }
        }

  
    }else if ArgVar.Pass !="" && ArgVar.UserList!=""{
        for _, user = range UserName {
            connect, err := ftp.Dial(DomainNet)
            defer connect.Quit()
            err = connect.Login(user,ArgVar.Pass)
            if err == nil {
                fmt.Print("\033[G\033[K")
                return user, ArgVar.Pass
            }else{
                  fmt.Print("\033[G\033[K")
                  fmt.Printf("❌ User Can not Login   USERNAME %s  PASSWORD  %s", user, ArgVar.Pass)
                }
        }

    } else if ArgVar.PassList != "" && ArgVar.UserList != "" {
        for _, user := range UserName {
            for _, pass := range PassWords {
                connect, err := ftp.Dial(DomainNet)
                defer connect.Quit()
                err = connect.Login(user, pass)
                if err == nil {
                    fmt.Print("\033[G\033[K")
                    return user, pass
                }else{
                  fmt.Print("\033[G\033[K")
                  fmt.Printf("❌ User Can not Login   USERNAME %s  PASSWORD  %s", user,pass)
                }
            
            }
        }
    }

    return "", ""
}