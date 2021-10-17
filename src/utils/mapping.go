package utils

import (
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
	"github.com/ducnguyen96/reddit-clone/graph/model"
)

func MapEntGoUserToGraphUser(u *ent.User) *model.User {
	return &model.User{
		ID:        Uint64ToString(u.ID),
		Username:  u.Username,
		Avatar:    u.AvatarURL,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}

func MapEntGoCommunityToGraphCommunity(c *ent.Community) *model.Community {
	return &model.Community{
		ID:        Uint64ToString(c.ID),
		Name:      c.Name,
		Slug:      c.Slug,
		Type:      EntGoCommunityTypeToGraphCommunityType(c.Type),
		IsAdult:   c.IsAdult,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}

func EntGoPostToGraphPost(p *ent.Post) *model.Post {
	return &model.Post{
		ID:          Uint64ToString(p.ID),
		Title:       p.Title,
		Slug:        p.Slug,
		Content:     p.Content,
		Type:        EntGoPostTypeToGraphPostType(p.Type),
		ContentMode: EntGoContentModeToGraphContentMode(p.ContentMode),
		UpVotes:     p.UpVotes,
		DownVotes:   p.DownVotes,
		CreatedAt:   p.CreatedAt.String(),
		UpdatedAt:   p.UpdatedAt.String(),
	}
}

func EntGoCommunityTypeToGraphCommunityType(t enums.CommunityType) model.CommunityType {
	switch t {
	case enums.Public:
		return model.CommunityTypePublic
	case enums.Restricted:
		return model.CommunityTypeRestricted
	default:
		return model.CommunityTypePrivate
	}
}

func EntGoPostTypeToGraphPostType(p enums.PostType) model.PostType {
	switch p {
	case enums.Post:
		return model.PostTypePost
	case enums.Image_Video:
		return model.PostTypeImageVideo
	default:
		return model.PostTypeLink
	}
}

func EntGoContentModeToGraphContentMode(c enums.InputContentMode) model.InputContentMode {
	switch c {
	case enums.MarkDown:
		return model.InputContentModeMarkDown
	default:
		return model.InputContentModeTextEditor
	}
}

func GraphCommunityTypeToCommunityType(t model.CommunityType) enums.CommunityType {
	switch t {
	case model.CommunityTypePublic:
		return enums.Public
	case model.CommunityTypeRestricted:
		return enums.Restricted
	default:
		return enums.Private
	}
}

func GraphPostTypeToEntPostType(t model.PostType) enums.PostType {
	switch t {
	case model.PostTypePost:
		return enums.Post
	case model.PostTypeImageVideo:
		return enums.Image_Video
	default:
		return enums.Link
	}
}

func GraphContentModeToEntContentMode(t model.InputContentMode) enums.InputContentMode {
	switch t {
	case model.InputContentModeMarkDown:
		return enums.MarkDown
	default:
		return enums.TextEditor
	}
}
