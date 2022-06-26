package matcher

import (
	"errors"
	"regexp"
	"strings"

	"github.com/phamdinhha/go-authorizer/internal/models"
)

// Fake data to get resource owner, could be replaced by query the actual resources from other services
func getUserResources(userID string) (userResources []models.UserResources, err error) {
	if userID == "user1@gmail.com" {
		return []models.UserResources{
			models.UserResources{
				BookID:       "book1",
				TemplateList: []string{"template1", "template2", "template3"},
			},
			models.UserResources{
				BookID:       "book3",
				TemplateList: []string{"template1", "template2", "template3"},
			},
		}, nil
	}
	if userID == "user2@gmail.com" {
		return []models.UserResources{
			models.UserResources{
				BookID:       "book2",
				TemplateList: []string{"template1", "template2", "template4"},
			},
		}, nil
	}
	return nil, errors.New("cannot get user's owned resources")
}

func ResourceMatch(userID, bookID, templateID string) bool {
	userResources, err := getUserResources(userID)
	if err != nil {
		return false
	}
	for _, resource := range userResources {
		if resource.BookID == bookID {
			for _, template := range resource.TemplateList {
				if template == templateID {
					return true
				}
			}
			return false
		}
	}
	return false
}

// From casbin KeyMatch4
func KeyMatch(userID, key1, key2 string) bool {

	key2 = strings.Replace(key2, "/*", "/.*", -1)

	tokens := []string{}

	re := regexp.MustCompile(`\{([^/]+)\}`)
	key2 = re.ReplaceAllStringFunc(key2, func(s string) string {
		tokens = append(tokens, s[1:len(s)-1])
		return "([^/]+)"
	})

	re = regexp.MustCompile("^" + key2 + "$")
	matches := re.FindStringSubmatch(key1)
	if matches == nil {
		return false
	}
	matches = matches[1:]
	if len(tokens) != len(matches) {
		panic(errors.New("KeyMatch: number of tokens is not equal to number of values"))
	}
	values := map[string]string{}
	for key, token := range tokens {
		if _, ok := values[token]; !ok {
			values[token] = matches[key]
		}
		if values[token] != matches[key] {
			return false
		}
	}
	resourceList := strings.Split(key2, "/")
	if resourceList[3] == ".*" && resourceList[5] == ".*" {
		return true
	}
	keyList := strings.Split(key1, "/")
	bookID := keyList[3]
	templateID := keyList[5]

	return ResourceMatch(userID, bookID, templateID)
}

func ResourceMatchFunc(args ...interface{}) (interface{}, error) {
	userID := args[0].(string)
	key1 := args[1].(string)
	key2 := args[2].(string)
	return (bool)(KeyMatch(userID, key1, key2)), nil
}
