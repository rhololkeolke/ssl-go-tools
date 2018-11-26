package persistence

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"github.com/emirpasic/gods/trees/btree"
	"github.com/emirpasic/gods/utils"	
	"github.com/pkg/errors"
	"io"
	"os"
)

const channelBufferSize = 100

type Reader struct {
	file   *os.File
	reader *bufio.Reader
	// TODO(dschwab): Move this to it's own class IndexedReader,
	// and use a Reader as an anonymous Field. This should allow
	// all the public methods to be called by the IndexedReader
	index  *btree.Tree
}

func NewReader(filename string) (logReader *Reader, err error) {
	logReader = new(Reader)
	logReader.file, err = os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Could not open log file: "+filename)
	}

	if filename[len(filename)-2:] == "gz" {
		gzipReader, err := gzip.NewReader(logReader.file)
		if err != nil {
			return nil, err
		}
		logReader.reader = bufio.NewReader(gzipReader)
	} else {
		logReader.reader = bufio.NewReader(logReader.file)
	}
	logReader.verifyLogFile()

	// create the time index to allow for querying and jumping to
	// specific time codes
	logReader.index = btree.NewWith(3, utils.Int64Comparator)

	for logReader.HasMessage() {
		currOffset, err := logReader.file.Seek(0, 1)
		if err != nil {
			return nil, errors.Wrap(err, "Could not create time index for file: "+filename)
		}
		msg, err := logReader.ReadMessage()
		if err != nil {
			// TODO(dschwab): On some of the log files I
			// get an unexpected EOF. Either the log files
			// are bad, or I'm doing something wrong
			
			// return nil, err
			break
		}
		logReader.index.Put(msg.Timestamp, currOffset)
	}
	if _, err := logReader.file.Seek(0, 0); err != nil {
		return nil, errors.Wrap(err, "Failed to reset file offset to 0 for file: "+filename)
	}

	return
}

func (l *Reader) Close() error {
	return l.file.Close()
}

func (l *Reader) NumMessages() int {
	return l.index.Size()
}

func (l *Reader) HasMessage() bool {
	_, err := l.reader.Peek(1)
	return err != io.EOF
}

func (l *Reader) ReadMessage() (msg *Message, err error) {
	msg = new(Message)
	msg.Timestamp, err = l.readInt64()
	if err != nil {
		return
	}
	messageId, err := l.readInt32()
	msg.MessageType.Id = MessageId(messageId)
	if err != nil {
		return
	}
	length, err := l.readInt32()
	if err != nil {
		return
	}
	if length < 0 {
		err = errors.New(fmt.Sprintf("Length is invalid: %d", length))
		return
	}
	msg.Message, err = l.readBytes(int(length))
	if err != nil {
		return
	}
	return
}

func (l *Reader) CreateChannel() (channel chan *Message) {
	channel = make(chan *Message, channelBufferSize)
	go l.readToChannel(channel)
	return
}

func (l *Reader) readToChannel(channel chan *Message) {
	for l.HasMessage() {
		msg, err := l.ReadMessage()
		if err != nil {
			break
		}

		channel <- msg
	}
	close(channel)
}

func (l *Reader) verifyLogFile() error {
	header, err := l.readString(12)
	if err != nil {
		return err
	}
	if header != "SSL_LOG_FILE" {
		return errors.New("header validation failed. Found: " + header)
	}

	version, err := l.readInt32()
	if err != nil {
		return err
	}
	if version != 1 {
		return errors.New(fmt.Sprintf("unsupported version: %d", version))
	}
	return err
}

func (l *Reader) readBytes(length int) ([]byte, error) {
	byteSlice := make([]byte, length)
	_, err := io.ReadAtLeast(l.reader, byteSlice, length)

	return byteSlice, err
}

func (l *Reader) readString(length int) (string, error) {
	s, err := l.readBytes(length)
	return string(s), err
}

func (l *Reader) readInt32() (ret int32, err error) {
	err = binary.Read(l.reader, binary.BigEndian, &ret)
	return
}

func (l *Reader) readInt64() (ret int64, err error) {
	err = binary.Read(l.reader, binary.BigEndian, &ret)
	return
}
