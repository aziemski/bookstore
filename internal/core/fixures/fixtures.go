package fixures

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/aziemski/bookstore/internal/core"
	"github.com/aziemski/bookstore/internal/x/xlog"
	"log/slog"
	"time"
)

//go:embed assets/fake_data.json
var fakeDataDir embed.FS

type (
	FakeData struct {
		string
		LoremIpsum string     `json:"lorem_ipsum"`
		Books      []FakeBook `json:"books"`
	}

	FakeBook struct {
		Title     string `json:"title"`
		Author    string `json:"author"`
		Summary   string `json:"summary"`
		ImageLink string `json:"image_link"`
		Category  string `json:"category"`
		Featured  bool   `json:"featured"`
	}
)

func readFakeData(log *slog.Logger) FakeData {
	result := FakeData{}

	file, err := fakeDataDir.Open("assets/fake_data.json")
	if err != nil {
		log.Error("unexpected os.Open() err", xlog.Err(err))
		return result
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		log.Error("unexpected decoder.Decode() err", xlog.Err(err))
		return result
	}

	return result
}

func Populate(repo *core.Repository, log *slog.Logger) {
	try := func() error {
		if c, err := repo.GetTotalCount(context.Background()); c > 0 && err == nil {
			log.Info("book records found ", "count", c)
			return nil
		}
		time.Sleep(3 * time.Second)

		fd := readFakeData(log)
		for i, fb := range fd.Books {
			log.Info("inserting book")
			_, err := repo.InsertNew(context.Background(), &core.Book{
				ID:          fmt.Sprintf("b%d", i),
				Title:       fb.Title,
				Author:      fb.Author,
				Summary:     fb.Summary,
				Description: fd.LoremIpsum,
				ImageLink:   fb.ImageLink,
				Category:    fb.Category,
				Featured:    fb.Featured,
			})
			if err != nil {
				log.Warn("cannot save book", xlog.Err(err))
				return err
			}
		}

		return nil
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			log.Warn("received cancellation signal due to deadline.")
			return
		default:
			if err := try(); err == nil {
				return
			}
		}
	}
}
