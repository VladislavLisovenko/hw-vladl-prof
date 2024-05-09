package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	easyjson "github.com/mailru/easyjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	domainStat := make(DomainStat)
	buffer := bufio.NewReader(r)
	lastLine := false
	for !lastLine {
		var line []byte

		line, err := buffer.ReadBytes(10)
		if err == io.EOF {
			lastLine = true
		} else if err != nil {
			return nil, err
		}

		var user User
		if err = easyjson.Unmarshal(line, &user); err != nil {
			return nil, err
		}
		email := user.Email

		if strings.HasSuffix(email, "."+domain) {
			emailParts := strings.SplitN(email, "@", 2)
			if len(emailParts) > 1 {
				d := strings.ToLower(emailParts[1])
				domainStat[d]++
			}
		}
	}

	return domainStat, nil
}
