package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks ozonva/ova-task-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks ozonva/ova-task-api/internal/flusher Flusher
