package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-demo/m3u8/tool"
	"net/url"
	"os"
	"regexp"
	"strings"
)
const (
	CryptMethodAES CryptMethod = "AES-128"
	CryptMethodNONE CryptMethod = "NONE"
)

var lineParameterPattern = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)
type CryptMethod string

type Key struct {
	URI string
	IV  string
	key    string
	Method CryptMethod
}

type Segment struct {
	URI string
	Key *Key
}

type Result struct {
	URL *url.URL
	M3u8 *M3u8
	Keys map[*Key]string
}

type M3u8 struct {
	Segments           []*Segment
	MasterPlaylistURIs []string
}

func parseLines(lines []string) (*M3u8,error)  {
	var (
		i       = 0
		lineLen = len(lines)
		m3u8    = &M3u8{}

		key *Key
		seg *Segment
	)
	for ; i < lineLen; i++ {
		line := strings.TrimSpace(lines[i])
		if i == 0 {
			if "#EXTM3U" != line {
				return nil, fmt.Errorf("invalid m3u8, missing #EXTM3U in line 1")
			}
			continue
		}
		switch {
		case line == "":
			continue
		// Master playlist 解析
		case strings.HasPrefix(line, "#EXT-X-STREAM-INF:"):
			i++
			m3u8.MasterPlaylistURIs = append(m3u8.MasterPlaylistURIs, lines[i])
			continue
		// TS URI 解析
		case !strings.HasPrefix(line, "#"):
			seg = new(Segment)
			seg.URI = line
			m3u8.Segments = append(m3u8.Segments, seg)
			seg.Key = key
			continue
		// 解密秘钥解析
		case strings.HasPrefix(line, "#EXT-X-KEY"):
			params := parseLineParameters(line)
			if len(params) == 0 {
				return nil, fmt.Errorf("invalid EXT-X-KEY: %s, line: %d", line, i+1)
			}
			key = new(Key)
			method := CryptMethod(params["METHOD"])
			if method != "" && method != CryptMethodAES && method != CryptMethodNONE {
				return nil, fmt.Errorf("invalid EXT-X-KEY method: %s, line: %d", method, i+1)
			}
			key.Method = method
			key.URI = params["URI"]
			key.IV = params["IV"]
		default:
			continue
		}
	}
	return m3u8, nil

}

func parseLineParameters(line string) map[string]string {
	r := lineParameterPattern.FindAllStringSubmatch(line,-1)
	params := make(map[string]string)
	for _,arr := range r{
		params[arr[1]] = strings.Trim(arr[2],"\"")
	}
	
	return params
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
			os.Exit(-1)
		}
	}()
	m3u8URL := "http://devimages.apple.com/iphone/samples/bipbop/bipbopall.m3u8"
	u, err := url.Parse(m3u8URL)
	if err != nil {
		panic(err)
	}
	m3u8URL = u.String()
	body, err := tool.Get(m3u8URL)
	if err != nil {
		panic(err)
	}
	//noinspection GoUnhandledErrorResult
	defer body.Close()
	s := bufio.NewScanner(body)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	m3u8, err := parseLines(lines)
	if err != nil {
		panic(err)
	}
	jsonBytes, err := json.MarshalIndent(m3u8, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}





