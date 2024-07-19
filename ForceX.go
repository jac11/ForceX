package main

import (
   // "bufio"
    "flag"
    "fmt"
//   "log"
//    "net"
  //  "os"
    //"strings"
   // "github.com/jlaffaye/ftp"
)

type ArgVar struct {
    Address   string
    Protocol  string
    UserList  string
    PassList  string
    Pass      string
    Port      string
    User      string
    Sleep     string
    Wordpress string
}

func (Argv *ArgVar) ParseArgs() {
    flag.StringVar(&Argv.Address, "Address", "", "IP/Domain Target BruteForce")
    flag.StringVar(&Argv.Protocol, "Protocol", "", "Protocol 'http,ssh,ftp..etc'")
    flag.StringVar(&Argv.UserList, "UserList", "", "List of users names")
    flag.StringVar(&Argv.PassList, "PassList", "", "Password list")
    flag.StringVar(&Argv.User, "User", "", "single username")
    flag.StringVar(&Argv.Pass, "Pass", "", "single Password")
    flag.StringVar(&Argv.Port, "Port", "", "Service Port")
    flag.StringVar(&Argv.Sleep, "Sleep", "", "Time to Sleep to wait between Requests")
    flag.StringVar(&Argv.Wordpress, "Wordpress", "", "BruteForce Wordpress login Page")
    flag.Parse()
}


func main() {
    Banner:=Logo()
    fmt.Println(Banner)
    var ArgVar ArgVar
    ArgVar.ParseArgs()
    user, pass := ArgVar.FTPConnect()
    if user != "" && pass != "" {
        fmt.Println("Successful login:", user, pass)
    } else {
        fmt.Println("No successful login found.")
    }
}
