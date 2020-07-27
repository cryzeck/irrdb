package irrdb

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

func doquery(connect, args string) ([]string, error) {
	version := "go-irrdb 0.1"

	CONNECT := connect
	c, err := net.Dial("tcp", CONNECT)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Fprintf(c, "!!\n")
	fmt.Fprintf(c, fmt.Sprintf("!n %s\n", version))
	r := bufio.NewReader(c)
	message, _ := r.ReadString('\n')

	if message != "C\n" {
		return nil, errors.New(fmt.Sprintf("Got: %s from server.", message))
	}

	fmt.Fprintf(c, args+"\n")
	for {
		req, _ := r.ReadString('\n')
		if req == "D\n" {
			return nil, errors.New("No prefies received.")
		}
		req, _ = r.ReadString('\n')
		return strings.Split(strings.TrimSuffix(req, "\n"), " "), nil
		log.Fatal(req)
	}
}

func getquery(res, ver string) (string, error) {
	var autnum = regexp.MustCompile(`^AS\d{1,10}$`)
	var asset = regexp.MustCompile(`^AS-[A-Za-z0-9]+$`)

	if autnum.MatchString(res) {
		if ver == "6" {
			return fmt.Sprintf("!6%s", strings.ToLower(res)), nil
		}
		return fmt.Sprintf("!g%s", strings.ToLower(res)), nil
	}
	if asset.MatchString(res) {
		if ver == "6" {
			return fmt.Sprintf("!a6%s", res), nil
		}
		return fmt.Sprintf("!a4%s", res), nil
	}
	return "0", errors.New(fmt.Sprintf("%s: invalid aut-num or as-set.", res))
}

func Query(s, res, ver string) ([]string, error) {

	if !(ver == "4" || ver == "6") {
		return nil, errors.New("IP version need to be 4 or 6")
	}

	query, err := getquery(res, ver)

	if err != nil {
		return nil, err
	}

	return doquery(s, query)
}
