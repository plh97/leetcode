package numUniqueEmails

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	// 处理

	for i := range emails {
		email := emails[i]

		// 处理 +
		indexPlus := strings.Index(email, "+")
		indexMail := strings.Index(email, "@")
		if indexPlus > -1 {
			email = email[:indexPlus] + email[indexMail:]
		}
		// 处理 .
		indexMail = strings.Index(email, "@")
		for i := range email {
			if i < indexMail && email[i] == '.' {
				email = email[:i] + email[i+1:]
				i--
			}
		}
		emails[i] = email
	}
	lastIndexOf([]string{"r3"}, "123")
	// 去重
	res := 0
	for i := range emails {
		if i == lastIndexOf(emails, emails[i]) {
			res++
		}
	}

	return res
}

func lastIndexOf(nums []string, val string) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
