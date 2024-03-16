package betav0

import (
	"context"
	"fmt"

	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetTag(tagId string) (*beta_v0.Tag, *beta_v0.RenTag, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  records(filter: { uuid:  "%s" }) {
			uuid
			ip_id
			tag
			deleted_at
			block_number
			block_timestamp
		  }
		}
    `, tagId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var tagRes beta_v0.RenTagsTheGraphResponse
		if err := c.client.Run(ctx, req, &tagRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get tags from the graph. error: %v", err)
		}

		tags := []*beta_v0.RenTag{}
		for _, tag := range tagRes.Tags {
			tag.ID = tag.UUID
			tags = append(tags, tag)
			tag.UUID = ""
		}

		return nil, tags[0], nil
	} else {
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
			return nil, nil, fmt.Errorf("failed to get tags from the graph. error: %v", err)
		}

		tags := []*beta_v0.Tag{}
		for _, tag := range tagRes.Tags {
			tag.ID = tag.UUID
			tags = append(tags, tag)
			tag.UUID = ""
		}

		return tags[0], nil, nil
	}

}

func (c *ServiceBetaImpl) ListTag(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Tag, []*beta_v0.RenTag, error) {
	whereString := c.buildWhereConditions(options)
	query := ""
	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)
		query = fmt.Sprintf(`
		query(%s) {
		  records(%s, filter:{%s}) {
			uuid
			ipId
			tag
			deletedAt
			blockNumber
			blockTimestamp
		  }
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)
		ctx := context.Background()
		var tagRes beta_v0.RenTagsTheGraphResponse
		if err := c.client.Run(ctx, req, &tagRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get tags from the graph. error: %v", err)
		}

		tags := []*beta_v0.RenTag{}
		for _, tag := range tagRes.Tags {
			tag.ID = tag.UUID
			tags = append(tags, tag)
			tag.UUID = ""

		}

		return nil, tags, nil
	} else {
		query = fmt.Sprintf(`
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
			return nil, nil, fmt.Errorf("failed to get tags from the graph. error: %v", err)
		}

		tags := []*beta_v0.Tag{}
		for _, tag := range tagRes.Tags {
			tag.ID = tag.UUID
			tags = append(tags, tag)
			tag.UUID = ""

		}

		return tags, nil, nil
	}

}
