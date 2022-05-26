package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	gkitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"aegis/internal/api/md"
	"aegis/internal/model"
	collectionDao "aegis/internal/module/collection_dao"
	"aegis/internal/response"
)

func makeAddToCollectionHandler(
	router *mux.Router,
	logger gkitLog.Logger,
	mysqlConn *sqlx.DB,
) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqPayload := md.PopulateJSONPayload(ctx)
		vars := mux.Vars(r)
		rCollectionId := vars["collection_id"]
		collectionId, err := strconv.Atoi(rCollectionId)
		if err != nil {
			panic(err)
		}

		rItem := reqPayload.Get("item")
		if !rItem.Exists() {
			_, _ = response.ResponseBadRequest.WriteTo(w)
			return
		}

		_ = level.Info(logger).Log("message", fmt.Sprintf("add %s to collection %d", rItem.String(), collectionId))

		tx, err := mysqlConn.BeginTxx(ctx, nil)
		if err != nil {
			panic(err)
		}
		err = collectionDao.CreateCollectionItem(ctx, tx, &model.CollectionItem{
			CollectionId: int64(collectionId),
			Value:        rItem.String(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		if err != nil {
			if err := tx.Rollback(); err != nil {
				_ = level.Error(logger).Log("message", "DB rollback failed", "error", err, "culprit", "api.AddToCollectionHandler")
			}

			panic(err)
		}
		if err := tx.Commit(); err != nil {
			_ = level.Error(logger).Log("message", "DB commit failed", "error", err, "culprit", "api.AddToCollectionHandler")
		}

		_, _ = response.Make().WriteTo(w)
	}

	var h http.Handler
	h = http.HandlerFunc(f)
	h = md.Chain(
		md.LogRequestDuration(logger, "add_to_collection"),
		md.RequireJSONPayload,
	)(h)

	router.Handle("/collections/{collection_id}/items", h).Methods("POST")

	return router
}
