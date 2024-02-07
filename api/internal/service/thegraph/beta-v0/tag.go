package beta_v0

import (
	"context"
	"fmt"

	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetTag(tagId string) (*beta_v0.Tag, error) {
	query := fmt.Sprintf(`
		query {
		  tags(where: { uuid:  "%s" }) {
			uuid
			ipId
			tag
			deletedAt
			blockNumber
			blockTimestamp
		  }
		}
    `, tagId)

	req := c.buildNewRequest(nil, query)
	ctx := context.Background()
	var tagRes beta_v0.TagsTheGraphResponse
	if err := c.client.Run(ctx, req, &tagRes); err != nil {
		return nil, fmt.Errorf("failed to get tags from the graph. error: %v", err)
	}

	tags := []*beta_v0.Tag{}
	for _, tag := range tagRes.Tags {
		tag.ID = tag.UUID
		tags = append(tags, tag)
		tag.UUID = ""
	}

	return tags[0], nil
}

func (c *ServiceBetaImpl) ListTag(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Tag, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s) {
	  tags(%s, where:{%s}) {
		uuid
		ipId
		tag
		deletedAt
		blockNumber
		blockTimestamp
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
		tag.ID = tag.UUID
		tags = append(tags, tag)
		tag.UUID = ""

	}

	return tags, nil
}
