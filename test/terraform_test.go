package test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformModule(t *testing.T) {
	// 환경 변수에서 민감 정보 가져오기
	projectID := os.Getenv("GCP_PROJECT_ID")

	apiServicesRaw := os.Getenv("API_SERVICES")
	apiServices := strings.Split(apiServicesRaw, ",")

	// Terraform 모듈 옵션 설정
	terraformOptions := &terraform.Options{
		TerraformDir: "../", // 모듈 경로
		Vars: map[string]interface{}{
			"project_id":  projectID,
			"api_services":  apiServices,
		},
	}

	// Terraform Destroy 실행 (테스트 완료 후 정리)
	defer terraform.Destroy(t, terraformOptions)
	// TestTerraformModule이 끝나면 실행됨.
	// 에러 상황 발생 시에도 실행됨.

	// Terraform Apply 실행
	terraform.InitAndApply(t, terraformOptions)

	// 출력값 검증
	// Terraform 출력값 가져오기
	output := terraform.Output(t, terraformOptions, "enabled_apis")

	// 출력값을 JSON 배열로 파싱
	var enabledAPIs []string
	err := json.Unmarshal([]byte(output), &enabledAPIs)
	if err != nil {
		t.Fatalf("출력값 파싱 실패: %v", err)
	}
	
	// 활성화한 API에 활성화하고자 한 API가 포함되어 있는지 확인하기
	assert.Subset(t, enabledAPIs, apiServices, "Terraform이 활성화한 API가 예상과 다릅니다.")
}
