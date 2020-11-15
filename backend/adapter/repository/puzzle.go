package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"

	"backend/model"
	"backend/usecase/repository"
)

type puzzleRepository struct {
	db *gorm.DB
}

func NewPuzzleRepository(db *gorm.DB) repository.PuzzleRepository {
	db.AutoMigrate(&model.Puzzlepiece{})
	return &puzzleRepository{db}
}

func (sr *puzzleRepository) Put(surveyId string, questionId string, pieces []*model.Puzzlepiece) ([]*model.Puzzlepiece, error) {

	updatedPieces := make([]*model.Puzzlepiece, 0)
	for _, p := range pieces {
		if p.Image == "" {
			err := sr.deletePiece(surveyId, questionId, p)
			if err != nil {
				log.Error().Err(err).Msg("unable to delete piece")
			}
			continue
		}
		err := sr.updatePiece(surveyId, questionId, p)
		if err != nil {
			log.Error().Err(err).Msg("unable to update piece")
		}
		updatedPieces = append(updatedPieces, p)
	}
	return updatedPieces, nil
}

func (sr *puzzleRepository) updatePiece(surveyId string, questionId string, piece *model.Puzzlepiece) error {
	var loadedPiece model.Puzzlepiece
	err := sr.db.Where("question_id = ? and position = ?", questionId, piece.Position).First(&loadedPiece).Error
	if err == nil {
		loadedPiece.Image = piece.Image
		loadedPiece.Tapped = piece.Tapped
		sr.db.Save(&loadedPiece)
		return nil
	}

	err = sr.addPiece(surveyId, questionId, piece)
	if err != nil {
		return err
	}
	return nil
}

func (sr *puzzleRepository) addPiece(surveyId string, questionId string, piece *model.Puzzlepiece) error {
	/*
		var loadedPiece model.Puzzlepiece
		err := sr.db.Where("question_id = ? and position = ?", questionId, piece.Position).First(&loadedPiece).Error
		if err != nil {
			return err
		}
	*/
	sr.db.Save(piece)
	return nil
}

func (sr *puzzleRepository) deletePiece(surveyId string, questionId string, piece *model.Puzzlepiece) error {
	var loadedPiece model.Puzzlepiece
	err := sr.db.Where("question_id = ? and position = ?", questionId, piece.Position).First(&loadedPiece).Error
	if err != nil {
		return err
	}
	sr.db.Delete(loadedPiece)
	return nil
}
