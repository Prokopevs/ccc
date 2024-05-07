build_images:
	docker build -t game/auth-api:latest -f ./auth/docker/Dockerfile ./auth
	docker build -t game/game-api:latest -f ./game/docker/Dockerfile ./game
	docker build -t game/users-api:latest -f ./users/docker/Dockerfile ./users
