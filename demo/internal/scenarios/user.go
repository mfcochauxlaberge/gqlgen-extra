package scenarios

import (
	"github.com/mfcochauxlaberge/gqlgen-extra/gqltest"
)

func init() {
	Scenarios = append(
		Scenarios,
		gqltest.Scenario{
			Name: "query user",
			Play: func(env *gqltest.Env) {
				_ = env.Rec.Query(
					"Get user",
					`
					query GetUser {
						user(input: {id: "u1"}) {
							__typename
							... on User {
								id
								username
							}
						}
					}
					`,
				)
			},
		},
		gqltest.Scenario{
			Name: "query user with relationships",
			Play: func(env *gqltest.Env) {
				_ = env.Rec.Query(
					"Get user",
					`
					query GetUser {
						user(input: {id: "u1"}) {
							__typename
							... on User {
								id
								username
								articles {
									id
									createdAt
									title
									content
									comments {
										content
										author {
											username
										}
									}
								}
								comments {
									id
									content
								}
								likes {
									title
								}
							}
						}
					}
					`,
				)
			},
		},
	)
}
