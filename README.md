# protocol-api
API service for core protocol data

### Note:
In order to generate a propertly formated api gateway openapi json:
- copy `info` object (line 3) from api/cmd/docs/swagger.json
- run `make compileswag`
- paste `info` in newly generated swagger.json
- copy swagger.json and replace old api_gateway_swagger.json
- copy  and replace `betav0.` and `options.` with empty space inside of the new api_gateway_swager.json
