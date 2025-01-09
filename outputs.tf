output "enabled_api_services" {
  description = "활성화된 GCP API 서비스의 목록"
  value       = keys(google_project_service.enabled_services)
}

output "enabled_api_ids" {
  description = "활성화된 GCP API 서비스의 고유 식별자 목록"
  value       = [for key in keys(google_project_service.enabled_services) : google_project_service.enabled_services[key].id]
}

output "enabled_api_names" {
  description = "활성화된 GCP API 서비스 이름들"
  value       = [for key in keys(google_project_service.enabled_services) : google_project_service.enabled_services[key].service]
}

output "project_id" {
  description = "활성화된 GCP API 서비스가 속한 프로젝트 ID"
  value       = var.project
}