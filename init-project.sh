#!/bin/bash

set -e

echo "🚀 Initializing user-admin DDD project layout..."

mkdir -p \
  api \
  application \
  domain/{entity,repository,service} \
  infra/{db,auth,logger} \
  routes \
  conf \
  di \
  devops \
  .github/workflows

touch \
  conf/app.yml \
  conf/route.yml \
  conf/config.go \
  di/wire.go \
  routes/router.go \
  main.go \
  devops/README.md \
  .env.example \
  README.md

echo "✅ Basic directories and files created."

## main
cat <<EOF > main.go
package main

func main() {

}
EOF

## package routes
cat <<EOF > routes/router.go
package routes

EOF

## wire
cat <<EOF > di/wire.go
//go:build wireinject

package di
EOF


# Add .gitignore
cat <<EOF > .gitignore
.env
.idea
bin/
*.log
*.exe
*.out
coverage.out
EOF

# Add .env.example
cat <<EOF > .env.example
# ┌────────────────────────────────────────────┐
# │           User Admin Service ENV           │
# └────────────────────────────────────────────┘

APP_ENV=prod
APP_PORT=8080
JWT_SECRET=your_very_strong_jwt_secret_here
JWT_EXPIRATION_MINUTES=60
DATABASE_DSN=user:password@tcp(127.0.0.1:3306)/user_admin?charset=utf8mb4&parseTime=True&loc=Local
LOG_LEVEL=info
EOF

# Initialize Go module
read -p "📦 Enter your module name (e.g. github.com/yourname/user-admin): " module_name
go mod init "$module_name"
go mod tidy
go get github.com/labstack/echo/v4
go get github.com/google/wire
go get github.com/golang-jwt/jwt/v5
go get github.com/sirupsen/logrus
go get gorm.io/gorm
go get github.com/sirupsen/logrus
go get gorm.io/driver/mysql


echo "📦 Go dependencies installed."

echo "📁 Project initialized. Next steps:"
echo "----------------------------------------"
echo "✅ Start coding in: api/, domain/, application/"
echo "✅ Edit config in: conf/app.yml, route.yml"
echo "✅ Setup GitHub Actions in: .github/workflows/"
echo "✅ Use systemd to deploy with .env"
echo "✅ Use 'wire' to generate di/wire_gen.go"

echo "🔥 You're ready to build!"
