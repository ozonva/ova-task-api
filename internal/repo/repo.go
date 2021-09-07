package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
	"ozonva/ova-task-api/internal/kafka"
	. "ozonva/ova-task-api/pkg/entities/tasks"
	"time"
)

type Repo interface {
	AddTasks(ctx context.Context, tasks []Task) error
	ListTasks(ctx context.Context, limit, offset uint64) ([]Task, error)
	DescribeTask(ctx context.Context, taskId uint64) (*Task, error)
	RemoveTask(ctx context.Context, taskId uint64) error
	UpdateTask(ctx context.Context, task Task) error
}

// todo move to utils
type sqlQuery interface {
	ToSql() (string, []interface{}, error)
}

type repo struct {
	db *sql.DB
}

func NewRepo(sqlDb *sql.DB) Repo {
	return &repo{
		db: sqlDb,
	}
}

func queryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func OpenDb(connectionString string) (*sql.DB, error) {
	return sql.Open("pgx", connectionString)
}

func (repo *repo) RemoveTask(ctx context.Context, taskId uint64) error {
	query := queryBuilder().Delete("tasks").Where(sq.Eq{"id": taskId})
	err := logQuery(query)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	result, err := query.RunWith(repo.db).ExecContext(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	logExecResult(result)

	err = kafka.SendMessageWithContext(ctx, "remove_task", taskId)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	return nil
}

func (repo *repo) AddTasks(ctx context.Context, tasks []Task) error {
	query := queryBuilder().Insert("tasks").Columns("userid", "description", "created_at")
	for _, task := range tasks {
		query = query.Values(task.UserId, task.Description, task.DateCreated)
	}
	err := logQuery(query)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	result, err := query.RunWith(repo.db).ExecContext(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	logExecResult(result)

	err = kafka.SendMessageWithContext(ctx, "add_tasks", tasks)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}

func (repo *repo) ListTasks(ctx context.Context, limit, offset uint64) ([]Task, error) {
	query := queryBuilder().
		Select("id", "userid", "description", "created_at").
		From("tasks").
		OrderBy("created_at").
		Limit(limit).
		Offset(offset)

	err := logQuery(query)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	rows, err := query.RunWith(repo.db).QueryContext(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error().Msg("cant close rows")
		}
	}(rows)
	tasks := make([]Task, 0, limit)
	for rows.Next() {
		task, err := scanTask(rows)
		if err != nil {
			log.Error().Err(err).Send()
			continue
		}
		tasks = append(tasks, *task)
	}
	if err := rows.Err(); err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return tasks, nil
}

func (repo *repo) DescribeTask(ctx context.Context, taskId uint64) (*Task, error) {
	query := queryBuilder().
		Select("id", "userid", "description", "created_at").
		From("tasks").
		Where(sq.Eq{"id": taskId})

	err := logQuery(query)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	rowScanner := query.RunWith(repo.db).QueryRowContext(ctx)
	task, err := scanTask(rowScanner)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	return task, nil
}

func (repo *repo) UpdateTask(ctx context.Context, task Task) error {
	query := queryBuilder().
		Update("tasks").
		Set("userid", task.UserId).
		Set("description", task.Description).
		Where(sq.Eq{"id": task.TaskId})

	err := logQuery(query)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	result, err := query.RunWith(repo.db).ExecContext(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	logExecResult(result)
	err = kafka.SendMessageWithContext(ctx, "update_task", task)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}

func logExecResult(result sql.Result) {
	log.Debug().Msgf("%v", result)
}

func logQuery(query sqlQuery) error {
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}
	log.Debug().Msgf("%v %v", sqlQuery, args)
	return nil
}

func scanTask(rows sq.RowScanner) (*Task, error) {
	var taskId, userId uint64
	var description string
	var createAt time.Time
	err := rows.Scan(&taskId, &userId, &description, &createAt)
	if err != nil {
		return nil, err
	}
	task := New(userId, taskId, description, createAt)
	return task, nil
}
