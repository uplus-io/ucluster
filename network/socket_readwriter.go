package network

import (
	"bufio"
	"net"
)

type SocketReadWriter struct {
	connection        net.Conn
	reader            *Reader
	writer            *Writer
	readerInitialized bool
	writerInitialized bool
}

func NewReadWriter(conn net.Conn) *SocketReadWriter {
	readWriter := &SocketReadWriter{
		connection: conn,
		reader:     &Reader{connection: conn},
		writer:     &Writer{connection: conn},
	}
	return readWriter
}

func (p *SocketReadWriter) Reader(spliter Splitable) *Reader {
	if !p.readerInitialized {
		p.reader.init(spliter)
		p.readerInitialized = true
	}
	return p.reader
}

func (this *SocketReadWriter) Writer() *Writer {
	if !this.writerInitialized {
		this.writerInitialized = true
	}
	return this.writer
}

type Reader struct {
	Spliter    Splitable
	connection net.Conn
	scanner    *bufio.Scanner
}

func (r *Reader) init(spliter Splitable) {
	r.Spliter = spliter
	scanner := bufio.NewScanner(r.connection)
	scanner.Split(r.Spliter.Split)
	r.scanner = scanner
}

func (r *Reader) HasNext() bool {
	return r.scanner.Scan()
}

func (r *Reader) Err() error {
	return r.scanner.Err()
}

func (r *Reader) Read() []byte {
	return r.scanner.Bytes()
}

type Writer struct {
	connection net.Conn
}

func (r *Writer) Write(bytes []byte) (int, error) {
	return r.connection.Write(bytes)
}
