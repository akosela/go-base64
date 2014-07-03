// Copyright 2014 Andy Kosela.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func wrap(s string) []string {
        var ret []string
        for len(s) > 76 {
                ret = append(ret, s[:76])
                s = s[76:]
        }
        return append(ret, s)
}

func main() {
	var str = flag.Bool("s", false, "string")
	var decode = flag.Bool("d", false, "decode")
	flag.Parse()

	if *str {
		if len(os.Args) != 3 {
			fmt.Fprintf(os.Stderr, "Usage: %s -s <string>\n",
				os.Args[0])
			os.Exit(1)
		}
		data := os.Args[2]
		enc := b64.StdEncoding.EncodeToString([]byte(data))
		fmt.Printf("BASE64 (\"%s\") = %s\n", data, enc)
		os.Exit(0)
	}

	data, _ := ioutil.ReadAll(os.Stdin)
	if *decode {
		dec, _ := b64.StdEncoding.DecodeString(string(data))
		fmt.Println(string(dec))
		os.Exit(0)
	}

	enc := b64.StdEncoding.EncodeToString(data)
	fmt.Println(strings.Join(wrap(enc), "\n"))
}
