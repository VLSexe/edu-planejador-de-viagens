package api

import (
	"context"
	"errors"
	"journey/internal/api/spec"
	"journey/internal/pgstore"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type store interface {
	GetParticipant(ctx context.Context, participantID uuid.UUID) (pgstore.Participant, error)
	ConfirmParticipant(ctx context.Context, participantID uuid.UUID) error
}

type API struct{
	store store
	logger *zap.Logger
}

// NewApi initializes a new API with the provided store and logger.
func NewApi(pool *pgxpool.Pool, looger *zap.Logger) API {
	return API{pgstore.New(pool), looger}
}

// Confirms a participant on a trip.
// (PATCH /participants/{participantId}/confirm)
func (api API) PatchParticipantsParticipantIDConfirm(w http.ResponseWriter, r *http.Request, participantID string) *spec.Response {
	id, err := uuid.Parse(participantID)
	if err!= nil {
        return spec.PatchParticipantsParticipantIDConfirmJSON400Response(spec.Error{Message: "uuid invalido"})
    }
	
	partcipant, err := api.store.GetParticipant(r.Context(), id)
	if err!= nil {
        if errors.Is(err, pgx.ErrNoRows) {
			return spec.PatchParticipantsParticipantIDConfirmJSON400Response(spec.Error{Message: "participant não encontrado"})
		}
		api.logger.Error("falha consultar participante", zap.Error(err), zap.String("participant_id",participantID))
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(
			spec.Error{Message: "Ocorreu um erro inesperado, tente novamente mais tarde."},
		)
    }
	if partcipant.IsConfirmed {
		return spec.GetTripsTripIDActivitiesJSON400Response(
			spec.Error{Message: "Participante já confirmado"},
		)
	}
	if err := api.store.ConfirmParticipant(r.Context(), id); err != nil {
		api.logger.Error("falha ao confirmar participante", zap.Error(err), zap.String("participant_id",participantID))
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(
			spec.Error{Message: "Ocorreu um erro inesperado, tente novamente mais tarde."},
		)
	}
	return spec.PatchParticipantsParticipantIDConfirmJSON204Response(nil)
}

// Create a new trip
// (POST /trips)
func (api API) PostTrips(w http.ResponseWriter, r *http.Request) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip details.
// (GET /trips/{tripId})
func (api API) GetTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Update a trip.
// (PUT /trips/{tripId})
func (api API) PutTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip activities.
// (GET /trips/{tripId}/activities)
func (api API) GetTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Create a trip activity.
// (POST /trips/{tripId}/activities)
func (api API) PostTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Confirm a trip and send e-mail invitations.
// (GET /trips/{tripId}/confirm)
func (api API) GetTripsTripIDConfirm(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Invite someone to the trip.
// (POST /trips/{tripId}/invites)
func (api API) PostTripsTripIDInvites(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip links.
// (GET /trips/{tripId}/links)
func (api API) GetTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Create a trip link.
// (POST /trips/{tripId}/links)
func (api API) PostTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}

// Get a trip participants.
// (GET /trips/{tripId}/participants)
func (api API) GetTripsTripIDParticipants(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	panic("not implemented") // TODO: Implement
}
