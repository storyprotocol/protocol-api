package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetTag(policyId string) ([]*beta_v0.Tag, error) {
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

	tags := []*beta_v0.Tag{}
	for _, tag := range tagRes.Tag {
		tags = append(tags, tag)
	}

	return tags, nil
}

func (c *ServiceBetaImpl) ListTags(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Tag, error) {
	query := fmt.Sprintf(`
	query(%s) {
		{
		  tags(id: "%s") {
			id
			ipId
			tag
			deletedAt
		  }
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)
	ctx := context.Background()
	var tagRes beta_v0.TagTheGraphResponse
	if err := c.client.Run(ctx, req, &tagRes); err != nil {
		return nil, fmt.Errorf("failed to get tags from the graph. error: %v", err)
	}

	tags := []*beta_v0.Tag{}
	for _, tag := range tagRes.Tag {
		tags = append(tags, tag)
	}

	return tags, nil
}
