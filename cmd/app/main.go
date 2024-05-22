package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Nomoke/wb-test-app/internal/config"
	"github.com/Nomoke/wb-test-app/internal/http-server/handler"
	"github.com/Nomoke/wb-test-app/internal/logger"
	"github.com/Nomoke/wb-test-app/internal/nats"
	"github.com/Nomoke/wb-test-app/internal/nats/subscriber"
	"github.com/Nomoke/wb-test-app/internal/service"
	"github.com/Nomoke/wb-test-app/internal/storage"
	"github.com/Nomoke/wb-test-app/internal/storage/cache"
	"github.com/go-chi/chi"
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func main() {
	// Создание корневого контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Настройка для обработки системных сигналов для корректного завершения работы
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение логгера
	log := logger.SetupLogger(cfg.Env)

	// Подключение БД
	conn, err := gorm.Open(postgres.Open(cfg.PostgresUrl), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	db := storage.New(conn, log)
	defer db.Close()

	// Инициализация кэша
	orderCache := cache.NewOrder(log)

	// Инициализация сервиса заказов
	orderService := service.New(db, orderCache)
	// Восстановление кэша из БД
	err = orderService.RecoverOrderCache(ctx)
	if err != nil {
		log.Error("failed to recover order cache", slog.Any("error", err))
	}

	// Подключение к nats серверу
	nc, err := nats.ConnectToServer(cfg.NustURL)
	if err != nil {
		log.Error("unable to connect to NATS", slog.Any("error", err))
		os.Exit(1)
	}
	defer nc.Close()

	// Подписка на поток
	subscriber.New(nc, orderService)

	// Инициализация HTTP - обработчика
	orderHandler := handler.New(*orderService)

	// Инициализация HTTP - сервера
	router := setupRouter(orderHandler, log.Logger)

	// Запуск HTTP-сервера
	server := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	go func() {
		log.Info("Запуск HTTP-сервера", slog.String("address", cfg.Address))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("HTTP server error", slog.Any("error", err))
			cancel()
		}
	}()

	// Ожидание сигнала для завершения работы
	<-sigs
	log.Info("Получен сигнал завершения, остановка сервера...")

	// Контекст для завершения сервера
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error("Ошибка при остановке сервера", slog.Any("error", err))
	}

	log.Info("Сервер остановлен")
}

func setupRouter(handler *handler.Order, log *slog.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetOrderByID(w, r, log)
	})
	return router
}
