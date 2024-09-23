package utils

import (
	"fmt"
	"strings"

	"github.com/saahil-mahato/authorization-service/internal/models"
)

func GeneratePrismaSchema(configs []*models.AuthorizationConfig) string {
	var sb strings.Builder

	sb.WriteString("model User {\n")
	sb.WriteString("  id    Int     @id @default(autoincrement())\n")
	sb.WriteString("  role  String\n")
	sb.WriteString("}\n\n")

	resourceMap := make(map[string]bool)
	for _, config := range configs {
		resourceMap[config.Resource] = true
	}

	for resource := range resourceMap {
		sb.WriteString(fmt.Sprintf("model %s {\n", strings.Title(resource)))
		sb.WriteString("  id    Int     @id @default(autoincrement())\n")
		sb.WriteString("  url   String\n")
		sb.WriteString("}\n\n")
	}

	sb.WriteString("model Permission {\n")
	sb.WriteString("  id        Int     @id @default(autoincrement())\n")
	sb.WriteString("  role      String\n")
	sb.WriteString("  resource  String\n")
	sb.WriteString("  canCreate Boolean\n")
	sb.WriteString("  canRead   Boolean\n")
	sb.WriteString("  canUpdate Boolean\n")
	sb.WriteString("  canDelete Boolean\n")
	sb.WriteString("}\n")

	return sb.String()
}
