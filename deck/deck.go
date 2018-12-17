package deck

import (
	"bufio"
	"encoding/base64"
	"encoding/binary"
	"io"
	"strings"
)

const (
	Version = 1

	MaxCard = 2
)

type Deck struct {
	Cards  []Card
	Heroes []Hero
	Format Format
}

type Format uint8

const (
	_ Format = iota
	UnknownFormat
	WildFormat
	StandardFormat
)

func DecodeString(s string) (d *Deck, err error) {
	return Decode(strings.NewReader(s))
}

func Decode(i io.Reader) (d *Deck, err error) {
	r := bufio.NewReader(
		base64.NewDecoder(base64.StdEncoding, i),
	)
	b, n, err := r.ReadRune()
	if err != nil {
		return nil, ErrInvalidDeckString.Wrap(err)
	}

	if n != 1 && b != '0' {
		return nil, ErrInvalidDeckString
	}

	version, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, ErrInvalidDeckString.Wrap(err)
	}

	if version != Version {
		return nil, ErrUnsupportedVersion
	}

	format, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, ErrInvalidDeckString.Wrap(err)
	}

	c, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, ErrInvalidDeckString.Wrap(err)
	}
	heroes := make([]Hero, c)
	for k, _ := range heroes {
		h, err := binary.ReadUvarint(r)
		if err != nil {
			return nil, ErrInvalidDeckString.Wrap(err)
		}
		heroes[k] = Hero(h)
	}

	var cards []Card
	for j := 1; j <= MaxCard+1; j++ {
		v, err := binary.ReadUvarint(r)
		if err != nil {
			return nil, ErrInvalidDeckString.Wrap(err)
		}

		for k := 0; k < int(v); k++ {
			v, err := binary.ReadUvarint(r)
			if err != nil {
				return nil, ErrInvalidDeckString.Wrap(err)
			}
			count, err := cardCount(j, r)
			if err != nil {
				return nil, ErrInvalidDeckString.Wrap(err)
			}
			cards = append(cards, Card{DbfID: v, Count: count})
		}
	}

	//read aditional xn data

	return &Deck{
		Heroes: heroes,
		Cards:  cards,
		Format: Format(format),
	}, nil
}

func cardCount(seq int, r io.ByteReader) (int, error) {
	if seq <= 2 {
		return seq, nil
	}

	v, err := binary.ReadUvarint(r)
	if err != nil {
		return seq, err
	}

	return int(v), nil
}
