## Add MCP Server

Below are curl examples

curl -X 'POST' \
  'https://ai.bitop.dev/v1/mcp/server' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "server_id": "string",
  "server_name": "string",
  "alias": "string",
  "description": "string",
  "transport": "sse",
  "spec_version": "2025-06-18",
  "auth_type": "none",
  "url": "string",
  "mcp_info": {
    "server_name": "string",
    "description": "string",
    "logo_url": "string",
    "mcp_server_cost_info": {
      "default_cost_per_query": 0,
      "tool_name_to_cost_per_query": {
        "additionalProp1": 0,
        "additionalProp2": 0,
        "additionalProp3": 0
      }
    }
  },
  "mcp_access_groups": [
    "string"
  ],
  "command": "string",
  "args": [
    "string"
  ],
  "env": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  }
}'

example Successful Response

{
  "server_id": "string",
  "server_name": "string",
  "alias": "string",
  "description": "string",
  "url": "string",
  "transport": "sse",
  "spec_version": "2024-11-05",
  "auth_type": "none",
  "created_at": "2025-08-10T15:27:35.211Z",
  "created_by": "string",
  "updated_at": "2025-08-10T15:27:35.211Z",
  "updated_by": "string",
  "teams": [
    {
      "additionalProp1": "string",
      "additionalProp2": "string",
      "additionalProp3": "string"
    }
  ],
  "mcp_access_groups": [
    "string"
  ],
  "mcp_info": {
    "server_name": "string",
    "description": "string",
    "logo_url": "string",
    "mcp_server_cost_info": {
      "default_cost_per_query": 0,
      "tool_name_to_cost_per_query": {
        "additionalProp1": 0,
        "additionalProp2": 0,
        "additionalProp3": 0
      }
    }
  },
  "status": "unknown",
  "last_health_check": "2025-08-10T15:27:35.211Z",
  "health_check_error": "string",
  "command": "string",
  "args": [
    "string"
  ],
  "env": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  }
}

## Edit MCP Server

Curl Example

curl -X 'PUT' \
  'https://ai.bitop.dev/v1/mcp/server' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "server_id": "string",
  "server_name": "string",
  "alias": "string",
  "description": "string",
  "transport": "sse",
  "spec_version": "2025-06-18",
  "auth_type": "none",
  "url": "string",
  "mcp_info": {
    "server_name": "string",
    "description": "string",
    "logo_url": "string",
    "mcp_server_cost_info": {
      "default_cost_per_query": 0,
      "tool_name_to_cost_per_query": {
        "additionalProp1": 0,
        "additionalProp2": 0,
        "additionalProp3": 0
      }
    }
  },
  "mcp_access_groups": [
    "string"
  ],
  "command": "string",
  "args": [
    "string"
  ],
  "env": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  }
}'

Example Successful Response

{
  "server_id": "string",
  "server_name": "string",
  "alias": "string",
  "description": "string",
  "url": "string",
  "transport": "sse",
  "spec_version": "2024-11-05",
  "auth_type": "none",
  "created_at": "2025-08-10T15:29:08.740Z",
  "created_by": "string",
  "updated_at": "2025-08-10T15:29:08.740Z",
  "updated_by": "string",
  "teams": [
    {
      "additionalProp1": "string",
      "additionalProp2": "string",
      "additionalProp3": "string"
    }
  ],
  "mcp_access_groups": [
    "string"
  ],
  "mcp_info": {
    "server_name": "string",
    "description": "string",
    "logo_url": "string",
    "mcp_server_cost_info": {
      "default_cost_per_query": 0,
      "tool_name_to_cost_per_query": {
        "additionalProp1": 0,
        "additionalProp2": 0,
        "additionalProp3": 0
      }
    }
  },
  "status": "unknown",
  "last_health_check": "2025-08-10T15:29:08.740Z",
  "health_check_error": "string",
  "command": "string",
  "args": [
    "string"
  ],
  "env": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  }
}

## Get all MCP Servers

curl -X 'GET' \
  'https://ai.bitop.dev/v1/mcp/server' \
  -H 'accept: application/json'

Example Successful Response

[
  {
    "server_id": "string",
    "server_name": "string",
    "alias": "string",
    "description": "string",
    "url": "string",
    "transport": "sse",
    "spec_version": "2024-11-05",
    "auth_type": "none",
    "created_at": "2025-08-10T15:30:09.074Z",
    "created_by": "string",
    "updated_at": "2025-08-10T15:30:09.074Z",
    "updated_by": "string",
    "teams": [
      {
        "additionalProp1": "string",
        "additionalProp2": "string",
        "additionalProp3": "string"
      }
    ],
    "mcp_access_groups": [
      "string"
    ],
    "mcp_info": {
      "server_name": "string",
      "description": "string",
      "logo_url": "string",
      "mcp_server_cost_info": {
        "default_cost_per_query": 0,
        "tool_name_to_cost_per_query": {
          "additionalProp1": 0,
          "additionalProp2": 0,
          "additionalProp3": 0
        }
      }
    },
    "status": "unknown",
    "last_health_check": "2025-08-10T15:30:09.074Z",
    "health_check_error": "string",
    "command": "string",
    "args": [
      "string"
    ],
    "env": {
      "additionalProp1": "string",
      "additionalProp2": "string",
      "additionalProp3": "string"
    }
  }
]

## Get specific MCP server

Curl example where MCP server id is 123

curl -X 'GET' \
  'https://ai.bitop.dev/v1/mcp/server/123' \
  -H 'accept: application/json'

Example Successful Response

{
  "server_id": "string",
  "server_name": "string",
  "alias": "string",
  "description": "string",
  "url": "string",
  "transport": "sse",
  "spec_version": "2024-11-05",
  "auth_type": "none",
  "created_at": "2025-08-10T15:31:08.637Z",
  "created_by": "string",
  "updated_at": "2025-08-10T15:31:08.637Z",
  "updated_by": "string",
  "teams": [
    {
      "additionalProp1": "string",
      "additionalProp2": "string",
      "additionalProp3": "string"
    }
  ],
  "mcp_access_groups": [
    "string"
  ],
  "mcp_info": {
    "server_name": "string",
    "description": "string",
    "logo_url": "string",
    "mcp_server_cost_info": {
      "default_cost_per_query": 0,
      "tool_name_to_cost_per_query": {
        "additionalProp1": 0,
        "additionalProp2": 0,
        "additionalProp3": 0
      }
    }
  },
  "status": "unknown",
  "last_health_check": "2025-08-10T15:31:08.637Z",
  "health_check_error": "string",
  "command": "string",
  "args": [
    "string"
  ],
  "env": {
    "additionalProp1": "string",
    "additionalProp2": "string",
    "additionalProp3": "string"
  }
}

## Remove MCP Server

Curl example with MCP server with id 123

curl -X 'DELETE' \
  'https://ai.bitop.dev/v1/mcp/server/123' \
  -H 'accept: application/json'

example Successful Response

"string"