package hw10programoptimization

import (
	"bufio"
	"fmt"
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
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (users, error) {
	var result users
	var err error

	buffer := bufio.NewReader(r)
	userCount := 0
	lastLine := false
	for !lastLine {
		var line []byte

		line, err = buffer.ReadBytes(10)
		if err == io.EOF {
			lastLine = true
		} else if err != nil {
			return result, nil
		}

		var user User
		if err = easyjson.Unmarshal(line, &user); err != nil {
			return result, err
		}
		result[userCount] = user
		userCount++
	}

	return result, nil
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {

		if strings.HasSuffix(user.Email, "."+domain) {
			d := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[d]++
		}
	}
	return result, nil
}
