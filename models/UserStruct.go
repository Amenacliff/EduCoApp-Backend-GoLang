package models

import "edu_app_backend/types"

type User struct {
	ID_                string                 `bson:"_id, omitempty"`
	UserName           string                 `bson:"userName, omitempty"`
	EmailAddress       string                 `bson:"emailAddress ,omitempty"`
	ProfileImage       string                 `bson:"profileImage, omitempty"`
	Password           string                 `bson:"password, omitempty"`
	Projects           []string               `bson:"projects, omitempty"`
	Courses            []string               `bson:"courses, omitiempty"`
	Following          []string               `bson:"following, omitiempty"`
	Followers          []string               `bson:"followers, omitiempty"`
	ProfileDescription string                 `bson:"profileDescription, omitempty"`
	SocialMediaLinks   []types.UserSocialLink `bson:"profileSocialLinks, omitempty"`
}
