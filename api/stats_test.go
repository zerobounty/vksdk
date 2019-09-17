package api

import (
	"strconv"
	"testing"
)

func TestVK_StatsGet(t *testing.T) {
	t.Skip("See https://vk.com/bug136096")
	needUserToken(t)

	_, err := vkUser.StatsGet(map[string]string{
		"group_id":        strconv.Itoa(vkGroupID),
		"interval":        "day",
		"intervals_count": "10",
		"extended":        "1",
	})
	if err != nil {
		t.Errorf("VK.StatsGet() err = %v", err)
	}
}

func TestVK_StatsTrackVisitor(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatsTrackVisitor(map[string]string{})
	if err != nil {
		t.Errorf("VK.StatsTrackVisitor() err = %v", err)
	}
}
