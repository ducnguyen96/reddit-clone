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

func EntGoPostToGraphPost(p *ent.Post, isUpVoted bool, isDownVoted bool) *model.Post {
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
		IsUpVoted: isUpVoted,
		IsDownVoted: isDownVoted,
	}
}

func EntUserActionToGraph(action *ent.Action) *model.UserAction {
	return &model.UserAction{
		ID:         Uint64ToString(action.ID),
		Type:       EntUserActionTypeToGraph(action.Type),
		Target:     Uint64ToString(action.Target),
		TargetType: EntUserActionTargetTypeToGraph(action.TargetType),
		CreatedAt:   action.CreatedAt.String(),
		UpdatedAt:   action.UpdatedAt.String(),
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

func GraphMediaTypeToEntMediaType(t model.MediaType) enums.MediaType {
	switch t {
	case model.MediaTypeVideo:
		return enums.Video
	default:
		return enums.Image
	}
}

func EntMediaTypeToGraph(t enums.MediaType) model.MediaType {
	switch t {
	case enums.Video:
		return model.MediaTypeVideo
	default:
		return model.MediaTypeImage
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

func GraphUserActionTypeToEnt(t model.UserActionType) enums.UserActionType {
	switch t {
	case model.UserActionTypeUpVote:
		return enums.UpVote
	default:
		return enums.DownVote
	}
}

func GraphUserActionTargetTypeToEnt(t model.UserActionTargetType) enums.UserActionTargetType {
	switch t {
	case model.UserActionTargetTypePost:
		return enums.POST
	default:
		return enums.COMMENT
	}
}

func EntUserActionTypeToGraph(actionType enums.UserActionType) model.UserActionType {
	switch actionType {
	case enums.UpVote:
		return model.UserActionTypeUpVote
	default:
		return model.UserActionTypeDownVote
	}
}

func EntUserActionTargetTypeToGraph(targetType enums.UserActionTargetType) model.UserActionTargetType {
	switch targetType {
	case enums.POST:
		return model.UserActionTargetTypePost
	default:
		return model.UserActionTargetTypeComment
	}
}

func EntCommentToGraph(comment *ent.Comment, isUpVoted bool, isDownVoted bool) *model.Comment {
	return &model.Comment{
		ID:          Uint64ToString(comment.ID),
		Content:     comment.Content,
		ContentMode: EntGoContentModeToGraphContentMode(comment.ContentMode),
		CreatedAt:   comment.CreatedAt.String(),
		UpdatedAt:   comment.UpdatedAt.String(),
		UpVotes:     comment.UpVotes,
		DownVotes:   comment.DownVotes,
		IsUpVoted:   isUpVoted,
		IsDownVoted: isDownVoted,
		PostID: Uint64ToString(comment.PostID),
	}
}

func EntMediaToGraph(media *ent.Media) *model.Media {
	return &model.Media{
		ID:        Uint64ToString(media.ID),
		URL:       media.URL,
		Type:      EntMediaTypeToGraph(media.Type),
		CreatedAt: media.CreatedAt.String(),
		UpdatedAt: media.UpdatedAt.String(),
	}
}