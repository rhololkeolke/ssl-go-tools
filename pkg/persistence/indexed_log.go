package persistence

import (
	"fmt"

	"github.com/emirpasic/gods/trees/btree"
	"github.com/emirpasic/gods/utils"
	"github.com/pkg/errors"
)

type IndexedLog struct {
	// In order to support scrubbing, will need to have the log
	// file in memory. This is because gzip and bufio do not
	// support seaking in a file. Since the file is being read
	// into memory anyways, might as well pre-parse the messages

	// maps time offset from initial message to actual messages
	// for seeking
	index *btree.Tree

	iterator btree.Iterator
}

func NewIndexedLog(filename string) (indexedLog *IndexedLog, err error) {
	fmt.Println("DEBUG: Making indexed reader for file ", filename)
	indexedLog = new(IndexedLog)

	// TODO(dschwab): The following block is basically copied from
	// reader, but using the raw file so it is seekable. Will want
	// to merge this back into the main reader class, after
	// figuring out which pieces are common between the
	// structs. probably make the reader methods use the
	// *io.Reader interface, so that both os.File and bufio.Reader
	// can be passed in transparently
	reader, err := NewReader(filename)
	if err != nil {
		return nil, err
	}

	// time index for quick iterating, searching and jumping
	indexedLog.index = btree.NewWith(3, utils.Int64Comparator)

	var firstMessageTimestamp int64 = -1
	for reader.HasMessage() {
		msg, err := reader.ReadMessage()

		if err != nil {
			// TODO(dschwab): On some of the log files I
			// get an unexpected EOF. Either log files are
			// bad, or I'm doing something wrong

			// return nil, err
			break
		}

		if firstMessageTimestamp < 0 {
			firstMessageTimestamp = msg.Timestamp
		}
		indexedLog.index.Put(msg.Timestamp-firstMessageTimestamp, msg)
	}
	indexedLog.iterator = indexedLog.index.Iterator()

	return
}

func (l *IndexedLog) NumMessages() int {
	return l.index.Size()
}

func (l *IndexedLog) Index() []interface{} {
	return l.index.Keys()
}

func (l *IndexedLog) Messages() []interface{} {
	return l.index.Values()
}

func (l *IndexedLog) Reset() {
	l.iterator = l.index.Iterator()
}

func (l *IndexedLog) Next() (timestamp int64, msg *Message, err error) {
	if l.iterator.Next() {
		timestamp := l.iterator.Key().(int64)
		msg := l.iterator.Value().(*Message)
		return timestamp, msg, nil
	} else {
		return 0, nil, errors.New("No more messages")
	}
}

// func (l *IndexedLog) Seek(offset int, whence int) (int, error) {
// 	new_offset := l.curr_message
// 	switch whence {
// 	case 0: // relative to origin
// 		new_offset = offset
// 	case 1: // relative to current offset
// 		new_offset = l.curr_message + offset
// 	case 2: // relative to end
// 		new_offset = l.NumMessages() + offset
// 	default:
// 		return l.curr_message, errors.New(fmt.Sprintf("Invalid `whence`. Must be [0, 1, 2]. Got %d", whence))
// 	}
// 	if new_offset < 0 && new_offset > l.NumMessages() {
// 		return l.curr_message, errors.New(fmt.Sprintf("Invalid offset %d. Out of range", new_offset))
// 	}
// 	seek_loc, ok := l.index.Values()[new_offset].(int64)
// 	fmt.Printf("Seek loc: %d\n", seek_loc)
// 	if !ok {
// 		return l.curr_message, errors.New("Could not convert index value to int64")
// 	}
// 	if _, err := l.reader.Seek(seek_loc, 0); err != nil {
// 		return l.curr_message, errors.Wrap(err, "Failed to seek file to specified offset")
// 	}
// 	l.curr_message = new_offset
// 	return l.curr_message, nil
// }

// func (l *IndexedLog) ReadIndexedMessage() (msg *Message, err error) {
// 	msg, err = l.ReadMessage()
// 	if err != nil {
// 		l.curr_message++
// 	}
// 	return
// }
