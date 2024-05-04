package core

import "context"

func (s *ServiceImpl) UpdateScore(ctx context.Context, id int, score int) (error) {
	err := s.db.UpdateScore(ctx, id, score)
	if err != nil {
		return err
	}

	return err
}

func (s *ServiceImpl) UpdateMultiplicator(ctx context.Context, id int, mType string) (error) {
	tasks, err := s.db.GetUserTasks(ctx, uid)
	if err != nil {
		return nil, err
	}

	return err
}