output "enabled_apis" {
  description = "활성화된 API 목록"
  value       = google_project_service.enabled_services[*].service
}