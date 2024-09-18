package message

import (
	"awesome-chat/gateway/internal/config"
	"awesome-chat/gateway/internal/model"
	"awesome-chat/gateway/internal/repository"
	"context"
	"github.com/gocql/gocql"
	"log"
)

var _ repository.MessageRepository = (*repo)(nil)

type repo struct {
	scyllaConfig config.ScyllaConfig
	cluster      *gocql.ClusterConfig
	session      *gocql.Session
}

//TODO [09.09.2024] put some scylla logic

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

func (r *repo) Create(ctx context.Context, message model.Message) error {
	println("create in repo")

	if err := r.session.Query(`INSERT INTO messages (id, msg_text) VALUES (?, ?)`,
		gocql.MustRandomUUID(), message.Text).WithContext(ctx).Exec(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (r *repo) Get(ctx context.Context) error {
	println("get")
	return nil
}
