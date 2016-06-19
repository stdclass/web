package main

import "os"
import "os/exec"
import "fmt"
import "io/ioutil"
import "syscall"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    if len(os.Args) < 3{
        fmt.Println("web(1) requires 2 or 3 arguments ( ./web domain /path [optional: domainAlias] )")
        os.Exit(1)
    }

    domain := os.Args[1]
    directory := os.Args[2]
    domainAlias := ""

    if len(os.Args) == 4 {
        domainAlias = os.Args[3]
    }

    var contentAsString = fmt.Sprintf(`
<VirtualHost *:80>

    ServerAdmin phillip@dornauer.cc
    ServerName %s
    ServerAlias %s
    DocumentRoot %s

    <Directory />
        AllowOverride All
    </Directory>
    <Directory %s>
        AllowOverride All
    </Directory>

    ErrorLog /var/log/apache2/%s.log
    LogLevel error
    CustomLog /var/log/apache2/%s.log custom

</VirtualHost>
`, domain, domainAlias, directory, directory, domain, domain)

    content := []byte(contentAsString)

    err := ioutil.WriteFile("/etc/apache2/sites-available/" + domain + ".conf", content, 0644)
    check(err)

    execCmd("a2ensite", []string{domain})
    execCmd("service", []string{"apache2", "reload"})
}

func execCmd(command string, args []string) {
    binary, lookErr := exec.LookPath(command)
    check(lookErr)

    env := os.Environ()
    args = append([]string{command}, args...)

    execErr := syscall.Exec(binary, args, env)
    check(execErr)
}
