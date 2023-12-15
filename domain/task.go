package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

type Packet struct {
	ID           primitive.ObjectID `bson:"_id" json:"-"`
	Timestamp    time.Time          `bson:"timestamp" form:"timestamp" binding:"required" json:"timestamp"`
	Length       int                `bson:"length" form:"length" binding:"required" json:"length"`
	SrcMAC       string             `bson:"srcMAC" form:"srcMAC" binding:"required" json:"srcMAC"`
	DstMAC       string             `bson:"dstMAC" form:"dstMAC" binding:"required" json:"dstMAC"`
	SrcIP        string             `bson:"srcIP" form:"srcIP" binding:"required" json:"srcIP"`
	DstIP        string             `bson:"dstIP" form:"dstIP" binding:"required" json:"dstIP"`
	SrcPort      uint16             `bson:"srcPort" form:"srcPort" binding:"required" json:"srcPort"`
	DstPort      uint16             `bson:"dstPort" form:"dstPort" binding:"required" json:"dstPort"`
	MatchedRules []string           `bson:"matchedRules" form:"matchedRules" binding:"required" json:"matchedRules"`
}

type PacketRepository interface {
	Create(c context.Context, task *Packet) error
	FetchAll(c context.Context) ([]Packet, error)
}

type PacketUsecase interface {
	Create(c context.Context, task *Packet) error
	FetchAll(c context.Context) ([]Packet, error)
}
