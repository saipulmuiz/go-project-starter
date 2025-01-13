package queries

const (
	InsertCategory = `
		INSERT INTO categories
			(
				category_name,
				created_at,
				updated_at
			)
		VALUES
			($1, $2, $3)
		RETURNING category_id;
	`

	GetCategories = `
		SELECT
			category_id,
			category_name,
			created_at,
			updated_at
		FROM
			categories
	`

	GetCategoryByID = `
		SELECT
			category_id,
			category_name,
			created_at,
			updated_at
		FROM
			categories
		WHERE
			category_id = $1
	`

	UpdateCategoryByID = `
		UPDATE categories
		SET
			category_name = $2,
			update_at = $3
		WHERE
			category_id = $1
		RETURNING category_id;
	`

	DeleteCategoryByID = `
		DELETE categories WHERE category_id = $1;
	`
)
