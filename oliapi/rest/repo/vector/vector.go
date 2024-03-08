package vector

import (
	"context"
	"oliapi/domain"
	"oliapi/domain/repository"

	"github.com/google/uuid"
	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
)

const vectorDistance = pb.Distance_Cosine

func NewVectorRepo(grpc *grpc.ClientConn) Repo {
	return Repo{
		grpc: grpc,
	}
}

type Repo struct {
	grpc *grpc.ClientConn
}

// SaveVector implements repository.VectorRepository.
func (r Repo) SaveVector(ctx context.Context, data repository.SaveVectorData) error {
	pointsClient := pb.NewPointsClient(r.grpc)
	waitForUpsert := true
	_, err := pointsClient.Upsert(
		ctx,
		&pb.UpsertPoints{ //nolint:exhaustruct
			CollectionName: data.BotID.String(),
			Wait:           &waitForUpsert,
			Points: []*pb.PointStruct{
				{
					Id: &pb.PointId{
						PointIdOptions: &pb.PointId_Uuid{Uuid: data.DocumentID.String()},
					},
					Vectors: &pb.Vectors{
						VectorsOptions: &pb.Vectors_Vector{
							Vector: &pb.Vector{ //nolint:exhaustruct
								Data: data.Vector,
							},
						},
					},
				},
			},
		},
	)

	return err
}

// CreateCollection implements repository.VectorRepository.
func (r Repo) CreateCollection(ctx context.Context, botID uuid.UUID, embeddingProvider domain.EmbeddingProvider) error {
	collectionsClient := pb.NewCollectionsClient(r.grpc)
	_, err := collectionsClient.Create(
		ctx,
		&pb.CreateCollection{ //nolint:exhaustruct
			CollectionName: botID.String(),
			VectorsConfig: &pb.VectorsConfig{
				Config: &pb.VectorsConfig_Params{
					Params: &pb.VectorParams{ //nolint:exhaustruct
						Size:     uint64(embeddingProvider.VectorSize),
						Distance: vectorDistance,
					},
				},
			},
		},
	)

	return err
}
