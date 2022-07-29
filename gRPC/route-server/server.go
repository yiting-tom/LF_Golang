package main

import (
	"context"

	pb "github.com/yiting-tom/LF_Golang/grpc/route"
	"google.golang.org/protobuf/proto"
)

// routeGuideServer is a mock implementation of RouteGuideServer.
type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	features []*pb.Feature
}

// dbServer is a simple example mock database.
func dbServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
			{
				Name: "Taiwan",
				Location: &pb.Point{
					Latitude:  37,
					Longitude: 122,
				},
			},
			{
				Name: "United States",
				Location: &pb.Point{
					Latitude:  37,
					Longitude: 95,
				},
			},
		},
	}
}

// GetFeature returns the feature at the given point.
func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, nil
}
