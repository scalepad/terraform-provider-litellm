# Example LiteLLM Vector Store Configuration

# Pinecone vector store with credential
resource "litellm_credential" "pinecone_cred" {
  credential_name = "pinecone-production"
  
  credential_info = {
    provider    = "pinecone"
    environment = "production"
    region      = "us-east-1"
  }
  
  credential_values = {
    api_key    = var.pinecone_api_key
    index_name = "document-embeddings"
  }
}

resource "litellm_vector_store" "pinecone_store" {
  vector_store_name        = "production-documents"
  custom_llm_provider      = "pinecone"
  litellm_credential_name  = litellm_credential.pinecone_cred.credential_name
  
  vector_store_description = "Production vector store for document embeddings"
  
  vector_store_metadata = {
    environment = "production"
    team        = "ai-team"
    purpose     = "document-search"
    version     = "v1.0"
  }
  
  litellm_params = {
    dimension = "1536"
    metric    = "cosine"
    pods      = "1"
    replicas  = "1"
    pod_type  = "p1.x1"
  }
}

# Weaviate vector store
resource "litellm_credential" "weaviate_cred" {
  credential_name = "weaviate-cluster"
  
  credential_info = {
    provider    = "weaviate"
    cluster_url = "https://your-cluster.weaviate.network"
  }
  
  credential_values = {
    api_key = var.weaviate_api_key
  }
}

resource "litellm_vector_store" "weaviate_store" {
  vector_store_name        = "weaviate-knowledge-base"
  custom_llm_provider      = "weaviate"
  litellm_credential_name  = litellm_credential.weaviate_cred.credential_name
  
  vector_store_description = "Weaviate vector store for knowledge base"
  
  vector_store_metadata = {
    class_name     = "Document"
    schema_version = "v1"
    embedding_model = "text-embedding-ada-002"
  }
  
  litellm_params = {
    class_name = "Document"
    vectorizer = "text2vec-openai"
  }
}

# Local Chroma vector store (no credentials needed)
resource "litellm_vector_store" "chroma_store" {
  vector_store_name   = "chroma-local-docs"
  custom_llm_provider = "chroma"
  
  vector_store_description = "Local Chroma vector store for development"
  
  litellm_params = {
    host            = "localhost"
    port            = "8000"
    collection_name = "documents"
  }
  
  vector_store_metadata = {
    environment     = "development"
    embedding_model = "sentence-transformers/all-MiniLM-L6-v2"
    chunk_size      = "512"
    chunk_overlap   = "50"
  }
}

# Qdrant vector store
resource "litellm_credential" "qdrant_cred" {
  credential_name = "qdrant-cluster"
  
  credential_info = {
    provider = "qdrant"
    host     = "your-qdrant-cluster.com"
    port     = "6333"
  }
  
  credential_values = {
    api_key = var.qdrant_api_key
  }
}

resource "litellm_vector_store" "qdrant_store" {
  vector_store_name        = "qdrant-semantic-search"
  custom_llm_provider      = "qdrant"
  litellm_credential_name  = litellm_credential.qdrant_cred.credential_name
  
  vector_store_description = "Qdrant vector store for semantic search"
  
  vector_store_metadata = {
    collection_name = "documents"
    distance_metric = "cosine"
    vector_size     = "1536"
  }
  
  litellm_params = {
    collection_name = "documents"
    distance        = "cosine"
    vector_size     = "1536"
  }
}

# Variables for sensitive values
variable "pinecone_api_key" {
  description = "Pinecone API key"
  type        = string
  sensitive   = true
}

variable "weaviate_api_key" {
  description = "Weaviate API key"
  type        = string
  sensitive   = true
}

variable "qdrant_api_key" {
  description = "Qdrant API key"
  type        = string
  sensitive   = true
}

# Outputs
output "pinecone_vector_store_id" {
  description = "ID of the Pinecone vector store"
  value       = litellm_vector_store.pinecone_store.vector_store_id
}

output "weaviate_vector_store_id" {
  description = "ID of the Weaviate vector store"
  value       = litellm_vector_store.weaviate_store.vector_store_id
}

output "chroma_vector_store_id" {
  description = "ID of the Chroma vector store"
  value       = litellm_vector_store.chroma_store.vector_store_id
}
