package mockdata

import "restapi.com/myuser/myapi/models"

var Users = []models.User{
	{
		ID:        1,
		FirstName: "Pooja",
		LastName:  "Varma",
		Email:     "poojavarma@example.com",
		Age:       25,
	},
	{
		ID:        2,
		FirstName: "pooja",
		LastName:  "varma",
		Email:     "varmapooja@example.com",
		Age:       30,
	},
}
