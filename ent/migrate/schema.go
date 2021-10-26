// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActionsColumns holds the columns for the "actions" table.
	ActionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "target", Type: field.TypeUint64},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"UP_VOTES", "DOWN_VOTES"}},
		{Name: "target_type", Type: field.TypeEnum, Enums: []string{"POST", "COMMENT"}},
	}
	// ActionsTable holds the schema information for the "actions" table.
	ActionsTable = &schema.Table{
		Name:       "actions",
		Columns:    ActionsColumns,
		PrimaryKey: []*schema.Column{ActionsColumns[0]},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
		{Name: "content_mode", Type: field.TypeEnum, Enums: []string{"MarkDown", "TextEditor"}},
		{Name: "up_votes", Type: field.TypeInt, Default: 0},
		{Name: "down_votes", Type: field.TypeInt, Default: 0},
		{Name: "post_id", Type: field.TypeUint64},
		{Name: "comment_children", Type: field.TypeUint64, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_comments_children",
				Columns:    []*schema.Column{CommentsColumns[8]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommunitiesColumns holds the columns for the "communities" table.
	CommunitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 300},
		{Name: "slug", Type: field.TypeString, Unique: true, Size: 300},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"Public", "Restricted", "Private"}},
		{Name: "is_adult", Type: field.TypeBool, Default: false},
	}
	// CommunitiesTable holds the schema information for the "communities" table.
	CommunitiesTable = &schema.Table{
		Name:       "communities",
		Columns:    CommunitiesColumns,
		PrimaryKey: []*schema.Column{CommunitiesColumns[0]},
	}
	// MediaColumns holds the columns for the "media" table.
	MediaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "url", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"Image", "Video"}},
	}
	// MediaTable holds the schema information for the "media" table.
	MediaTable = &schema.Table{
		Name:       "media",
		Columns:    MediaColumns,
		PrimaryKey: []*schema.Column{MediaColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Unique: true, Size: 300},
		{Name: "slug", Type: field.TypeString, Unique: true, Size: 400},
		{Name: "content", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"Post", "Image_Video", "Link"}},
		{Name: "content_mode", Type: field.TypeEnum, Enums: []string{"MarkDown", "TextEditor"}},
		{Name: "up_votes", Type: field.TypeInt, Default: 0},
		{Name: "down_votes", Type: field.TypeInt, Default: 0},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "value", Type: field.TypeString, Size: 50},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 150},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// CommunityUsersColumns holds the columns for the "community_users" table.
	CommunityUsersColumns = []*schema.Column{
		{Name: "community_id", Type: field.TypeUint64},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// CommunityUsersTable holds the schema information for the "community_users" table.
	CommunityUsersTable = &schema.Table{
		Name:       "community_users",
		Columns:    CommunityUsersColumns,
		PrimaryKey: []*schema.Column{CommunityUsersColumns[0], CommunityUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "community_users_community_id",
				Columns:    []*schema.Column{CommunityUsersColumns[0]},
				RefColumns: []*schema.Column{CommunitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "community_users_user_id",
				Columns:    []*schema.Column{CommunityUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CommunityAdminsColumns holds the columns for the "community_admins" table.
	CommunityAdminsColumns = []*schema.Column{
		{Name: "community_id", Type: field.TypeUint64},
		{Name: "user_id", Type: field.TypeUint64},
	}
	// CommunityAdminsTable holds the schema information for the "community_admins" table.
	CommunityAdminsTable = &schema.Table{
		Name:       "community_admins",
		Columns:    CommunityAdminsColumns,
		PrimaryKey: []*schema.Column{CommunityAdminsColumns[0], CommunityAdminsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "community_admins_community_id",
				Columns:    []*schema.Column{CommunityAdminsColumns[0]},
				RefColumns: []*schema.Column{CommunitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "community_admins_user_id",
				Columns:    []*schema.Column{CommunityAdminsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CommunityPostsColumns holds the columns for the "community_posts" table.
	CommunityPostsColumns = []*schema.Column{
		{Name: "community_id", Type: field.TypeUint64},
		{Name: "post_id", Type: field.TypeUint64},
	}
	// CommunityPostsTable holds the schema information for the "community_posts" table.
	CommunityPostsTable = &schema.Table{
		Name:       "community_posts",
		Columns:    CommunityPostsColumns,
		PrimaryKey: []*schema.Column{CommunityPostsColumns[0], CommunityPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "community_posts_community_id",
				Columns:    []*schema.Column{CommunityPostsColumns[0]},
				RefColumns: []*schema.Column{CommunitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "community_posts_post_id",
				Columns:    []*schema.Column{CommunityPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostTagsColumns holds the columns for the "post_tags" table.
	PostTagsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeUint64},
		{Name: "tag_id", Type: field.TypeUint64},
	}
	// PostTagsTable holds the schema information for the "post_tags" table.
	PostTagsTable = &schema.Table{
		Name:       "post_tags",
		Columns:    PostTagsColumns,
		PrimaryKey: []*schema.Column{PostTagsColumns[0], PostTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_tags_post_id",
				Columns:    []*schema.Column{PostTagsColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_tags_tag_id",
				Columns:    []*schema.Column{PostTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostCommentsColumns holds the columns for the "post_comments" table.
	PostCommentsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeUint64},
		{Name: "comment_id", Type: field.TypeUint64},
	}
	// PostCommentsTable holds the schema information for the "post_comments" table.
	PostCommentsTable = &schema.Table{
		Name:       "post_comments",
		Columns:    PostCommentsColumns,
		PrimaryKey: []*schema.Column{PostCommentsColumns[0], PostCommentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_comments_post_id",
				Columns:    []*schema.Column{PostCommentsColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_comments_comment_id",
				Columns:    []*schema.Column{PostCommentsColumns[1]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPostsColumns holds the columns for the "user_posts" table.
	UserPostsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "post_id", Type: field.TypeUint64},
	}
	// UserPostsTable holds the schema information for the "user_posts" table.
	UserPostsTable = &schema.Table{
		Name:       "user_posts",
		Columns:    UserPostsColumns,
		PrimaryKey: []*schema.Column{UserPostsColumns[0], UserPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_posts_user_id",
				Columns:    []*schema.Column{UserPostsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_posts_post_id",
				Columns:    []*schema.Column{UserPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserCommentsColumns holds the columns for the "user_comments" table.
	UserCommentsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "comment_id", Type: field.TypeUint64},
	}
	// UserCommentsTable holds the schema information for the "user_comments" table.
	UserCommentsTable = &schema.Table{
		Name:       "user_comments",
		Columns:    UserCommentsColumns,
		PrimaryKey: []*schema.Column{UserCommentsColumns[0], UserCommentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_comments_user_id",
				Columns:    []*schema.Column{UserCommentsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_comments_comment_id",
				Columns:    []*schema.Column{UserCommentsColumns[1]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserActionsColumns holds the columns for the "user_actions" table.
	UserActionsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "action_id", Type: field.TypeUint64},
	}
	// UserActionsTable holds the schema information for the "user_actions" table.
	UserActionsTable = &schema.Table{
		Name:       "user_actions",
		Columns:    UserActionsColumns,
		PrimaryKey: []*schema.Column{UserActionsColumns[0], UserActionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_actions_user_id",
				Columns:    []*schema.Column{UserActionsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_actions_action_id",
				Columns:    []*schema.Column{UserActionsColumns[1]},
				RefColumns: []*schema.Column{ActionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActionsTable,
		CommentsTable,
		CommunitiesTable,
		MediaTable,
		PostsTable,
		TagsTable,
		UsersTable,
		CommunityUsersTable,
		CommunityAdminsTable,
		CommunityPostsTable,
		PostTagsTable,
		PostCommentsTable,
		UserPostsTable,
		UserCommentsTable,
		UserActionsTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = CommentsTable
	CommunityUsersTable.ForeignKeys[0].RefTable = CommunitiesTable
	CommunityUsersTable.ForeignKeys[1].RefTable = UsersTable
	CommunityAdminsTable.ForeignKeys[0].RefTable = CommunitiesTable
	CommunityAdminsTable.ForeignKeys[1].RefTable = UsersTable
	CommunityPostsTable.ForeignKeys[0].RefTable = CommunitiesTable
	CommunityPostsTable.ForeignKeys[1].RefTable = PostsTable
	PostTagsTable.ForeignKeys[0].RefTable = PostsTable
	PostTagsTable.ForeignKeys[1].RefTable = TagsTable
	PostCommentsTable.ForeignKeys[0].RefTable = PostsTable
	PostCommentsTable.ForeignKeys[1].RefTable = CommentsTable
	UserPostsTable.ForeignKeys[0].RefTable = UsersTable
	UserPostsTable.ForeignKeys[1].RefTable = PostsTable
	UserCommentsTable.ForeignKeys[0].RefTable = UsersTable
	UserCommentsTable.ForeignKeys[1].RefTable = CommentsTable
	UserActionsTable.ForeignKeys[0].RefTable = UsersTable
	UserActionsTable.ForeignKeys[1].RefTable = ActionsTable
}
