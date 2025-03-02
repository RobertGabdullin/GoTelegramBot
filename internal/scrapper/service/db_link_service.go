package service

import (
	"fmt"
	"github.com/RobertGabdullin/GoTelegramBot/internal/scrapper/storage"
)

type DBLinkService struct {
	linkTrackerStorage storage.LinkTrackerStorage
}

func NewDBLinkService(linkTrackerStorage storage.LinkTrackerStorage) *DBLinkService {
	return &DBLinkService{
		linkTrackerStorage: linkTrackerStorage,
	}
}

func (s *DBLinkService) Register(chatId int) error {
	exist, err := s.linkTrackerStorage.IsChatPresent(chatId)
	if err != nil {
		return fmt.Errorf("error while checking if chat is present: %w", err)
	}

	if exist {
		return fmt.Errorf("chat is already present")
	}

	err = s.linkTrackerStorage.AddChat(chatId)
	if err != nil {
		return fmt.Errorf("error while adding chat: %w", err)
	}
	return nil
}

func (s *DBLinkService) Track(chatId int, link string) error {
	exist, err := s.linkTrackerStorage.IsChatPresent(chatId)
	if err != nil {
		return fmt.Errorf("error while checking if chat is present: %w", err)
	}
	if !exist {
		return fmt.Errorf("chat is not present")
	}
	exist, err = s.linkTrackerStorage.IsLinkPresent(link)
	if err != nil {
		return fmt.Errorf("error while checking if link is present: %w", err)
	}
	if !exist {
		err = s.linkTrackerStorage.AddLink(link)
		if err != nil {
			return fmt.Errorf("error while adding link: %w", err)
		}
	}
	linkId, err := s.linkTrackerStorage.GetIdByLink(link)
	if err != nil {
		return fmt.Errorf("error while getting link id: %w", err)
	}

	exist, err = s.linkTrackerStorage.IsChatLinkPresent(chatId, linkId)
	if err != nil {
		return fmt.Errorf("error while checking if chat link is present: %w", err)
	}

	if !exist {
		err = s.linkTrackerStorage.AddChatLink(chatId, linkId)
		if err != nil {
			return fmt.Errorf("error while adding chat link: %w", err)
		}
	} else {
		return fmt.Errorf("chat link is already present")
	}

	return nil
}

func (s *DBLinkService) Untrack(chatId int, link string) error {
	exist, err := s.linkTrackerStorage.IsChatPresent(chatId)
	if err != nil {
		return fmt.Errorf("error while checking if chat is present: %w", err)
	}
	if !exist {
		return fmt.Errorf("chat is not present")
	}
	return nil
}

func (s *DBLinkService) List(chatId int) ([]string, error) {
	return nil, nil
}
