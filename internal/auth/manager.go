package auth

import (
	"fmt"

	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type AuthManager struct {
	ProjectType typeutils.ProjectType
}

func NewAuthManager(projectType typeutils.ProjectType) *AuthManager {
	return &AuthManager{
		ProjectType: projectType,
	}
}

func (am *AuthManager) AddAuth() error {
	configData, err := config.ReadConfigFileData()
	if err != nil {
		return err
	}

	if configData.Auth {
		return fmt.Errorf("❌ auth already exists in this project")
	}

	strategy := GetAuthStrategy(am.ProjectType)
	if err := strategy.AddAuth(); err != nil {
		return err
	}

	if !config.DomainNameAlreadyExistsInRoutes(configData.Routes, "auth") {
		if err := config.AddNewRouteInConfigFile("auth", "auth", configData); err != nil {
			return err
		}
	}

	if err := config.SetAuthInConfigFile(configData); err != nil {
		return err
	}

	am.displaySuccess()

	return nil
}

func (am *AuthManager) displaySuccess() {
	fmt.Println()
	logutils.Logger{}.Success("✅ Auth added successfully!")
	fmt.Println()
	logutils.Logger{}.Info("📁 Files:")

	if am.ProjectType == typeutils.LiteProjectType {
		fmt.Printf("  ✓ src/config/jwt.config.ts\n")
		fmt.Printf("  ✓ src/types/auth.types.ts\n")
		fmt.Printf("  ✓ src/types/express.d.ts\n")
		fmt.Printf("  ✓ src/services/jwt.service.ts\n")
		fmt.Printf("  ✓ src/services/auth.service.ts\n")
		fmt.Printf("  ✓ src/middlewares/auth.middleware.ts\n")
		fmt.Printf("  ✓ src/controllers/auth/Login.controller.ts\n")
		fmt.Printf("  ✓ src/controllers/auth/Register.controller.ts\n")
		fmt.Printf("  ✓ src/routes/auth.routes.ts\n")
		fmt.Printf("  ✓ src/routes/index.ts (updated)\n")
		fmt.Println()
		logutils.Logger{}.Info("📋 Next steps:")
		fmt.Printf("  1. Fill in RegisterBody and LoginBody in src/types/auth.types.ts\n")
		fmt.Printf("  2. Implement login and register logic in src/services/auth.service.ts\n")
		fmt.Printf("  3. Install dependencies:\n")
		fmt.Printf("     npm install jsonwebtoken\n")
		fmt.Printf("     npm install -D @types/jsonwebtoken\n")
	} else {
		fmt.Printf("  ✓ src/config/jwt.config.ts\n")
		fmt.Printf("  ✓ src/entities/types/auth.types.ts\n")
		fmt.Printf("  ✓ src/useCases/gateway/JwtManager.interface.ts\n")
		fmt.Printf("  ✓ src/adapters/gateway/JwtManager.ts\n")
		fmt.Printf("  ✓ src/frameworks/tsoa/authentication.ts\n")
		fmt.Printf("  ✓ src/adapters/controllers/auth/login/Login.controller.ts\n")
		fmt.Printf("  ✓ src/adapters/controllers/auth/login/Login.dto.ts\n")
		fmt.Printf("  ✓ src/adapters/controllers/auth/register/Register.controller.ts\n")
		fmt.Printf("  ✓ src/adapters/controllers/auth/register/Register.dto.ts\n")
		fmt.Printf("  ✓ src/useCases/auth/login/Login.useCase.ts\n")
		fmt.Printf("  ✓ src/useCases/auth/login/Login.request.ts\n")
		fmt.Printf("  ✓ src/useCases/auth/login/Login.response.ts\n")
		fmt.Printf("  ✓ src/useCases/auth/register/Register.useCase.ts\n")
		fmt.Printf("  ✓ src/useCases/auth/register/Register.request.ts\n")
		fmt.Printf("  ✓ src/useCases/auth/register/Register.response.ts\n")
		fmt.Printf("  ✓ tsoa.json (updated)\n")
		fmt.Println()
		logutils.Logger{}.Info("📋 Next steps:")
		fmt.Printf("  1. Fill in RegisterRequest and LoginRequest interfaces\n")
		fmt.Printf("  2. Implement login and register logic in use cases\n")
		fmt.Printf("  3. Register JwtManager in your IoC container (src/frameworks/tsoa/services.ts):\n")
		fmt.Printf("     { token: JwtManager, class: JwtManager }\n")
		fmt.Printf("  4. Install dependencies:\n")
		fmt.Printf("     npm install jsonwebtoken\n")
		fmt.Printf("     npm install -D @types/jsonwebtoken\n")
	}

	fmt.Println()
	fmt.Printf("  ✓ .env (updated)\n")
	fmt.Printf("  ✓ epseconfig.json (updated)\n")
	fmt.Println()
}
