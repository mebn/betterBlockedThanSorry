package blocker

import (
	"fmt"
	"os"
	"strings"
)

type EtcHosts struct {
	blocklist []string
	filename  string
}

func NewEtcHosts(filename string, blocklist []string) EtcHosts {
	return EtcHosts{
		blocklist: generateBlocklist(blocklist),
		filename:  filename,
	}
}

func generateBlocklist(blocklist []string) []string {
	res := []string{}

	for _, entry := range blocklist {
		res = append(res, fmt.Sprintf("127.0.0.1 %s", entry))
		res = append(res, fmt.Sprintf("127.0.0.1 www.%s", entry))
		res = append(res, fmt.Sprintf("::1 %s", entry))
		res = append(res, fmt.Sprintf("::1 www.%s", entry))
	}

	return res
}

func (e *EtcHosts) AddBlock() {
	var str strings.Builder

	for _, entry := range e.blocklist {
		str.WriteString(fmt.Sprintf("\n%s", entry))
	}

	file, _ := os.OpenFile(e.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(str.String())
}

func (e *EtcHosts) RemoveBlock() {
	contentsB, _ := os.ReadFile(e.filename)
	contents := string(contentsB)

	for _, entry := range e.blocklist {
		contents = strings.ReplaceAll(contents, fmt.Sprintf("\n%s", entry), "")
	}

	file, _ := os.OpenFile(e.filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(contents)
}

func (e *EtcHosts) IsTamperedWith() bool {
	set := make(map[string]struct{})
	for _, value := range e.blocklist {
		set[value] = struct{}{}
	}

	contentsB, _ := os.ReadFile(e.filename)
	contents := string(contentsB)

	for _, line := range strings.Split(contents, "\n") {
		delete(set, line)
	}

	return len(set) != 0
}
