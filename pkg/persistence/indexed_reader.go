package persistence

import (
	"fmt"

	"github.com/emirpasic/gods/trees/btree"
	"github.com/emirpasic/gods/utils"
	"github.com/pkg/errors"
)

type IndexedReader struct {
	Reader
	index *btree.Tree

	curr_message int
}

func NewIndexedReader(filename string) (indexedLogReader *IndexedReader, err error) {
	indexedLogReader = new(IndexedReader)

	logReader, err := NewReader(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Could not open log file: "+filename)
	}
	indexedLogReader.file = logReader.file
	indexedLogReader.reader = logReader.reader

	// time index for quick iterating, searching and jumping
	indexedLogReader.index = btree.NewWith(3, utils.Int64Comparator)

	for indexedLogReader.HasMessage() {
		currOffset, err := indexedLogReader.file.Seek(0, 1)
		if err != nil {
			return nil, errors.Wrap(err, "Could not create time index for file: "+filename)
		}
		msg, err := indexedLogReader.ReadMessage()
		if err != nil {
			// TODO(dschwab): On some of the log files I
			// get an unexpected EOF. Either log files are
			// bad, or I'm doing something wrong

			// return nil, err
			break
		}
		indexedLogReader.index.Put(msg.Timestamp, currOffset)
	}

	if _, err := logReader.file.Seek(0, 0); err != nil {
		return nil, errors.Wrap(err, "Failed to reset file offset to 0 for file: "+filename)
	}
	return
}

func (l *IndexedReader) NumMessages() int {
	return l.index.Size()
}

func (l *IndexedReader) CurrMessage() int {
	return l.curr_message
}

func (l *IndexedReader) Seek(offset int, whence int) (int, error) {
	new_offset := l.curr_message
	switch whence {
	case 0: // relative to origin
		new_offset = offset
	case 1: // relative to current offset
		new_offset = l.curr_message + offset
	case 2: // relative to end
		new_offset = l.NumMessages() + offset
	default:
		return l.curr_message, errors.New(fmt.Sprintf("Invalid `whence`. Must be [0, 1, 2]. Got %d", whence))
	}
	if new_offset < 0 && new_offset > l.NumMessages() {
		return l.curr_message, errors.New(fmt.Sprintf("Invalid offset %d. Out of range", new_offset))
	}
	seek_loc, ok := l.index.Values()[new_offset].(int64)
	if !ok {
		return l.curr_message, errors.New("Could not convert index value to int64")
	}
	if _, err := l.file.Seek(seek_loc, 0); err != nil {
		return l.curr_message, errors.Wrap(err, "Failed to seek file to specified offset")
	}
	l.curr_message = new_offset
	return l.curr_message, nil
}

func (l *IndexedReader) ReadIndexedMessage() (msg *Message, err error) {
	msg, err = l.ReadMessage()
	if err != nil {
		l.curr_message++
	}
	return
}
