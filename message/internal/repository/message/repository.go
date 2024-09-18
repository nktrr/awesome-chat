package message

import (
	"awesome-chat/message/internal/config"
	"awesome-chat/message/internal/model"
	"awesome-chat/message/internal/repository"
	"awesome-chat/message/internal/repository/message/converter"
	modelRepo "awesome-chat/message/internal/repository/message/model"
	"context"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"log"
)

var _ repository.MessageRepository = (*repo)(nil)

type repo struct {
	scyllaConfig config.ScyllaConfig
	cluster      *gocql.ClusterConfig
	session      *gocql.Session
}

func NewRepository(scyllaConfig config.ScyllaConfig) (*repo, error) {
	cluster := gocql.NewCluster(scyllaConfig.GetAddress())
	cluster.Keyspace = scyllaConfig.GetKeyspace()
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &repo{
		scyllaConfig: scyllaConfig,
		cluster:      cluster,
		session:      session,
	}, nil
}

func (r *repo) Create(ctx context.Context, message *model.Message) error {
	println("create in repo")

	if err := r.session.Query(`INSERT INTO messages (id, msg_text) VALUES (?, ?)`,
		gocql.MustRandomUUID(), message.Text).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}

	return
}

func (r *repo) Get(ctx context.Context, id uuid.UUID) (*model.Message, error) {
	var msg modelRepo.Message

	err := r.session.Query(`SELECT * FROM messages WHERE id == ?`, id).WithContext(ctx).Scan(&msg.ID, msg.Text)
	if err != nil {
		return nil, err
	}

	return converter.ToMessageFromRepo(&msg), nil
}
