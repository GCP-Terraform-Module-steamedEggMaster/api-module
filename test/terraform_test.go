package test

import (
	"os"
	"strings"
	"strconv"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
)

func TestTerraformModule(t *testing.T) {
	// 환경 변수에서 민감 정보 가져오기
	projectID := os.Getenv("GCP_PROJECT_ID")

	apiServicesRaw := os.Getenv("API_SERVICES")
	apiServices := strings.Split(apiServicesRaw, ",")

	disableOnDestroyRaw := os.Getenv("DISABLE_ON_DESTROY")
	disableOnDestroy := strconv.ParseBool(disableOnDestroyRaw)
	
	// Terraform 모듈 옵션 설정
	terraformOptions := &terraform.Options{
		TerraformDir: "../", // 모듈 경로
		Vars: map[string]interface{}{
			"project_id":  projectID,
			"api_services":  apiServices,
			"disable_on_destroy":  disableOnDestroy,
		},
	}

	// Terraform Destroy 실행 (테스트 완료 후 정리)
	defer terraform.Destroy(t, terraformOptions)
	// TestTerraformModule이 끝나면 실행됨.
	// 에러 상황 발생 시에도 실행됨.

	// Terraform Apply 실행
	terraform.InitAndApply(t, terraformOptions)

	// 출력값 검증
	// Terraform이 활성화한 API 목록 가져오기
	enabledAPIs := terraform.OutputList(t, terraformOptions, "enabled_apis")
	// 활성화한 API에 활성화하고자 한 API가 포함되어 있는지 확인하기
	assert.Subset(t, enabledAPIs, apiServices, "Terraform이 활성화한 API가 예상과 다릅니다.")
}