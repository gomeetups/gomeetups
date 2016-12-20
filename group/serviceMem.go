package group

import (
	"errors"
	"strings"

	fixtures "bitbucket.org/devmach/gomeetups/fixtures"
	models "bitbucket.org/devmach/gomeetups/models"
)

type GroupServiceMem struct {
	Limit int64
}

func (*GroupServiceMem) SearchGroups(filter *models.GroupSearchValidParams) (groups map[string]models.Group, err error) {
	groups = make(map[string]models.Group)

	for uuid, group := range fixtures.Groups {

		if filter.Name != "" || filter.Description != "" {
			var doesMatch = false

			if filter.Name != "" && strings.Contains(strings.ToLower(group.Name), strings.ToLower(filter.Name)) {
				doesMatch = true
			}

			if filter.Description != "" && strings.Contains(strings.ToLower(group.Name), strings.ToLower(filter.Description)) {
				doesMatch = true
			}

			if doesMatch {
				groups[uuid] = group
			}

		} else {
			groups[uuid] = group
		}
	}

	return groups, nil
}

func (*GroupServiceMem) GroupDetails(groupId string) (group models.Group, err error) {
	if val, ok := fixtures.Groups[groupId]; ok {
		return val, nil
	}

	return models.Group{}, errors.New("Not found!")
}
