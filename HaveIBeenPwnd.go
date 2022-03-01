package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    inFile, err := os.Open("emails.txt")
    if err != nil {
        fmt.Println(err.Error() + `: ` + "emails.txt")
        return
    }
    defer inFile.Close()

    scanner := bufio.NewScanner(inFile)
    for scanner.Scan() {
        call("https://haveibeenpwned.com/api/v3/breachedaccount/"+scanner.Text(), "GET")
        time.Sleep(2 * time.Second)
    }

}
func call(url, method string) {
    client := &http.Client{
        Timeout: time.Second * 10,
    }
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        //return fmt.Errorf("Got error %s", err.Error())
    }
    req.Header.Set("accept", "application/json")
    req.Header.Add("hibp-api-key", "<YOUR_KEY>")
    response, err := client.Do(req)
    if err != nil {
        //return fmt.Errorf("Got error %s", err.Error())
    }

    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)

    if err != nil {
        log.Fatal(err)
        //return fmt.Errorf("Got error %s", err.Error())

    }

    fmt.Println(string(body))
    //return fmt.Errorf("Got error %s", err.Error())
}
