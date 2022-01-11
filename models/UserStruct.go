package models

import "edu_app_backend/types"

type User struct {
	ID_                string                 `bson:"_id, omitempty"`
	UserName           string                 `bson:"UserName, omitempty"`
	EmailAddress       string                 `bson:"EmailAddress ,omitempty"`
	ProfileImage       string                 `bson:"ProfileImage, omitempty"`
	PassHash           string                 `bson:"PassHash, omitempty"`
	Projects           []string               `bson:"Projects, omitempty"`
	Courses            []string               `bson:"Courses, omitiempty"`
	Following          []string               `bson:"Following, omitiempty"`
	Followers          []string               `bson:"Followers, omitiempty"`
	ProfileDescription string                 `bson:"ProfileDescription, omitempty"`
	SocialMediaLinks   []types.UserSocialLink `bson:"SocialMediaLinks, omitempty"`
}
