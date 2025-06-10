package service

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/escape-ship/paymentsrv/config"
	"github.com/escape-ship/paymentsrv/internal/infras/sqlc/postgresql"
	"github.com/escape-ship/paymentsrv/internal/provider/kakao"
	"github.com/escape-ship/paymentsrv/pkg/kafka"
	"github.com/escape-ship/paymentsrv/pkg/postgres"
	pb "github.com/escape-ship/paymentsrv/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PaymentServer implements the PaymentService gRPC service
type PaymentServer struct {
	pb.UnimplementedPaymentServiceServer
	kakaoClient *kakao.Client
	pg          postgres.DBEngine
	kafka       kafka.Engine
}

// NewPaymentServer creates a new payment service server
func NewPaymentServer(cfg *config.Config, pg postgres.DBEngine, kafkaEngine kafka.Engine) *PaymentServer {
	kakaoConfig := kakao.Config{
		BaseURL:   cfg.Kakao.BaseURL,
		SecretKey: cfg.Kakao.SecretKey,
		CID:       cfg.Kakao.CID,
	}
	return &PaymentServer{
		kakaoClient: kakao.NewClient(kakaoConfig),
		pg:          pg,
		kafka:       kafkaEngine,
	}
}

// KakaoReady initiates a payment process with Kakao
func (s *PaymentServer) KakaoReady(ctx context.Context, req *pb.KakaoReadyRequest) (*pb.KakaoReadyResponse, error) {
	slog.Info("Received KakaoReady request",
		"partner_order_id", req.PartnerOrderId,
		"partner_user_id", req.PartnerUserId,
		"item_name", req.ItemName,
		"total_amount", req.TotalAmount)

	if req.TotalAmount <= 0 {
		return nil, status.Error(codes.InvalidArgument, "total_amount must be greater than 0")
	}

	readyReq := kakao.ReadyRequest{
		PartnerOrderID: req.PartnerOrderId,
		PartnerUserID:  req.PartnerUserId,
		ItemName:       req.ItemName,
		Quantity:       req.Quantity,
		TotalAmount:    req.TotalAmount,
		TaxFreeAmount:  req.TaxFreeAmount,
		ApprovalURL:    "http://localhost:8080/success",
		FailURL:        "http://localhost:8080/fail",
		CancelURL:      "http://localhost:8080/cancel",
	}

	resp, err := s.kakaoClient.Ready(ctx, readyReq)
	if err != nil {
		slog.Error("Failed to process KakaoReady", "error", err)
		return nil, status.Error(codes.Internal, "failed to process payment")
	}

	db := s.pg.GetDB()
	querier := postgresql.New(db)

	tx, err := db.Begin()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get db transaction")
	}

	qtx := querier.WithTx(tx)

	params := postgresql.CreateKakaoParams{
		Tid:            resp.TID,
		PartnerOrderID: req.PartnerOrderId,
		PartnerUserID:  req.PartnerUserId,
		ItemName:       req.ItemName,
		Quantity:       req.Quantity,
		TotalAmount:    req.TotalAmount,
		TaxFreeAmount:  req.TaxFreeAmount,
	}

	tid, err := qtx.CreateKakao(ctx, params)
	if err != nil {
		slog.Error("Failed to create Kakao transaction", "error", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			slog.Error("Failed to rollback transaction", "error", rollbackErr)
		}
		return nil, status.Error(codes.Internal, "failed to create Kakao transaction")
	}
	err = tx.Commit()
	if err != nil {
		slog.Error("Failed to commit transaction", "error", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			slog.Error("Failed to rollback transaction", "error", rollbackErr)
		}
		return nil, status.Error(codes.Internal, "failed to commit transaction")
	}
	slog.Info("KakaoReady transaction committed successfully", "tid", tid)

	return &pb.KakaoReadyResponse{
		Tid:                   tid,
		NextRedirectAppUrl:    resp.NextRedirectAppURL,
		NextRedirectMobileUrl: resp.NextRedirectMobileURL,
		NextRedirectPcUrl:     resp.NextRedirectPCURL,
		AndroidAppScheme:      resp.AndroidAppScheme,
		IosAppScheme:          resp.IOSAppScheme,
	}, nil
}

// KakaoApprove approves a payment process with Kakao
func (s *PaymentServer) KakaoApprove(ctx context.Context, req *pb.KakaoApproveRequest) (*pb.KakaoApproveResponse, error) {
	slog.Info("Received KakaoApprove request",
		"tid", req.Tid,
		"partner_order_id", req.PartnerOrderId,
		"partner_user_id", req.PartnerUserId,
		"pg_token", req.PgToken)

	if req.PartnerOrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "partner_order_id is required")
	}
	if req.PgToken == "" {
		return nil, status.Error(codes.InvalidArgument, "pg_token is required")
	}

	approveReq := kakao.ApproveRequest{
		TID:            req.Tid,
		PartnerOrderID: req.PartnerOrderId,
		PartnerUserID:  req.PartnerUserId,
		PgToken:        req.PgToken,
	}

	resp, err := s.kakaoClient.Approve(ctx, approveReq)
	if err != nil {
		slog.Error("Failed to process KakaoApprove", "error", err)
		return nil, status.Error(codes.Internal, "failed to process payment approval")
	}

	db := s.pg.GetDB()
	querier := postgresql.New(db)

	tx, err := db.Begin()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get db transaction")
	}

	qtx := querier.WithTx(tx)
	params := postgresql.UpdateKakaoApproveParams{
		Tid:        resp.TID,
		Aid:        sql.NullString{String: resp.AID, Valid: true},
		ApprovedAt: sql.NullTime{Time: resp.ApprovedAt, Valid: true},
	}
	err = qtx.UpdateKakaoApprove(ctx, params)
	if err != nil {
		slog.Error("Failed to update Kakao transaction", "error", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			slog.Error("Failed to rollback transaction", "error", rollbackErr)
		}
		return nil, status.Error(codes.Internal, "failed to update Kakao transaction")
	}
	err = tx.Commit()
	if err != nil {
		slog.Error("Failed to commit transaction", "error", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			slog.Error("Failed to rollback transaction", "error", rollbackErr)
		}
		return nil, status.Error(codes.Internal, "failed to commit transaction")
	}
	slog.Info("KakaoApprove transaction committed successfully")

	// Publish Kafka message to 'kakao-approve' topic
	if s.kafka != nil {
		producer := s.kafka.Producer()
		if producer != nil {
			msgValue := []byte(req.PartnerOrderId)
			err := producer.Publish(ctx, []byte("kakao-approve"), msgValue)
			if err != nil {
				slog.Error("Failed to publish kakao-approve kafka message", "error", err)
			}
			slog.Info("Published kakao-approve message to Kafka", "partner_order_id", req.PartnerOrderId)
		}
	}

	return &pb.KakaoApproveResponse{
		PartnerOrderId: resp.PartnerOrderID,
	}, nil
}

// KakaoCancel cancels a payment with Kakao
func (s *PaymentServer) KakaoCancel(ctx context.Context, req *pb.KakaoCancelRequest) (*pb.KakaoCancelResponse, error) {
	slog.Info("Received KakaoCancel request",
		"order_id", req.PartnerOrderId,
		"cancel_amount", req.CancelAmount,
		"cancel_tax_free_amount", req.CancelTaxFreeAmount)

	if req.PartnerOrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "order_id is required")
	}

	// TODO: Implement Kakao payment cancellation logic using the kakaoClient
	return &pb.KakaoCancelResponse{
		PartnerOrderId: req.PartnerOrderId,
	}, nil
}
