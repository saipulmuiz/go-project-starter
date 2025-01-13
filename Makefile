generate-mocks:
	# repositories
	@mockgen -destination=./service/repository/mocks/mock_user_repository.go -package=mocks github.com/saipulmuiz/go-project-starter/service UserRepository
	@mockgen -destination=./service/repository/mocks/mock_category_repository.go -package=mocks github.com/saipulmuiz/go-project-starter/service CategoryRepository