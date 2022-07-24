package sidebar

import (
	"fmt"
	"piennews/models"
)

func GetUserSidebar(c *models.CustomerModel, exist bool) (string, string) {
	var name, profile string
	if exist {
		name = fmt.Sprintf("%v", c.Name)
		profile = c.Image
	} else {
		name = "<a href=\"/auth\" class=\"btn-xs btn-secondary hide-on-collapse  text-sm\">Loginx</a>"
		profile = "/assets/img/user-default.jpg"
	}

	return name, profile

}
