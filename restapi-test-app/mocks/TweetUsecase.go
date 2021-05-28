// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entities "restapi-tested-app/entities"

	mock "github.com/stretchr/testify/mock"
)

// TweetUsecase is an autogenerated mock type for the TweetUsecase type
type TweetUsecase struct {
	mock.Mock
}

// CreateTweet provides a mock function with given fields: tweet
func (_m *TweetUsecase) CreateTweet(tweet *entities.Tweet) error {
	ret := _m.Called(tweet)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Tweet) error); ok {
		r0 = rf(tweet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTweet provides a mock function with given fields: id
func (_m *TweetUsecase) DeleteTweet(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTweets provides a mock function with given fields:
func (_m *TweetUsecase) GetAllTweets() (*[]entities.Tweet, error) {
	ret := _m.Called()

	var r0 *[]entities.Tweet
	if rf, ok := ret.Get(0).(func() *[]entities.Tweet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.Tweet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTweetByID provides a mock function with given fields: id
func (_m *TweetUsecase) GetTweetByID(id int) (*entities.Tweet, error) {
	ret := _m.Called(id)

	var r0 *entities.Tweet
	if rf, ok := ret.Get(0).(func(int) *entities.Tweet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Tweet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchTweetByText provides a mock function with given fields: text
func (_m *TweetUsecase) SearchTweetByText(text string) (*[]entities.Tweet, error) {
	ret := _m.Called(text)

	var r0 *[]entities.Tweet
	if rf, ok := ret.Get(0).(func(string) *[]entities.Tweet); ok {
		r0 = rf(text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.Tweet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTweet provides a mock function with given fields: tweet
func (_m *TweetUsecase) UpdateTweet(tweet *entities.Tweet) error {
	ret := _m.Called(tweet)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Tweet) error); ok {
		r0 = rf(tweet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}