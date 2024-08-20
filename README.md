# ForceX
## BruteForce Attack Tool

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
| `--Protocol`  | Protocol (`http`, `ssh`, `ftp`, etc.)  |
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

#### FTP Brute Force
```bash
../ForceX  --Address 192.168.1.1 --Protocol ftp --UserList users.txt --PassList passwords.txt --Port 21 
```

#### SSH Brute Force
```bash
./ForceX  --Address 192.168.1.2 --Protocol ssh --UserList users.txt --PassList passwords.txt --Port 22 
```

#### WordPress Brute Force
```bash
./ForceX --URL http://example.com/wp-login.php --Protocol wordpress --UserList users.txt --PassList passwords.txt 
```

## Output

The tool provides detailed output for each attempt, including:

- Target IP and Protocol
- User and Password Lists or Single Credentials
- Connection Port
- Success or Failure of login attempts

Successful login details are clearly displayed with user-friendly icons.

## Notes

- This tool is designed for educational purposes only. Unauthorized access to computer systems is illegal and unethical.
- Use responsibly and ensure you have permission to test the systems you target.


## Contributing

Contributions are welcome! Please fork the repository and create a pull request.

## Contact

For issues, questions, or suggestions, please create an issue in the repository.
