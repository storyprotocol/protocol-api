package beta_v0

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetTag(policyId string) (*beta_v0.Tag, error) {
	query := fmt.Sprintf(`
		query {
		  tag(id: "%s") {
			id
			ipId
			tag
			deletedAt
		  }
		}
    `, policyId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var tagRes beta_v0.TagTheGraphResponse
	if err := c.client.Run(ctx, req, &tagRes); err != nil {
		return nil, fmt.Errorf("failed to get tag from the graph. error: %v", err)
	}

	tagIdHashed := b64.StdEncoding.EncodeToString([]byte(tagRes.Tag.ID))
	tagRes.Tag.ID = tagIdHashed

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
		tagIdHashed := b64.StdEncoding.EncodeToString([]byte(tag.ID))
		tag.ID = tagIdHashed
		tags = append(tags, tag)
	}

	return tags, nil
}
