package hw10_program_optimization //nolint:golint,stylecheck

import (
	"io"
	"strings"
	"unicode"
)

type DomainStat map[string]int

/* Example:
{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
{"Id":2,"Name":"Justin Oliver Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"oPerez","Email":"MelissaGutierrez@Twinte.biz","Phone":"106-05-18","Password":"f00GKr9i","Address":"Oak Valley Lane 19"}
{"Id":3,"Name":"Brian Olson","Username":"non_quia_id","Email":"FrancesEllis@Quinu.edu","Phone":"237-75-34","Password":"cmEPhX8","Address":"Butterfield Junction 74"}
{"Id":4,"Name":"Jesse Vasquez Jr. Sr. I II III IV V MD DDS PhD DVM","Username":"qRichardson","Email":"mLynch@Dabtype.name","Phone":"9-373-949-64-00","Password":"SiZLeNSGn","Address":"Fulton Hill 80"}
{"Id":5,"Name":"Clarence Olson","Username":"RachelAdams","Email":"RoseSmith@Browsecat.com","Phone":"988-48-97","Password":"71kuz3gA5w","Address":"Monterey Park 39"}
{"Id":6,"Name":"Gregory Reid","Username":"tButler","Email":"5Moore@Teklist.net","Phone":"520-04-16","Password":"r639qLNu","Address":"Sunfield Park 20"}
{"Id":7,"Name":"Janice Rose","Username":"KeithHart","Email":"nulla@Linktype.com","Phone":"146-91-01","Password":"acSBF5","Address":"Russell Trail 61"}
*/

const (
	StateScan             = 0
	StateReadKey          = 1
	StateReadEmailUser    = 2
	StateReadEmailDomain1 = 3
	StateReadEmailDomain2 = 4
)

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	// what we do here is a simple parsing, without any json correctness checks
	// normally you won't do that, but it should be fine for the task goals

	result := make(DomainStat)
	buf := make([]byte, 4096)
	state := StateScan
	// how many characters to skip
	skip := 0

	// full part of the domain: "gmail.com", "yahoo.com", etc
	var emailDomain1 strings.Builder
	emailDomain1.Grow(64)
	// only 1-st level domain: "com", "gov", tec
	var emailDomain2 strings.Builder
	emailDomain2.Grow(16)

	var last5, last4, last3, last2, last1, c rune

	for {
		read, err := r.Read(buf)
		if read == 0 {
			break
		}
		if err != nil && err != io.EOF { // nolint: errorlint
			return nil, err // nolint: wrapcheck
		}

		for i := 0; i < read; i++ {
			last5 = last4
			last4 = last3
			last3 = last2
			last2 = last1
			last1 = c
			c = rune(buf[i])

			// skip some bytes, if needed
			if skip > 0 {
				skip--
				continue
			}

			switch state {
			case StateScan:
				if c == '"' {
					state = StateReadKey
				}
			case StateReadKey:
				if c == '"' {
					// check if the json key matches "Email"
					if last5 == 'E' &&
						last4 == 'm' &&
						last3 == 'a' &&
						last2 == 'i' &&
						last1 == 'l' {
						skip = 2
						state = StateReadEmailUser
					} else {
						state = StateScan
					}
				}
			case StateReadEmailUser:
				if c == '@' {
					// found @, so now we're reading domain part
					state = StateReadEmailDomain1
				}
			case StateReadEmailDomain1:
				emailDomain1.WriteRune(unicode.ToLower(c))
				if c == '.' {
					// "gmail.com"
					//       ^ found this dot in the domain
					state = StateReadEmailDomain2
				}
			case StateReadEmailDomain2:
				if c == '"' {
					// done reading full email
					domainStr := emailDomain2.String()
					// check if the 1-st level domain matches the desired value
					if domainStr == domain {
						result[emailDomain1.String()]++
					}
					state = StateScan
					emailDomain1.Reset()
					emailDomain2.Reset()
				} else {
					lc := unicode.ToLower(c)
					emailDomain1.WriteRune(lc)
					emailDomain2.WriteRune(lc)
				}
			}
		}
	}
	return result, nil
}
