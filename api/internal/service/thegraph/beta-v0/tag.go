package beta_v0

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	encoder "github.com/storyprotocol/protocol-api/api/internal/helpers"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetTag(tagId string) (*beta_v0.Tag, error) {
	id, err := encoder.Decrypt(tagId)
	if err != nil {
		return nil, fmt.Errorf("failed to encode id. error: %v", err)
	}

	query := fmt.Sprintf(`
		query {
		  tag(id: "%s") {
			id
			ipId
			tag
			deletedAt
		  }
		}
    `, id)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var tagRes beta_v0.TagTheGraphResponse
	if err := c.client.Run(ctx, req, &tagRes); err != nil {
		return nil, fmt.Errorf("failed to get tag from the graph. error: %v", err)
	}

	tagRes.Tag.ID = tagId

	return tagRes.Tag, nil
}

func (c *ServiceBetaImpl) ListTag(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Tag, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s) {
	  tags(%s, where:{%s}) {
		id
		ipId
		tag
		deletedAt
	  }
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

	req := c.buildNewRequest(options, query)
	ctx := context.Background()
	var tagRes beta_v0.TagsTheGraphResponse
	if err := c.client.Run(ctx, req, &tagRes); err != nil {
		return nil, fmt.Errorf("failed to get tags from the graph. error: %v", err)
	}

	tags := []*beta_v0.Tag{}
	for _, tag := range tagRes.Tags {
		id, err := encoder.Encrypt(tag.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to encode id. error: %v", err)
		}

		tag.ID = id

		tags = append(tags, tag)
	}

	return tags, nil
}
