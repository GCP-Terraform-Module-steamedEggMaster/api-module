output "enabled_api_services" {
  description = "모듈에서 활성화된 GCP API 서비스의 목록"
  value       = module.gcp_services.enabled_api_services
}