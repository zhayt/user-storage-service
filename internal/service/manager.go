package service

import (
	"context"
	"github.com/zhayt/user-storage-service/internal/model"
	"github.com/zhayt/user-storage-service/internal/storage"
	"go.uber.org/zap"
)

type IUserService interface {
	GetUserByID(ctx context.Context, userId int) (model.User, error)
	GetUserByEmail(ctx context.Context, login model.UserLogin) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (int, error)
	UpdateUserFIO(ctx context.Context, user model.UserUpdateFIO) (int, error)
	UpdateUserPassword(ctx context.Context, userUP model.UserUpdatePassword) (int, error)
	DeleteUser(ctx context.Context, userId int) error
}

type IBookService interface {
	CreateBook(ctx context.Context, book model.Book) (int, error)
	GetBookByID(ctx context.Context, bookId int) (model.Book, error)
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	UpdateBook(ctx context.Context, book model.Book) (int, error)
	DeleteBook(ctx context.Context, bookId int) error
}

type IBIHistoryService interface {
	CreateBIHistory(ctx context.Context, history model.BIHistory) (int, error)
	GetCurrentBorrowedBooks(ctx context.Context) ([]model.BorrowedBooks, error)
	GetBIHistoryLastMonth(ctx context.Context) ([]model.BorrowedBooks, error)
	UpdateBIHistory(ctx context.Context, bIHistoryID int) (int, error)
	DeleteBIHistory(ctx context.Context, bIHistoryID int) error
}

type Service struct {
	IUserService
	IBookService
	IBIHistoryService
}

func NewService(logger *zap.Logger, storage *storage.Storage) *Service {
	return &Service{
		IUserService:      NewUserService(logger, storage),
		IBookService:      NewBookService(logger, storage),
		IBIHistoryService: NewBIHistory(logger, storage),
	}
}
