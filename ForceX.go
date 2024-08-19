package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
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
    HTMLUSER  string
    HTMLPASS  string
    Cookies   string
    URL       string
}

// ParseArgs parses command-line arguments into the ArgVar struct
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
    flag.StringVar(&Argv.HTMLUSER, "HTMLUSER", "", "User Field In HTML")
    flag.StringVar(&Argv.HTMLPASS, "HTMLPASS", "", "Password Field In HTML")
    flag.StringVar(&Argv.URL, "URL", "", "Url Login Page admin/..etc")
    flag.StringVar(&Argv.Cookies, "Cookies", "", "Web Cookies")
    flag.Parse()
}

func printInputInfo(Argv *ArgVar) {
    fmt.Println("Input Info:")
    fmt.Println(strings.Repeat("=", 15))
    if Argv.Address != "" {
        fmt.Println("🌏 TargetIP         -----------| > ", Argv.Address)
    }
    if Argv.URL != "" {
        fmt.Println("🌏 WordpressURL     -----------| > ", Argv.URL)
    }
    fmt.Println("🌏 Protocol         -----------| > ", Argv.Protocol)
    if Argv.User != "" {
        fmt.Println("🌏 TargetUser       -----------| > ", Argv.User)
    }
    if Argv.Pass != "" {
        fmt.Println("🌏 TargetPassword   -----------| > ", Argv.Pass)
    }
    if Argv.UserList != "" {
        fmt.Println("🌏 UserList         -----------| > ", Argv.UserList)
    }
    if Argv.PassList != "" {
        fmt.Println("🌏 PasswordList     -----------| > ", Argv.PassList)
    }
    if Argv.Port != "" {
        fmt.Println("🌏 ConnectionPort   -----------| > ", Argv.Port)
    }
    fmt.Println(strings.Repeat("=", 40))
}

func main() {
    var ArgVar ArgVar

    Banner := Logo()
    fmt.Println(Banner)
    ArgVar.ParseArgs()
    CheckArg(&ArgVar)
    printInputInfo(&ArgVar)

    switch ArgVar.Protocol {
    case "ftp":
        user, pass := ArgVar.FTPConnect()
        if user != "" && pass != "" {
            fmt.Println("FTP Successful login:", user, pass)
        } else {
            fmt.Println("No successful login found.")
        }
    case "ssh":
        user, pass := ArgVar.SSHConnect()
        if user != "" && pass != "" {
            fmt.Println("SSH Successful login:", user, pass)
        } else {
            fmt.Println("No successful login found.")
        }
    case "wordpress", "http":
        user, pass := ArgVar.WordpressLogin()
        if user != "" && pass != "" {
            fmt.Println("Wordpress Successful login:", user, pass)
        } else {
            fmt.Println("No successful login found.")
        }
    default:
        fmt.Println("Usage --Protocol [ssh-ftp-wordpress]")
        os.Exit(1)
    }
}
