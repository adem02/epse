package auth

import (
	"github.com/adem02/epse/internal/utils/osutils"
)

type LiteAuthStrategy struct{}

func (s *LiteAuthStrategy) AddAuth() error {
	if err := CreateAllLiteAuthFiles(); err != nil {
		return err
	}

	if err := AddAuthRouteInIndexFile(); err != nil {
		return err
	}

	if err := osutils.AppendToFile(GetEnvFilePath(), "\n# JWT\nJWT_SECRET=your_secret_key_here\nJWT_EXPIRES_IN=24h\n"); err != nil {
		return err
	}

	return nil
}
